package main

import (
	"encoding/json"
	"fmt"
	"github.com/bartholdbos/golegram"
	"io/ioutil"
	"os"
)

type Configuration struct {
	TelegramUserId   int32  `json:"chat_id"`
	TelegramBotToken string `json:"token"`
}

var Config Configuration

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func getConfig() {
	raw, err := ioutil.ReadFile(os.Getenv("HOME") + "/.config/sendtgmsg.json")
	checkError(err)
	json.Unmarshal(raw, &Config)
}

func main() {
	getConfig()
	var bytes []byte
	var err error
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		fmt.Println("data is being piped to stdin")
		bytes, err = ioutil.ReadAll(os.Stdin)
		checkError(err)
	} else {
		fmt.Println("Please send data to this beautiful tool by using ' | sendtgmsg' after your command")
		fmt.Println("Checking configuration:")
		fmt.Println(Config.TelegramUserId)
		fmt.Println(Config.TelegramBotToken)
		os.Exit(2)
	}

	bot, err := golegram.NewBot(Config.TelegramBotToken)
	if err == nil {
		bot.SendMessage(Config.TelegramUserId, string(bytes))
	} else {
		fmt.Println(err.Error())
	}
}

