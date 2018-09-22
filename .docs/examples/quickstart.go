package main

import (
	"log"

	verify "github.com/nmrshll/go-verify"
)

type Argument string

func doSomething(arg1 Argument, arg2 *Argument) {}

func FunctionToCall(arg1 Argument, arg2 *Argument) error {
	verify.That(string(arg1) != "", "arg 1 can't be empty")
	verify.That(arg2 != nil, "arg2 can't be nil")
	if err := verify.Error(); err != nil {
		return err
	}
	verify.That(*arg2 == Argument("hello world"), "arg2 must be \"hello world\"")
	if err := verify.Error(); err != nil {
		return err
	}

	// perform your function here
	doSomething(arg1, arg2)

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
