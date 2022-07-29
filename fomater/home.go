package main

import (
	"nuxui.org/nuxui/log"
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/ui"

	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	}

	style: {
		radio: {
			type: ui.Radio,
			textColor: #000000,
			font: {size: 14},
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
				width: 100%,
				height: 50%,
				children: [
					{
						type: ui.Column,
						width: 80%,
						height: 98%,
						margin: {left: 1wt, right: 1wt, top: 2px},
						align: {horizontal: center},
						background: {
							type: ui.ShapeDrawable,
							shape: {
								name: rect,
								stroke: #d3d3d3,
								strokeWidth: 2px,
								cornerRadius: 20px,
								dash: [5, 5],
							},
						},
						children: [
							{
								type: ui.Image,
								width: 50px,
								height: 1:1,
								src: "assets/ic_add_light.svg",
								margin: {top: 4wt, bottom: 3wt},
							},{
								type: ui.Text,
								text: "Choose an image to convert",
								font: {size: 20},
								textColor: #8b8b8b,
								margin: {bottom: 1wt},
							}
						],
					},{
						id: img_preview,
						type: ui.Image,
						width: 100%,
						height: 100%,
					}
				],
			},
			{
				type: ui.Column,
				width: 100%,
				height: 50%,
				children: [
					{
						type: ui.Row,
						width: 100%,
						align: {vertical: center},
						children: [
							{
								type: ui.Text,
								text: "Convert To:",
							},{
								id: options,
								type: ui.Options,
								single: true,
								content: {
									type: ui.Column,
									children: [
										{
											type: ui.Row,
											children: [
												{style: [radio], text: bmp},
												{style: [radio], text: jpg},
												{style: [radio], text: png},
												{style: [radio], text: tif},
												{style: [radio], text: gif},
											],
										},{
											type: ui.Row,
											children: [
												{style: [radio], text: webp},
												{style: [radio], text: pcx},
												{style: [radio], text: tga},
												{style: [radio], text: ico},
											],
										}
									],
								},
							}
						],
					},{
						type: ui.Row,
						width: 100%,
						margin: {top: 20px},
						align: {horizontal: center},
						children: [
							{
								type: ui.Text,
								text: "Save To:",

							},{
								id: txt_saveto,
								type: ui.Text,
								width: 57%,
								text: "/users/nux/Desktop"
								margin: {left: 20px},
							},{
								id: btn_change,
								type: ui.Button,
								text: Change,
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
						id: btn_convert,
						type: ui.Button,
						text: Convert,
						margin: {top: 20px, left: 1wt, right: 1wt},
						theme: [btn, btn_primary],
					},{
						id: txt_info,
						type: ui.Text,
						margin: {top: 20px},
					}
				],
			}
		],
	}
}
`
}

func (me *home) OnMount() {
	img_preview := nux.FindChild(me, "img_preview").(ui.Image)
	btn_convert := nux.FindChild(me, "btn_convert").(ui.Button)
	btn_change := nux.FindChild(me, "btn_change").(ui.Button)
	txt_saveto := nux.FindChild(me, "txt_saveto").(ui.Text)
	txt_info := nux.FindChild(me, "txt_info").(ui.Text)
	options := nux.FindChild(me, "options").(ui.Options)
	pick_image := nux.FindChild(me, "pick_image")

	dir, err := os.UserHomeDir()
	if err != nil {
		dir, _ = filepath.Abs(".")
	}
	txt_saveto.SetText(dir)

	nux.OnTap(pick_image, func(detail nux.GestureDetail) {
		nux.PickFileDialog().
			SetDirectory(dir).
			SetExtensionFilters(map[string][]string{
				"images": {"bmp", "jpg", "jpeg", "png", "tif", "tiff", "gif", "webp", "pcx", "tga", "ico"},
			}).
			AllowsChooseFiles().
			AllowsCreateFolders().
			ShowModal(func(ok bool, ret []string) {
				log.I("fomater", "%t, %s", ok, ret)
				if ok && len(ret) > 0 {
					me.pickedFile = ret[0]
					img_preview.SetSrc(me.pickedFile)
					img_preview.SetBackgroundColor(nux.White)

					if me.convertTo == "" {
						txt_saveto.SetText(filepath.Dir(me.pickedFile))
					} else {
						me.saveTo = filepath.Dir(me.pickedFile) + "/" +
							me.getFileBaseNameWidthoutExt(me.pickedFile) + "." + me.convertTo
						txt_saveto.SetText(me.saveTo)
					}
				}
			})
	})

	options.SetOnSelectionChanged(func(w ui.Options, fromUser bool) {
		if w.Selected() {
			me.convertTo = w.Values()[0]
		}

		if fromUser && me.pickedFile != "" {
			if w.Selected() {
				me.saveTo = filepath.Dir(me.pickedFile) + "/" +
					me.getFileBaseNameWidthoutExt(me.pickedFile) + "." + me.convertTo

				txt_saveto.SetText(me.saveTo)
			} else {
				me.saveTo = ""
				txt_saveto.SetText(filepath.Dir(me.pickedFile) + "/" +
					me.getFileBaseNameWidthoutExt(me.pickedFile))
			}
		}
	})

	nux.OnTap(btn_convert, func(detail nux.GestureDetail) {
		if me.pickedFile == "" {
			txt_info.SetTextColor(nux.Red)
			txt_info.SetText("Please choose an image file")
			return
		}

		if me.saveTo == "" || me.convertTo == "" {
			txt_info.SetTextColor(nux.Red)
			txt_info.SetText("Please choose convert to")
			return
		}

		cmd := exec.Command(ffmpeg_name(), "-i", me.pickedFile, me.saveTo)
		err := cmd.Run()
		if err != nil {
			txt_info.SetTextColor(nux.Red)
			txt_info.SetText(err.Error())
			return
		} else {
			txt_info.SetTextColor(nux.Green)
			txt_info.SetText("Convert Successed")
		}
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
			if me.pickedFile != "" {
				saveName = me.getFileBaseNameWidthoutExt(me.pickedFile) + "." + me.convertTo
			} else {
				saveName = saveName + "." + me.convertTo
			}
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
		nux.LoadNativeCursor(nux.CursorFinger).Set()
	})
	nux.OnHoverExit(txt_saveto, func(detail nux.GestureDetail) {
		nux.LoadNativeCursor(nux.CursorArrow).Set()
	})

	nux.OnTap(txt_saveto, func(detail nux.GestureDetail) {
		if me.saveTo != "" {
			nux.ViewFileDialog().
				SetDirectory(filepath.Dir(me.pickedFile)).
				SetActiveFileNames([]string{me.saveTo}).
				Show()
		}
	})
}

func (me *home) getFileBaseNameWidthoutExt(filename string) string {
	arr := strings.Split(filepath.Base(filename), filepath.Ext(filename))
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
