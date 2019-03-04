package placenamegen

import (
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
		"i": 1,
		"o": 1,
		"u": 1,
	}

	return utility.RandomItemFromThresholdMap(vowels)
}

// Generate generates a random place name
func Generate() string {
	c1 := randomConsonant()
	c2 := randomConsonant()
	v1 := randomVowel()
	v2 := randomVowel()

	name := strings.Title(c1 + v1 + c2 + c1 + v2 + c2)

	return name
}
