package config

var ServerHTTPPort string
var ServerMode string

func ServerParams() {
	ServerHTTPPort = envString("NF_PORT", "")
	ServerMode = envString("NF_MODE", "debug")
}
