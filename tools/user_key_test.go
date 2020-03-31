package tools

import (
	"fmt"
	"testing"
)

func TestUserKey(t *testing.T) {
	fmt.Println(UserKey("123456"))
}

func TestCreateUserKey(t *testing.T) {
	fmt.Println(CreateUserKey("admin", "123456"))
}
