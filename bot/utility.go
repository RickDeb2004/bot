package bot
import(
	"fmt"
	"github.com/bwmarrin/discordgo"
)
func DelMessages(session *discordgo.Session, message *discordgo.MessageCreate){
	channelID:=message.ChannelID
	limit:=100
	 messageContent:=fmt.Sprintf("last messgaes %d can be purged", limit)
	messages,err:=session.ChannelMessageSend(channelID,messageContent)
if err!=nil{
	session.ChannelMessageSend(channelID,"Failed")
	return
}
var messageIDs []string
for _,message:=range messages{
	messageIDs=append(messageIDs,message.ID)
	}
	err=session.ChannelMessagesBulkDelete(channelID,messageIDs)
	if err!=nil{
		session.ChannelMessageSend(channelID,"Failed")
		return
	}
	session.ChannelMessageSend(channelID,"Successfully delete messages")




}