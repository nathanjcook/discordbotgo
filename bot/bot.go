package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var BotId string

func Start() {

	goBot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	user, err := goBot.User("@me")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	BotId = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	// if m.content contains botid (Mentions) and "ping" then send "pong!"
	if m.Content == "<@"+BotId+"> ping" || m.Content == os.Getenv("BOT_PREFIX")+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}

}
