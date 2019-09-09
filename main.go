package main

import (
	newpkg "awesomeProject18/go-newpkg"
	packer "awesomeProject18/mydir"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	packer.DemoFunc()
	newpkg.DemoFunc()
	_ = yaml.Decoder{}
}
