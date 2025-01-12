package problems

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && !isPrefix(prefix, strs[i]) {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func isPrefix(prefix, str string) bool {
	if len(prefix) > len(str) {
		return false
	}
	return str[:len(prefix)] == prefix
}
