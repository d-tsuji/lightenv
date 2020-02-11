# lightenv [![Actions Status](https://github.com/d-tsuji/lightenv/workflows/CI/badge.svg)](https://github.com/d-tsuji/lightenv/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/d-tsuji/lightenv)](https://goreportcard.com/report/github.com/d-tsuji/lightenv) ![License MIT](https://img.shields.io/badge/license-MIT-blue.svg) [![GoDoc](https://godoc.org/github.com/d-tsuji/lightenv?status.svg)](https://godoc.org/github.com/d-tsuji/lightenv)

The lightenv is a lightweight library that handles environment variables. The thin wrapper using the standard [os](https://golang.org/pkg/os) library. This library is inspired by [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig).

The purpose is to allow the user to handle environment variables properly by returning to the caller as an error when the environment variable is not set. The feature is that: 

- You can set whether it is required as a struct tag like `required:"true"`.
- You can set a default value when no environment variables are set.

## USAGE

It is assumed that the following environment variables have been set in advance.

```
export APP_URL="http://example.com"
export PORT=8888
export CONCURRENCY_NUM=100
```

```go
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
	var s Sample
	if err := lightenv.Process(&s); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", s)
}
```

You will see the following output:

```bash
{Url:http://example.com PORT:8888 ConcurrencyNum:100}
```

## Struct Tag Support

It supports the tags shown in the example above.

| #   | Tag                    | Detail                                                                                                                                                                                      | Default value            |
| --- | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------ |
| 1   | `name:"ENV_PARAM_NAME` | Get the parameter set as ENV_PARAM_NAME from environment variables. <br> If nothing is set, the structure field name UPPER_CASE is used.                                                    | UPPER_CASE of field name |
| 2   | `required:"true"`      | Requires that the environment variable is set. Requires that the environment variable is set. <br>  If this parameter is true and the environment variable is not set, an error will occur. | "false"                  |
| 3   | `default:"8888"`       | Sets the default value if the environment variable has not been set.                                                                                                                        | -                        |

## Struct Type Support

- string
- int, int64, int32, int16, int8
- float64, float32

## Author

Tsuji Daishiro

## LICENSE

This software is licensed under the MIT license, see [LICENSE](https://github.com/d-tsuji/lightenv/blob/master/LICENSE) for more information.
