package main

import (
	"fmt"
	"main/handlers"
	"os"
	//"os/signal"
	//"syscall"
	"github.com/bwmarrin/discordgo"
    "github.com/joho/godotenv"
)

var(
    Token = "Bot " + os.Getenv("TOKEN") //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	ChannelID = os.Getenv("CHANNEL_ID")
    stopBot = make(chan bool)
    vcsession *discordgo.VoiceConnection
)

func main() {
    //Discordのセッションを作成
    errr := godotenv.Load()
    Token = "Bot " + os.Getenv("TOKEN") //"Bot"という接頭辞がないと401 unauthorizedエラーが起きます
	ChannelID = os.Getenv("CHANNEL_ID")
    if errr != nil {
        fmt.Println("Error loading .env file")
        os.Exit(1)
    }
    discord, err := discordgo.New(Token)
    discord.Token = Token
    if err != nil {
        fmt.Println("Error logging in")
        fmt.Println(err)
    }

    handlers.RegisterHandlers(discord)
	//discord.AddHandler(onMessageCreate) //全てのWSAPIイベントが発生した時のイベントハンドラを追加
    // websocketを開いてlistening開始
    err = discord.Open()
    if err != nil {
        fmt.Println(err)
    }
    defer discord.Close()

    fmt.Println("Listening...")

    // Wait until Ctrl+C or another signal is received
	fmt.Println("The bot is now running. Press Ctrl+C to exit.")
	//sc := make(chan os.Signal, 1)
	//signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-stopBot //プログラムが終了しないようロック

	err = discord.Close()

    return
}

// メッセージの受信
func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	u := m.Author
	fmt.Printf("%20s %20s(%20s) > %s\n", m.ChannelID, u.Username, u.ID, m.Content)

    if(m.Author.Bot == false){
        s.ChannelMessageSend(ChannelID, m.Content)
    }
}
