package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Be carefull with the error: %v", v)
		}
	}()

	f, err := os.Create("new_file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = fmt.Errorf("An error occurred at: %v", time.Now())
	if err != nil {
		panic(err)
	}

}
