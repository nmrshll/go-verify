package verify

import (
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("pass by value parameters", func(t *testing.T) {
		functionToCall := func(stringArg string, intArg int) error {
			if err := All(
				That(stringArg != "", "stringArg can't be empty"),
				That(intArg != 0, "intArg can't be nil"),
			); err != nil {
				return err
			}
			return nil
		}
		type args struct {
			stringArg string
			intArg    int
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"correct parameters", args{"something not empty", 1}, false},
			{"invalid string (empty)", args{"", 1}, true},
			{"invalid int (zero-value)", args{"something not empty", 0}, true},
			{"both values invalid (zero-values)", args{"", 0}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := functionToCall(tt.args.stringArg, tt.args.intArg); (err != nil) != tt.wantErr {
					t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})

	t.Run("pointer parameters", func(t *testing.T) {
		functionToCall := func(stringPointerArg *string) error {
			if err := All(
				That(stringPointerArg != nil, "stringPointerArg can't be nil"),
				That(valueOfStringPointer(stringPointerArg) != "", "stringPointerArg can't be pointer to empty string"),
			); err != nil {
				return err
			}
			return nil
		}
		type args struct {
			stringPointerArg *string
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"correct parameter", args{pointerToString("something not empty")}, false},
			{"invalid pointer (nil pointer)", args{nil}, true},
			{"invalid string (pointer to empty string)", args{pointerToString("")}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := functionToCall(tt.args.stringPointerArg); (err != nil) != tt.wantErr {
					t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
}

func pointerToString(str string) *string { return &str }
func valueOfStringPointer(strp *string) string {
	if strp != nil {
		return *strp
	}
	return ""
}
