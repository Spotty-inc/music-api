package main

import (
	"math/rand"
	"strings"
	"time"
)

func randomString(n int, alphabet []rune) string {

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
	  ch := alphabet[rand.Intn(alphabetSize)]
	  sb.WriteRune(ch)
	}

	s := sb.String()
	return s
  }

func RandomLetter()string{
	rand.Seed(time.Now().UnixNano())

  	var alphabet []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

  	rs := randomString(1, alphabet)
	return rs
}
