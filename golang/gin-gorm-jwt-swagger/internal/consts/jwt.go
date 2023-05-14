package consts

import "github.com/go-hl/os/env"

// JWTSECRETKEY is the JWT secretkey token.
var JWTSECRETKEY = env.Get("JWTSECRETKEY", "")
