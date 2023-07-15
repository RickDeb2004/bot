
# Discord-Bot

This bot project is designed to create a Discord bot with various features. The project consists of multiple files, including main.go, bot.go, config.go, and additional files for utility, poll, quiz and custom embed functionality.



##  Special Features
1.**main.go**: The main.go file serves as the entry point for the bot application. It reads the configuration from a JSON file, starts the bot, and keeps the application running.

2.**bot.go**: The bot.go file contains the core functionality of the bot. It initializes the DiscordGo session, handles the bot's start-up process, and includes a message handler function to respond to user messages. It also integrates the utility, poll, and custom embed  and quiz features.

3.**config.go**: The config.go file is responsible for reading the bot's configuration from a JSON file. It defines a configStruct type to hold the configuration values and provides a ReadConfig function to read and parse the configuration file.

4.**utility.go**:The utility functionality is added in a separate utility.go file. This includes features like spam prevention.

5.**poll.go**: The poll functionality is implemented in a poll.go file. It allows the bot to create and manage polls in the server, sending questions and recording user votes.

6.**quiz.go**:The quiz feature is added in a separate quiz.go file.This includes feature like asking question about teach stacks and for fun and entertainment.

**Custom embeded**   *enables the bot to send custom embedded messages with various fields like title, description, timestamp, color, and more.*


**In the config.json file you have to put your bot token and bot prefix.**


## Run Locally

Clone the project

```bash
  git clone https://github.com/RickDeb2004/bot
```

Go to the project directory

```bash
  cd discord-bot
```

Install dependencies

```bash
  go mod init
  go mod tidy
  go get "github.com/bwmarrin/discordgo"
```

Start the server

```bash
  go build
  go run main.go
```


## Tech Stack
**Go Lang**

