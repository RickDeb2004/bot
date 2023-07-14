package bot

import (
	"discord-bot/config"
	"fmt"

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
func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == BotID {
		return
	}
	if message.Content == "hello" {
		_, _ = session.ChannelMessageSend(message.ChannelID, "hello, welcome to hackNITR discord server")
	}
	 else if message.Content=="!startquiz" {
		StartQuiz(session, message.ChannelID)
	}
	else {
		ananswerQuestion(session,message)

	}
	  
}
