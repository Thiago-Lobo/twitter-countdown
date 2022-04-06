package logic

import (
	"fmt"
	
	"twitter-countdown/pkg/twitter"
)

func run(client *twitter.TwitterClient, userId string, tweetTemplate string, targetDate string, postTime string) {

	taskClosure := func () bool {
		fmt.Println(">> Starting task!")
		remainingDays := GetDaysToEvent(targetDate)
		fmt.Println("Remaining days: ", remainingDays)
		
		if (remainingDays <= 0) {
			return true
		}
		
		formattedTweet := fmt.Sprintf(tweetTemplate, remainingDays)
		fmt.Println("Formatted tweet: ", formattedTweet)

		// TODO: add time range to query!
		query := BuildQuery(formattedTweet, userId)
		fmt.Println("Query: ", query)
		tweets := client.LookupRecentTweets(query)

		if (len(tweets) == 0) {
			fmt.Println("Will tweet!")
		} else {
			fmt.Println("Won't tweet!")
		}

		return false
	}

	// ScheduleTask(taskClosure, )

}

func Initialize(oauthToken string, oauthTokenSecret string, tweetTemplate string, targetDate string, postTime string) {

	/*
		Receive parameters
			oauth_token
			oauth_secret

			tweetTemplate
			targetDate
			postTime
	*/
	
	client := twitter.New(oauthToken, oauthTokenSecret)
	client.Initialize()
	client.Test()

	userId := client.LookupAuthenticatedUserInfo()
	
	run(client, userId, tweetTemplate, targetDate, postTime)

	// query builder
	// fmt.Println(os.Args[1])
	// fmt.Println(fmt.Sprintf(os.Args[1], 240))

	// fmt.Println(logic.GetDaysToEvent("2023-01-01"))
	// fmt.Println(logic.GetDaysToEvent("2022-04-04"))
	// fmt.Println(logic.GetDaysToEvent("2022-04-05"))
	// fmt.Println(logic.GetDaysToEvent("2022-04-06"))
	
	// twitter

	// client.PostTweet("posted from golang - works fine :)")
	
	// query := logic.BuildQuery("posted from golang - works fine :)", userId)
	// fmt.Println(query)
	// tweets := client.LookupRecentTweets(query)

	// fmt.Printf("%d", len(tweets))
}


