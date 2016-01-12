package main

// http://blog.sgmansfield.com/2016/01/the-hidden-dangers-of-default-rand/

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/mahata/HipChabot/hipchat"
)

var msgFileName = flag.String("message-file", "messages.json", "Source message file to post to HipChat")
var confFileName = flag.String("conf-file", "conf.json", "Configuration file to post to HipChat")

type config struct {
	Token string `json:"token`
	Room  string `json:"room`
}

func readConf(confFileName string) (*config, error) {
	confStr, err := ioutil.ReadFile(confFileName)
	if err != nil {
		return nil, err
	}

	conf := &config{}
	err = json.Unmarshal(confStr, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func readMessages(msgFileName string) ([]string, error) {
	messageStr, err := ioutil.ReadFile(msgFileName)
	if err != nil {
		return nil, err
	}

	var messages []string
	err = json.Unmarshal(messageStr, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func pickMessage(msgFileName string) (string, error) {
	messages, err := readMessages(msgFileName)
	if err != nil {
		return "", err
	}
	rand.Seed(time.Now().UnixNano())
	return messages[rand.Intn(len(messages)-1)], nil
}

func main() {
	flag.Parse()

	conf, err := readConf(*confFileName)
	if err != nil {
		fmt.Println("Conf File Error: ", err)
		return
	}

	msg, err := pickMessage(*msgFileName)
	if err != nil {
		fmt.Println(os.Stderr, "Messages File Error: ", err)
		return
	}

	fmt.Println("MSG:", msg, "ROOM:", conf.Room, "TOKEN:", conf.Token)
	return

	hipchat.Post(conf.Room, conf.Token, msg)
}
