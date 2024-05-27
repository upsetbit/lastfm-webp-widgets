//go:build exec_lambda

package main

import (
	// standard
	"fmt"

	// 3rd-party
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() error {
	var err error = nil

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			fmt.Println(err)
		}
	}()

	doRoutine()
	return err
}

func main() {
	lambda.Start(Handler)
}
