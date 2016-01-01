package main

import (
	"../hipchat"
	"flag"
)

func main() {
	var msgFileName = flag.String("message-file", "messages.json", "Source message file to post to HipChat")
	flag.Parse()

	hipchat.Post(*msgFileName)
}
