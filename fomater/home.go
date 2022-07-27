package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"nuxui.org/nuxui/log"
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/ui"
)

type Home interface {
	nux.Component
}

type home struct {
	*nux.ComponentBase

	pickedFile string
	convertTo  string
	saveTo     string
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
		padding: 20px,
		children: [
			{
				id: pick_image,
				type: ui.Layer,
				height: 50%,
				width: 100%,
				children: [
					{
						type: ui.Column,
						width: 80%, height: 96%,
						align: {horizontal: center},
						margin: "1wt, 2px, 1wt, 0px",
						background: {
							type: ui.ShapeDrawable,
							shape: {
								name: rect,
								stroke: #D3D3D3,
								strokeWidth: 2px,
								cornerRadius: 20px,
								dash: [5, 5],
							},
						},
						children:[
							{
								id: image,
								type: ui.Image,
								width: 50px,
								height: 50px,
								src: "./assets/ic_add_light.svg",
								margin: {top: 4wt},
							},{
								type: ui.Text,
								text: "Choose an image to convert",
								textColor: #8B8B8B,
								font: {size: 20},
								margin: {top: 3wt, bottom: 1wt},
							}
						]
					},{
						id: preview,
						type: ui.Image,
						width: 81%,
						height: 100%,
						margin: {left: 1wt, right: 1wt},
					}
				],
			},{
				id: options,
				type: ui.Options,
				single: true,
				content: {
					type: ui.Row,
					align: {vertical: center},
					margin: {top: 20px},
					children:[
						{
							type: ui.Text,
							text: "Convert To:",
						},{
							type: ui.Column,
							margin: {left: 10px},
							children:[
								{
									type: ui.Row,
									width: 100%,
									align: {vertical: center},
									children:[
										{style: [radio], text: bmp, margin:{right:5px}},
										{style: [radio], text: jpg, margin:{right:5px}},
										{style: [radio], text: png, margin:{right:5px}},
										{style: [radio], text: tif, margin:{right:5px}},
										{style: [radio], text: gif, margin:{right:5px}},
									],
								},{
									type: ui.Row,
									width: 100%,
									align: {vertical: center},
									children:[
										{style: [radio], text: webp, margin:{right:5px}},
										{style: [radio], text: pcx, margin:{right:5px}},
										{style: [radio], text: tga, margin:{right:5px}},
										{style: [radio], text: ico, margin:{right:5px}},
									],
								}
							]
						}
					],
				},
			},{
				type: ui.Row,
				width: 100%,
				margin: {top: 20px},
				align: {horizontal: center},
				children: [
					{
						type: ui.Text,
						text: "Save to:",
					},
					{
						id: save_to,
						type: ui.Text,
						width: 57%,
						margin: {left: 20px},
					},{
						id: change,
						type: ui.Button,
						text: "Change",
						// theme: [btn, btn_primary],
						margin: {left: 20px},
						icon: {
							left: {
								width: 1em, height: 1em,
								type: ui.Image,
								src: "assets/folder-open-fill.svg",
							},
						},
					}
				],
			},{
				id: convert, type: ui.Button, text: Convert,
				margin: {top: 20px, left:1wt, right:1wt},
				theme: [btn, btn_primary],
			},{
				id: info, 
				type: ui.Text,
				margin: {top: 20px},
				textColor: #ff0000,
			}
		],
	}
}`
}

func (me home) OnMount() {
	btn_change := nux.FindChild(me, "change").(ui.Button)
	btn_convert := nux.FindChild(me, "convert").(ui.Button)
	txt_saveto := nux.FindChild(me, "save_to").(ui.Text)
	txt_info := nux.FindChild(me, "info").(ui.Text)
	pick_image := nux.FindChild(me, "pick_image")
	preview := nux.FindChild(me, "preview").(ui.Image)
	options := nux.FindChild(me, "options").(ui.Options)

	dir, err := os.UserHomeDir()
	if err != nil {
		dir, _ = filepath.Abs(".")
	}
	txt_saveto.SetText(dir)

	options.SetOnSelectionChanged(func(w ui.Options, fromUser bool) {
		log.I("fomater", "SetOnSelectionChanged %t", w.Selected())
		if w.Selected() {
			me.convertTo = w.Values()[0]
		}
		if fromUser && me.pickedFile != "" {
			if w.Selected() {
				me.saveTo = filepath.Dir(me.pickedFile) + "/" +
					me.getFileBaseNameWithoutExt(me.pickedFile) + "." + me.convertTo
				txt_saveto.SetText(me.saveTo)
			} else {
				me.saveTo = ""
				txt_saveto.SetText(filepath.Dir(me.pickedFile) + "/" +
					me.getFileBaseNameWithoutExt(me.pickedFile))
			}
		}
	})

	nux.OnTap(btn_convert, func(detail nux.GestureDetail) {
		if me.pickedFile == "" {
			txt_info.SetTextColor(nux.Red)
			txt_info.SetText("Please choose an image file")
			return
		}
		if me.saveTo == "" {
			txt_info.SetTextColor(nux.Red)
			txt_info.SetText("Please choose save to")
			return
		}

		cmd := exec.Command("./ffmpeg", "-i", me.pickedFile, me.saveTo)
		err := cmd.Run()
		if err != nil {
			txt_info.SetTextColor(nux.Red)
			txt_info.SetText(err.Error())
		} else {
			txt_info.SetTextColor(nux.Black)
			txt_info.SetText("Convert successed")
		}
	})

	nux.OnTap(pick_image, func(detail nux.GestureDetail) {
		log.I("nuxui", "pick_image ontap")

		dir, err := os.UserHomeDir()
		if err != nil {
			dir = "."
		}

		nux.PickFileDialog().
			SetDirectory(dir).
			SetExtensionFilters(map[string][]string{
				"images": {"bmp", "jpg", "jpeg", "png", "tif", "tiff", "gif", "pcx", "tga", "webp", "ico"}}).
			AllowsChooseFiles().
			AllowsCreateFolders().
			ShowModal(func(ok bool, ret []string) {
				log.I("fomater", "%t, %s", ok, ret)
				if ok && len(ret) > 0 {
					me.pickedFile = ret[0]
					preview.SetSrc(me.pickedFile)
					preview.SetBackgroundColor(0xffffffff)
					if me.convertTo == "" {
						txt_saveto.SetText(filepath.Dir(me.pickedFile))
					} else {
						me.saveTo = filepath.Dir(me.pickedFile) + "/" + me.getFileBaseNameWithoutExt(me.pickedFile) + "." + me.convertTo
						txt_saveto.SetText(me.saveTo)
					}
				}
			})
	})

	nux.OnTap(btn_change, func(detail nux.GestureDetail) {
		dir := txt_saveto.Text()
		if me.pickedFile != "" {
			dir = filepath.Dir(me.pickedFile)
		}

		if !isDir(dir) {
			dir = filepath.Dir(dir)
		}

		saveName := "untitled"
		if me.convertTo != "" {
			saveName = saveName + "." + me.convertTo
		}
		if me.convertTo != "" && me.pickedFile != "" {
			saveName = me.getFileBaseNameWithoutExt(me.pickedFile) + "." + me.convertTo
		}

		nux.SaveFileDialog().
			SetDirectory(dir).
			SetSaveName(saveName).
			ShowModal(func(ok bool, ret string) {
				if ok {
					me.saveTo = ret
					txt_saveto.SetText(ret)
				}
			})
	})

	nux.OnHoverEnter(txt_saveto, func(detail nux.GestureDetail) {
		log.I("nuxui", "OnHoverEnter")
		nux.LoadNativeCursor(nux.CursorFinger).Set()
	})
	nux.OnHoverExit(txt_saveto, func(detail nux.GestureDetail) {
		nux.LoadNativeCursor(nux.CursorArrow).Set()
	})
	nux.OnTap(txt_saveto, func(detail nux.GestureDetail) {
		if me.saveTo != "" {
			nux.ViewFileDialog().
				SetDirectory(filepath.Dir(me.saveTo)).
				SetActiveFileNames([]string{me.saveTo}).
				Show()
		}
	})
}

func (me *home) getFileBaseNameWithoutExt(fileName string) string {
	arr := strings.Split(filepath.Base(fileName), filepath.Ext(fileName))
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
