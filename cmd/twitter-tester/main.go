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

	logic.Initialize(os.Getenv(constants.OAUTH_TOKEN), os.Getenv(constants.OAUTH_TOKEN_SECRET), os.Args[1], "2023-01-01", "23:40:00")

}
