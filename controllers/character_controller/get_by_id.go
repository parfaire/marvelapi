package character_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/parfaire/marvelapi/models"
)

func (c CharacterController) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) (apiError *models.ApiError) {
	character, apiError := c.getById(params.ByName("id"))
	if apiError != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(character)

	return
}

func (c CharacterController) getById(id string) (character *models.Character, apiError *models.ApiError) {
	// get character from channel (third party)
	resp, err := c.MarvelChannel.GetCharacterById(id)
	if err != nil {
		return nil, models.NewApiError(err.Error())
	}
	if resp.StatusCode > 299 {
		log.Print(resp)
		return nil, models.ErrorInternalServer()

	}

	// normalise character JSON to model
	characterDW := models.CharacterDataWrapper{}
	err = json.NewDecoder(resp.Body).Decode(&characterDW)

	if err != nil {
		return nil, models.NewApiError(err.Error())
	} else if len(characterDW.Data.Results) < 1 {
		return nil, models.ErrorNotFound()
	}
	return &characterDW.Data.Results[0], nil
}
