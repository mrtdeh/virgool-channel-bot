package main

import (
	"log"
	"time"
	
	"gopkg.in/telegram-bot-api.v4"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("508437240:AAG_TAtrPPBEv5fkXV0xEkg8VO4yY-WOjm8")
	if err != nil {
		log.Panic(err)
	}

	//url := "https://virgool.io"
	chatId :=  "@VirgoolChannel"

	bot.Debug = true

	//log.Printf("Authorized on account %s", bot.Self.UserName)


	



	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	//updates, err := bot.GetUpdatesChan(u)

	lastUrl, url := "",""

	for  {

		url = GetNewPostUrl()

		if(lastUrl != "" && lastUrl == url){
			continue;
		}
		
		lastUrl = url;

		msg := tgbotapi.NewMessageToChannel(chatId, url )
		bot.Send(msg)
		
		time.Sleep(time.Second * 5)
	}
}


func GetNewPostUrl() string {

	doc, _ := goquery.NewDocument("https://virgool.io")
	post := doc.Find(".card.card-post .post-content").First()
	link, _ := post.Find("a").First().Attr("href")

	doc, _ = goquery.NewDocument( link )	
	shorturl := doc.Find(".shorturl-text").Text()
	
	return shorturl

  }
  