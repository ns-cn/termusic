package model

type SearchResponse struct {
	Code   int16                `json:"code"`   // 响应结果
	Result SearchResponseResult `json:"result"` // 结果详情
}
type SearchResponseResult struct {
	Songs     []SearchResponseResultSong `json:"songs"`     // 歌曲集合
	SongCount int64                      `json:"songCount"` // 歌曲总数
}

type SearchResponseResultSong struct {
	Id      int64            `json:"id"`   // id
	Name    string           `json:"name"` // 歌曲名称
	Artists []ResponseArtist `json:"ar"`   // 歌手集合
}
