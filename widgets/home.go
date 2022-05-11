// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"nuxui.org/nuxui/log"
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/ui"
	_ "nuxui.org/nuxui/ui"
)

type Home interface {
	nux.Component
}

type home struct {
	*nux.ComponentBase

	content nux.Widget
}

func NewHome(attr nux.Attr) Home {
	me := &home{}
	me.ComponentBase = nux.NewComponentBase(me, attr)
	me.content = nux.InflateLayout(me, me.template())
	return me
}

func (me *home) template() string {
	return `
{
  import: {
    ui: nuxui.org/nuxui/ui,
    // theme: nuxui.org/nuxui/// theme,
  },

  styles: {
	  btnDark:{
		  drawable: ui.StateDrawable,
		  pressed: {
			  drawable: ui.ShapeDrawable,
		  },
		  default:{
			drawable: ui.ShapeDrawable,
		  }
	  }
  }

  layout: {
	id: "root",
	widget: ui.Column,
	width: 100%,
	height: 100%,
	background: #303030,
	padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
	children:[
		{
			widget: main.Buttons,
			width: 100%,
			height: auto,
		},
	]
  }
}
  `
}

func (me *home) Mount(parent nux.Widget) {
	// nux.SetTheme()
	// ui.NewButton(nil, nux.ThemeStyle("button.large"))
	col := nux.Find(me, "root").(ui.Column)
	nux.OnTap(col, func(detail nux.GestureDetail) {
		log.V("nuxui", "root tap")
	})

	btn := nux.Find(me, "xxx").(ui.Button)
	nux.OnTap(btn, func(detail nux.GestureDetail) {
		log.V("nuxui", "xxx tap")
	})
}
