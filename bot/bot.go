package bot

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	addcommand "github.com/nathanjcook/discordbotgo/add_command"
	deletecommand "github.com/nathanjcook/discordbotgo/delete_command"
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

	if strings.Contains(m.Content, os.Getenv("BOT_PREFIX")) {
		cmdsplit := strings.Split(m.Content, " ")
		fmt.Println(cmdsplit)

		if cmdsplit[1] == "add" {
			p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Error Running Command")
				panic(err)
			}
			if len(cmdsplit) < 5 || len(cmdsplit) > 6 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid Amount Of Args Provided")
			} else {
				if p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
					_, _ = s.ChannelMessageSend(m.ChannelID, addcommand.Add(cmdsplit[2], cmdsplit[3], cmdsplit[4]))
				} else {
					_, _ = s.ChannelMessageSend(m.ChannelID, "Only Admins Can Add MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Added")
				}
			}
		}

		if cmdsplit[1] == "delete" {
			p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Error Running Command")
				panic(err)
			}
			if len(cmdsplit) < 3 || len(cmdsplit) > 4 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid Amount Of Args Provided")
			} else {
				if p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
					_, _ = s.ChannelMessageSend(m.ChannelID, deletecommand.Delete(cmdsplit[2]))
				} else {
					_, _ = s.ChannelMessageSend(m.ChannelID, "Only Admins Can Add MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Added")
				}
			}
		}
	}
}
