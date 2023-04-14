package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type request struct {
	input string `json:"input"`
}

type response struct {
	output string `json:"output"`
}

func HandleFunc(input request) (response, error) {
	var res response
	res.output = input.input + " recieved"
	return res, nil
}

func main() {
	lambda.Start(HandleFunc)
}
