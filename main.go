package main

import (
	"discord-history-bot/commands"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create Discord session
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	// Register command handlers
	discord.AddHandler(commands.HandleCommands)

	// Open connection to Discord
	err = discord.Open()
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	// Get test guild ID from env
	testGuildID := os.Getenv("TEST_GUILD_ID")

	// Register slash command
	command := &discordgo.ApplicationCommand{
		Name:        "history",
		Description: "Get chat history and save to file",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "messages",
				Description: "Number of messages to retrieve (max 1000)",
				Required:    true,
				MinValue:    &[]float64{1}[0],
				MaxValue:    1000,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        "save",
				Description: "Save the history file locally",
				Required:    false,
			},
		},
	}

	// Change the command registration to use test guild if available
	guildID := "" // Empty string means global command
	if testGuildID != "" {
		guildID = testGuildID // Use test guild for faster updates during development
	}

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, guildID, command)
	if err != nil {
		log.Fatal("Error creating slash command:", err)
	}

	fmt.Println("Bot is running. Press Ctrl+C to exit.")

	// Wait for interrupt signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	// Close Discord session
	discord.Close()
} 