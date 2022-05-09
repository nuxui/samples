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
	width: 100%,
	height: 100%,
	background: #303030,
	children:[
		{
			widget: ui.Text,
			text: Hello NuxUI,
			textSize: 30,
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