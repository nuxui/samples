// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"nuxui.org/nuxui/nux"
)

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

func init() {
	nux.RegisterWidget((*Home)(nil), func(attr nux.Attr) nux.Widget { return NewHome(attr) })
}

func main() {
	nux.Init(manifest)
	nux.Run()
}
