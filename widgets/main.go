// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"nuxui.org/nuxui/log"
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/theme"
)

const manifest = `
{
  import: {
	main: main,
  }, 

  manifest: {
	  name: "Whizz",
	  package: "nuxui.org/whizz"	//android will be 'github.com.nuxui.whizz'
	  version: "1.0.1",
	  icon: "app-icon",
	  appId: "ac29bdde-def5-5b06-a0c4-19ea611a865f",
	  versionId: "47b97dd6-d5e7-560f-bb73-1b4ffdabae86",
	  appMixins:[main.Application],
	  mainWndowMixins:[main.Window],
	  main: main.Home,
	  // https://developer.android.com/reference/android/Manifest.permission
	  permissions: [
		  "access_internet",
		  "read_sms",
		  "send_sms",
		  "camera",
		  "read_external_storage",
		  "write_external_storage",
		  "nfc",
	  ],
	//   window: {	//
	// 	width: 50%,		// screen.width * 50%
	// 	height: 50%,	// screen.height * 50%
	// 	margin: {left:1wt, top: 1wt, right: 1wt, bottom: 1wt}, // default align is center
	// 	mixins:[main.Window],
	// 	useLastStatus: false,	// default is true
	//   }, //

	  router: {
		path: "/",
		window: {
			width: 61.8%,		// screen.width * 61.8%
			height: 61.8%,	// screen.height * 61.8%
			margin: {left:1wt, top: 1wt, right: 1wt, bottom: 1wt}, // default align is center
			mixins:[main.Window],
		},
		widget: main.Root,
		children: [
			{
				path: "/home",
				widget: main.Rect,
			},
			{
				path: "/login",
				widget: main.Login,
			},
		],
	  },

  },
}
`

func init() {
	nux.UseTheme(&theme.Material{})
	nux.RegisterWidget((*Home)(nil), func(attr nux.Attr) nux.Widget { return NewHome(attr) })
	nux.RegisterWidget((*Buttons)(nil), func(attr nux.Attr) nux.Widget { return NewButtons(attr) })
}

func main() {
	// nux.TestDraw = TestDraw
	log.SetLevel(log.VERBOSE)
	nux.Init(manifest)
	nux.Run()
}

var count int

func TestDraw(canvas nux.Canvas) {
	if count > 0 {
		return
	}

	canvas.Save()
	canvas.Translate(0, float32(600))
	canvas.Scale(1, -1)
	paint := nux.NewPaint()
	paint.SetColor(0xffffff00)
	canvas.ClipRect(0, 0, 8, 8)
	canvas.DrawRect(0, 0, 100000, 100000, paint)
	paint.SetColor(0xffff0000)
	canvas.DrawRect(0, 0, 4, 4, paint)
	canvas.DrawRect(8, 0, 12, 4, paint)

	// canvas.Translate(0, 0)
	// canvas.Translate(-0.5, -0.5)
	paint.SetStyle(nux.PaintStyle_Stroke)
	paint.SetWidth(1)
	paint.SetColor(0x88ff0000)
	canvas.DrawRect(4, 4, 6, 6, paint)

	// canvas.DrawRect(0, 0, 800, 600, paint)
	// canvas.DrawRoundRect(0, 0, 800, 600, 10, paint)
	canvas.Restore()
	// count++
}
