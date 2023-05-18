package main

import (
	"context"
	"fmt"
	"go-db-cache/pkg/db/model"
)

func main() {
	optioon, err := NewOptioon()
	if err != nil {
		panic(err)
	}
	err = optioon.Complete()
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	user, err := optioon.Factory.User().Create(ctx, &model.User{
		Name:   "chenze",
		Age:    0,
		Gender: "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
