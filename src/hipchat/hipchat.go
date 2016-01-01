package hipchat

import (
	"fmt"
	"net/http"
	"bytes"
	"time"
	"encoding/json"
	"io/ioutil"
	"os"
	"math/rand"
)

const (
	hipchatEndpoint = "https://api.hipchat.com/v2/room/%s/message?auth_token=%s"
)

type Config struct {
	Token string `json:"token`
	Room string `json:"room`
}

func Post(msgFileName string) {
	conf := readConf()
	message := pickMessage(msgFileName)

	postErr := HttpPost(
		conf.Room,
		conf.Token,
		message,
	)
	if postErr != nil {
		fmt.Println(os.Stderr, "Post Error: ", postErr)
	}
}

func HttpPost(room, token, msg string) error {
	jsonStr, _ := json.Marshal(map[string]string{"message": msg})

	req, httpErr := http.NewRequest(
		"POST",
		fmt.Sprintf(hipchatEndpoint, room, token),
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if httpErr != nil {
		fmt.Println(os.Stderr, "HipChat API error: ", httpErr)
		return httpErr
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{ Timeout: time.Duration(15 * time.Second) }
	resp, clientErr := client.Do(req)
	defer resp.Body.Close()

	return clientErr
}

func readConf() Config {
	confStr, fileErr := ioutil.ReadFile("conf.json")
	if fileErr != nil {
		fmt.Println(os.Stderr, "File Open Error: ", fileErr)
	}

	var conf Config
	jsonErr := json.Unmarshal(confStr, &conf)
	if jsonErr != nil {
		fmt.Println(os.Stderr, "Json Parse Error: ", jsonErr)
	}

	return conf
}

func readMessages(msgFileName string) []string {
	messageStr, fileErr := ioutil.ReadFile(msgFileName)
	if fileErr != nil {
		fmt.Println(os.Stderr, "File Open Error: ", fileErr)
	}

	var messages []string
	jsonErr := json.Unmarshal(messageStr, &messages)
	if jsonErr != nil {
		fmt.Println(os.Stderr, "Json Parse Error: ", jsonErr)
	}

	return messages
}

func pickMessage(msgFileName string) string {
	messages := readMessages(msgFileName)
	rand.Seed(time.Now().UnixNano())
	return messages[rand.Intn(len(messages) -1)]
}
