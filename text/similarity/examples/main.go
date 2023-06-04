package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"similarity/internal"

	"github.com/antzucaro/matchr"
	"github.com/xrash/smetrics"
)

// Define test cases
var testCases = []struct {
	str1, str2 string
}{
	{"", "abcdefg"},
	{"abcdefg", ""},
	{"abcdefg", "abcdefg"},
	{"abcdefg", "bcdefgh"},
	{"D N H Enterprises Inc", "D & H Enterprises, Inc."},
	{"abcdefg", "gfedcba"},
	{"abcdefg", "ghijklm"},
	{"abcdefg", "0123456"},
	{"fly", "ant"},
	{"foo", "foo  "},
	{"foo", "foo "},
	{"hippo", "zzzzzzzz"},
	{"abcdefg", "ABCDEFG"},
	{"ABC Corporation", "ABC Corp"},
	{"My Gym Children's Fitness Center", "My Gym. Childrens Fitness"},
	{"the quick brown fox jumps over the lazy dog", "the quick brown cat jumps over the lazy dog"},
	{"abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyza"},
	{"", "cars"},
	{"cars", ""},
	{"car", "cars"},
	{"dixon", "dicksonx"},
	{"martha", "marhta"},
	{"dwayne", "duane"},
	{"martüa", "marüta"},
	{"dr", "driveway"},
}

func main() {
	runJaroWinklerSimilarityBenchmark()
}

func runJaroWinklerSimilarityBenchmark() {
	// Seed random number generator for generating test cases
	rand.Seed(time.Now().UnixNano())

	// Generate additional random test cases
	for i := 0; i < 10; i++ {
		str1 := randomString(20)
		str2 := randomString(20)
		testCases = append(testCases, struct{ str1, str2 string }{str1, str2})
	}

	fmt.Println("Running JaroWinkler Similarity Benchmark:")

	for _, testCase := range testCases {
		str1 := testCase.str1
		str2 := testCase.str2

		MatchrJaroWinkler, SmetricsJaroWinkler, InternalSimilarity := 0.0, 0.0, 0.0

		matchrResult := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MatchrJaroWinkler = matchr.JaroWinkler(str1, str2, true)
			}
		})

		smetricsResult := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SmetricsJaroWinkler = smetrics.JaroWinkler(str1, str2, 0.7, 4)
			}
		})

		internalSimilarity := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				InternalSimilarity = internal.Similarity(str1, str2)
			}
		})

		fmt.Printf("String 1: %s\n", str1)
		fmt.Printf("String 2: %s\n", str2)
		fmt.Printf("Matchr JaroWinkler Similarity:   \t benchmark =  %s \t | accuracy = %f \n", matchrResult.String(), MatchrJaroWinkler)
		fmt.Printf("Smetrics JaroWinkler Similarity: \t benchmark =  %s \t | accuracy = %f \n", smetricsResult.String(), SmetricsJaroWinkler)
		fmt.Printf("Internal JaroWinkler Similarity: \t benchmark =  %s \t | accuracy = %f \n", internalSimilarity.String(), InternalSimilarity)
		fmt.Println("--------------------------------------------------------------------------------------------------------------")
	}
}

// Helper function to generate random string
func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
