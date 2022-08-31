package model

type ResponseArtist struct {
	Id        int64  `json:"id"`        // id
	Name      string `json:"name"`      // 歌手名称
	AlbumSize int64  `json:"albumSize"` // 专辑数量
}
