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

func TestCheckEnglishOnly(t *testing.T) {
	tests := []Test{
		{
			input:   "starting server 8080",
			wantMsg: "",
			wantOk:  true,
		},
		{
			input:   "запуск сервера",
			wantMsg: "message must contain only English letters, digits and spaces",
			wantOk:  false,
		},
	}

	for _, test := range tests {
		msg, ok := checkEnglishOnly(test.input)
		if msg != test.wantMsg || ok != test.wantOk {
			t.Errorf("checkEnglishOnly(%q) = (%q, %v), want(%q, %v)", test.input, msg, ok, test.wantMsg, test.wantOk)
		}
	}
}

func TestCheckNoEmojiOrSpecial(t *testing.T) {
	tests := []Test{
		{
			input:   "blinding lights",
			wantMsg: "",
			wantOk:  true,
		},
		{
			input:   "⚡️blinding lights⚡️",
			wantMsg: "message must not contain emoji or special symbols",
			wantOk:  false,
		},
	}

	for _, test := range tests {
		msg, ok := checkNoEmojiOrSpecial(test.input)
		if msg != test.wantMsg || ok != test.wantOk {
			t.Errorf("checkNoEmojiOrSpecial(%q) = (%q, %v), want(%q, %v)", test.input, msg, ok, test.wantMsg, test.wantOk)
		}
	}
}

func TestCheckNoSensitive(t *testing.T) {
	tests := []Test{
		{
			input:   "user authenticated successfully",
			wantMsg: "",
			wantOk:  true,
		},
		{
			input:   "user password: 123",
			wantMsg: "message must not contain sensitive data",
			wantOk:  false,
		},
		{
			input:   "token: aoaoaoao",
			wantMsg: "message must not contain sensitive data",
			wantOk:  false,
		},
		{
			input:   "apiKey is abcdef",
			wantMsg: "message must not contain sensitive data",
			wantOk:  false,
		},
	}

	for _, test := range tests {
		msg, ok := checkNoSensitive(test.input)
		if msg != test.wantMsg || ok != test.wantOk {
			t.Errorf("checkNoSensitive(%q) = (%q, %v), want(%q, %v)", test.input, msg, ok, test.wantMsg, test.wantOk)
		}
	}
}
