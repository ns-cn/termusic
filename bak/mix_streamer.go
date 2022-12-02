package bak

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	"time"
)

type MixStreamer struct {
	Original *beep.StreamSeekCloser
	Ctrl     *beep.Ctrl
	Volume   *effects.Volume
	Built    *beep.Resampler
	Progress chan float32
	isInit   bool
}

func (s *MixStreamer) Err() error {
	return s.Built.Err()
}

func (s *MixStreamer) Play(input *beep.StreamSeekCloser, format beep.Format) error {
	s.Progress = make(chan float32, 1)
	if s.Original != nil {
		err := (*s.Original).Close()
		if err != nil {
			return err
		}
	}
	speaker.Clear()
	s.Original = input
	s.Ctrl = &beep.Ctrl{Streamer: *s.Original, Paused: false}
	s.Volume = &effects.Volume{
		Streamer: s.Ctrl,
		Base:     100,
		Volume:   0,
		Silent:   false,
	}
	s.Built = beep.ResampleRatio(1, 1, s.Volume)
	if !s.isInit {
		err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			return err
		}
		s.isInit = true
	}
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		s.Progress <- 100
	})))
	return nil
}
func (s *MixStreamer) Stream(samples [][2]float64) (n int, ok bool) {
	n, ok = s.Built.Stream(samples)
	if !ok {
		//s.Progress <- 100
	}
	return
}

func (s *MixStreamer) Pause() {
	if s.Ctrl != nil {
		s.Ctrl.Paused = true
	}
}

func (s MixStreamer) Continue() {
	if s.Ctrl != nil {
		s.Ctrl.Paused = false
	}
}
