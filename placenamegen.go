package placenamegen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/random"
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

	return random.Item(consonants)
}

func randomConsonantGroup(size int) []string {
	var consonants = []string{}

	for i := 0; i < size; i++ {
		consonants = append(consonants, randomConsonant())
	}

	return consonants
}

func randomVowel() string {
	vowels := map[string]int{
		"a": 3,
		"e": 4,
		"i": 2,
		"o": 1,
		"u": 1,
	}

	return random.ItemFromThresholdMap(vowels)
}

func randomSyllable(consonants []string) string {
	c1 := random.Item(consonants)
	c2 := random.Item(consonants)

	v := randomVowel()

	syllable := c1 + v

	i := rand.Intn(10)
	if i < 3 {
		syllable += c2
	}

	return syllable
}

func randomSuffix() string {
	suffixes := map[string]int{
		"a":   6,
		"an":  2,
		"en":  1,
		"ia":  9,
		"ion": 1,
		"is":  2,
		"ya":  1,
	}

	return random.ItemFromThresholdMap(suffixes)
}

// Generate generates a random place name
func Generate() string {
	group1 := randomConsonantGroup(3)
	group2 := randomConsonantGroup(3)
	s1 := randomSyllable(group1)
	s2 := randomSyllable(group2)
	s3 := randomSyllable(group2)

	name := s1 + s2

	i := rand.Intn(10)
	if i < 3 {
		name += s3
	}

	name += randomSuffix()

	name = strings.Title(name)

	return name
}
