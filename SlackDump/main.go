package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"./structure"
)

func main() {

	setting := readSetting()
	chMapsPublic := getChannelMap("public_channel", setting.Token)
	chMapsPrivate := getChannelMap("private_channel", setting.Token)
	chMapsMpin := getChannelMap("mpim", setting.Token)
	chMapsIm := getChannelMap("im", setting.Token)

	fmt.Print(chMapsPublic, chMapsPrivate, chMapsMpin, chMapsIm)
}

func readSetting() structure.Settings {
	r, err := ioutil.ReadFile("./setting.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var s structure.Settings

	json.Unmarshal(r, &s)

	return s
}

func getChannelMap(chType string, token string) map[string]string {
	m := map[string]string{}
	u := "https://slack.com/api/conversations.list?token=" +
		token +
		"&types=" +
		chType +
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
		if v.User != "" {
			m[v.ID] = v.User
			fmt.Println(v.ID, v.User)
		} else {
			m[v.ID] = v.Name
			fmt.Println(v.ID, v.Name)
		}

	}

	return m
}
