package utils_test

import (
	"github.com/jessaleks/simple-todo-server-go/utils"
	"testing"
)

func TestDB(t *testing.T) {
	_, err := utils.ConnectToTheDatabase()
	if err != nil {
		t.Fail()
	}

}
