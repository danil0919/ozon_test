package model

import "testing"

func TestLink(t *testing.T) *Link {
	return &Link{
		Link:  "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/",
		Token: "lol",
	}
}
