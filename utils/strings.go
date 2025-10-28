package utils

import (
	"math/rand"
	"strings"
)

// ShortUUID accepts a string and returns a substring before the first '-'
//
// Example: xxxx-yy-zzzzzzzz would return xxxxx
func ShortUUID(uuid string) string {
	parts := strings.Split(uuid, "-")
	if len(parts) > 0 {
		return strings.ToUpper(parts[0])
	}
	return strings.ToUpper(uuid)
}

func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// function for splitting strings into small forms, with specific guidelines
func IndexString(s string) []string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	if len(words) == 0 {
		return []string{}
	}

	resultMap := make(map[string]bool)
	var result []string

	for i, word := range words {
		if len(result) >= 35 {
			break
		}

		if !resultMap[word] {
			resultMap[word] = true
			result = append(result, word)
			if len(result) >= 35 {
				break
			}
		}

		if i != 0 && len(word) < 3 {
			continue
		}

		var limit int
		switch i {
		case 0:
			limit = min(len(word), 10)
		case 1:
			limit = min(len(word), 5)
		default:
			limit = min(len(word), 3)
		}

		for j := 1; j <= limit; j++ {
			prefix := word[:j]
			if !resultMap[prefix] {
				resultMap[prefix] = true
				result = append(result, prefix)
				if len(result) >= 35 {
					break
				}
			}
		}
	}

	if len(words) > 1 {
		full := strings.Join(words, " ")
		if !resultMap[full] {
			result = append(result, full)
		}
	}

	return result
}
