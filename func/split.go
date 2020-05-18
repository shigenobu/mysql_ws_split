/*
 split
 */
package _func

import (
	"bytes"
	"encoding/json"
	"regexp"
)

func Split(pattern string, subject string) string {
	if pattern == "" {
		return "[]"
	}
	if subject == "" {
		return "[]"
	}

	result := regexp.MustCompile(pattern).Split(subject, -1)

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(result)
	if err != nil {
		return "[]"
	}
	return string(buffer.Bytes())
}