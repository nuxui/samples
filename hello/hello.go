// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"nuxui.org/nuxui/nux"
	_ "nuxui.org/nuxui/ui"
)

func main() {
	nux.Run(func(){
		nux.NewWindow(nux.Attr{
			"width":  "15%", // screen_width x 15%
			"height": "2:1", // width : height = 2 : 1
			"title":  "hello",
			"content": nux.Attr{
				"type": "nuxui.org/nuxui/ui.Text",
				"text": `Hello nuxui`,
			},
		})
	})

}
