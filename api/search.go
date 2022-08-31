package api

import (
	"fmt"
	"termusic/api/model"
)

const (
	SEARCH_SONG   int64 = 1
	SEARCH_ARTIST int64 = 100
)

func SearchSong(keyword string, offset int64) (model.SearchResponse, error) {
	requestUrl := fmt.Sprintf(API+"/cloudsearch?keywords=%s&type=%d&offset=%d", keyword, SEARCH_SONG, offset)
	var data model.SearchResponse
	err := Get(requestUrl, &data)
	return data, err
}
