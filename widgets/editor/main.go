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
	nux.InflateLayout(me, me.layout(), nil)
	return me
}

func (me *home) layout() string {
	return `
{
    import: {
        ui: nuxui.org/nuxui/ui,
    },

    layout: {
        type: ui.Column,
        width: 100%,
        height: 100%,
        background: #f7f7f7,
        padding: 10px,
        children:[
            {
                type: ui.Editor,
                text: input words,
            }
        ],
    }
}
  `
}

func (me *home) OnMount() {
}

func init() {
	nux.RegisterType((*Home)(nil), func(attr nux.Attr) any { return NewHome(attr) })
}

func main() {
	nux.App().Init(manifest)
	nux.App().Run()
}
