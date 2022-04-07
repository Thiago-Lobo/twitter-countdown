# Twitter Countdown Automator

## Introduction

Have you ever wanted a tool that reminds your followers about when your a\*\*hole president will go down?\* Or maybe all you want is for your followers to know how many days are left until the next episode of your favorite saga comes out? **Good news**: I made a tool which will satisfy your needs!

*\* Note that the president-related use case will only work if you are based on a democratic country - enjoy it while you can.*

## Requirements

- Something running docker with docker-compose available (although you could build the binary manually and run it yourself)
- A Twitter account with developer access ([here's some information on how to get that](https://developer.twitter.com/en/support/twitter-api/developer-account))
- A project + app created within your Twitter account developer dashboard
  - You will need the app's **API key**, **API key secret**, **Access Token** and **Access Token Secret**
## Setup

1. Clone the repo
2. Edit the `docker-compose.yml` file:

        version: "3.9"
        services:
        twitter-countdown-thaigo:
            restart: always
            build: .
            environment:
            GOTWI_API_KEY: "<API_KEY_HERE>"
            GOTWI_API_KEY_SECRET: "<API_KEY_SECRET_HERE>"
            GOTWI_OAUTH_TOKEN: "<ACCESS_TOKEN_HERE>"
            GOTWI_OAUTH_TOKEN_SECRET: "<ACCESS_TOKEN_SECRET_HERE>"
            TWEET_TEMPLATE: "👉 The tweet text template. It needs a %d which will be replaced by the countdown."
            POST_TIME: "HH:MM:SS" # this is the timestamp after which the daily tweet will be posted

    Note that the tweet will be posted as soon as possible with respect to the `POST_TIME` - if you run the program after that time, then it will post right away, otherwise it will just wait for that time on a daily basis and post the tweet. 

3. Run the container with something like:

        docker-compose up -d --build

