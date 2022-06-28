package constants

const GOTWI_API_KEY = "GOTWI_API_KEY"
const GOTWI_API_KEY_SECRET = "GOTWI_API_KEY_SECRET"

const OAUTH_TOKEN = "GOTWI_OAUTH_TOKEN"
const OAUTH_TOKEN_SECRET = "GOTWI_OAUTH_TOKEN_SECRET"

const END_DATE = "END_DATE"
const TWEET_TEMPLATE = "TWEET_TEMPLATE"
const POST_TIME = "POST_TIME"

const REPLY_SETTING = "REPLY_SETTING"

// Tweet reply settings

type TweetReplySetting int

const (
	NONE TweetReplySetting = iota
	FOLLOWING
	MENTIONED
)

var tweetReplySettingStrings = [...]string{"", "following", "mentionedUsers"}

func (setting TweetReplySetting) String() *string {
	if setting == NONE {
		return nil
	}

	return &tweetReplySettingStrings[setting]
}

var (
	tweetReplySettingMap = map[string]TweetReplySetting{
		"":               NONE,
		"following":      FOLLOWING,
		"mentionedUsers": MENTIONED,
	}
)

func StringToTweetReplySetting(str string) TweetReplySetting {
	setting := tweetReplySettingMap[str]
	return setting
}
