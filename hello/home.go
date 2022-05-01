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
  },

  layout: {
	id: "root",
	widget: ui.Column,
	width: 100%,
	height: 100%,
	background: #303030,
	padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
	children:[
		{
			widget: ui.Text,
			text: hello,
			textSize: 20,
		}
	]
  }
}
  `
}
