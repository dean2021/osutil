package main

import (
	"fmt"
	"github.com/dean2021/osutil/user"
)

func main() {
	users, err := user.List()
	if err != nil {
		return
	}
	for _, u := range users {
		fmt.Println(u.GetValue("Name"))
	}
}
