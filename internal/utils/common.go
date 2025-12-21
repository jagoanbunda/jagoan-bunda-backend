package utils

func StringToPtr(s string) *string {
	if s != "" {
		return &s
	}
	return nil
}
