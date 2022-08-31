package api

import (
	"fmt"
	"testing"
)

func TestGetUrl(t *testing.T){
	url, err := GetUrl(1970559943)
	if err != nil {
		return
	}
	fmt.Println(url)
}