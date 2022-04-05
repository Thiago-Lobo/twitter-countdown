package main

import (
	"fmt"
	"os"
	"twitter-countdown/pkg/twitter"
	"twitter-countdown/internal/constants"
	"twitter-countdown/internal/logic"
)

func main() {

	fmt.Println("Hello world!")
	
	fmt.Println(os.Args[0])

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
