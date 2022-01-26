package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/TKMAX777/LightNovelChecker/slack_webhook"
)

var Settings []struct {
	Delta  int
	Hour   int
	Minute int
}

func init() {
	b, err := ioutil.ReadFile("settings.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &Settings)
	if err != nil {
		panic(err)
	}
}

func main() {
	calender, err := NewLightNovelCalender()
	if err != nil {
		log.Println(err.Error())
	}

	for {
		var date = time.Now()
		for _, s := range Settings {
			if date.Hour() != s.Hour || date.Minute() != s.Minute {
				continue
			}

			books := calender.DayBooks(date.Day())

			var slack = slack_webhook.New(os.Getenv("SLACK_TOKEN"))
			_, err = slack.Send(slack_webhook.Message{
				Channel:     os.Getenv("SLACK_CHANNEL"),
				Text:        date.Format("1月2日出版のラノベ"),
				Blocks:      NewMessageBlock(books, date),
				UnfurlLinks: false,
			})

			if err != nil {
				log.Println(err.Error())
			}
		}

		time.Sleep(time.Minute)
	}
}
