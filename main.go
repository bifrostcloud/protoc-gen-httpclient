package main

import (
	modulego "github.com/bifrostcloud/protoc-gen-httpclient/modules/go"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	plugin := pgs.Init(pgs.DebugEnv("DEBUG"))
	// Registering go codegen module
	plugin.RegisterModule(modulego.New())
	// TODO : This may cause issues when other language targets are added
	plugin.RegisterPostProcessor(pgsgo.GoFmt()).Render()
}
