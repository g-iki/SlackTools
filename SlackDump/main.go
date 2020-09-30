package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"./structure"
)

var token string
var now = strconv.FormatInt(time.Now().Unix(), 10)
var userList map[string]string

func main() {

	readSetting()
	readSetting()
	// chMapsPublic := getChannelMap("public_channel")
	// for k, v := range chMapsPublic {
	// 	getConversationHistory(k, v)
	// }
	// chMapsPrivate := getChannelMap("private_channel")
	// for k, v := range chMapsPrivate {
	// 	getConversationHistory(k, v)
	// }
	// chMapsMpin := getChannelMap("mpim")
	// for k, v := range chMapsMpin {
	// 	getConversationHistory(k, v)
	// }

	userList = getUserList()
	chMapsIm := getChannelMap("im")
	for k, v := range chMapsIm {
		userName := userList[v]
		if userName == "" {
			userName = v
		}
		getConversationHistory(k, userName)
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

func getConversationHistory(ch string, chname string) error {
	u := "https://slack.com/api/conversations.history?token=" +
		token +
		"&channel=" + ch +
		"&limit=1000&pretty=1"
	res, _ := http.Get(u)
	defer res.Body.Close()

	ba, _ := ioutil.ReadAll(res.Body)

	jb := ([]byte)(ba)
	fileWeite(ch, chname, "main", *(*string)(unsafe.Pointer(&jb)))

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

func getUserList() map[string]string {
	m := map[string]string{}

	u := "https://slack.com/api/users.list?token=" +
		token +
		"&limit=1000&pretty=1"
	res, _ := http.Get(u)
	defer res.Body.Close()
	ba, _ := ioutil.ReadAll(res.Body)

	jb := ([]byte)(ba)
	d := new(structure.Users)

	if err := json.Unmarshal(jb, d); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return m
	}

	for _, v := range d.Members {
		var s string
		if v.RealName != "" {
			s = v.RealName
		} else {
			s = v.Name
		}
		s = strings.ReplaceAll(s, "/", "_")
		s = strings.ReplaceAll(s, " ", "")
		m[v.ID] = s
	}

	return m
}

func fileWeite(ch string, chname string, filename string, data string) error {
	dir := "output/" + now + "/" + chname + "/"

	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	path := dir + filename + ".json"
	err := ioutil.WriteFile(path, []byte(data), 0777)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
