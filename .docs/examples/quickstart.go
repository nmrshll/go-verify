package main

import (
	"log"

	verify "github.com/nmrshll/go-verify"
)

type Argument string

func FunctionToCall(arg1 Argument, arg2 *Argument) error {
	if err := verify.All(
		verify.That(string(arg1) != "", "arg 1 can't be empty"),
		verify.That(arg2 != nil, "arg2 can't be nil"),
		verify.That(verify.ValueOf(arg2) == Argument("hello world"), "arg2 must be \"hello world\""),
	); err != nil {
		return err
	}

	// perform your function here

	return nil
}

func main() {
	arg1 := Argument("hola mundo")
	arg2 := Argument("hola mundo") // incorrect argument as "hello world is expected here"
	err := FunctionToCall(arg1, &arg2)
	if err != nil {
		log.Fatal(err)
	}
}
