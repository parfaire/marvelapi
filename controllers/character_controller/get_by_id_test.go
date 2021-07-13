package character_controller

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"encoding/json"

	"github.com/parfaire/marvelapi/channels/marvel"
	"github.com/parfaire/marvelapi/models"
)

func TestGetById(t *testing.T) {
	// prepare mock
	marvelChannel := new(marvel.MockInterface)
	characterController := New(nil, marvelChannel)
	expectedChar := models.Character{
		Id:          1,
		Name:        "Hero",
		Description: "Superhero!",
	}
	mockRespBodyObj :=
		models.CharacterDataWrapper{
			Data: models.CharacterDataContainer{
				Results: []models.Character{
					expectedChar,
				},
			},
		}
	mockRespBodyJson, _ := json.Marshal(mockRespBodyObj)
	mockId := "1"
	resp := http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(mockRespBodyJson)),
	}

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Success",
			mock: func() {
				marvelChannel.On("GetCharacterById", mockId).Return(&resp, nil).Once()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			gotChar, err := characterController.getById(mockId)
			if (err != nil) != tt.wantErr {
				t.Errorf("get error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotChar.Description != expectedChar.Description ||
				gotChar.Name != expectedChar.Name || gotChar.Id != expectedChar.Id {
				t.Errorf("get = %+v, want = %+v", gotChar, expectedChar)
			}
		})
	}
}
