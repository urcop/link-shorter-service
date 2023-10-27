package helpers

import "math/rand"

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomUrl(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = runes[rand.Intn(length)]
	}
	return string(result)
}
