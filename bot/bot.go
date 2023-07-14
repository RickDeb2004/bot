package bot

import (
	"discord-bot/config"
	"fmt"
	_ "discord-bot/bot/quiz"
	"github.com/bwmarrin/discordgo"
	_ "discord-bot/bot/poll"
	"strings"
)

var BotID string
var goBot *discordgo.Session
var (
	pollMessageID  string
	pollOptions    []string
	reactionEmojis = []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣"}
)


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
	goBot.AddHandler(handleReactionAdd) 
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
	else if strings.HasPrefix(message.Content, "!poll") {
		args := strings.Split(message.Content, " ")
		if len(args) < 3 {
			session.ChannelMessageSend(message.ChannelID, "Invalid poll command. Usage: `!poll <question> <option1> <option2> ...`")
			return
		}

		question := args[1]
		options := args[3:]

		createPoll(session, message.ChannelID, question, options)
	}

	else {
		answerQanswerQuestion(session,message)
	}
	  
}
