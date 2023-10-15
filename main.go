package main

import (
	_ "github.com/coredns/caddy/onevent"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
	_ "github.com/coredns/coredns/plugin/acl"
	_ "github.com/coredns/coredns/plugin/any"
	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/bufsize"
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/cancel"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/dnssec"
	_ "github.com/coredns/coredns/plugin/errors"
	_ "github.com/coredns/coredns/plugin/forward"
	_ "github.com/coredns/coredns/plugin/header"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/loadbalance"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/loop"
	_ "github.com/coredns/coredns/plugin/metadata"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/minimal"
	_ "github.com/coredns/coredns/plugin/nsid"
	_ "github.com/coredns/coredns/plugin/ready"
	_ "github.com/coredns/coredns/plugin/reload"
	_ "github.com/coredns/coredns/plugin/rewrite"
	_ "github.com/coredns/coredns/plugin/secondary"
	_ "github.com/coredns/coredns/plugin/sign"
	_ "github.com/coredns/coredns/plugin/timeouts"
	_ "github.com/coredns/coredns/plugin/tls"
	_ "github.com/coredns/coredns/plugin/tsig"
	_ "github.com/coredns/coredns/plugin/view"
	_ "github.com/michaelbeaumont/coredns-tailscale"
	_ "github.com/michaelbeaumont/coredns-libvirt"
)

func main() {
	dnsserver.Directives = []string{
		"metadata",
		"cancel",
		"tls",
		"timeouts",
		"reload",
		"nsid",
		"bufsize",
		"bind",
		"debug",
		"ready",
		"health",
		"prometheus",
		"errors",
		"log",
		"acl",
		"any",
		"loadbalance",
		"tsig",
		"cache",
		"rewrite",
		"header",
		"dnssec",
		"minimal",
		"secondary",
		"loop",
		"forward",
		"on",
		"sign",
		"view",
		"tailscale",
		"libvirt",
	}
	coremain.Run()
}
