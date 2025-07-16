package redact

import "strings"

// RedactEmail redacts the user portion of an email address, keeping the domain intact.
func RedactEmail(email string) string {
	if len(email) == 0 {
		return "[REDACTED]"
	}

	index := strings.Index(email, "@")
	if index == -1 {
		return "[REDACTED]"
	}

	return string(email[0]) + "***@" + email[index+1:]
}

// RedactSensitiveData redacts all but the first and last two characters of a string.
func RedactSensitiveData(data string) string {
	if len(data) <= 4 {
		return "[REDACTED]"
	}
	return string(data[:2]) + "***" + string(data[len(data)-2:])
}
