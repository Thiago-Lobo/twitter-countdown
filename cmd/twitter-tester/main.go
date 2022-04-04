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
}
