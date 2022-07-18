package main

import (
	"os/exec"

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
	nux.InflateLayout(me, me.layout(), nux.InflateStyle(me.style()))
	return me
}

func (me *home) style() string {
	return `
{
	import: {
		ui: nuxui.org/nuxui/ui,
	},

	style: {
		radio: {
			type: ui.Radio,
			textSize: 14,
			textColor: #000000,
			padding: {left: 0px, top: 8px, right: 0px, bottom: 8px},
			icon:{
				left: {
					type: ui.Image,
					width: 1.2em,
					height: 1.2em,
					margin:{right: 3px },
					src: {
						type: ui.ImageDrawable,
						states:[
							{state:"default", src: "assets/ic_radio_default.png"},
							{state:"checked", src: "assets/ic_radio_checked.png"},
						],
					},
				},
			},
		}
	}
}	
`
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
		children: [
			{
				type: ui.Layer,
				width: 50%,
				height: 50%,
				children:[
					{
						type: ui.Image, id: image,
						width: 100%, height: 100%,
						src: "./Snap.png",
						background: #00000033,
					},{
						type: ui.Column,
						width: 100%, height: 100%,
						children:[
							{
								type: ui.Text,
								text: "Supported bmp/jpg/jpeg/png/tif/gif/pcx/tga/web",
							},
							{
								type: ui.Button,
								text: "Choose a image to convert",
							}
						]
					}
				]
			},{
				id: options,
				type: ui.Options,
				single: true,
				content: {
					type: ui.Row,
					width: 100%,
					align: {vertical: center},
					children:[
						{
							type: ui.Text,
							text: "Convert To:",
						},
						{style: [radio], text: bmp, margin:{right:5px}},
						{style: [radio], text: jpg, margin:{right:5px}},
						{style: [radio], text: png, margin:{right:5px}},
						{style: [radio], text: tif, margin:{right:5px}},
						{style: [radio], text: gif, margin:{right:5px}},
						{style: [radio], text: pcx, margin:{right:5px}},
						{style: [radio], text: tga, margin:{right:5px}},
						{style: [radio], text: webp, margin:{right:5px}},
					],
				},
			},{
				type: ui.Row,
				width: 100%,
				children: [
					{
						type: ui.Text,
						text: "Save to:",
					},
					{
						type: ui.Editor,
						width: 50%,
					},{
						type: ui.Button,
						text: "Change",
					}
				],
			},{
				type: ui.Row,
				width: 100%,
				children: [
					{
						id: convert,
						type: ui.Button,
						text: Convert,
					},{
						id: show,
						type: ui.Button,
						text: Show,
					},{
						id: save,
						type: ui.Button,
						text: Save,
					}
				]
			}
		],
	}
}`
}

func (me home) OnMount() {
	img := nux.FindChild(me, "image").(ui.Image)
	btn_convert := nux.FindChild(me, "convert").(ui.Button)
	btn_show := nux.FindChild(me, "show").(ui.Button)
	btn_save := nux.FindChild(me, "save").(ui.Button)
	options := nux.FindChild(me, "options").(ui.Options)
	nux.OnTap(btn_convert, func(detail nux.GestureDetail) {
		if options.Selected() {
			log.I("formater", "%s", options.Values())
			cmd := exec.Command("./ffmpeg", "-i", "./Snap.png", "./Snap."+options.Values()[0])
			err := cmd.Run()
			if err != nil {
				log.E("fomater", "%s", err.Error())
			}
		}

	})

	nux.OnTap(btn_show, func(detail nux.GestureDetail) {
		nux.ViewFileDialog().
			// SetDirectory("/Users/mustodo/Documents/go/src/github.com/nuxui/examples/testcode/test/coa").
			// SetActiveFileNames([]string{
			// 	"/Users/mustodo/Documents/go/src/github.com/nuxui/examples/testcode/test/coa/out.out",
			// 	"/Users/mustodo/Documents/go/src/github.com/nuxui/examples/testcode/test/coa/nux.m"}).
			SetDirectory("D:\\Downloads").
			// SetActiveFileNames([]string{"D:\\Downloads\\dxwebsetup.exe",
			// 	"D:\\Downloads\\depends22_x86.zip"}).
			Show()
	})

	nux.OnTap(img, func(detail nux.GestureDetail) {
		log.I("nuxui", "img ontap")
		nux.PickFileDialog().
			// SetDirectory(".").
			SetDirectory("D:\\Workspace\\golang\\src\\github.com").
			SetExtensionFilters(map[string][]string{
				"images": {"bmp","jpg", "jpeg", "png", "tif", "gif", "pcx", "tga", "web"},
				"text": {"txt"}}).
			AllowsChooseFiles().
			AllowsChooseFolders().
			AllowsMultipleSelection().
			AllowsCreateFolders().
			ShowModal(func(ok bool, ret []string){
				log.I("fomater", "%t, %s",ok, ret)
			})
	})

	nux.OnTap(btn_save, func(detail nux.GestureDetail){
		nux.SaveFileDialog().
			SetDirectory("D:\\Downloads").
			SetSaveName("kkk.png").
			// SetDirectory("/Users/mustodo/Downloads").
			ShowModal(func(ok bool, ret string){
				log.I("fomater", "%t, %s",ok, ret)
			})
	})
}
