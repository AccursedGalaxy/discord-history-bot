package commands

import (
	"discord-history-bot/utils"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func HandleCommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	switch i.ApplicationCommandData().Name {
	case "history":
		handleHistory(s, i)
	}
}

func handleHistory(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Acknowledge the command immediately
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	// Get the number of messages to retrieve
	numMessages := int(i.ApplicationCommandData().Options[0].IntValue())
	if numMessages <= 0 || numMessages > 100 {
		content := "Please request between 1 and 100 messages."
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		})
		return
	}

	// Get messages
	messages, err := s.ChannelMessages(i.ChannelID, numMessages, "", "", "")
	if err != nil {
		errMsg := fmt.Sprintf("Error retrieving messages: %s", err.Error())
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &errMsg,
		})
		return
	}

	// Format messages for file
	var content string
	for i := len(messages) - 1; i >= 0; i-- {
		msg := messages[i]
		// Skip messages from bots and empty messages
		if msg.Author.Bot || msg.Content == "" {
			continue
		}
		timestamp := msg.Timestamp
		content += fmt.Sprintf("[%s] %s: %s\n", 
			timestamp.Format(time.RFC822), 
			msg.Author.Username, 
			msg.Content,
		)
	}

	// Add a check if no valid messages were found
	if content == "" {
		noMsgContent := "No valid messages found (excluding bot messages and empty messages)"
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &noMsgContent,
		})
		return
	}

	// Get save option from command
	var shouldSave bool
	if len(i.ApplicationCommandData().Options) > 1 {
		shouldSave = i.ApplicationCommandData().Options[1].BoolValue()
	}

	// Generate filename with timestamp
	filename := fmt.Sprintf("chat_history_%s.txt", time.Now().Format("2006-01-02_15-04-05"))

	// Save to temporary file
	err = utils.SaveToFile(filename, content)
	if err != nil {
		errMsg := fmt.Sprintf("Error saving file: %s", err.Error())
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &errMsg,
		})
		return
	}

	// Send file
	file := &discordgo.File{
		Name:   filename,
		Reader: utils.GetFileReader(filename),
	}

	finalMsg := fmt.Sprintf("Here's the last %d messages from this channel:", numMessages)
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &finalMsg,
		Files:   []*discordgo.File{file},
	})

	// Close and remove the file if we shouldn't save it
	if !shouldSave {
		if f, ok := file.Reader.(*os.File); ok {
			f.Close()
		}
		os.Remove(filename)
	}
} 