// +build linux,cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp6315

/*
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/v6.3.15_20190220_api_tradeapi_se_linux64 -Wl,-rpath,${SRCDIR}/v6.3.15_20190220_api_tradeapi_se_linux64 -lthostmduserapi6315 -lthosttraderapi6315 -lstdc++
#cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/v6.3.15_20190220_api_tradeapi_se_linux64
*/
import "C"
