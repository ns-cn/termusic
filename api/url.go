package api

import (
	"fmt"
	"termusic/api/model"
)

const (
	API = "http://api.tangyujun.com"
)


func GetUrl(id int32) (model.UrlResponse, error) {
	requestUrl := fmt.Sprintf(API+"/song/url?id=%d&realIP=116.25.146.177", id)
	var data model.UrlResponse
	err := Get(requestUrl, &data)
	return data, err
}
