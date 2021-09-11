package equalcmp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_equalCmpMatcher_Match(t *testing.T) {
	type args struct {
		expected interface{}
		options  cmp.Options
		actual   interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
		wantErr   bool
	}{
		{
			name: "The actual value matches the expected value.",
			args: args{
				expected: "string",
				actual:   "string",
			},
			wantMatch: true,
			wantErr:   false,
		},
		{
			name: "The actual value doesn't match the expected value",
			args: args{
				expected: "string",
				actual:   "string2",
			},
			wantMatch: false,
			wantErr:   false,
		},
		{
			name: "If the actual value and the expected value are both nil, make an error",
			args: args{
				expected: nil,
				actual:   nil,
			},
			wantMatch: false,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := EqualCmp(tt.args.expected, tt.args.options)

			gotSuccess, err := matcher.Match(tt.args.actual)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSuccess != tt.wantMatch {
				t.Errorf("Match() gotSuccess = %v, want %v", gotSuccess, tt.wantMatch)
			}
		})
	}
}
