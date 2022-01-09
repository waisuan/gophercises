package main

import "math/rand"

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

const sequenceLength = 6

type SequenceGenerator struct{}

type Generator interface {
	GenerateSequence() string
}

func NewSequenceGenerator() Generator {
	return &SequenceGenerator{}
}

func (s *SequenceGenerator) GenerateSequence() string {
	b := make([]rune, sequenceLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
