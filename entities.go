package entities

import "fmt"

type user struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func (user User) ToString () string {
	return fmt.Fprintf("id: %s\nName: %\Password: %s\n",user.Id, user.Name, user.Password)
}

