# go-verify
Defensive programming utilities for Go

Spare yourself and your colleagues some debugging time, validate function parameters !

### Why ?

While it's usual practice to validate requests when they come from an untrusted party (e.g. a request from the front-end to the back-end), it's less common to validate input from inside the same program.

However, it takes minimal effort and can be really useful, and avoid bugs and long debugging times. If someone tries to use a function you wrote with incorrect parameters, they will be warned with a helpful error instead of wondering where the problem comes from and having to dig into your code.

Having learned recently that this concept is actually a thing and that it's called defensive programming, I set out to build a library to make it easier/less verbose. No excuses now !

### How ?

Install with `go get github.com/nmrshll/go-verify`

Then use this way:
[embedmd]:# (.docs/examples/quickstart.go /func FunctionToCall/ $)

```go
func FunctionToCall(arg1 Argument, arg2 *Argument) error {
	if err := verify.All(
		verify.That(string(arg1) != "", "arg 1 can't be empty"),
		verify.That(arg2 != nil, "arg2 can't be nil"),
		verify.That(*arg2 == Argument("hello world"), "arg2 must be \"hello world\""),
	); err != nil {
		return err
	}
}

func main() {
	arg1 := Argument("hola mundo")
	arg2 := Argument("hola mundo")
	err := FunctionToCall(arg1, &arg2)
	if err != nil {
		log.Fatal(err)
	}
}
```

#### Gotchas
- Always assert that a pointer is not nil before asserting anything else about its value

### License
[MIT](.docs/LICENSE)
