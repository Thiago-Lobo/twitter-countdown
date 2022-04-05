package main

import (
	"fmt"
	"os"
	"time"
	"twitter-countdown/pkg/twitter"
	"twitter-countdown/internal/constants"
	"twitter-countdown/internal/logic"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}

func main() {

	fmt.Println("Hello world!")
	
	fmt.Println(os.Args[1])
	fmt.Println(fmt.Sprintf(os.Args[1], 240))

	fmt.Println(logic.GetDaysToEvent("2023-01-01"))
	fmt.Println(logic.GetDaysToEvent("2022-04-04"))

	client := twitter.New(os.Getenv(constants.OAUTH_TOKEN), os.Getenv(constants.OAUTH_TOKEN_SECRET))
	client.Initialize()
	client.Test()

	// client.PostTweet("posted from golang - works fine :)")
	userId := client.LookupAuthenticatedUserInfo()
	
	query := logic.BuildQuery("posted from golang - works fine :)", userId)
	fmt.Println(query)
	tweets := client.LookupRecentTweets(query)

	fmt.Printf("%d", len(tweets))
	
}
