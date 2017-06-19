# -*- coding: utf-8 -*-

import pandas as pd
import tweepy

def auth():
    id = pd.read_csv('id.csv')
    auth = tweepy.OAuthHandler(id['consumer_key'][0], id['consumer_secret'][0])
    auth.set_access_token(id['access_token'][0], id['access_token_secret'][0])
    return tweepy.API(auth)

def get_tl(tw):
    for i, tweet in enumerate(tw.home_timeline()):
        print("{:02d}: {}".format(i, tweet.text))

def get_user(tw, name):
    usr = tw.get_user(name)
    print(usr.screen_name)
    for f in usr.friends():
        print(f.screen_name)


if __name__ == '__main__':
    tw = auth()
    get_user(tw, "hoge")
    get_tl(tw)
