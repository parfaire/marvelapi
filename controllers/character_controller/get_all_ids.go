package character_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/parfaire/marvelapi/models"
	"github.com/parfaire/marvelapi/util"
)

func (c CharacterController) GetAllIDs(writer http.ResponseWriter, request *http.Request, params httprouter.Params) (apiError *models.ApiError) {
	ids, apiError := c.getAllIDs()
	if apiError != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(ids)

	return
}

func (c CharacterController) getAllIDs() (ids []int, apiError *models.ApiError) {

	// get from redis, if exists return right away
	ids, err := c.getIDsFromRedis()
	if err == nil && len(ids) > 0 {
		return
	}

	// non-redis flow, manual fetch to channel (third party)
	// first batch
	idsInBatch, total, apiError := c.getCharacterIDsByBatch(1)
	if apiError != nil {
		return
	}
	ids = append(ids, idsInBatch...)

	// following batches
	totalBatch := total / util.MARVEL_MAX_GET
	if total%util.MARVEL_MAX_GET > 0 {
		totalBatch++ // last batch for the remaining
	}
	// get batches in concurrent-manner
	var wg sync.WaitGroup
	for i := 2; i <= totalBatch; i++ {
		wg.Add(1)
		go func(batchNo int) {
			defer wg.Done()
			idsInBatch, _, apiError = c.getCharacterIDsByBatch(batchNo)
			if apiError != nil {
				log.Println(apiError.Error()) // non-breaking error
			}
			ids = append(ids, idsInBatch...)
		}(i)
	}
	wg.Wait()

	// set to redis
	c.setIDsToRedis(ids)
	return
}

func (c CharacterController) getIDsFromRedis() (ids []int, err error) {
	idsString, err := c.RedisClient.Get(util.REDIS_KEY_ALL_CHARACTER_IDS).Result()
	if err != nil {
		return
	}
	bytes := []byte(idsString)
	err = json.Unmarshal(bytes, &ids)
	return
}

func (c CharacterController) setIDsToRedis(ids []int) {
	idsJson, _ := json.Marshal(ids)
	err := c.RedisClient.Set(util.REDIS_KEY_ALL_CHARACTER_IDS, idsJson, 8*time.Hour).Err()
	if err != nil {
		log.Println(err.Error()) // non-breaking error
	}
}

// batchNo starts from 1
func (c CharacterController) getCharacterIDsByBatch(batchNo int) (ids []int, total int, apiError *models.ApiError) {
	offset := (batchNo - 1) * util.MARVEL_MAX_GET
	resp, err := c.MarvelChannel.GetCharacters(offset, util.MARVEL_MAX_GET)
	if err != nil {
		return nil, 0, models.NewApiError(err.Error())
	}
	if resp.StatusCode > 299 {
		log.Print(resp)
		return nil, 0, models.ErrorInternalServer()

	}

	// normalise character JSON to model
	characterDW := models.CharacterDataWrapperIdOnly{}
	err = json.NewDecoder(resp.Body).Decode(&characterDW)
	if err != nil {
		return nil, 0, models.NewApiError(err.Error())
	}

	// strip result to IDs only
	ids = make([]int, len(characterDW.Data.Results))
	for index, character := range characterDW.Data.Results {
		ids[index] = character.Id
	}

	// get total character ids
	total = characterDW.Data.Total
	return
}
