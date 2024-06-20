package bot

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nathanjcook/discordbotgo/bot/commands"
	dbconfig "github.com/nathanjcook/discordbotgo/config"
	"go.uber.org/zap"
)

type Microservice struct {
	MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
	MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
	MicroserviceUrl     string `gorm:"column:microservice_url;"`
	MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var title string
	var msg string
	var is_help bool

	var query Microservice

	// check if sender is self, and don't reply if true
	if m.Author.ID == BotId {
		return
	}

	if strings.Contains(m.Content, os.Getenv("BOT_PREFIX")) {
		cmdsplit := strings.Split(m.Content, " ")

		p, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
		if err != nil {
			zap.L().Error("Error", zap.Error(err))
		}
		adminCheck := p & discordgo.PermissionAdministrator

		messageContent := m.Content

		if cmdsplit[1] == "add" {
			title, msg = Add_Handler(int(adminCheck), cmdsplit)

		} else if cmdsplit[1] == "delete" {
			title, msg = Delete_Handler(int(adminCheck), cmdsplit)

		} else if cmdsplit[1] == "help" {
			title, msg, is_help = Help_Handler(cmdsplit)

		} else if cmdsplit[1] == "info" {
			title, msg = Info_Handler(cmdsplit)

		} else {
			host := dbconfig.DB.Table("microservices").Where("microservice_name = ?", string(cmdsplit[1])).Scan(&query)
			if host.RowsAffected < 0 {
				title = "Microservice Command Error"
				msg = "Microservice Name Does Not Exist"
			} else {
				title, msg = Microservice_Handler(query, cmdsplit, messageContent)
			}
		}
		if is_help {
			embed := discordgo.MessageEmbed{
				Title: "Help!!!",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  commands.AddTitle,
						Value: commands.AddMsg,
					},
					{
						Name:  commands.DeleteTitle,
						Value: commands.DeleteMsg,
					},
					{
						Name:  commands.InfoTitle,
						Value: commands.InfoMsg,
					},
					{
						Name:  commands.MicroserviceTitle,
						Value: commands.MicroserviceMsg,
					},
				},
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		} else {
			embed := discordgo.MessageEmbed{
				Title:       title,
				Description: msg,
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}
	}
}

func Add_Handler(adminCheck int, cmdsplit []string) (string, string) {
	var title string
	var msg string

	if len(cmdsplit) < 5 || len(cmdsplit) > 6 {
		title := "Add Command Error"
		msg := "Invalid Amount Of Args Provided"
		return title, msg
	} else if adminCheck == 0 {
		title = "Add Command Error"
		msg = "Only Admins Can Add MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Added"
		return title, msg
	} else {
		title, msg = commands.Add(cmdsplit[2], cmdsplit[3], cmdsplit[4])
		return title, msg
	}
}

func Delete_Handler(adminCheck int, cmdsplit []string) (string, string) {
	var title string
	var msg string

	if len(cmdsplit) < 3 || len(cmdsplit) > 4 {
		title := "Delete Command Error"
		msg := "Invalid Amount Of Args Provided"
		return title, msg
	} else if adminCheck == 0 {
		title = "Delete Command Error"
		msg = "Only Admins Can Delete MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Deleted"
		return title, msg
	} else {
		title, msg = commands.Delete(cmdsplit[2])
		return title, msg
	}
}

func Help_Handler(cmdsplit []string) (string, string, bool) {
	var title string
	var msg string
	var is_help bool

	if len(cmdsplit) > 2 {
		title := "Help Command Error"
		msg := "Invalid Amount Of Args Provided"
		return title, msg, is_help
	} else {
		is_help = true
		return title, msg, is_help
	}
}

func Info_Handler(cmdsplit []string) (string, string) {
	var title string
	var msg string

	if len(cmdsplit) > 2 {
		title := "Info Command Error"
		msg := "Invalid Amount Of Args Provided"
		return title, msg
	} else {
		title, msg = commands.Info()
		return title, msg
	}
}

func Microservice_Handler(query Microservice, cmdsplit []string, messageContent string) (string, string) {
	var title string
	var msg string

	if len(cmdsplit) < 3 {
		title = "Microservice Command Error"
		msg = "Invalid Amount Of Args Provided"
		return title, msg
	} else {
		txt, str := Body_Parser(messageContent)
		if str != "" {
			title = "Pre Microservice JSON Body Error"
			msg = str
			return title, msg
		} else {
			body := bytes.NewBuffer(txt)
			urls := (query.MicroserviceUrl + "/api/" + cmdsplit[2])
			resp, err := http.Post(urls, "application/json", body)
			if err != nil {
				title = cmdsplit[1] + "error"
				msg = "Error Connecting To Microservice"
				return title, msg
			} else {
				if resp.StatusCode == 404 {
					if cmdsplit[2] == "help" {
						title = cmdsplit[1] + " No Help"
						msg = "The Microservice " + cmdsplit[1] + "Does Not Have A Help Section! Report This To An Admin"
						return title, msg
					} else {
						title = cmdsplit[1] + " Endpoint Not Found"
						helper, txt := commands.Get_Help((query.MicroserviceUrl + "/api/help"))
						if txt != "" {
							msg = txt
						} else {
							msg = Body_Reader(helper)
						}
					}
					return title, msg
				} else {
					defer resp.Body.Close()
					body, err := io.ReadAll(resp.Body)

					if err != nil {
						title = cmdsplit[1] + "error"
						msg = "Error Reading Response Body"
						return title, msg
					} else {
						title = cmdsplit[1]
						msg = Body_Reader(body)
						return title, msg
					}
				}
			}
		}
	}
}

// This function will be called a new shard connects
func onConnect(s *discordgo.Session, evt *discordgo.Connect) {
	fmt.Printf("[INFO] Shard #%v connected.\n", s.ShardID)
}
