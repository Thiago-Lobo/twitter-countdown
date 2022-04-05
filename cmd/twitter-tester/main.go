package main

import (
	"fmt"
	"os"
	"twitter-countdown/pkg/twitter"
	"twitter-countdown/internal/constants"
)

func main() {

	fmt.Println("Hello world!")
	
	client := twitter.New(os.Getenv(constants.OAUTH_TOKEN), os.Getenv(constants.OAUTH_TOKEN_SECRET))
	client.Initialize()
	client.Test()

	// client.PostTweet("posted from golang - works fine :)")
	user_id := client.LookupAuthenticatedUserInfo()
	client.LookupRecentTweets(fmt.Sprintf("from:%s", user_id))
	
}
