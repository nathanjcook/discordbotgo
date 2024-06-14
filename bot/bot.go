package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	addcommand "github.com/nathanjcook/discordbotgo/add_command"
	deletecommand "github.com/nathanjcook/discordbotgo/delete_command"
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

	// Listen for messages in channels bot is member of
	Mgr.RegisterIntent(discordgo.IntentsGuildMessages)
	// Listen for DM messages
	Mgr.RegisterIntent(discordgo.IntentsDirectMessages)

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
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Shut down bot
	fmt.Println("[INFO] Stopping shard manager...")
	Mgr.Shutdown()
	fmt.Println("[SUCCESS] Shard manager stopped. Bot is shut down.")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
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

// This function will be called a new shard connects
func onConnect(s *discordgo.Session, evt *discordgo.Connect) {
	fmt.Printf("[INFO] Shard #%v connected.\n", s.ShardID)
}
