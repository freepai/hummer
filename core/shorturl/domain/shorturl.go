package domain

import "time"

type ShortUrl struct {
	Namespace string
	Code string
	LongUrl string
	CreatedAt time.Time
}

func NewShortUrl(ns string, code string, longUrl string) *ShortUrl {
	return &ShortUrl{
		Namespace: ns,
		Code: code,
		LongUrl: longUrl,
		CreatedAt: time.Now(),
	}
}