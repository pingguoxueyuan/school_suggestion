package main

import (
	"github.com/pingguoxueyuan/school_suggestion/logic"
	"fmt"
)

func Init() (err error){

	err = logic.Init("./data/school.dat")
	if err != nil {
		return
	}

	fmt.Printf("Init succ\n")
	return
}