// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/nuxui/nuxui/log"
	"github.com/nuxui/nuxui/nux"
	"github.com/nuxui/nuxui/ui"
)

const manifest = `
{
  import: {
	main: main,
  }, 

  manifest: {
	  name: "Whizz",
	  package: "github.com/nuxui/samples/hello_world"	//android will be 'github.com.nuxui.samples.hello_world'
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

/*
window creating(){
	Owner().(Window).SetWidth(50%)
	Owner().(Window).SetHeight(50%)
	nux.GetScreen().Width()
	nux.GetScreen().Height()
}

*/

// type windowDelegate struct {
// }

// func (me *windowDelegate) Creating(attr nux.Attr) {

// }

// func (me *windowDelegate) Created() {

// }

// func (me *windowDelegate) Measured() {

// }

// func (me *windowDelegate) Layout() {

// }

// func (me *windowDelegate) Draw(canvas nux.Canvas) {

// }

func init() {
	nux.Use(ui.NewRow, ui.NewColumn)
	nux.RegisterWidget((*Home)(nil), func(ctx nux.Context, attr ...nux.Attr) nux.Widget { return NewHome(ctx, attr...) })

}

func main() {
	defer log.Close()
	log.V("main", "hello")
	nux.Init(manifest)
	// window is not created before call Run, how did set it?
	// nux.App().Window().Decor().SetBackgroundColor(nux.Transparent)
	// nux.AddMainWindowMixins(&windowDelegate{})
	// registerActivityLifecycleCallbacks(Application.ActivityLifecycleCallbacks
	// app.Mixins()
	// nux.MainWindow()
	// nux.AddAppMixins()
	nux.Run() // application.run()
}
