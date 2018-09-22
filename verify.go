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

// That performs a verification
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

// Error returns the verifier error and triggers cleanup of the verifier.
// The cleanup happens after enough time to return the error (5 seconds)
func Error() error {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic(fmt.Errorf("failed evaluating runtime.caller"))
	}
	verifier, upFunctionLocationUint := verifierForWrappingPC(pc)

	// cleanup in map of verifiers
	go func() {
		time.Sleep(10 * time.Second)
		delete(verifiers, upFunctionLocationUint)
	}()

	return verifier.Error()
}
