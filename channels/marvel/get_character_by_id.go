package marvel

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c Channel) GetCharacterById(id string) (resp *http.Response, err error) {
	target := fmt.Sprintf("http://gateway.marvel.com/v1/public/characters/%s", id)
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	c.AppendGETAuthentication(targetURL)
	fmt.Println(targetURL.String())
	return http.Get(targetURL.String())
}
