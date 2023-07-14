package bot

import (
	"math/rand"
	"time"
	
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	questions = []Question{
		{Question: "what is your name",
			Answer: "my name is bot"},
		{Question: "what is your age",
			Answer: "my age is 1 year"},
	}
	CurrentQuestion   Question
	quizActive        bool
	participants      map[string]bool
	quizTimer         *time.Timer
	quizTimerDuration = 20 * time.Second
)

type Question struct {
	Question string
	Answer   string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func StartQuiz(session *discordgo.Session, channelID string){
	if quizActive {
		session.ChannelMessageSend(channelID, "A quiz is already active. Please wait for the current one to finish.")
		return
	}
	CurrentQuestion = questions[rand.Intn(len(questions))] //randomly selecting questions
	session.ChannelMessageSend(channelID, "A new quiz has been started")
	session.ChannelMessageSend(channelID, "Question"+CurrentQuestion.Question)
	quizActive=true
	participants=make(map[string]bool)
	quizTimer=time.AfterFunc(quizTimerDuration, func(){EndQuiz(session,channelID)})
}
func EndQuiz(session *discordgo.Session, channelID string){
	quizActive=false
	// Check if anyone participated
	if len(participants) == 0 {
		session.ChannelMessageSend(channelID, "No one participated in the quiz. Better luck next time!")
		return
}
 	// Select a random winner
	winners:=make([]string, 0, len(participants))
	for participant:=range participants{
		winners=append(winners,participant)
	}
	winner := winners[rand.Intn(len(winners))]
	session.ChannelMessageSend(channelID, "Time's up!")
	session.ChannelMessageSend(channelID, "The correct answer was: "+CurrentQuestion.Answer)
	session.ChannelMessageSend(channelID, "The winner is: <@"+winner+">")

}
func clearQuizData() {
	CurrentQuestion = Question{}
	participants = nil
	if quizTimer != nil {
		quizTimer.Stop()
		quizTimer = nil
	}
}
func answerQuestion(session *discordgo.Session, message *discordgo.MessageCreate){
	if !quizActive {
		return
	}
	//strings.EqualFold is a string comparison function that checks if two strings are equal while ignoring the case sensitivity. It compares the message.Content (content of the received message) and currentQuestion.Answer (the correct answer to the current question).

	if strings.EqualFold(message.Content, CurrentQuestion.Answer) {
		participants[message.Author.ID] = true
	}

}

