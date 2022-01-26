package main

import (
	"fmt"
	"time"

	"github.com/TKMAX777/LightNovelChecker/slack_webhook"
)

func NewMessageBlock(books []LightNovelBookInfo, date time.Time) []slack_webhook.BlockBase {
	var blocks = []slack_webhook.BlockBase{}

	var title = slack_webhook.SectionBlock()
	title.Text = slack_webhook.MrkdwnElement(date.Format("1月2日出版のラノベ"))

	blocks = append(blocks, title)

	for _, tb := range books {
		var section = slack_webhook.SectionBlock()
		section.Text = slack_webhook.MrkdwnElement(fmt.Sprintf("【%s】\n<%s|%s>\n著者: %s", tb.Type, tb.AmazonURI, tb.Title, tb.Author))
		section.Accessory = slack_webhook.ImageElement(tb.ThumbURI, tb.Title)

		blocks = append(blocks, section)
	}

	if books == nil {
		blocks[0].Text.Text += "はありません"
	}

	return blocks
}
