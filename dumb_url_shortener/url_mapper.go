package main

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

func (u *UrlMapper) GenerateShortUrlToken(originalUrl string) (shortUrlToken string) {
	if u.isAtMaxCapacity() {
		return
	}

	for {
		shortUrlToken = (*u.generator).GenerateSequence()

		if _, ok := u.shortUrlTokenToUrl[shortUrlToken]; !ok {
			u.shortUrlTokenToUrl[shortUrlToken] = originalUrl
			break
		}
	}

	return
}

func (u *UrlMapper) GetUrlByShortUrl(shortUrl string) string {
	return u.shortUrlTokenToUrl[shortUrl]
}

func (u *UrlMapper) isAtMaxCapacity() bool {
	return len(u.shortUrlTokenToUrl) == u.maxCapacity
}
