package soundcloud

import (
		//"fmt"
		"net/http"
		"encoding/json"
)

type Request struct {
		Id int
}

var (
		client  *http.Client = &http.Client{
				//CheckRedirect: noRedirects,
				//CheckRedirect: redirectPolicyFunc,
		}
)

func (r *Request)Get(e *EndPoint, model interface{}) error {

		url, err := http.NewRequest(e.Verb, e.Url, nil);

		if err != nil {
				return err
		}

		resp, err := client.Do(url)
		if err != nil {
				return err
		}

		jsonErr := json.NewDecoder(resp.Body).Decode(model)
		
		//fmt.Println(err);
		//fmt.Println(e.Url);
		//fmt.Println(resp.Body);
		//fmt.Println(resp);
		//fmt.Println(model);

		return jsonErr
}


