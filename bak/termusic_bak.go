package bak

import (
	"fmt"
	"log"
	"net/http"
	"termusic/api"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main2() {
	mixer := MixStreamer{}
	song, err := api.SearchSong("梦 温淑娴", 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
	res, err := http.Get("https://music.163.com/song/media/outer/url?id=1872856513.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	err = mixer.Play(&streamer, format)
	if err != nil {
		fmt.Println(err)
		return
	}
	//go func() {
	//	pause := true
	//	for {
	//		time.Sleep(2 * time.Second)
	//		if pause {
	//			mixer.Pause()
	//		} else {
	//			mixer.Continue()
	//		}
	//		pause = !pause
	//	}
	//}()

	go func() {
		time.Sleep(3 * time.Second)
		nextMusic, err := http.Get("https://music.163.com/song/media/outer/url?id=1872856513.mp3")
		if err != nil {
			log.Fatal(err)
		}
		nextStreamer, format, err := mp3.Decode(nextMusic.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = mixer.Play(&nextStreamer, format)
		if err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		for {
			select {
			case <-mixer.Progress:
				fmt.Println("done")
				return
			case <-time.After(time.Second):
				speaker.Lock()
				fmt.Println(format.SampleRate.D((*mixer.Original).Position()).Round(time.Second))
				speaker.Unlock()
			}
		}
	}()
	for {
		progress := <-mixer.Progress
		fmt.Println(progress)
	}
}
