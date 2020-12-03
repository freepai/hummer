package leveldb

import "github.com/freepai/hummer/core/shorturl/domain"

type LevelDB struct {

}

func (this *LevelDB) Save(ns string, code string, longUrl string) (*domain.ShortUrl, error) {
	return nil, nil
}

func (this *LevelDB) Get(ns string, code string) (*domain.ShortUrl, error) {
	return nil, nil
}