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
	ResultId           string    `json:"result_id" db:"result_id"`
	PieceId            string    `json:"piece_id" db:"piece_id"`
	ScrapingId         string    `json:"scraping_id" db:"scraping_id"`
	AppId              string    `json:"app_id" db:"app_id"`
	DeviceId           string    `json:"device_id" db:"device_id"`
	DateScraped        time.Time `json:"date_scraped" db:"date_scraped"`
	AveragePrizeRent   float64   `json:"date_scraped" db:"average_prize_rent"`
	NumberOfAdsRent    int       `json:"number_of_ads_rent" db:"number_of_ads_rent"`
	AveragePrizeBuy    float64   `json:"average_prize_buy" db:"average_prize_buy"`
	NumberOfAdsBuy     int       `json:"number_of_ads_buy" db:"number_of_ads_buy"`
	ExtraData          string    `json:"extra_data" db:"extra_data"`
	PieceName          string    `json:"piece_name" db:"piece_name"`
	CityName           string    `json:"city_name" db:"city_name"`
	Scraped            bool      `json:"scraped" db:"scraped"`
	BoundingBox1X      float64   `json:"bounding_box1_x" db:"bounding_box1_x"`
	BoundingBox1Y      float64   `json:"bounding_box1_y" db:"bounding_box1_y"`
	BoundingBox2X      float64   `json:"bounding_box2_x" db:"bounding_box2_x"`
	BoundingBox2Y      float64   `json:"bounding_box2_y" db:"bounding_box2_y"`
	CenterPointX       float64   `json:"center_point_x" db:"center_point_x"`
	CenterPointY       float64   `json:"center_point_y" db:"center_point_y"`
	GeojsonCoordinates string    `json:"geojson_coordinates" db:"geojson_coordinates"`
	Method             string    `json:"method" db:"method"`
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
