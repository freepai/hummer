package api

import "github.com/freepai/hummer/core/domain"

type IDStore interface {
	Save(ns string, code string, longUrl string) (*domain.ShortUrl, error)
	Get(ns string, code string) (*domain.ShortUrl, error)
}
