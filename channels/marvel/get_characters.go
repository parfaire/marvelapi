package marvel

import (
	"net/http"
	"net/url"
)

func (c Channel) GetCharacters(offset int, limit int) (resp *http.Response, err error) {
	target := "http://gateway.marvel.com/v1/public/characters"
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	c.AppendOffsetLimit(targetURL, offset, limit)
	c.AppendGETAuthentication(targetURL)

	return http.Get(targetURL.String())
}
