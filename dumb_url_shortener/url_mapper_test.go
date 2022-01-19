package main

import (
	"math/rand"
	"testing"
)

func TestUrlMapper_GenerateShortUrl(t *testing.T) {
	t.Run("generates a custom URL string", func(t *testing.T) {
		u := NewUrlMapper(NewSimpleMockGenerator(), maxCapacity)
		got, _ := u.GenerateShortUrlToken("https://google.com")
		want := "abcdef"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("stores the generated custom URL string", func(t *testing.T) {
		u := NewUrlMapper(NewSimpleMockGenerator(), maxCapacity)
		shortUrl, _ := u.GenerateShortUrlToken("https://google.com")

		got := u.GetUrlByShortUrl(shortUrl)
		want := "https://google.com"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("always generate a unique custom URL string", func(t *testing.T) {
		u := NewUrlMapper(NewComplexMockGenerator(), maxCapacity)
		got1, _ := u.GenerateShortUrlToken("https://google.com")
		got2, _ := u.GenerateShortUrlToken("https://google.com")

		if got1 == got2 {
			t.Errorf("got duplicates %s", got1)
		}
	})

	t.Run("does not generate if already at max capacity", func(t *testing.T) {
		u := NewUrlMapper(NewComplexMockGenerator(), 1)
		_, err := u.GenerateShortUrlToken("https://google.com")
		if err != nil {
			t.Errorf("Did not expect an error")
		}

		_, err = u.GenerateShortUrlToken("https://google.com")
		if err == nil {
			t.Errorf("Expected an error")
		}
	})
}

type SimpleMockGenerator struct{}

func NewSimpleMockGenerator() Generator {
	return &SimpleMockGenerator{}
}

func (m *SimpleMockGenerator) GenerateSequence() string {
	return "abcdef"
}

type ComplexMockGenerator struct {
	sequences []string
}

func NewComplexMockGenerator() Generator {
	return &ComplexMockGenerator{
		sequences: []string{"abcdef", "aaaaaa"},
	}
}

func (cm *ComplexMockGenerator) GenerateSequence() string {
	return cm.sequences[rand.Intn(len(cm.sequences))]
}
