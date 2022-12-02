package bak

import "fmt"

const (
	LOOP_NONE = iota
	LOOP_SINGLE
	LOOP_ALL
	LOOP_CYCLE
)

type Player struct {
	queue []Music // 播放队列
	loop  int     // 循环方式
}

// AddMusic 添加歌曲
func (p *Player) AddMusic(music Music) error {
	return nil
}

// RemoveMusic 移除指定的歌曲
func (p *Player) RemoveMusic(music Music) error {
	return nil
}

// RemoveAllMusic 移除所有的歌曲
func (p *Player) RemoveAllMusic(music Music) error {
	return nil
}

// Play 播放
func (p *Player) Play(loop int) error {
	return nil
}

// Pause 暂停
func (p *Player) Pause() error {
	return nil
}

// PlayBackward 播放上一个
func (p *Player) PlayBackward() error {
	return p.playNear(false)
}

// PlayForward 播放下一个
func (p *Player) PlayForward() error {
	return p.playNear(true)
}

// playNear 播放前一个或则后一个,参数指定向前还是向后
func (p *Player) playNear(forward bool) error {
	return nil
}

// SetLoop 设置循环方式
func (p *Player) SetLoop(loop int) error {
	if loop < LOOP_NONE || loop > LOOP_CYCLE {
		return fmt.Errorf("wrong loop type!")
	}
	p.loop = loop
	return nil
}
