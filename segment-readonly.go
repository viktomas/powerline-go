// +build !windows

package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func segmentPerms(p *powerline) {
	cwd := p.cwd
	if cwd == "" {
		cwd, _ = os.LookupEnv("PWD")
	}
	if unix.Access(cwd, unix.W_OK) != nil {
		p.appendSegment("perms", segment{
			content:    fmt.Sprintf(" %s ", p.symbolTemplates.Lock),
			foreground: p.theme.ReadonlyFg,
			background: p.theme.ReadonlyBg,
		})
	}
}
