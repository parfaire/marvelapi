package marvel

import "net/http"

type Interface interface {
	GetCharacters(offset int, limit int) (resp *http.Response, err error)
	GetCharacterById(id string) (resp *http.Response, err error)
}
