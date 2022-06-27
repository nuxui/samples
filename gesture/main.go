package main

import (
	"nuxui.org/nuxui/log"
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/ui"
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
		id: "root",
		type: ui.Layer,
		width: 100%,
		height: 100%,
		background: #f7f7f7,
		padding: 10px,
		children:[
			{
				id: "tap", 
				type: ui.Text,
				text: tap gesture,
				background: #ff0000,
				padding: 20px,
			},
			{
				id: "langtap", 
				type: ui.Text,
				text: lang tap gesture,
				background: #ff0000,
				padding: 20px,
				margin: {left: 200px},
			},{
				id: "pan", 
				type: ui.Text,
				text: "pan gesture",
				background: #ff0000,
				padding: 20px,
				margin: {top: 200px},
			}
		],
	}
}
  `
}

func (me *home) Mount() {
	root := nux.FindChild(me, "root").(ui.Layer)
	txtTap := nux.FindChild(me, "tap").(ui.Text)
	txtPan := nux.FindChild(me, "pan").(ui.Text)
	langtap := nux.FindChild(me, "langtap").(ui.Text)

	// nux.OnHover(root, func(detail nux.GestureDetail) {
	// 	log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : root OnHover")
	// })

	nux.OnHover(langtap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : langtap OnHover")
	})

	nux.OnHoverEnter(langtap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : langtap OnHoverEnter")
	})

	nux.OnHoverExit(langtap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : langtap OnHoverExit")
	})

	nux.OnTapDown(root, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : tap down")
	})

	nux.OnTapUp(root, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : tap up")
	})

	nux.OnTapCancel(root, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : tap cancel")
	})

	nux.OnTapDown(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : tap down")
	})

	nux.OnTapUp(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : tap up")
	})

	nux.OnTapCancel(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : tap cancel")
	})

	nux.OnSecondaryTapDown(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Secondary tap down")
	})

	nux.OnSecondaryTapUp(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Secondary tap up")
	})

	nux.OnSecondaryTapCancel(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Secondary tap cancel")
	})

	nux.OnOtherTapDown(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Other tap down")
	})

	nux.OnOtherTapUp(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Other tap up")
	})

	nux.OnOtherTapCancel(txtTap, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Other tap cancel")
	})

	nux.OnOtherTapDown(root, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Other tap down")
	})

	nux.OnOtherTapUp(root, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Other tap up")
	})

	nux.OnOtherTapCancel(root, func(detail nux.GestureDetail) {
		log.I("nxuui", "id:'"+detail.Target().Info().ID+"' : Other tap cancel")
	})

	nux.OnPanDown(txtPan, func(detail nux.GestureDetail) {
		log.I("nxuui", "%T:%s pan down", detail.Target(), detail.Target().Info().ID)
	})

	nux.OnPanUpdate(txtPan, func(detail nux.GestureDetail) {
		log.I("nxuui", "%T:%s pan update", detail.Target(), detail.Target().Info().ID)
	})
}

func init() {
	nux.RegisterType((*Home)(nil), func(attr nux.Attr) any { return NewHome(attr) })
}

func main() {
	w := nux.NewWindow(nux.Attr{
		"title": "gesture",
		"content": nux.Attr{
			"type": "main.Home",
		},
	})
	nux.Run(w)
}
