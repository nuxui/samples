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
        widget: ui.Column,
        width: auto,
        height: 100px,
        background: #eae5ec,
        padding: 10px,
        children:[
            {
                widget: ui.Button,
                text: normal button,
            },{
                widget: ui.Button,
                text: icon button,
                margin:{top: 10px},
                icon:{
                    left: {
                        widget: ui.Image,
                        width: 1em,
                        height: 1em,
                        margin:{right: 3px },
                        src: "delete.png",
                    },
                },
            }
        ]
    }
}
  `
}

func init() {
	nux.RegisterWidget((*Home)(nil), func(attr nux.Attr) nux.Widget { return NewHome(attr) })
}

const manifest = `
{
  import: {
    main: main,
  }, 

  manifest: {
      main: main.Home,
  },
}
`

func main() {
	nux.Init(manifest)
	nux.Run()
}
