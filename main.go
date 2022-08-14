package main

import (
	_ "github.com/mattn/go-sqlite3"
	x90 "github.com/nopfault/x90-discobot/src"
)

func main() {

	x90.DatabaseInit()
	x90.BotStart()

	<-make(chan struct{})
	return
}
