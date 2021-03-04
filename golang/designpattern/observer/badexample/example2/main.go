package main

import (
	"golang/designpattern/observer/badexample/example2/player"
)

func main() {
	p := player.NewPlayer("xie4ever")
	_ = p.Attack()
}
