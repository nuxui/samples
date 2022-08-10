// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"nuxui.org/nuxui/nux"
	_ "nuxui.org/nuxui/ui"
)

type Home interface {
	nux.Component
}

type home struct {
	*nux.ComponentBase
}

func NewHome(attr nux.Attr) Home {
	me := &home{}
	me.ComponentBase = nux.NewComponentBase(me, attr)
	nux.InflateLayout(me, me.template(), nil)
	return me
}

func (me *home) template() string {
	return `
{
	import: {
		ui: nuxui.org/nuxui/ui,
	},

	layout: {
		type: ui.Row,
		width: 100%,
		height: 100%,
		background: #303030,
		children:[
			{
				type: ui.Text,
				width: 1wt,
				text: left text,
				background: #ff0000,
			},{
				type: ui.Text,
				height: 100%,
				text: right text,
				background: #ff00ff,
			},
		]
	}
}
  `
}

func init() {
	nux.RegisterType((*Home)(nil), func(attr nux.Attr) any { return NewHome(attr) })
}

func main() {
	nux.App().Init(manifest)
	nux.App().Run()
}
