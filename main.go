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

	// Register slash command
	command := &discordgo.ApplicationCommand{
		Name:        "history",
		Description: "Get chat history and save to file",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "messages",
				Description: "Number of messages to retrieve",
				Required:    true,
			},
		},
	}

	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "", command)
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