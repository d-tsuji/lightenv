package lightenv_test

import (
	"fmt"
	"os"

	"github.com/d-tsuji/lightenv"
)

func Example() {
	os.Setenv("APP_URL", "http://localhost")
	os.Setenv("PORT", "8888")
	os.Setenv("CONCURRENCY_NUM", "100")

	type Sample struct {
		Url            string `name:"APP_URL" required:"true"`
		PORT           string `required:"true"`
		ConcurrencyNum int    `name:"CONCURRENCY_NUM" required:"true"`
	}

	var s Sample
	if err := lightenv.Process(&s); err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Printf("%+v", s)

	// Output: {Url:http://localhost PORT:8888 ConcurrencyNum:100}
}

func ExampleProcess() {

	type Sample struct {
		Env string `name:"APP_ENV" required:"true"`
	}

	var s Sample
	if err := lightenv.Process(&s); err != nil {
		fmt.Printf("%+v", err)
	}

	// Output: required key APP_ENV missing value
}
