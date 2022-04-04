package twitter

import (
	"context"
	"fmt"
	"github.com/michimani/gotwi"
	// "github.com/michimani/gotwi/tweets"
	// "github.com/michimani/gotwi/tweets/types"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/users"
	"github.com/michimani/gotwi/users/types"
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

	p := &types.UserLookupByUsernameParams{
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

// func main() {
// 	in := &gotwi.NewGotwiClientInput{
// 		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
// 		OAuthToken:           "your-twitter-acount-oauth-token",
// 		OAuthTokenSecret:     "your-twitter-acount-oauth-token-secret",
// 	}

// 	c, err := gotwi.NewGotwiClient(in)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	p := &types.ManageTweetsPostParams{
// 		Text: gotwi.String("This is a test tweet with poll."),
// 		Poll: &types.ManageTweetsPostParamsPoll{
// 			DurationMinutes: gotwi.Int(5),
// 			Options: []string{
// 				"Cyan",
// 				"Magenta",
// 				"Yellow",
// 				"Key plate",
// 			},
// 		},
// 	}

// 	res, err := tweets.ManageTweetsPost(context.Background(), c, p)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
// }