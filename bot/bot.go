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
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == "hello" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "hello, welcome to hackNITR discord server")
	}
	else if m.Content=="!startquiz" {
		StartQuiz(s, m.ChannelID)
	}
	else{

	}
	  
}
