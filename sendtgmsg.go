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
		fmt.Println(e.Error())
	}
}

func getConfig() {
	raw, err := ioutil.ReadFile(os.Getenv("HOME") + "/.config/sendtgmsg.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(raw, &Config)
}

func main() {
	getConfig()
	var bytes []byte
	var err error
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
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

		s := splitter(string(bytes))
		for _, piece := range s {
			msg, err := bot.SendMessage(Config.TelegramUserId, piece)
			if err != nil {
				fmt.Println(err)
				fmt.Println(msg)
			}
		}
	} else {
		fmt.Println(err.Error())
	}
}

func splitter(inputStr string) []string {
	a := []rune(inputStr)

	stringLength := len(a)
	maxChars := 4000
	amountOfSplits := (stringLength / maxChars)

	s := make([]string, amountOfSplits+1)

	for j := 0; j <= amountOfSplits; j++ {

		if j*maxChars < stringLength-maxChars {
			composed := string(a[j*maxChars : (j+1)*maxChars])
			s[j] = composed
		} else {
			s[j] = string(a[j*maxChars : stringLength])
		}
	}

	return s
}
