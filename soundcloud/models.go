package soundcloud

import ()

type User struct {
		Id                   int
		AvatarUrl            string
		City                 string
		Country              string
		Description          string
		DiscogsName          string
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
		Plan                 string
		PlaylistCount        int
		PublicFavoritesCount int
		TrackCount           int
		UriHttps             string
		UserName             string
		WebsiteHttp          string
		WebsiteTitle         string
		//subscriptions        [] UserSubscriptions TODO
}

type UserSubscriptions    struct {
		Subscriptions string
}

var user User
var users []User
var endPoint = EndPoint{}
var request = Request{}

func (u *User) GetUser(userId int, clientId string) (User, error) {
		e := endPoint.User(userId, clientId);
		err := request.Get(e, &user);
		return user, err
}

