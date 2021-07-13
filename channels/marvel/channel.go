package marvel

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Channel struct {
	key    string
	secret string
}

func New() *Channel {
	return &Channel{
		key:    os.Getenv("MARVEL_PUBLIC_KEY"),
		secret: os.Getenv("MARVEL_PRIVATE_KEY"),
	}
}

// Authentication for get method
// will append auth to ANY given URL
func (c *Channel) AppendGETAuthentication(targetURL *url.URL) {
	timestamp := time.Now().Unix()
	compositeKey := fmt.Sprintf("%v%s%s", timestamp, c.secret, c.key)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(compositeKey)))

	params := targetURL.Query()
	params.Add("ts", strconv.FormatInt(timestamp, 10))
	params.Add("apikey", c.key)
	params.Add("hash", hash)

	targetURL.RawQuery = params.Encode()
}

func (c *Channel) AppendOffsetLimit(targetURL *url.URL, offset int, limit int) {
	params := targetURL.Query()
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))
	targetURL.RawQuery = params.Encode()
}
