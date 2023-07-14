package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func createPoll(session *discordgo.Session, channelID string, question string, options []string) {
	//Format the question and options
	pollContent := fmt.Sprintf("**%s**\n\n", question)
	var option string
	var i int
	for i, option = range options {
		pollContent += fmt.Sprintf("%d.%s\n", i+1, option)
	}
	// Send the poll message
	pollMessage, err := session.ChannelMessageSend(channelID, pollContent)
	if err != nil {
		session.ChannelMessageSend(channelID, "Failed to create the poll.")
		return
	}
	// Add reactions as options
	for i := range options {
		emoji := reactionEmojis[i]
		err := session.MessageReactionAdd(channelID, pollMessage.ID, emoji)
		if err != nil {
			session.ChannelMessageSend(channelID, "Failed to add reactions to the poll message.")
			return
		}
	}

	session.ChannelMessageSend(channelID, "Poll created! Vote by reacting to the options.")
}
func handleReactionAdd(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	if reaction.UserID == session.State.User.ID {
		return
	}

	// Check if the reaction is added to a poll message
	if reaction.MessageID != pollMessageID {
		return
	}

	// Get the selected option based on the reaction emoji
	selectedOption := ""
	for i, emoji := range reactionEmojis {
		if emoji == reaction.Emoji.Name {
			selectedOption = pollOptions[i]
			break
		}
	}

	if selectedOption != "" {
		// Send a confirmation message
		user, _ := session.User(reaction.UserID)
		response := fmt.Sprintf("Your vote for option **%s** has been recorded, %s!", selectedOption, user.Mention())
		session.ChannelMessageSend(reaction.ChannelID, response)
	}
}
