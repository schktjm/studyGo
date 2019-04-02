// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	b := 0
	if _hasSign(s) {
		b = 1
		buf.WriteByte('-')
	}

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		fmt.Fprintf(&buf, "%s", _intComma(s[b:dot]))
		buf.WriteByte('.')
		fmt.Fprintf(&buf, "%s", _decimalComma(s[dot+1:]))
	} else {
		fmt.Fprintf(&buf, "%s", _intComma(s[b:]))
	}
	return buf.String()
}

func _hasSign(s string) bool {
	return s != "" && s[0:1] == "-"
}

func _intComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return _intComma(s[:n-3]) + "," + s[n-3:]
}

func _decimalComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:3] + "," + _decimalComma(s[3:])
}

//!-
