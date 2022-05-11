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

type Buttons interface {
	nux.Component
}

type buttons struct {
	*nux.ComponentBase

	content nux.Widget
}

func NewButtons(attr nux.Attr) Buttons {
	me := &buttons{}
	me.ComponentBase = nux.NewComponentBase(me, attr)
	me.content = nux.InflateLayout(me, me.template())
	return me
}

func (me *buttons) template() string {
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
	id: "root123",
	widget: ui.Column,
	width: 100%,
	height: auto,
	background: #303030,
	// padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
	children:[
		{
			widget: ui.Text,
			width: auto,
			height: auto,
			text: Buttons,
			textSize: 30,
			textColor: #ffffff,
			styles: "ui.h1",
		},{
			widget: ui.Text,
			width: auto,
			height: auto,
			text: "Contained Buttons",
			textSize: 20,
			textColor: #ffffff,
			styles: "ui.h1",
			margin:{left: 50px, top: 20px},
		},{
			widget: ui.Row,
			width: 1wt,
			height: auto,
			margin:{left: 50px, top: 10px, },
			children: [
				{
					widget: ui.Button,
					theme: {name: "material", kind: "light", style: "button"},
					style: @style.button.small,
					width: auto,
					height: auto,
					text: "DEFAULT",
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "PRIMARY",
					textSize: 14,
					textColor: #ffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								solid: #3f51b5,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #2b397e,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
						]
					}
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "SECONDARY",
					textSize: 14,
					textColor: #ffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								solid: #f50057,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #ab003c,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
						],
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "DISABLE",
					disable: true,
					textSize: 14,
					textColor: #4cffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					background: {
						drawable: ui.ShapeDrawable,
						shape:{
							shape: rect,
							solid: #1effffff,
							cornerRadius: 4px,
							shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
						},
					}
				},{
					id: "xxx",
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "LINKs",
					textSize: 14,
					textColor: #ffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					// background: #3f51b5,
					background: {
						drawable: ui.ColorDrawable,
						// color: #3f51b5,
						states:[
							{state:"default", color: #3f51b5},
							{state:"pressed", color: #ff0000},
							{state:"hovered", color: #ffff00},
						],
					}
				}
			],
		},{
			widget: ui.Text,
			width: auto,
			height: auto,
			text: "Text Buttons",
			textSize: 20,
			textColor: #ffffff,
			styles: "ui.h1",
			margin:{left: 50px, top: 20px},
		},{
			widget: ui.Row,
			width: 1wt,
			height: auto,
			margin:{left: 50px, top: 10px, },
			children: [
				{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "DEFAULT",
					textSize: 14,
					textColor: #ffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "PRIMARY",
					textSize: 14,
					textColor: #3f51b5,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "SECONDARY",
					textSize: 14,
					textColor: #e30044,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "DISABLE",
					textSize: 14,
					textColor: #575757,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "LINK",
					textSize: 14,
					textColor: #3f51b5,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
				}
			],
		},{
			widget: ui.Text,
			width: auto,
			height: auto,
			text: "Outlined Buttons",
			textSize: 20,
			textColor: #ffffff,
			styles: "ui.h1",
			margin:{left: 50px, top: 20px},
		},{
			widget: ui.Row,
			width: 1wt,
			height: auto,
			margin:{left: 50px, top: 10px },
			children: [
				{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "DEFAULT",
					textSize: 14,
					textColor: #ffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								stroke: #3affffff,
								strokeWidth: 1px,
								cornerRadius: 4px,
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #797979,,
								stroke: #3affffff,
								strokeWidth: 1px,
								cornerRadius: 4px,
							}},
						],
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "PRIMARY",
					textSize: 14,
					textColor: #3f51b5,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					background: {
						drawable: ui.ShapeDrawable,
						states: [
							{state:"default", shape: {
								shape: rect,
								stroke: #3f51b5,
								cornerRadius: 4px,
							}},
							{state:"pressed", shape: {
								shape: rect,
								solid: #353c60,
								stroke: #3f51b5,
								cornerRadius: 4px,
							}},
						],
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "SECONDARY",
					textSize: 14,
					textColor: #f50357,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								stroke: #952744,
								cornerRadius: 4px,
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #782b3f,
								stroke: #952744,
								cornerRadius: 4px,
							}},
						],
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					disable: true,
					text: "DISABLE",
					textSize: 14,
					textColor: #575757,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					background: {
						drawable: ui.ShapeDrawable,
						shape:{
							shape: rect,
							stroke: #494949,
							cornerRadius: 4px,
						},
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "LINK",
					textSize: 14,
					textColor: #3f51b5,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					// background: @styles.btnDark,
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								stroke: #3f51b5,
								cornerRadius: 4px,
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #353c60,
								stroke: #3f51b5,
								cornerRadius: 4px,
							}},
						],
					},
				}
			],
		},{
			widget: ui.Text,
			width: auto,
			height: auto,
			text: "Buttons with icons and label",
			textSize: 20,
			textColor: #ffffff,
			styles: "ui.h1",
			margin:{left: 50px, top: 20px},
		},{
			widget: ui.Row,
			width: 1wt,
			height: auto,
			margin:{left: 50px, top: 10px },
			children: [
				{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "DELETE",
					textSize: 14,
					textColor: #ffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					icon:{
						left: {
							widget: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_delete.png",
						},
					},
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								solid: #f50357,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #ab003c,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
						],
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "UPLOAD",
					textSize: 14,
					textColor: #1d1d1d,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					icon:{
						left: {
							widget: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_upload.png",
						},
					},
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								solid: #e0e0e0,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #d5d5d5,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
						],
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "TALK",
					textSize: 14,
					textColor: #ababab,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					icon:{
						left: {
							widget: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_mic.png",
						},
					},
					background: {
						drawable: ui.ShapeDrawable,
						shape:{
							shape: rect,
							solid: #dcdcdc,
							cornerRadius: 4px,
							shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
						},
					},
				},{
					widget: ui.Button,
					width: auto,
					height: auto,
					text: "SAVE",
					textSize: 14,
					textColor: #ffffff,
					padding: {left: 16px, top: 8px, right: 16px, bottom: 8px},
					margin:{left: 15px },
					icon:{
						left: {
							widget: ui.Image,
							width: 1em,
							height: 1em,
							margin:{right: 6px },
							src: "assets/ic_save.png",
						},
					},
					background: {
						drawable: ui.ShapeDrawable,
						states:[
							{state:"default", shape:{
								shape: rect,
								solid: #3f51b5,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
							{state:"pressed", shape:{
								shape: rect,
								solid: #2b397e,
								cornerRadius: 4px,
								shadow:{color: #88000000, x: 0, y: 1px, blur: 3px},
							}},
						],
					}
				}
			],
		},{
			widget: ui.Text,
			width: auto,
			height: auto,
			text: "Options",
			textSize: 20,
			textColor: #ffffff,
			styles: "ui.h1",
			margin:{left: 50px, top: 20px},
		},{
			widget: ui.Row,
			id: "options",
			width: 1wt,
			height: auto,
			margin:{left: 50px, top: 10px },
			children: [
				{
					widget: ui.Options,
					content: {
						id: "col4CheckBox"
						widget: ui.Column,
						width: 1wt,
						height: auto,
						children:[
							{
								widget: ui.Check,
								width: auto,
								height: auto,
								text: "Gilad Gray",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										widget: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											drawable: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_opt_default.png"},
												{state:"checked", src: "assets/ic_opt_checked.png"},
											],
										},
									},
								},
							},{
								widget: ui.Check,
								width: auto,
								height: auto,
								text: "Jason Killian",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										widget: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											drawable: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_opt_default.png"},
												{state:"checked", src: "assets/ic_opt_checked.png"},
											],
										},
									},
								},
							},{
								widget: ui.Check,
								width: auto,
								height: auto,
								text: "Antoine Llorca",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										widget: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											drawable: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_opt_default.png"},
												{state:"checked", src: "assets/ic_opt_checked.png"},
											],
										},
									},
								},
							}
						],
					}
				},{
					widget: ui.Options,
					radio: true,
					content: {
						widget: ui.Column,
						width: 1wt,
						height: auto,
						children:[
							{
								widget: ui.Radio,
								width: auto,
								height: auto,
								text: "Gilad Gray",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										widget: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											drawable: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_radio_default.png"},
												{state:"checked", src: "assets/ic_radio_checked.png"},
											],
										},
									},
								},
							},{
								widget: ui.Radio,
								width: auto,
								height: auto,
								text: "Jason Killian",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										widget: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											drawable: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_radio_default.png"},
												{state:"checked", src: "assets/ic_radio_checked.png"},
											],
										},
									},
								},
							},{
								widget: ui.Radio,
								width: auto,
								height: auto,
								text: "Antoine Llorca",
								textSize: 14,
								textColor: #ffffff,
								margin: {left: 16px},
								padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
								icon:{
									left: {
										widget: ui.Image,
										width: 1.2em,
										height: 1.2em,
										margin:{right: 6px },
										src: {
											drawable: ui.ImageDrawable,
											states:[
												{state:"default", src: "assets/ic_radio_default.png"},
												{state:"checked", src: "assets/ic_radio_checked.png"},
											],
										},
									},
								},
							}
						],
					}
				},{
					widget: ui.Column,
					width: 1wt,
					height: auto,
					bacground: #ff0000,
					clipChildren: true,
					margin: {left: 16px},
					children: [
						{
							widget: ui.Switch,
							id: "switch",
							width: auto,
							height: auto,
							text: "Left Switch",
							textSize: 14,
							textColor: #ffffff,
							icon:{
								left: {
									widget: ui.Image,
									width: 1.5em,
									height: 1.5em,
									margin:{right: 6px },
									src: {
										drawable: ui.ImageDrawable,
										states:[
											{state:"default", src: "assets/ic_switch_on.png"},
											{state:"checked", src: "assets/ic_switch_off.png"},
										],
									},
								}
							},
						},						{
							widget: ui.Switch,
							id: "switch",
							width: auto,
							height: auto,
							text: "Right Switch",
							textSize: 14,
							textColor: #ffffff,
							icon:{
								right: {
									widget: ui.Image,
									width: 1.5em,
									height: 1.5em,
									margin:{left: 6px },
									src: {
										drawable: ui.ImageDrawable,
										states:[
											{state:"default", src: "assets/ic_switch_on.png"},
											{state:"checked", src: "assets/ic_switch_off.png"},
										],
									},
								}
							},
						}
					],
				}
			],
		}
	]
  }
}
  `
}

func (me *buttons) Mount(parent nux.Widget) {
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
