package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const LightNovelCalenderURI = "https://calendar.gameiroiro.com/litenovel.php"

type LightNovelCalender [][]LightNovelBookInfo

type LightNovelBookInfo struct {
	Day       int
	Type      string
	Title     string
	AmazonURI string
	ThumbURI  string
	Author    string
	Publisher string
}

func NewLightNovelCalender() (LightNovelCalender, error) {
	resp, err := http.Get(LightNovelCalenderURI)
	if err != nil {
		return nil, errors.Wrap(err, "Request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ReadBody")
	}

	const AmazonDBURI = "https://www.amazon.co.jp/dp"
	const AmazonThumbURI = "https://m.media-amazon.com/images/I"

	var Calender = make(LightNovelCalender, 0)
	sep := strings.Split(string(body), `class="day-td"`)

	for _, day := range sep {
		// Get date
		tmps := strings.SplitN(day, `name="`, 2)

		if len(tmps) < 2 {
			continue
		}

		var dayString = strings.SplitN(tmps[1], `"`, 2)[0]
		d, err := strconv.Atoi(dayString)
		if err != nil {
			continue
		}

		var DayBooks = []LightNovelBookInfo{}

		// Parse book info
		for _, book := range strings.Split(day, `class="div-wrap"`) {
			var dayBook LightNovelBookInfo

			dayBook.Day = d

			// parse book type
			tmps = strings.SplitN(book, `<span>`, 2)
			if len(tmps) < 2 {
				continue
			}

			dayBook.Type = strings.TrimSpace(strings.SplitN(tmps[1], `</span>`, 2)[0])

			tmps = strings.Split(book, AmazonDBURI)
			if len(tmps) < 3 {
				continue
			}

			// parse book URI
			dayBook.AmazonURI = strings.TrimSpace(strings.SplitN(tmps[2], "?tag=game-calendar", 2)[0])
			dayBook.AmazonURI = AmazonDBURI + dayBook.AmazonURI

			// parse book Title
			tmps2 := strings.Split(tmps[2], `target="_blank">`)
			if len(tmps2) < 2 {
				continue
			}

			dayBook.Title = strings.TrimSpace(strings.Split(tmps2[1], "</a>")[0])

			// parse book ThumbURI
			tmps2 = strings.Split(tmps[1], AmazonThumbURI)
			if len(tmps2) < 2 {
				continue
			}

			dayBook.ThumbURI = AmazonThumbURI + strings.TrimSpace(strings.Split(tmps2[1], `"`)[0])

			// parse book Author and Publisher
			tmps2 = strings.Split(tmps[2], "<span>")
			if len(tmps2) < 3 {
				continue
			}

			dayBook.Publisher = strings.TrimSpace(strings.Split(tmps2[1], "</span>")[0])
			dayBook.Author = strings.TrimSpace(strings.Split(tmps2[2], "</span>")[0])

			DayBooks = append(DayBooks, dayBook)
		}

		Calender = append(Calender, DayBooks)
	}

	return Calender, nil
}

func (l LightNovelCalender) DayBooks(day int) []LightNovelBookInfo {
	for _, dayBooks := range l {
		if len(dayBooks) == 0 || dayBooks[0].Day != day {
			continue
		}
		return dayBooks
	}
	return nil
}
