package asciiart

// handles instances of "\t" in the input
func Tab(s string) string {
	for i, char := range s {
		if (char == '\\' && i == 0) && s[i+1] == 't' {
			s = "        " + s[i+2:]
		}
		if char == '\\' && s[i+1] == 't' && i != 0 {
			s = s[:i] + "        " + s[i+2:]
		}
	}
	return s
}

// handles instances of "\b" in the input
func Backspace(s string) string {
	for i, char := range s {
		if char == '\\' && s[i+1] == 'b' && i == 0 {
			s = s[i+2:]
		}
		if char == '\\' && s[i+1] == 'b' && i != 0 {
			s = s[:i-1] + s[i+2:]
		}
	}
	return s
}
