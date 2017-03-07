package config

import "flag"

var ServerHTTPAddr *string
var ServerMode *string

func ServerParams() {
	envServerHTTPPort := envString("NF_PORT", "")
	envServerMode := envString("NF_MODE", "debug")

	ServerHTTPAddr = flag.String("server.httpAddr", ":"+envServerHTTPPort, "HTTP listen address")
	ServerMode = flag.String("server.mode", envServerMode, "Server mode")
}
