package main

import (
	"./log"
)

// func vary(a ...int) {
// 	log.Debug("sdf %t", a)
// 	log.Debug("sdf %d", len(a))
// }
func changestring(p *string) {
	*p = "new string"
}
func main() {
	str := "old string"
	p := &str
	changestring(p)
	log.Trace(str)
	// var a string
	log.Trace("hello")
	log.Debug("hello")
	log.Info("hello")
	log.Warn("hello")
	log.Error("hello")
	log.Fatal("hello")
	// vary(1, 2, 3)
	// log.Debug("sdf %T", a)
	// log.Debug("%d", log.l)
}
