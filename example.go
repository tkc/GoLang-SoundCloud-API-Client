package main

import (
		"./client/models"
		"fmt"
)

var clientId = "<clientId>";

func main() {
		user := models.User{};
		res := user.Get(3207, clientId);
		fmt.Println(res.Id);
}
