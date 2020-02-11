package main

import (
	"fmt"
	"log"

	"github.com/d-tsuji/lightenv"
)

type Sample struct {
	Url            string `name:"APP_URL" required:"true"`
	PORT           string `required:"true"`
	ConcurrencyNum int    `name:"CONCURRENCY_NUM" required:"true"`
}

func main() {
	// It is assumed that the following environment variables have been set in advance.
	//
	// export APP_URL="http://example.com"
	// export PORT=8888
	// export CONCURRENCY_NUM=100

	var s Sample
	if err := lightenv.Process(&s); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", s)

	// You will see the following output:
	//
	// {Url:http://example.com PORT:8888 ConcurrencyNum:100}
}
