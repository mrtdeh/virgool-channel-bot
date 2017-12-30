package main

import (
	"log"
//	"time"
	"strings"
	"os"
	
	"gopkg.in/telegram-bot-api.v4"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	bot_api_token := os.Getenv("BOT_API_TOKEN")

	bot, err := tgbotapi.NewBotAPI(bot_api_token)
	if err != nil {
		log.Panic(err)
	}

	chatId :=  "@VirgoolChannel"

	bot.Debug = true

	lastUrl, url := "",""

	for  {

		url = GetNewPostUrl()
		
		if(!strings.Contains(url, "vrgl.ir")){
		
			log.Printf("NOt contain")
			continue;
		
		}
		

		if(lastUrl != "" && lastUrl == url){
			continue;
		}
		
		lastUrl = url;

		msg := tgbotapi.NewMessageToChannel(chatId, url )
		bot.Send(msg)
		
		//time.Sleep(time.Second * 5)
	}
}


func GetNewPostUrl() string {

	doc,_ := goquery.NewDocument("https://virgool.io")
	post := doc.Find(".card.card-post .post-content").First()
	link, _ := post.Find("a").First().Attr("href")

	doc, _ = goquery.NewDocument( link )	
	
	shorturl := ""
	a := doc.Find(".shorturl-text")
	if a != nil{
		shorturl = a.Text()
	}
		
	return shorturl

  }
  
