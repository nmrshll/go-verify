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

// type verification struct {
// 	condition    bool
// 	errorMessage string
// }

// type verifFunc func() verification

// func That(condition bool, errorMessage string) verifFunc {
// 	return func() verification {
// 		return verification{condition, errorMessage}
// 	}
// }

func That(condition bool, errorFormat string, args ...interface{}) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		// return fmt.Errorf("failed evaluating runtime.caller")
		panic(fmt.Errorf("failed evaluating runtime.caller"))
	}
	// function := runtime.FuncForPC(pc)
	verifier, _ := verifierForWrappingPC(pc)
	// spew.Dump(verifier.Error())
	// // runtime.FuncForPC(pc)
	if !condition {
		// return fmt.Errorf(errorFormat, args...)
		verifier.err = fmt.Errorf(errorFormat, args...)
	}
	// return nil
}

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

// func All(verifFuncs ...verifFunc) error {
// 	for _, verifFunc := range verifFuncs {
// 		v := verifFunc()
// 		if !v.condition {
// 			return stacktrace.NewError(v.errorMessage)
// 		}
// 	}
// 	return nil
// }

// func StringValueOf(in *string) (out string) {
// 	if in == nil {
// 		return
// 	}
// 	return *in
// }

// func IntValueOf(in *int) (out int) {
// 	if in == nil {
// 		return
// 	}
// 	return *in
// }

// func I64ValueOf(in *int64) (out int64) {
// 	if in == nil {
// 		return
// 	}
// 	return *in
// }
