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

const btns_style = `
{
	import: {
		ui: nuxui.org/nuxui/ui,
	},

	style: {}
}
`

type Buttons interface {
	nux.Component
}

type buttons struct {
	*nux.ComponentBase
}

func NewButtons(attr nux.Attr) Buttons {
	me := &buttons{}
	me.ComponentBase = nux.NewComponentBase(me, attr)
	nux.InflateLayout(me, me.layout(), nux.InflateStyle(btns_style))
	return me
}

func (me *buttons) layout() string {
	return `
{
  import: {
    ui: nuxui.org/nuxui/ui,
  },

  layout: {
	id: "root",
	type: ui.Column,
	width: 100%,
	background: #303030,
	// padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
	children:[
		{
			type: ui.Text,
			text: Buttons,
			textSize: 30,
			textColor: #ffffff,
		},{
			type: ui.Text,
			text: "Contained Buttons",
			textSize: 20,
			textColor: #ffffff,
			margin:{left: 50px, top: 20px},
		},{
			type: ui.Row,
			width: 1wt,
			margin:{left: 50px, top: 10px, },
			children: [
				{theme: [btn, btn_default],   text: "DEFAULT", },
				{theme: [btn, btn_primary],   text: "PRIMARY",   margin:{left: 15px} },
				{theme: [btn, btn_secondary], text: "SECONDARY", margin:{left: 15px} },
				{theme: [btn, btn_disable],   text: "DISABLE",   margin:{left: 15px} },
			],
		},{
			type: ui.Text,
			text: "Text Buttons",
			textSize: 20,
			textColor: #ffffff,
			margin:{left: 50px, top: 20px},
		},{
			type: ui.Row,
			width: 1wt,
			margin:{left: 50px, top: 10px, },
			children: [
				{theme: [btn, btn_default_text],   text: "DEFAULT", },
				{theme: [btn, btn_primary_text],   text: "PRIMARY",   margin:{left: 15px} },
				{theme: [btn, btn_secondary_text], text: "SECONDARY", margin:{left: 15px} },
				{theme: [btn, btn_disable_text],   text: "DISABLE",   margin:{left: 15px} },
			],
		},{
			type: ui.Text,
			text: "Outlined Buttons",
			textSize: 20,
			textColor: #ffffff,
			margin:{left: 50px, top: 20px},
		},{
			type: ui.Row,
			width: 1wt,
			margin:{left: 50px, top: 10px },
			children: [
				{theme: [btn, btn_default_outline],   text: "DEFAULT" },
				{theme: [btn, btn_primary_outline],   text: "PRIMARY",   margin:{left: 15px} },
				{theme: [btn, btn_secondary_outline], text: "SECONDARY", margin:{left: 15px} },
				{theme: [btn, btn_disable_outline],   text: "DISABLE",   margin:{left: 15px} },
			],
		},{
			type: ui.Text,
			text: "Buttons with icons and label",
			textSize: 20,
			textColor: #ffffff,
			margin:{left: 50px, top: 20px},
		},{
			type: ui.Row,
			width: 1wt,
			margin:{left: 50px, top: 10px },
			children: [
				{
					theme: [btn, btn_secondary],
					text: "DELETE",
					icon: {
						left: {
							type: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_delete.png",
						},
					},
				},{
					theme: [btn, btn_default],
					text: "UPLOAD",
					margin:{left: 15px },
					icon:{
						left: {
							type: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_upload.png",
						},
					},
				},{
					theme: [btn, btn_disable],
					text: "TALK",
					margin:{left: 15px },
					icon:{
						left: {
							type: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_mic.png",
						},
					},
				},{
					theme: [btn, btn_primary],
					text: "SAVE",
					margin:{left: 15px },
					icon:{
						left: {
							type: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_save.png",
						},
					},
				}
			],
		},{
			type: ui.Text,
			text: "Options",
			textSize: 20,
			textColor: #ffffff,
			margin:{left: 50px, top: 20px},
		},{
			type: ui.Row,
			id: "options",
			width: 1wt,
			margin:{left: 50px, top: 10px },
			children: [
				{
					type: ui.Options,
					content: {
						id: "col4CheckBox"
						type: ui.Column,
						width: 1wt,
						children:[
							{
								type: ui.Check,
								text: "Gilad Gray",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										type: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											type: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_opt_default.svg"},
												{state:"checked", src: "assets/ic_opt_checked.svg"},
											],
										},
									},
								},
							},{
								type: ui.Check,
								text: "Jason Killian",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										type: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											type: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_opt_default.svg"},
												{state:"checked", src: "assets/ic_opt_checked.svg"},
											],
										},
									},
								},
							},{
								type: ui.Check,
								text: "Antoine Llorca",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										type: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											type: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_opt_default.svg"},
												{state:"checked", src: "assets/ic_opt_checked.svg"},
											],
										},
									},
								},
							}
						],
					}
				},{
					type: ui.Options,
					single: true,
					content: {
						type: ui.Column,
						width: 1wt,
						children:[
							{
								type: ui.Radio,
								text: "Gilad Gray",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										type: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											type: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_radio_default.svg"},
												{state:"checked", src: "assets/ic_radio_checked.svg"},
											],
										},
									},
								},
							},{
								type: ui.Radio,
								text: "Jason Killian",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										type: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											type: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_radio_default.svg"},
												{state:"checked", src: "assets/ic_radio_checked.svg"},
											],
										},
									},
								},
							},{
								type: ui.Radio,
								text: "Antoine Llorca",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										type: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											type: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_radio_default.svg"},
												{state:"checked", src: "assets/ic_radio_checked.svg"},
											],
										},
									},
								},
							}
						],
					}
				},{
					type: ui.Column,
					width: 1wt,
					bacground: #ff0000,
					clipChildren: true,
					margin: {left: 16px},
					children: [
						{
							type: ui.Switch,
							id: "switch",
							text: "Left Switch",
							textSize: 14,
							textColor: #ffffff,
							icon:{
								left: {
									type: ui.Image,
									width: 1.5em,
									height: 1.5em,
									margin:{right: 6px },
									src: {
										type: ui.ImageDrawable,
										states:[
											{state:"default", src: "assets/ic_switch_off.svg"},
											{state:"checked", src: "assets/ic_switch_on.svg"},
										],
									},
								}
							},
						},{
							type: ui.Switch,
							id: "switch",
							text: "Right Switch",
							textSize: 14,
							textColor: #ffffff,
							checked: true,
							icon:{
								right: {
									type: ui.Image,
									width: 1.5em,
									height: 1.5em,
									margin:{left: 6px },
									src: {
										type: ui.ImageDrawable,
										states:[
											{state:"default", src: "assets/ic_switch_off.svg"},
											{state:"checked", src: "assets/ic_switch_on.svg"},
										],
									},
								}
							},
						}
					],
				}
			],
		},{
			type: ui.Editor,
			text: "Input",
			textSize: 20,
			textColor: #ffffff,
			margin:{left: 50px, top: 20px},
		},
	]
  }
}
  `
}

func (me *buttons) OnMount(parent nux.Widget) {
	// nux.SetTheme()
	// ui.NewButton(nil, nux.ThemeStyle("button.large"))
	col := nux.FindChild(me, "root").(ui.Column)
	nux.OnTap(col, func(detail nux.GestureDetail) {
		log.V("nuxui", "root tap")
	})

	btn := nux.FindChild(me, "xxx").(ui.Button)
	nux.OnTap(btn, func(detail nux.GestureDetail) {
		log.V("nuxui", "xxx tap")
	})
}
