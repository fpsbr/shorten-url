package data

import (
	"database/sql"
	"fmt"
	"math/rand"
)

type URLShortener struct {
	DB      *sql.DB
	baseUrl string
}

func NewURLShortener(db *sql.DB, baseUrl string) *URLShortener {
	return &URLShortener{
		DB:      db,
		baseUrl: baseUrl,
	}
}

func (us *URLShortener) generateRandomUrl(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, length)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}
	return string(key)
}

func (us *URLShortener) ShortenURL(longUrl string) (string, error) {
	urlCode := us.generateRandomUrl(10)
	shortUrl := fmt.Sprintf("%s/%s", us.baseUrl, urlCode)
	_, err := us.DB.Exec("INSERT INTO url(url_code, long_url, short_url) VALUES (?, ?, ?)", urlCode, longUrl, shortUrl)
	return shortUrl, err
}

func (us *URLShortener) GetUrl(urlCode string) (string, error) {
	var longUrl string
	err := us.DB.QueryRow("SELECT long_url FROM url WHERE url_code = ?", urlCode).Scan(&longUrl)
	url := fmt.Sprintf("%s", longUrl)
	return url, err
}
