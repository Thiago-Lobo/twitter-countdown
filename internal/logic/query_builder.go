package logic

import (
	"fmt"
	"strings"
	"regexp"
)

func stripRegex(in string) string {
    reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
    return reg.ReplaceAllString(in, "")
}

func BuildSubQueryFromTweetTemplate(tweetTemplate string) string {
	tweetTemplate = stripRegex(tweetTemplate)
	tokens := strings.Fields(tweetTemplate)
	result := strings.Join(tokens[:], " OR ")

	return fmt.Sprintf("(%s)", result)
}

func BuildSubQueryFromUserId(userId string) string {
	return fmt.Sprintf("from:%s", userId)
}

func BuildQuery(tweetTemplate string, userId string) string {
	return fmt.Sprintf("%s %s", BuildSubQueryFromTweetTemplate(tweetTemplate), BuildSubQueryFromUserId(userId))
}
