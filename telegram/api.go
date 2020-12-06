package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const APIURL string = "https://api.telegram.org/bot"

type Api struct {
	ApiUrl string
}

func (this Api) New(token string) Api {
	this.ApiUrl = APIURL + token
	return this
}

func (this Api) SendMessage(message Message) {
	jsonMessage, _ := json.Marshal(message)
	resp, err := http.Post(this.ApiUrl+"/sendMessage", "application/json", strings.NewReader(string(jsonMessage)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body := getContents(*resp)

	fmt.Println(body)

	if resp.StatusCode != 200 {
		os.Exit(1)
	}
}

func getContents(resp http.Response) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(body)
}
