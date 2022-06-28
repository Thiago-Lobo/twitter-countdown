package main

import (
	"fmt"
	"os"
	"time"
	"twitter-countdown/internal/constants"
	"twitter-countdown/internal/logic"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}

func getEnvVarOrDie(name string) string {

	result := os.Getenv(name)

	if result == "" {
		fmt.Printf("Please provide the environment variable \"%s\"\n", name)
		os.Exit(1)
	}

	return result

}

func main() {

	getEnvVarOrDie(constants.GOTWI_API_KEY)
	getEnvVarOrDie(constants.GOTWI_API_KEY_SECRET)

	logic.Initialize(
		getEnvVarOrDie(constants.OAUTH_TOKEN),
		getEnvVarOrDie(constants.OAUTH_TOKEN_SECRET),
		getEnvVarOrDie(constants.TWEET_TEMPLATE),
		getEnvVarOrDie(constants.END_DATE),
		getEnvVarOrDie(constants.POST_TIME),
		constants.StringToTweetReplySetting(os.Getenv(constants.REPLY_SETTING)),
	)

}
