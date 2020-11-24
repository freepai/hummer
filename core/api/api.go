package api

import "github.com/freepai/hummer/core/domain"

type HummerAPI interface {
	Post(ns string, longUrl string) (*domain.ShortUrl, error)
	Get(ns string, code string) (*domain.ShortUrl, error)
}
