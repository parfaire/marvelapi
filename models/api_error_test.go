package models

import (
	"net/http"
	"testing"

	"github.com/parfaire/marvelapi/util"
)

func TestApiErrorGetMessage(t *testing.T) {
	errorMsg := "Kabooomm... Error detected."
	apiError := NewApiError(errorMsg)

	if apiError.Error() != errorMsg {
		t.Errorf("Got = %v | Expected =%v", apiError.Error(), errorMsg)
	}
	if apiError.Message != errorMsg {
		t.Errorf("Got = %v | Expected =%v", apiError.Message, errorMsg)
	}
	if apiError.HTTPStatus != http.StatusInternalServerError {
		t.Errorf("Got = %v | Expected =%v", apiError.HTTPStatus, http.StatusInternalServerError)
	}
}

func TestCustomApiError(t *testing.T) {
	customApiError := ErrorNotFound()
	if customApiError.Message != util.ERROR_NOT_FOUND {
		t.Errorf("Got = %v | Expected =%v", customApiError.Message, util.ERROR_NOT_FOUND)
	}
	if customApiError.HTTPStatus != http.StatusNotFound {
		t.Errorf("Got = %v | Expected =%v", customApiError.HTTPStatus, http.StatusNotFound)
	}

	customApiError = ErrorInternalServer()
	if customApiError.Message != util.ERROR_INTERNAL_SERVER {
		t.Errorf("Got = %v | Expected =%v", customApiError.Message, util.ERROR_INTERNAL_SERVER)
	}
	if customApiError.HTTPStatus != http.StatusInternalServerError {
		t.Errorf("Got = %v | Expected =%v", customApiError.HTTPStatus, http.StatusInternalServerError)
	}
}
