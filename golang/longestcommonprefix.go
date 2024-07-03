package golang

import "strings"

// longestCommonPrefix finds the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
// https://leetcode.com/problems/longest-common-prefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ss := strings.Split(strs[0], "") // compare characters of first word
	prefix := ""
	for i := 0; i < len(ss); i++ {
		prefix += ss[i]
		for j := 0; j < len(strs); j++ {
			if strings.HasPrefix(strs[j], prefix) {
				continue
			} else {
				return strings.TrimSuffix(prefix, ss[i])
			}
		}
	}
	return prefix
}
