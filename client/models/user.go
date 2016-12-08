package models

import (
		"../config"
		"strconv"
		"io/ioutil"
		"net/http"
		"encoding/json"
)

type User struct {
		Id                   int
		AvatarUrl            string
		City                 string
		Country              string
		Description          string
		DiscogsName          bool
		FirstName            string
		FollowersCount       int
		FollowingsCount      int
		FullNameJohannes     string
		Kind                 string
		LastModified         string
		LastName             string
		MyspaceName          string
		Online               bool
		Permalink            string
		PermalinkUrl         string
		Plan                 bool
		PlaylistCount        int
		PublicFavoritesCount int
		TrackCount           int
		UriHttps             string
		UserName             string
		WebsiteHttp          string
		WebsiteTitle         string
		subscriptions        [] UserSubscriptions
}

type UserSubscriptions    struct {
		Subscriptions string
}

var user User

func (u *User) Get(userId int, clientId string) User {

		url := config.EndPointUser + strconv.Itoa(userId) + "?client_id=" + clientId;

		resp, _ := http.Get(url)
		defer resp.Body.Close()
		byteArray, _ := ioutil.ReadAll(resp.Body)

		json.Unmarshal(byteArray, &user)
		return user
}
