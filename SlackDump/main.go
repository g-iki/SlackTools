package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"unsafe"

	"./structure"
)

var token string

func main() {

	readSetting()
	readSetting()
	chMapsPublic := getChannelMap("public_channel")
	fmt.Println("public===============================")
	for k, _ := range chMapsPublic {
		getConversationHistory(k)
	}
	fmt.Println("private===============================")
	chMapsPrivate := getChannelMap("private_channel")
	for k, v := range chMapsPrivate {
		fmt.Println(k, v)
	}
	fmt.Println("mpim===============================")
	chMapsMpin := getChannelMap("mpim")
	for k, v := range chMapsMpin {
		fmt.Println(k, v)
	}
	fmt.Println("im===============================")
	chMapsIm := getChannelMap("im")
	for k, v := range chMapsIm {
		fmt.Println(k, v)
	}

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

func getChannelMap(chType string) map[string]string {
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
		} else {
			m[v.ID] = v.Name
		}

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

	fmt.Println(*(*string)(unsafe.Pointer(&jb)))

	// d := new(structure.Conversations)

	// if err := json.Unmarshal(jb, d); err != nil {
	// 	fmt.Println("JSON Unmarshal error:", err)
	// 	return err
	// }

	// for _, v := range d.Messages {
	// 	m[v.ID] = v.Name
	// 	fmt.Println(v.ID, v.Name)
	// }

	return nil

}

// func fileWrite(filename string, jb ([]byte)(ba)) {
//     file, err := os.OpenFile(filename, os.O_CREATE 0666)
//     if err != nil {
//         //エラー処理
//         log.Fatal(err)
//     }
//     defer file.Close()
//     fmt.Fprintln(file, "書き込み〜")
// }
