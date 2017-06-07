package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ23456789" //remove characters that look alike i.e l,1,0,O
	letterIdxBits = 6                                                            // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1                                         // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits                                           // # of letter indices fitting in 63 bits
)

// RandString - generates a random string using bytes and masks
func RandString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func main() {

	flagStrLen := flag.Int("length", 12, "Length of string to generate.")
	flagSeparator := flag.String("separator", "-", "Separator to use between characters.")
	flagSeparateAt := flag.Int("separate-at", 3, "Insert separator every X characters.")
	flag.Parse()
	strLen := *flagStrLen
	separateAt := *flagSeparateAt
	separator := *flagSeparator

	// generate a string from the given length
	randString := RandString(strLen)

	// insert separators at the given break points
	for i := separateAt; i < len(randString); i += separateAt + 1 {
		randString = fmt.Sprintf("%s%s%s", randString[:i], separator, randString[i:])
	}

	fmt.Println(randString)

}
