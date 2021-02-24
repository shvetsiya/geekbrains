package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.Create("new_file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Be carefull with the error: %v that happend at time: %v", v, time.Now())
		}
	}()
	var a int
	fmt.Println(1 / a)

}
