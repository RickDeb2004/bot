package bot

import (
	"encoding/json"
	"fake-bot/config"
	"fmt"
	"io"
	"net/http"
	"github.com/bwmarrin/discordgo"
)
type GithubWebPayLoad struct{

	Repository struct{
		Name string `json:"name"`
	}`json:"repository"`
}
func WebhookHandler(w http.ResponseWriter, r *http.Request){
	body,err:=io.ReadAll(r.Body)
	if err!=nil{

		http.Error(w,"Failed to read",http.StatusInternalServerError)
		return
	}
var payLoad GithubWebPayLoad
	err=json.Unmarshal(body,&payLoad)
	if err!=nil{
		http.Error(w,"Failed to unmarshal",http.StatusInternalServerError)
		return
	}
	session,err:=discordgo.New("Bot "+ config.Token)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	message:=fmt.Sprintf("New push to GitHub Repository :%s",payLoad.Repository.Name)
	session.ChannelMessageSend("1128709087958421596",message)
	session.Close()
	w.WriteHeader(http.StatusOK)




}
