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
		widget: ui.Row,
		width: 200px,
		height: 200px,
		background: #303030,
		children:[
			{
				widget: ui.Text,
				width: 1wt,
				text: left text,
				background: #ff0000,
			},{
				widget: ui.Text,
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
