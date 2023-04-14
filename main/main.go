package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type request struct {
	Input string `json:"input"`
}

type response struct {
	Output string `json:"output"`
}

func HandleFunc(input request) (response, error) {
	return response{Output: fmt.Sprintf("%s recieved", input.Input)}, nil
}

func main() {
	lambda.Start(HandleFunc)
}
