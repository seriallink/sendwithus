# SendWithUs SDK in Go 

This is a simple package to interface with [SendWithUs](https://sendwithus.com) using Golang.

## Installation

```bash
$ go get github.com/seriallink/sendwithus
```

## Example

This is a brief example on how to send an email. You can find more examples by looking at [the test cases](https://github.com/seriallink/sendwithus/blob/master/swu_test.go).

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/seriallink/sendwithus"
)

func main() {
	api := swu.New(os.Getenv("SWU_API_KEY"))
	email := &swu.Email{
		Id: os.Getenv("SWU_TEMPLATE_ID"),
		Sender: &swu.Sender{
			Recipient: &swu.Recipient{
				Name:    "NoReply",
				Address: os.Getenv("SWU_SENDER_EMAIL"),
			},
		},
		Recipient: &swu.Recipient{
			Name:    "John Doe",
			Address: os.Getenv("SWU_CUSTOMER_EMAIL"),
		},
		EmailData: map[string]string{
			"first_name": "John",
			"last_name":  "Doe",
		},
	}
	log, err := api.Send(email)
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("log id: %s", log.Id)
}
```

## [API Documentation](https://support.sendwithus.com/api/)

## MIT License

Enjoy! Feel free to send pull requests or submit issues :)
