package golang

var openCloseBracketMap = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

var closeOpenBracketMap = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

// isValidParentheses given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
// https://leetcode.com/problems/valid-parentheses
func isValidParentheses(s string) bool {
	if len(s) < 2 {
		return false
	}
	ss := []rune(s)
	var openStack []rune
	wantedCloseBrackets := '_'
	for _, v := range ss {
		closeBrackets, isOpen := openCloseBracketMap[v]
		if isOpen {
			wantedCloseBrackets = closeBrackets
			openStack = append(openStack, v)
			continue
		}
		if v == wantedCloseBrackets {
			openStack = openStack[:len(openStack)-1]
			wantedCloseBrackets = '_'
			// get another wanted close bracket
			if len(openStack) == 0 {
				continue
			}
			wantedCloseBrackets = openCloseBracketMap[openStack[len(openStack)-1]]
			continue
		} else {
			if _, isClose := closeOpenBracketMap[v]; isClose {
				return false // found unexpected close bracket
			}
		}
	}

	if len(openStack) == 0 {
		return true
	}
	return false
}
