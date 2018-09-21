package verify

import (
	"github.com/palantir/stacktrace"
)

type verification struct {
	condition    bool
	errorMessage string
}

type verifFunc func() verification

func That(condition bool, errorMessage string) verifFunc {
	return func() verification {
		return verification{condition, errorMessage}
	}
}

func All(verifFuncs ...verifFunc) error {
	for _, verifFunc := range verifFuncs {
		v := verifFunc()
		if !v.condition {
			return stacktrace.NewError(v.errorMessage)
		}
	}
	return nil
}

func StringValueOf(in *string) (out string) {
	if in == nil {
		return
	}
	return *in
}

func IntValueOf(in *int) (out int) {
	if in == nil {
		return
	}
	return *in
}

func I64ValueOf(in *int64) (out int64) {
	if in == nil {
		return
	}
	return *in
}
