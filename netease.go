package main

import "termusic/api"

type Music struct {
	Id   int32    // id
	Name string   // 歌曲名称
	Auth []string // 作者集合
	Url  string   // 歌曲地址
}

func (m *Music) UpdateUrl() error {
	api.GetUrl(m.Id)
	return nil
}
