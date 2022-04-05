package twitter

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweets"
	"github.com/michimani/gotwi/tweets/types"
	tt "github.com/michimani/gotwi/tweets/types"
	"github.com/michimani/gotwi/users"
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

	fmt.Println(tc.OAuthToken)
	fmt.Println(tc.OAuthTokenSecret)

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
		Expansions: fields.ExpansionList{
			fields.ExpansionPinnedTweetID,
		},
		UserFields: fields.UserFieldList{
			fields.UserFieldCreatedAt,
		},
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldCreatedAt,
		},
	}

	u, err := users.UserLookupByUsername(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID: ", gotwi.StringValue(u.Data.ID))
	fmt.Println("Name: ", gotwi.StringValue(u.Data.Name))
	fmt.Println("Username: ", gotwi.StringValue(u.Data.Username))
	fmt.Println("CreatedAt: ", u.Data.CreatedAt)

	if u.Includes.Tweets != nil {
		for _, t := range u.Includes.Tweets {
			fmt.Println("PinnedTweet: ", gotwi.StringValue(t.Text))
		}
	}
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
	
	p := &ut.UserLookupMeParams{
		
	}

	res, err := users.UserLookupMe(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fmt.Printf("Authenticated user's ID is [%s]\n", gotwi.StringValue(res.Data.ID))
	
	return *res.Data.ID
}

func (tc *TwitterClient) LookupRecentTweets(username string, maxResults int) {
	
	p := &types.SearchTweetsRecentParams{
		Query: "from:elonmusk",
		TweetFields: fields.TweetFieldList{
			fields.TweetFieldText,
		},
		MaxResults: 1,
	}

	res, err := tweets.SearchTweetsRecent(context.Background(), tc.Client, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, t := range res.Data {
		fmt.Printf("Found tweet: [%s]\n", gotwi.StringValue(t.Text))
	}


}