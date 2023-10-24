package models

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mrz1836/go-sanitize"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	LongUrl  string        `json:"long_url"`
	ShortUrl string        `json:"short_url"`
	Expiry   time.Duration `json:"expiry_time"`
}

func (u *URL) Sanitize() error {
	cleanUrl := sanitize.URL(u.LongUrl)
	cleanUrl = strings.Replace(cleanUrl, "http://", "", 1)
	cleanUrl = strings.Replace(cleanUrl, "https://", "", 1)
	cleanUrl = strings.Replace(cleanUrl, "www.", "", 1)
	cleanUrl = strings.Split(cleanUrl, "/")[0]
	if cleanUrl == os.Getenv("DOMAIN") {
		return errors.New("can not shorten self")
	}
	u.LongUrl = cleanUrl
	log.Printf("Sanitized long URL: %s", cleanUrl)
	return nil
}
