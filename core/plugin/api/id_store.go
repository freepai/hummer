package api

type IDStore interface {
	Save(ns string, code string, longUrl string) error
	Get(ns string, code string) (string, error)
}
