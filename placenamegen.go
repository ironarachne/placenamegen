package placenamegen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/utility"
)

func randomConsonant() string {
	consonants := []string{
		"b",
		"c",
		"d",
		"f",
		"g",
		"h",
		"j",
		"k",
		"l",
		"m",
		"n",
		"p",
		"r",
		"s",
		"t",
		"w",
		"y",
		"z",
	}

	return utility.RandomItem(consonants)
}

func randomVowel() string {
	vowels := map[string]int{
		"a": 3,
		"e": 4,
		"i": 2,
		"o": 1,
		"u": 1,
	}

	return utility.RandomItemFromThresholdMap(vowels)
}

func randomSuffix() string {
	suffixes := []string{
		"a",
		"an",
		"en",
		"ia",
		"is",
	}

	return utility.RandomItem(suffixes)
}

// Generate generates a random place name
func Generate() string {
	c1 := randomConsonant()
	c2 := randomConsonant()
	c3 := randomConsonant()

	cons := []string{c1, c2, c3}

	v1 := randomVowel()
	v2 := randomVowel()

	vowels := []string{v1, v2}

	s := randomSuffix()

	name := ""

	nameLength := rand.Intn(4) + 3

	for i := 0; i < nameLength; i++ {
		if i%2 == 0 {
			name += utility.RandomItem(cons)
		} else {
			name += utility.RandomItem(vowels)
		}
	}

	name += s

	name = strings.Title(name)

	return name
}
