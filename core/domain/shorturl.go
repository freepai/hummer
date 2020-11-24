package domain

import "time"

type ShortUrl struct {
	NamespaceId string
	Code string
	LongUrl string
	createdAt time.Time
}