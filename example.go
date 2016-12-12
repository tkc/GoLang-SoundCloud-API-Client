package main

import (
		"./soundcloud"
		"fmt"
)

func main() {

		const clientId = "***********";

		user := soundcloud.User{};
		res, err := user.GetUser(3207, clientId);

		if err != nil {
				fmt.Println(err);
		}

		fmt.Println(res.Id);
		fmt.Println(res.UserName);
		fmt.Println(res.City);
}

