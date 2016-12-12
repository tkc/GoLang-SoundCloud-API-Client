package soundcloud

import (
		"fmt"
		"strconv"
)

const BaseUrl = "http://api.soundcloud.com/"

type EndPoint struct {
		Verb    string
		Url     string
		AuthReq bool
}

func (self *EndPoint) Initialize() {
}

var VERB_GET = "GET";

var VERB_POST = "POST";

var VERB_PUT = "PUT";

var VERB_DELETE = "DELETE";

func (e *EndPoint)User(userId int, clientId string) *EndPoint {
		Url := BaseUrl + fmt.Sprintf("users/%s?client_id=%s", strconv.Itoa(userId), clientId)
		endPoint := &EndPoint{
				Verb:VERB_GET,
				Url:Url,
				AuthReq:false,
		}
		return endPoint
}

//func UserTracks(userId int, clientId string) *EndPoint {
//		Url := BaseUrl + fmt.Sprintf("users/%s/tracks?client_id=%s", strconv.Itoa(userId), clientId)
//		endPoint := &EndPoint{
//				Verb:VERB_GET,
//				Url:Url,
//				AuthReq:false,
//		}
//		return endPoint
//}

//func UserPlaylists(userId int, clientId string) *EndPoint {
//		Url := BaseUrl + fmt.Sprintf("users/%s/playlists?client_id=%s", strconv.Itoa(userId), clientId)
//		endPoint := &EndPoint{
//				Verb:VERB_GET,
//				Url:Url,
//				AuthReq:false,
//		}
//		return endPoint
//}
//
//func UserFollowings(userId int, clientId string, Verb string) EndPoint {
//		Url := BaseUrl + fmt.Sprintf("users/%s/followings?client_id=%s", strconv.Itoa(userId), clientId)
//		e.Verb = Verb
//		e.AuthReq = false
//		e.Url = Url
//		return *e
//}
//
//func (e *EndPoint) UserFollowers(userId int, clientId string) EndPoint {
//		Url := BaseUrl + fmt.Sprintf("users/%s/followers?client_id=%s", strconv.Itoa(userId), clientId)
//		e.Verb = VERB_GET
//		e.Url = Url
//		e.AuthReq = false
//		return *e
//}

//func (e *EndPoint) UserComments(userId int, clientId string) EndPoint {
//		Url := BaseUrl + fmt.Sprintf("users/%s/commnet?client_id=%s", strconv.Itoa(userId), clientId)
//		e.Verb = VERB_GET
//		e.Url = Url
//		e.AuthReq = false
//		return *e
//}
//
//func (e *EndPoint) UserFavorites(userId int, clientId string) EndPoint {
//		Url := BaseUrl + fmt.Sprintf("users/%s/favorites?client_id=%s", strconv.Itoa(userId), clientId)
//		e.Verb = VERB_GET
//		e.Url = Url
//		e.AuthReq = false
//		return *e
//}
//
//func (e *EndPoint) WebProfiles(userId int, clientId string, Verb string) EndPoint {
//		Url := BaseUrl + fmt.Sprintf("users/%s/web-profiles?client_id=%s", strconv.Itoa(userId), clientId)
//		e.Verb = Verb
//		e.Url = Url
//		e.AuthReq = false
//		return *e
//}
