package main

import (
		"./soundcloud"
		"fmt"
)

var clientId = "";

func main() {
		user := soundcloud.User{};
		res, err := user.GetUser(3207, clientId);

		if err != nil {
				fmt.Println(err);
		}

		fmt.Println(res.Id);
		fmt.Println(res.UserName);
		fmt.Println(res.City);
}

