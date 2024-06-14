package bot

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	addcommand "github.com/nathanjcook/discordbotgo/add_command"
	deletecommand "github.com/nathanjcook/discordbotgo/delete_command"
	helpcommand "github.com/nathanjcook/discordbotgo/help_command"
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
	var title string
	var msg string

	if m.Author.ID == BotId {
		return
	}

	if strings.Contains(m.Content, os.Getenv("BOT_PREFIX")) {
		cmdsplit := strings.Split(m.Content, " ")
		fmt.Println(cmdsplit)

		if cmdsplit[1] == "add" {
			p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
			if err != nil {
				title = "Add Command Error"
				msg = "Error Running Command"
				panic(err)
			}
			if len(cmdsplit) < 5 || len(cmdsplit) > 6 {
				title = "Add Command Error"
				msg = "Invalid Amount Of Args Provided"
				_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid Amount Of Args Provided")
			} else {
				if p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
					title, msg = addcommand.Add(cmdsplit[2], cmdsplit[3], cmdsplit[4])
				} else {
					title = "Add Command Error"
					msg = "Only Admins Can Add MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Added"
				}
			}
			embedAdd := discordgo.MessageEmbed{
				Title:       title,
				Description: msg,
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embedAdd)
		}

		if cmdsplit[1] == "delete" {
			p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
			if err != nil {
				title = "Delete Command Error"
				msg = "Error Running Command"
				panic(err)
			}
			if len(cmdsplit) < 3 || len(cmdsplit) > 4 {
				title = "Delete Command Error"
				msg = "Invalid Amount Of Args Provided"
			} else {
				if p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
					title, msg = deletecommand.Delete(cmdsplit[2])
				} else {
					title = "Delete Command Error"
					msg = "Only Admins Can Delete MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Deleted"
				}
			}
			embedDelete := discordgo.MessageEmbed{
				Title:       title,
				Description: msg,
			}

			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embedDelete)

		}

		if cmdsplit[1] == "help" {
			if len(cmdsplit) > 2 {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid Amount Of Args Provided")
			} else {
				var at, a, dt, d, it, i, mt, micro = helpcommand.Help()
				embedAdd := discordgo.MessageEmbed{
					Title:       at,
					Description: a,
				}
				embedDelete := discordgo.MessageEmbed{
					Title:       dt,
					Description: d,
				}
				embedInfo := discordgo.MessageEmbed{
					Title:       it,
					Description: i,
				}
				embedMicro := discordgo.MessageEmbed{
					Title:       mt,
					Description: micro,
				}
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embedAdd)
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embedDelete)
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embedInfo)
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embedMicro)
			}
		}
	}
}
