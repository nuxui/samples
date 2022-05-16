// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"nuxui.org/nuxui/log"
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/ui"
)

type Home interface {
	nux.Component
}

type home struct {
	*nux.ComponentBase

	content      nux.Widget
	pictureNames []string
	pictureIndex int
	pictureDir   string
	picture      ui.Image
	rootView     ui.Column
}

func NewHome(attr nux.Attr) Home {
	me := &home{
		picture:      ui.NewImage(attr),
		pictureNames: []string{},
	}
	me.ComponentBase = nux.NewComponentBase(me, attr)
	me.content = nux.InflateLayout(me, me.template(), nil)
	return me
}

func (me *home) template() string {
	return `
{
  import: {
    ui: nuxui.org/nuxui/ui,
  },

  layout: {
    id: "root-row",
    type: ui.Column,
    width: 1wt,
    height: 1wt,
    background: #00ff00,
    padding: {left: 10%, top: 10px, right: 10%, bottom: 10px},
    children:[
		{
			id: "title",
			type: ui.Text,
			textSize: 25,
			text: "show current directory pictures",
			font: {family: "Consolas, Courier New, monospace", size: 14, color: #000000 },
		},
      {
        id: "img-preview",
        type: ui.Image,
        width: 1wt,
        height: 1wt,
        background: #ffffff,
        src: "a1.jpg",
        scaleType: "centerInside",
      }
    ],
  }
}
  
  `
}

func (me *home) Mount() {
	// nux.App().MainWindow().SetAlpha(0.5)

	me.picture = nux.FindChild(me, "img-preview").(ui.Image)
	nux.OnScrollY(me.picture, me.onScrollY)
	// me.picture.SetSrc("~/Downloads")
	me.getDirAllPicWitchOnePicName("./a")
	nux.App().MainWindow().SetTitle("win32 title ok ")
	log.V("home", "========= window tilte = %s", nux.App().MainWindow().Title())

	me.rootView = nux.FindChild(me, "root-row").(ui.Column)
	nux.OnTap(me.rootView, me.onTap)
}

func (me *home) onScrollY(detail nux.GestureDetail) {
	log.V("home", "onScrollY %f", detail.ScrollY())
	// me.picture.SetScaleType()
	// var scale float32 = 0.1
	// if detail.ScrollY() < 0 {
	// 	scale = -0.1
	// }

	// me.picture.SetScale(me.picture.Scale() + scale)
}

func (me *home) getDirAllPicWitchOnePicName(src string) {
	// log.V("home", "getDirAllPicWitchOnePicName dir = %s", filepath.Dir(src))
	abs, _ := filepath.Abs(src)
	srcBase := filepath.Base(src)
	me.pictureDir = filepath.Dir(abs)

	// log.V("home", "@@@@@@@@ getDirAllPicWitchOnePicName src=%s, abs = %s, base = %s, clean= %s dir=%s",
	// src, abs, filepath.Base(src), filepath.Clean(src), filepath.Dir(abs))

	me.pictureNames = []string{}
	dirs, err := ioutil.ReadDir(filepath.Dir(src))
	if err != nil {
		log.E("home", "%s", err.Error())
		return
	}

	i := 0
	for _, p := range dirs {
		ext := strings.ToLowerSpecial(unicode.TurkishCase, filepath.Ext(p.Name()))
		// if !p.IsDir() && (ext == ".png") {
		if !p.IsDir() && (ext == ".jpg" || ext == ".jpeg" || ext == ".png") {
			// if !p.IsDir() && (ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".svg") {
			log.V("home", "dirs = %s", p.Name())
			me.pictureNames = append(me.pictureNames, p.Name())
			if strings.Compare(srcBase, p.Name()) == 0 {
				me.pictureIndex = i
			}
			i++
		}
	}

	me.picture.SetSrc(fmt.Sprintf("%s%c%s", me.pictureDir, filepath.Separator, me.pictureNames[0]))

	log.I("gallery", "pictureNames: %s", me.pictureNames)

}

func (me *home) onTap(detail nux.GestureDetail) {
	// sx, sy := nux.GetCursorScreenPosition()
	// tsX, tsY := nux.CursorPositionWindowToScreen(nux.App().MainWindow(), detail.X(), detail.Y())
	// twX, twY := nux.CursorPositionScreenToWindow(nux.App().MainWindow(), sx, sy)
	// x:1680,y:1050
	// log.V("home", "onTap = %s, wX=%f, wY=%f, sX=%f, sY=%f, tsX=%f, tsY=%f, twX=%f, twY=%f", detail.Target().ID(),
	// detail.X(), detail.Y(), sx, sy, tsX, tsY, twX, twY)
	nux.App().MainWindow().SetTitle("click to set title")
}

func (me *home) onDoubleTap(w nux.Widget) {
	// log.V("home", "onDoubleTap = %s", w.ID())
}

func (me *home) OnKeyEvent(event nux.KeyEvent) bool {
	log.V("home", "OnKeyEvent Rune: %c , code: 0x%X, %s", event.Rune(), uint32(event.KeyCode()), event.KeyCode())
	if event.Action() == nux.Action_Down {
		switch event.KeyCode() {
		case nux.Key_Left:
			log.V("home", "OnKeyEvent = KeyCode_Left %d", me.pictureIndex)
			if me.pictureIndex > 0 {
				me.pictureIndex--
				me.picture.SetSrc(fmt.Sprintf("%s%c%s", me.pictureDir, filepath.Separator, me.pictureNames[me.pictureIndex]))
			}
			log.V("home", "OnKeyEvent = src = %s", me.picture.Src())
			runtime.GC()
			me.rootView.SetBackgroundColor(nux.Color(rand.Uint32()))
		case nux.Key_Right:
			log.V("home", "OnKeyEvent = Key_Right %d", me.pictureIndex)
			if me.pictureIndex < len(me.pictureNames)-1 {
				me.pictureIndex++
				imgpath := fmt.Sprintf("%s%c%s", me.pictureDir, filepath.Separator, me.pictureNames[me.pictureIndex])
				me.picture.SetSrc(imgpath)
				log.V("home", "OnKeyEvent = src = %s, %s", me.picture.Src(), imgpath)
			}
			runtime.GC()
		}

	}
	return false
}
