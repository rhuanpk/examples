package consts

import "github.com/go-hl/os/env"

// APIPORT is the port of API.
var APIPORT = env.Get("APIPORT", ":9999")
