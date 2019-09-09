package main

import (
	"gopkg.in/yaml.v2"

	"awesomeProject18/go-newpkg"
	packer "awesomeProject18/mydir"
)

func main() {
	packer.DemoFunc()
	newpkg.DemoFunc()
	_ = yaml.Decoder{}
}
