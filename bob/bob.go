// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) < 1 {
		return "Fine. Be that way!"
	}
	shout, question := false, false
	matched, err := regexp.Match("[A-Za-z]", []byte(remark))
	if err != nil {
		return ""
	}
	if remark == strings.ToUpper(remark) && matched {
		shout = true
	}
	switch remark[len(remark)-1] {
	// case '!':
	// 	shout = true
	// 	break
	case '?':
		question = true
	}

	if shout && question {
		return "Calm down, I know what I'm doing!"
	} else if shout {
		return "Whoa, chill out!"
	} else if question {
		return "Sure."
	}

	return "Whatever."
}
