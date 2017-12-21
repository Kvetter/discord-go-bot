package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type spotify struct {
	Key    string `json:"access_token"`
	Typ    string `json:"token_type"`
	Expire int    `json:"expires_in"`
	Scope  string `json:"scope"`
}

// Variables used for command line parameters
var (
	Toke string
)

/*
func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}
*/
func main() {

	spotifyConnect("Kasper")
	/*
		discord, err := discordgo.New("Bot " + Token)
		// if there is an error, we print it to the console
		if err != nil {
			fmt.Println("error creating Discord session,", err)
			return
		}

		discord.AddHandler(messageCreate)
		// Open a websocket connection to Discord and begin listening.
		err = discord.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
			return
		}

		// Wait here until CTRL-C or other term signal is received.
		fmt.Println("Bot is now running.  Press CTRL-C to exit.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		// Cleanly close down the Discord session.
		discord.Close()
	*/

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		spotifyConnect("Kasper")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func spotifyConnect(name string) {

	//Creating a client to make a http call
	client := &http.Client{}

	//Creating data values to get the format application/x-www-form-urlencoded
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	t := spotify{}
	//Creating a http reqeust and formats the data
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))

	//Setting headers and encrypting the client and secret
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("")))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//Making the http call
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	readJson(resp.Body, &t)

	//Printing out the body


	fmt.Println(t.Key)

}

func readJson ([]byte body, interface *t)  {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err2 := json.Unmarshal([]byte(bodyBytes), &t)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
