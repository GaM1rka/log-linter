package analyzer

import "testing"

type Test struct {
	input   string
	wantMsg string
	wantOk  bool
}

func TestCheckStartsWithLowercase(t *testing.T) {
	tests := []Test{
		{
			input:   "Server started",
			wantMsg: "message must start with lowercase letter",
			wantOk:  false,
		},
		{
			input:   "worker is running",
			wantMsg: "",
			wantOk:  true,
		},
	}

	for _, test := range tests {
		msg, ok := checkStartsWithLowercase(test.input)
		if msg != test.wantMsg || ok != test.wantOk {
			t.Errorf("checkStartsWithLowercase(%q) = (%q, %v), want(%q, %v)", test.input, msg, ok, test.wantMsg, test.wantOk)
		}
	}
}
