package model

type UrlResponse struct {
	Code int16     `json:"code"` // 响应编码
	Data []Url `json:"data"` // 地址明细
}

type Url struct {
	Id  int64  `json:"id"`  // 歌曲地址id
	Br  int64  `json:"br"`  // 歌曲编码
	Url string `json:"url"` // 歌曲地址
}
