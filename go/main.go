package main

import (
	"./log"
	"fmt"
	// "io"
)

func main() {
	// var out io.Writer
	// out.Write("I am io Writer")
	// p := Person{"Jeson", 23}
	// g := p
	// p.name = "others"
	// log.Trace(g.name)
	// log.Level = log.ERROR
	// str := "old string"
	// p := &str
	// changestring(p)
	// log.Trace(str)
	// var a string
	str := fmt.Sprintf("%d", log.TRACE)
	log.Trace(str)
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
