package logic

import (
	"fmt"
	"time"
	"strings"
	"twitter-countdown/pkg/twitter"
	"github.com/michimani/gotwi/resources"
)

func run(client *twitter.TwitterClient, userId string, tweetTemplate string, targetDate string, postTime string) {

	taskClosure := func () bool {
		fmt.Println(">> Starting task!")
		
		remainingDays := GetDaysToEvent(targetDate)
		fmt.Println("Remaining days: ", remainingDays)
		
		if (remainingDays <= 0) {
			return true
		}

		currentTime := time.Now()
		postTime := stringToTimestamp("15:04:05 -0700", fmt.Sprintf("%s -0300", postTime))
		postTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), postTime.Hour(), postTime.Minute(), postTime.Second(), 0, currentTime.Location())

		fmt.Println("Current time: ", currentTime)
		fmt.Println("Post time: ", postTime)

		if (currentTime.After(postTime)) {
			tweetTemplateForQuery := strings.Replace(tweetTemplate, "%d", "", -1)
			fmt.Println("Tweet template for query: ", tweetTemplateForQuery)
			
			queryStartTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
			queryEndTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 999999999, currentTime.Location())
			fmt.Println("Query startTime: ", queryStartTime)
			fmt.Println("Query endTime: ", queryEndTime)

			query := BuildQuery(tweetTemplateForQuery, userId)
			fmt.Println("Query: ", query)
			
			tweets := client.LookupRecentTweets(query)
			fmt.Println("Found tweets: ", len(tweets))
			
			validTweets := []resources.Tweet{}
			
			for _, tweet := range tweets {
				fmt.Println(*tweet.Text)
				fmt.Println(tweet.CreatedAt)
				
				if tweet.CreatedAt.After(queryStartTime.UTC()) && tweet.CreatedAt.Before(queryEndTime.UTC()) {
					validTweets = append(validTweets, tweet)
				}
			}
			
			fmt.Println("Found tweets after date filtering: ", len(validTweets))
			
			if (len(validTweets) == 0) {
				formattedTweet := fmt.Sprintf(tweetTemplate, remainingDays)
				fmt.Println("Formatted tweet: ", formattedTweet)
				fmt.Println("Will tweet!")
				client.PostTweet(formattedTweet)
			} else {
				fmt.Println("Won't tweet because already tweeted!")
			}
		} else {
			fmt.Println("Won't tweet because it's not time to tweet!")
		}

		return false
	}

	ScheduleTask(taskClosure, 40 * time.Second, true)

}

func Initialize(oauthToken string, oauthTokenSecret string, tweetTemplate string, targetDate string, postTime string) {
	
	client := twitter.New(oauthToken, oauthTokenSecret)
	client.Initialize()
	client.Test()

	fmt.Println()

	userId := client.LookupAuthenticatedUserInfo()
	
	run(client, userId, tweetTemplate, targetDate, postTime)

}


