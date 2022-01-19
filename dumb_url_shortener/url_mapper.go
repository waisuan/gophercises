package main

import "errors"

type UrlMapper struct {
	shortUrlTokenToUrl map[string]string
	generator          *Generator
	maxCapacity        int
}

func NewUrlMapper(generator Generator, maxCapacity int) *UrlMapper {
	return &UrlMapper{
		shortUrlTokenToUrl: make(map[string]string),
		generator:          &generator,
		maxCapacity:        maxCapacity,
	}
}

func (u *UrlMapper) GenerateShortUrlToken(originalUrl string) (string, error) {
	if u.isAtMaxCapacity() {
		return "", errors.New("not enough space to generate token")
	}

	var shortUrlToken string
	for {
		shortUrlToken = (*u.generator).GenerateSequence()

		if _, ok := u.shortUrlTokenToUrl[shortUrlToken]; !ok {
			u.shortUrlTokenToUrl[shortUrlToken] = originalUrl
			break
		}
	}

	return shortUrlToken, nil
}

func (u *UrlMapper) GetUrlByShortUrl(shortUrl string) string {
	return u.shortUrlTokenToUrl[shortUrl]
}

func (u *UrlMapper) isAtMaxCapacity() bool {
	return len(u.shortUrlTokenToUrl) == u.maxCapacity
}
