package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log := &logger{logrus.New()}
		log.Panic("missing required environment variable " + name)
	}
	return v
}

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)
	v := url.Values{}
	stream := api.UserStream(v)
	log.Print("UserStream Started Successfully")

	for t := range stream.C {
		switch v := t.(type) {
		case anaconda.EventFollow:
			log.Print("🎉🎉🎉Hurray! You have a new follower🎉🎉🎉")
			name := v.Source.Name
			sn := v.Source.ScreenName
			log.Print("Sending Auto Direct Message...", name)
			_, err := api.PostDMToScreenName("Hi "+name+",\nThank you so much for following me. Have a nice day!😊\n- Umesh(It's an auto generated msg delivered by My DM Bot.)", sn)
			if err != nil {
				log.Errorf("Failed to send dm due to %s", err)
				continue
			}
			log.Info("Send a msg to %s", name)
		}

	}

}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
func (log *logger) Fatal(args ...interface{})                    { log.Fatal(args...) }
func (log *logger) Fatalf(format string, args ...interface{})    { log.Fatalf(format, args...) }
func (log *logger) Panic(args ...interface{})                    { log.Panic(args...) }
func (log *logger) Panicf(format string, args ...interface{})    { log.Panicf(format, args...) }
