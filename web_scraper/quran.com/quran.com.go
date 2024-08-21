package quran_dot_com

import (
	"github.com/gocolly/colly"
)

type SurahInfo struct {
	SurahNumber    int
	Name           string
	TranslatedName string
	AyahCount      string
}

type Quran_dot_com struct {
	collector *colly.Collector
}

func NewQuranDotComScraper(col *colly.Collector) *Quran_dot_com {
	return &Quran_dot_com{
		collector: col,
	}
}

func (q *Quran_dot_com) GetQuranIndex() []SurahInfo {
	var surahIndex []SurahInfo
	surahCount := 1

	q.collector.OnHTML(SurahIndexClassNames["mainContainerName"], func(e *colly.HTMLElement) {
		surahInfo := SurahInfo{}
		surahInfo.Name = e.ChildText(SurahIndexClassNames["childSurahName"])
		surahInfo.TranslatedName = e.ChildText(SurahIndexClassNames["childSurahName"])
		surahInfo.TranslatedName = e.ChildText(SurahIndexClassNames["childTranslatedSurahName"])
		surahInfo.AyahCount = e.ChildText(SurahIndexClassNames["childAyahCount"])
		surahInfo.SurahNumber = surahCount
		surahCount++
		surahIndex = append(surahIndex, surahInfo)
	})

	q.collector.Visit("https://quran.com")

	return surahIndex
}
