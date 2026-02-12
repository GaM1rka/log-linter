package analyzer

import (
	"strings"
)

func checkStartsWithLowercase(s string) (string, bool) {
	s = strings.TrimLeft(s, " \t")
	if len(s) == 0 {
		return "", true
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return "", true
	}
	if s[0] >= 'A' && s[0] <= 'Z' {
		return "message must start with lowercase letter", false
	}
	return "", true
}

func checkEnglishOnly(s string) (string, bool) {
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == ' ' {
			continue
		}
		return "message must contain only English letters, digits and spaces", false
	}
	return "", true
}

func checkNoEmojiOrSpecial(s string) (string, bool) {
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == ' ' {
			continue
		}
		return "message must not contain emoji or special symbols", false
	}
	return "", true
}

func checkNoSensitive(s string) (string, bool) {
	sensitiveWords := []string{"password", "secret", "token", "apikey", "key", "credentials"}
	words := strings.Fields(strings.ToLower(s))
	for _, word := range words {
		word = strings.Trim(word, ".,:;!?()[]{}\"'")
		if contains(sensitiveWords, word) {
			return "message must not contain sensitive data", false
		}
	}
	return "", true
}

func contains(slice []string, item string) bool {
	for _, sensWord := range slice {
		if sensWord == item {
			return true
		}
	}
	return false
}
