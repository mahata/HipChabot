package hipchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	hipchatEndpoint = "https://api.hipchat.com/v2/room/%s/message?auth_token=%s"
)

func Post(room, token, msg string) {
	postErr := httpPost(
		room,
		token,
		msg,
	)
	if postErr != nil {
		fmt.Println(os.Stderr, "Post Error: ", postErr)
	}
}

func httpPost(room, token, msg string) error {
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

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, clientErr := client.Do(req)
	defer resp.Body.Close()

	return clientErr
}
