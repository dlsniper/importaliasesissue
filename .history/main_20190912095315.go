package main

import (
	"awesomeProject18/go-newpkg"
	packr "awesomeProject18/mydir"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

func main() {
	packr.DemoFunc()
	newpkg.DemoFunc()
	_ = yaml.Decoder{}

	_, _ = homedir.Dir()
}
