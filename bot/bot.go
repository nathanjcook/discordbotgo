package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
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

	// if m.content contains botid (Mentions) and "ping" then send "pong!"
	if m.Content == "<@"+BotId+"> ping" || m.Content == os.Getenv("BOT_PREFIX")+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
	}

}

// This function will be called a new shard connects
func onConnect(s *discordgo.Session, evt *discordgo.Connect) {
	fmt.Printf("[INFO] Shard #%v connected.\n", s.ShardID)
}
