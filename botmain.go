package main

import (
	"log"
	"strings"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rhnvrm/lyric-api-go"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1046861512:AAGOnyKtCMcRVebYosmBNO9ZUQbwDVyGcsU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = `type 
							/musiclyrics <artist> = <music-title> to get music lyrics`
							//playmusic <music-title> to listen music`
			case "start":
				msg.Text = `Hi!, Bitoven is here to serve you, 
							/musiclyrics <artist> = <music-title> to get music lyrics`
							//playmusic <music-title> to listen music`
			case "status":
				msg.Text = "I'm ok."
			case "playmusic":
				music:= update.Message.CommandArguments()
				msg.Text = music
			case "musiclyrics":
				str := update.Message.CommandArguments()
				if str != ""{
					retString := strings.SplitAfter(str,"=")
					msg.Text = GetLyrics(retString[0],retString[1])
				}else
				{
					msg.Text = "The wasn't correct, enter command followed by artist name and music title with '=' in between"
				}
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}

	}
}

func GetLyrics(artist string, song string) string {
	
	l := lyrics.New(lyrics.WithoutProviders(), lyrics.WithGeniusLyrics("NCkYiHGGf8FWdPbxM1Zt57-ssu8z5yZHHV8FWww1db2vr4wiSY7v2RXib2PwV5ux"))
	lyric, err := l.Search(artist, song)

	if err != nil {
		return ("Lyrics for "+artist+" - "+song+" were not found")
	}
	return (lyric)
}

