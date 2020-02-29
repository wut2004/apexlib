package strutil

func UnquoteBytes(data []byte) string {
	// If the value is quoted, strip the quotes
	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	return string(data)
}
