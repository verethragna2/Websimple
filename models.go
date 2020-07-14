package Models
import (
	"GolangRestApi/entities"
	"errors"
)
var (
	listUser = make([]*entities.User) bool {
func CreateUser(user *entities.User) bool {
	if user.Id != ""&& user.Name !="" && user.Password !=""{
		if userF,_:= FindUser (user.Id); userF == nil {
			listUser = append(listUser, user)
			return true
		}
	}
	return flase
}
func UpdateUser (user *entities.User) bool {
	for índex, user := range listUser {
		if user.Id == eUser.ID {
			listUser [index] = eUser
			return true
		}
	}
	return false
}

func FindUser (id string) (*entities.User,error) {
	for _, user:= range listUser {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, error.New("User does not exit")
}
func DeleteUser(id string) bool {
	for index, user := range listUser {
		if user.Id == id {
			copy(listUser[index:], listUser[index+1:])
			listUser[len(listUser)-1] = &entities.User{}
			listUser = listUser[:len(listUser)-1]
			return true
		}
	}
	return false
}

func GetAllUser() []*entities.User {
return listUser
}