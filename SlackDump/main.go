package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"./structure"
)

var token string

func main() {

	readSetting()
	chMaps := getChannelMap()
	getConversationHistory(chMaps)
}

func readSetting() {
	r, err := ioutil.ReadFile("./setting.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var s structure.Settings

	json.Unmarshal(r, &s)
	token = s.Token

}

func getChannelMap() map[string]string {
	m := map[string]string{}
	u := "https://slack.com/api/conversations.list?token=" +
		token +
		"&limit=1000&pretty=1"
	res, _ := http.Get(u)
	defer res.Body.Close()
	ba, _ := ioutil.ReadAll(res.Body)

	jb := ([]byte)(ba)
	d := new(structure.ChannelLists)

	if err := json.Unmarshal(jb, d); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return m
	}

	for _, v := range d.Channels {
		m[v.ID] = v.Name
		fmt.Println(v.ID, v.Name)
	}

	return m
}

func getConversationHistory(ch string) error {
	u := "https://slack.com/api/conversations.history?token=" +
		token +
		"&channel=" + ch +
		"&limit=1000&pretty=1"
	res, _ := http.Get(u)
	defer res.Body.Close()

	ba, _ := ioutil.ReadAll(res.Body)

	jb := ([]byte)(ba)
	d := new(structure.Conversations)

	if err := json.Unmarshal(jb, d); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return err
	}

	for _, v := range d.Messages {
		m[v.ID] = v.Name
		fmt.Println(v.ID, v.Name)
	}

	return nil

}
