package bot

import (
	"discord-bot/config"
	"fmt"
	_ "discord-bot/bot/quiz"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotID = u.ID
	goBot.AddHandler(messageHandler)

	goBot.AddHandler(answerQuestion)
	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running")

}

	func sendCustomEmbed(session *discordgo.Session, channelID string) {
		embed := &discordgo.MessageEmbed{
			URL:         "",
			Type:        "",
			Title:       "Custom Embed",
			Description: "This is a custom embed message.",
			Timestamp:   "",
			Color:       0x00ff00, // Green color
			Footer:      nil,
			Image:       nil,
			Thumbnail:   nil,
			Video:       nil,
			Provider:    nil,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Field 1",
					Value:  "Value 1",
					Inline: false,
				},
				{
					Name:   "Field 2",
					Value:  "Value 2",
					Inline: false,
				},
			},
		}
}
// Send the embedded message using Channel Message with Embedded (CME).
var _ error
 _,err:=session.ChannelMessageAndEmbded(channelID,embed)
 // Check for errors and handle them appropriately
if err!=nil{
	fmt.Println(err.Error())
	return 
}


func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID {
		return
	}
	if message.Content == "hello" {
		_, _ = session.ChannelMessageSend(message.ChannelID, "hello, welcome to hackNITR discord server")
	}
	 else {
		if message.Content=="!startquiz" {
			StartQuiz(session, message.ChannelID)
		}
		

	}
	else if message.Content=="!customembeded"{
		sendCustomEmbed(session,message.ChannelID)
	}
	else {
		answerQanswerQuestion(session,message)
	}
	  
}
