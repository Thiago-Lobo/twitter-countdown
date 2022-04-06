package twitter

import (
	"context"
	"fmt"
	"time"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/tweets"
	"github.com/michimani/gotwi/users"
	tt "github.com/michimani/gotwi/tweets/types"
	ut "github.com/michimani/gotwi/users/types"
)

type TwitterClient struct {
	OAuthToken 			string
	OAuthTokenSecret 	string
	Client				*gotwi.GotwiClient
}

// Constructor

func New(OAuthToken_ string, OAuthTokenSecret_ string) *TwitterClient {
	result := new(TwitterClient)

	result.OAuthToken = OAuthToken_
	result.OAuthTokenSecret = OAuthTokenSecret_

	return result
}

func (tc *TwitterClient) Initialize() {

	in := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           tc.OAuthToken,
		OAuthTokenSecret:     tc.OAuthTokenSecret,
	}

	c, err := gotwi.NewGotwiClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	tc.Client = c
	
}

func (tc *TwitterClient) Test() {

	p := &ut.UserLookupByUsernameParams{
		Username: "elonmusk",
	}

	
	res, err := users.UserLookupByUsername(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("TEST: [elonmusk]'s user ID is [%s]\n", *res.Data.ID)

}

func (tc *TwitterClient) PostTweet(text string) {
	p := &tt.ManageTweetsPostParams{
		Text: gotwi.String(text),
	}

	res, err := tweets.ManageTweetsPost(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Posted tweet: [%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}

func (tc *TwitterClient) LookupAuthenticatedUserInfo() string {
	
	p := &ut.UserLookupMeParams{}

	res, err := users.UserLookupMe(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fmt.Printf("Authenticated user's ID is [%s]\n", gotwi.StringValue(res.Data.ID))
	
	return *res.Data.ID
}

func (tc *TwitterClient) LookupRecentTweets(query string) []resources.Tweet {
	
	p := &tt.SearchTweetsRecentParams{
		Query: query,
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldText,
			fields.TweetFieldCreatedAt,
		},
	}

	res, err := tweets.SearchTweetsRecent(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return res.Data

}

// TODO: figure out why this one doesn't work :(
func (tc *TwitterClient) LookupRecentTweetsWithTimeRange(query string, startTime time.Time, endTime time.Time) []resources.Tweet {
	utcStartTime := startTime.UTC()
	utcEndTime := endTime.UTC()
	
	p := &tt.SearchTweetsRecentParams{
		Query: query,
		StartTime: &utcStartTime,
		EndTime: &utcEndTime,
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldText,
			fields.TweetFieldCreatedAt,
		},
	}

	res, err := tweets.SearchTweetsRecent(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return res.Data

}