package bot

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/nathanjcook/discordbotgo/bot/commands"
	dbconfig "github.com/nathanjcook/discordbotgo/config"
	"github.com/nathanjcook/discordbotgo/contentparser"
	"github.com/servusdei2018/shards"
)

var Mgr *shards.Manager
var BotId string

func Start() {
	var err error

	// Create a new shard manager using the provided bot token.
	Mgr, err = shards.New("Bot " + os.Getenv("BOT_TOKEN"))
	fmt.Println(err)
	if err != nil {
		log.Fatal("[ERROR] Error creating manager,", err)
		return
	}

	// Get details of bot to check if created
	user, err := Mgr.Gateway.User("@me")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	// Set BotId
	BotId = user.ID

	// Register handler for messages
	Mgr.AddHandler(messageCreate)
	// Register handler for sharding
	Mgr.AddHandler(onConnect)

	// Listen for messages in channels bot is member of Listen for DM messages & Listen for Messages
	Mgr.RegisterIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsGuilds | discordgo.IntentsMessageContent)

	fmt.Println("[INFO] Starting shard manager...")

	// Start all of our shards and begin listening.
	err = Mgr.Start()
	if err != nil {
		fmt.Println("[ERROR] Error starting manager,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("[SUCCESS] Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Shut down bot
	fmt.Println("[INFO] Stopping shard manager...")
	Mgr.Shutdown()
	fmt.Println("[SUCCESS] Shard manager stopped. Bot is shut down.")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var title string
	var msg string

	// check if sender is self, and don't reply if true
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
			} else {
				if p&discordgo.PermissionAdministrator != 0 {
					title, msg = commands.Add(cmdsplit[2], cmdsplit[3], cmdsplit[4])
				} else {
					title = "Add Command Error"
					msg = "Only Admins Can Add MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Added"
				}
			}
			embed := discordgo.MessageEmbed{
				Title:       title,
				Description: msg,
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		} else if cmdsplit[1] == "delete" {
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
				if p&discordgo.PermissionAdministrator != 0 {
					title, msg = commands.Delete(cmdsplit[2])
				} else {
					title = "Delete Command Error"
					msg = "Only Admins Can Delete MicroServices! Please Contact Any Administrators If You Want A Particular Microservice Deleted"
				}
			}
			embed := discordgo.MessageEmbed{
				Title:       title,
				Description: msg,
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		} else if cmdsplit[1] == "help" {
			if len(cmdsplit) > 2 {
				embederr := discordgo.MessageEmbed{
					Title:       "Help Command Error",
					Description: "Invalid Amount Of Args Provided",
				}
				_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embederr)
			} else {
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
			}
		} else if cmdsplit[1] == "info" {
			if len(cmdsplit) > 2 {
				title = "Info Command Error"
				msg = "Invalid Amount Of Args Provided"
			} else {
				title, msg = commands.Info()
			}
			embed := discordgo.MessageEmbed{
				Title:       title,
				Description: msg,
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		} else {
			type Microservice struct {
				MicroserviceId      int    `gorm:"column:microservice_id;unique;primaryKey;autoIncrement"`
				MicroserviceName    string `gorm:"column:microservice_name;size:25;"`
				MicroserviceUrl     string `gorm:"column:microservice_url;"`
				MicroserviceTimeout int    `gorm:"column:microservice_timeout;size:4;"`
			}

			var query Microservice

			if len(cmdsplit) < 3 {
				title = "Microservice Command Error"
				msg = "Invalid Amount Of Args Provided"
			} else {
				fmt.Println(string(cmdsplit[1]))
				host := dbconfig.DB.Table("microservices").Where("microservice_name = ?", string(cmdsplit[1])).Scan(&query)
				if host.RowsAffected > 0 {
					if host.RowsAffected > 0 {
						body := bytes.NewBuffer(contentparser.Body_Parser(m.Content))

						urls := (query.MicroserviceUrl + "/api/" + cmdsplit[2])

						resp, err := http.Post(urls, "application/json", body)
						if err != nil {
							fmt.Println("Err Placeholder")
						} else {
							if resp.StatusCode == 404 {
								title = cmdsplit[1] + "error"
								msg = "Endpoint Not Found"
							} else {
								defer resp.Body.Close()

								body, err := io.ReadAll(resp.Body)
								if err != nil {
									fmt.Println("Err Placeholder")
								} else {
									title = cmdsplit[1]
									msg = string(body)
								}
							}

						}
					}
				} else {
					title = "Microservice Command Error"
					msg = "Microservice Name Does Not Exist"
				}
			}
			embed := discordgo.MessageEmbed{
				Title:       title,
				Description: msg,
			}
			_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}
	}
}

// This function will be called a new shard connects
func onConnect(s *discordgo.Session, evt *discordgo.Connect) {
	fmt.Printf("[INFO] Shard #%v connected.\n", s.ShardID)
}
