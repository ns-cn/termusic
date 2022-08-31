package api

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	url, err := SearchSong("海阔天空", 0)
	if err != nil {
		return
	}
	fmt.Println(url)
}
