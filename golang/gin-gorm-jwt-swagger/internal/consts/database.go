package consts

import "github.com/go-hl/os/env"

// Database variables.
var (
	DBUSER     = env.Get("DBUSER", "root")
	DBPASSWORD = env.Get("DBPASSWORD", "root")
	DBPROTOCOL = env.Get("DBPROTOCOL", "tcp")
	DBHOST     = env.Get("DBHOST", "localhost")
	DBPORT     = env.GetShouldInt("DBPORT", 3306)
	DBNAME     = env.Get("DBNAME", "test")
)
