package character_controller

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"encoding/json"

	"bou.ke/monkey"
	"github.com/parfaire/marvelapi/channels/marvel"
	"github.com/parfaire/marvelapi/models"
	"github.com/parfaire/marvelapi/util"
)

func TestGetAllIdsSingleBatch(t *testing.T) {
	// prepare mock
	marvelChannel := new(marvel.MockInterface)
	characterController := New(nil, marvelChannel)
	expectedIds := []int{100, 200, 300}
	arrCharacterIdOnly := []models.CharacterIdOnly{
		{Id: expectedIds[0]},
		{Id: expectedIds[1]},
		{Id: expectedIds[2]},
	}
	mockRespBodyObj :=
		models.CharacterDataWrapperIdOnly{
			Data: models.CharacterDataContainerIdOnly{
				Results: arrCharacterIdOnly,
				Total:   len(expectedIds),
			},
		}
	mockRespBodyJson, _ := json.Marshal(mockRespBodyObj)
	resp := http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(mockRespBodyJson)),
	}

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Success - non redis - single batch",
			mock: func() {
				marvelChannel.On("GetCharacters", 0, util.MARVEL_MAX_GET).Return(&resp, nil)

				monkey.Patch(CharacterController.getIDsFromRedis, func(CharacterController) (ids []int, err error) {
					return nil, errors.New("a")
				})
				monkey.Patch(CharacterController.setIDsToRedis, func(CharacterController, []int) {})
			},
			wantErr: false,
		},
		{
			name: "Success - with redis - single batch",
			mock: func() {
				marvelChannel.On("GetCharacters", 0, util.MARVEL_MAX_GET).Return(&resp, nil)

				monkey.Patch(CharacterController.getIDsFromRedis, func(CharacterController) (ids []int, err error) {
					return expectedIds, nil
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			gotIds, err := characterController.getAllIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("get error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(gotIds) != len(expectedIds) {
				t.Errorf("get = %+v, want = %+v", len(gotIds), len(expectedIds))
			}
			for i, v := range gotIds {
				if v != expectedIds[i] {
					t.Errorf("get = %+v, want = %+v", v, expectedIds[i])
				}
			}
			monkey.UnpatchAll()
		})
	}
}

func TestGetAllIdsMultipleBatch(t *testing.T) {
	// prepare mock
	marvelChannel := new(marvel.MockInterface)
	characterController := New(nil, marvelChannel)
	total := 180
	expectedIds := make([]int, total)
	arrCharacterIdOnly := make([]models.CharacterIdOnly, total)
	for i := 0; i < total; i++ {
		expectedIds[i] = i
		arrCharacterIdOnly[i] = models.CharacterIdOnly{Id: i}
	}
	mockRespBodyObj1 :=
		models.CharacterDataWrapperIdOnly{
			Data: models.CharacterDataContainerIdOnly{
				Results: arrCharacterIdOnly[:util.MARVEL_MAX_GET],
				Total:   total,
			},
		}
	mockRespBodyObj2 :=
		models.CharacterDataWrapperIdOnly{
			Data: models.CharacterDataContainerIdOnly{
				Results: arrCharacterIdOnly[util.MARVEL_MAX_GET:],
				Total:   total,
			},
		}
	mockRespBodyJson1, _ := json.Marshal(mockRespBodyObj1)
	mockRespBodyJson2, _ := json.Marshal(mockRespBodyObj2)
	resp1 := http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(mockRespBodyJson1)),
	}
	resp2 := http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(mockRespBodyJson2)),
	}

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Success - non redis - multiple batch",
			mock: func() {
				marvelChannel.On("GetCharacters", 0, util.MARVEL_MAX_GET).Return(&resp1, nil)
				marvelChannel.On("GetCharacters", util.MARVEL_MAX_GET, util.MARVEL_MAX_GET).Return(&resp2, nil)

				monkey.Patch(CharacterController.getIDsFromRedis, func(CharacterController) (ids []int, err error) {
					return nil, errors.New("a")
				})
				monkey.Patch(CharacterController.setIDsToRedis, func(CharacterController, []int) {})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			gotIds, err := characterController.getAllIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("get error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(gotIds) != len(expectedIds) {
				t.Errorf("get = %+v, want = %+v", len(gotIds), len(expectedIds))
			}
			for i, v := range gotIds {
				if v != expectedIds[i] {
					t.Errorf("get = %+v, want = %+v", v, expectedIds[i])
				}
			}
			monkey.UnpatchAll()
		})
	}
}
