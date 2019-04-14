package models

import (
	"fmt"
	"time"
)

type ScrapingResults struct {
	ResultId         string    `json:"result_id" db:"result_id"`
	PieceId          string    `json:"piece_id" db:"piece_id"`
	ScrapingId       string    `json:"scraping_id" db:"scraping_id"`
	AppId            string    `json:"app_id" db:"app_id"`
	DeviceId         string    `json:"device_id" db:"device_id"`
	DateScraped      time.Time `json:"date_scraped" db:"date_scraped"`
	AveragePrizeRent float64   `json:"date_scraped" db:"average_prize_rent"`
	NumberOfAdsRent  int       `json:"number_of_ads_rent" db:"number_of_ads_rent"`
	AveragePrizeBuy  float64   `json:"average_prize_buy" db:"average_prize_buy"`
	NumberOfAdsBuy   int       `json:"number_of_ads_buy" db:"number_of_ads_buy"`
	ExtraData        string    `json:"extra_data" db:"extra_data"`
}
type ScrapingResultForCity struct {
	ScrapingResults     ScrapingResults     `db:"scraping_results",prefix=t.`
	ScrapingPiecesIndex ScrapingPiecesIndex `db:"scraping_pieces_index",prefix=s.`
}

func GetScrapingResultsForCity(cityName string, scrapingId string) (scrapingResultForCity []ScrapingResultForCity) {
	sql := "select t.*, s.* from scraping_results t ,scraping_pieces_index s where t.piece_id = s.piece_id and t.scraping_id = '" + scrapingId + "' and s.city_name = '" + cityName + "';"

	db := GetDb()
	rows, err := db.Queryx(sql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for rows.Next() {
		var item ScrapingResultForCity
		err = rows.StructScan(&item)
		if err != nil {
			fmt.Println(err)
		}
		scrapingResultForCity = append(scrapingResultForCity, item)
	}
	return scrapingResultForCity

}
