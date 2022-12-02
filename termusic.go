package main

import (
	"fmt"
	"github.com/CaoYnag/gocui"
	"log"
	"termusic/api"
	"time"
)

var (
	playing bool = false
)

func main() {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Cursor = true
	g.Mouse = true
	g.InputEsc = true
	g.SetManagerFunc(layout)

	go func() {
		for {
			select {
			case time := <-time.After(time.Second):
				g.Update(func(gui *gocui.Gui) error {
					timebar, err := gui.View("timebar")
					if err != nil {
						return err
					}
					timebar.Clear()
					_, _ = fmt.Fprint(timebar, fmt.Sprintf("00:%s", time.Format("05")))
					return nil
				})
			}

		}
	}()
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "termusic"
	}
	if _, err := g.SetView("network", 1, 3, maxX/2-1, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	if _, err := g.SetView("status", 1, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	pre, _ := g.SetView("pre", 1, maxY-3, 5, maxY-1)
	play, _ := g.SetView("play", 5, maxY-3, 11, maxY-1)
	next, _ := g.SetView("next", 11, maxY-3, 15, maxY-1)
	g.SetView("timebar", 15, maxY-3, maxX-1, maxY-1)
	pre.Clear()
	next.Clear()
	fmt.Fprint(pre, "<<")
	fmt.Fprint(next, ">>")
	debug, _ := g.SetView("debug", maxX-40, maxY-3, maxX-1, maxY-1)
	debug.Frame = false
	if search, err := g.SetView("search", 1, 1, maxX/2-1, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_ = g.SetKeybinding("", gocui.KeyCtrlF, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
			if _, err := g.SetCurrentView("search"); err != nil {
				return err
			}
			return nil
		})
		_ = g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
			searchText := view.ViewBuffer()
			network, err := g.View("network")
			if err != nil {
				return err
			}
			network.Clear()
			search.Title = fmt.Sprintf("搜索：%s", searchText)
			songs, _ := api.SearchSong(searchText, 0)
			for _, song := range songs.Result.Songs {
				_, _ = fmt.Fprintln(network, fmt.Sprintf("%13d\t%-20s\t\t%v", song.Id, song.Name, song.Artists))
			}
			_, _ = g.SetCurrentView("network")
			return nil
		})
		search.Title = "搜索"
		search.Editable = true
	}
	_ = g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		_, _ = gui.SetCurrentView(view.Name())
		debug, err := gui.View("debug")
		if err != nil {
			return err
		}
		debug.Clear()
		_, err = fmt.Fprint(debug, fmt.Sprintf("debug:%s:%v", view.Name(), time.Now()))
		if err != nil {
			return err
		}
		if view.Name() == "play" {
			playing = !playing
			play.Clear()
			if playing {
				fmt.Fprint(play, "pause")
			} else {
				fmt.Fprint(play, "play")
			}
		}
		return nil
	})
	_ = g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		playing = !playing
		play.Clear()
		if playing {
			fmt.Fprint(play, "pause")
		} else {
			fmt.Fprint(play, "play")
		}
		return nil
	})
	if local, err := g.SetView("local", maxX/2, 1, maxX-1, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		local.Title = "待播清单"
	}
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
