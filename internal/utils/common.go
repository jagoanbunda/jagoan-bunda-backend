// Package utils is mostly for helpers, like creating and assessing jwt tokens, converting string to pointer, hashing password, file upload,  etc
package utils

func StringToPtr(s string) *string {
	if s != "" {
		return &s
	}
	return nil
}
