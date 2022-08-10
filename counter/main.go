// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"nuxui.org/nuxui/log"
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/ui"
	"nuxui.org/nuxui/theme"
)

type Home interface {
	nux.Component
}

type home struct{
	*nux.ComponentBase
}

func NewHome(attr nux.Attr) Home {
	me := &home{}
	me.ComponentBase = nux.NewComponentBase(me, attr)
	nux.InflateLayout(me, me.layout(), nil)
	return me
}

func (me *home) layout() string {
	return `
{
	import: {
		ui: nuxui.org/nuxui/ui,
	}

	layout: {
		type: ui.Column,
		width: 100%,
		height: 100%,
		align: {horizontal:center, vertical: center},
		children: [
			{
				id: txt,
				type: ui.Text,
				text: "click me 0 times",
				font: {size: 16},
			},{
				id: btn,
				theme: [btn, btn_primary],
				text: "Add Count",
				margin: {top: 20px},
				font: {size: 16},
			}
		],
	}
}
`
}

func (me *home) OnMount() {
	log.I("nuxui", "home OnMount")
	txt := nux.FindChild(me, "txt").(ui.Text)
	btn := nux.FindChild(me, "btn").(ui.Button)
	count := 0
	nux.OnTap(btn, func(detail nux.GestureDetail){
		count ++
		s := fmt.Sprintf("click me %d times", count)
		txt.SetText(s)
		log.I("nuxui", "text %s", txt.Text())
	})
}

func init(){
	nux.RegisterType( (*Home)(nil), func(attr nux.Attr) any { return NewHome(attr) } )
}

func main() {
	nux.ApplyTheme(nux.ThemeLight, theme.BootstrapLight)
	nux.App().Init(manifest)
	nux.App().Run()
}