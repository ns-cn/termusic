package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	mixer := MixStreamer{}
	res, err := http.Get("https://music.163.com/song/media/outer/url?id=33894312.mp3")
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
		nextMusic, err := http.Get("https://music.163.com/song/media/outer/url?id=33894312.mp3")
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

func test() {
	res, err := http.Get("https://music.163.com/song/media/outer/url?id=33894312.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     100,
		Volume:   0,
		Silent:   false,
	}
	speedy := beep.ResampleRatio(1, 1, volume)
	speaker.Play(speedy)

	for {
		fmt.Print("Press [ENTER] to pause/resume. ")
		fmt.Scanln()

		speaker.Lock()
		//ctrl.Paused = !ctrl.Paused
		volume.Volume += 0.05
		speedy.SetRatio(speedy.Ratio() + 0.01)
		speaker.Unlock()
	}
}
