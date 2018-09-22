package verify

import (
	"fmt"
	"runtime"
	"time"
)

var verifiers = make(map[uint]*verifier)

type verifier struct {
	err error
}

func newVerifier() *verifier            { return &verifier{} }
func (v *verifier) Error() (rErr error) { return v.err }

func verifierForWrappingPC(pc uintptr) (_ *verifier, upFunctionLocationUint uint) {
	upFunctionLocationUint = uint(runtime.FuncForPC(pc).Entry())
	if _, ok := verifiers[upFunctionLocationUint]; !ok {
		verifiers[upFunctionLocationUint] = newVerifier()
	}
	return verifiers[upFunctionLocationUint], upFunctionLocationUint
}

func That(condition bool, errorFormat string, args ...interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic(fmt.Errorf("failed evaluating runtime.caller"))
	}
	verifier, _ := verifierForWrappingPC(pc)

	if !condition {
		verifier.err = fmt.Errorf(errorFormat, args...)
	}
}

func Error() error {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic(fmt.Errorf("failed evaluating runtime.caller"))
	}
	verifier, upFunctionLocationUint := verifierForWrappingPC(pc)

	// cleanup in map of verifiers
	go func() {
		time.Sleep(1 * time.Second)
		delete(verifiers, upFunctionLocationUint)
	}()

	return verifier.Error()
}
