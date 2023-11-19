package main

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"

	_ "github.com/novoleg/rrl/plugins/rrl"

)

func init() {
	dnsserver.Directives = append([]string{"rrl"}, dnsserver.Directives...)
}

func main() {
	coremain.Run()
}
