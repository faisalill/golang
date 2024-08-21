package main

import (
	"log"
	quran_dot_com "web_scraper/quran.com"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	quranDotComScraper := quran_dot_com.NewQuranDotComScraper(c)

	surahIndex := quranDotComScraper.GetQuranIndex()

	log.Println(surahIndex)
}
