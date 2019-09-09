package main

import (
	newpkg "awesomeProject18/go-newpkg"
	packr "awesomeProject18/mydir"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	packr.DemoFunc()
	newpkg.DemoFunc()
	_ = yaml.Decoder{}
}
