package main

import (
	"os"
	"time"
	"twitter-countdown/internal/logic"
	"twitter-countdown/internal/constants"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}

func main() {

	logic.Initialize(os.Getenv(constants.OAUTH_TOKEN), os.Getenv(constants.OAUTH_TOKEN_SECRET), os.Getenv(constants.TWEET_TEMPLATE), os.Getenv(constants.END_DATE), os.Getenv(constants.POST_TIME))

}
