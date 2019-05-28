package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	g := pgs.Init(pgs.DebugEnv("DEBUG"))
	g.RegisterModule(NewGatewayModule())
	g.RegisterPostProcessor(pgsgo.GoFmt())
	g.Render()
}
