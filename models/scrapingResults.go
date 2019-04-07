package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ScrapingPiecesIndex struct {
	gorm.Model
	PieceId string `json:"piece_id"`
	PieceName string `json:"piece_name"`
	CityName string `json:"city_name"`
	DeviceId string `json:"device_id"`
	Scraped bool `json:"scraped"`
	BoundingBox1X float64 `json:"bounding_box1_x"`
	BoundingBox1Y float64 `json:"bounding_box1_y"`
	BoundingBox2X float64 `json:"bounding_box2_x"`
	BoundingBox2Y float64 `json:"bounding_box2_y"`
	CenterPointX float64 `json:"center_point_x"`
	CenterPointY float64 `json:"center_point_y"`
	GeojsonCoordinates string `json:"geojson_coordinates"`
	Method string `json:"method"`
}

type ScrapingResults struct {
	gorm.Model
	ResultId string `json:"result_id"`
	PieceId string `json:"piece_id"`
	ScrapingId string `json:"scraping_id"`
	AppId string `json:"app_id"`
	DeviceId string `json:"device_id"`
	DateScraped time.Time `json:"scraped"`
	AveragePrizeRent float64 `json:"average_prize_rent"`
	NumberOfAdsRent int `json:"number_of_ads_rent"`
	AveragePrizeBuy float64 `json:"average_prize_buy"`
	NumberOfAdsBuy int `json:"number_of_ads_buy"`
	ExtraData string `json:"extra_data"`
}

type ScrapedCities struct {
	ScrapedCities []ScrapedCity
}

type ScrapedCity struct {
	CityName string `json:"city_name"`
}

type ScrapingExecutionLog struct {
	gorm.Model
	ScrapingId string `json:"scraping_id"`
	LastPiece string `json:"last_piece"`
	LastResult string `json:"last_result"`
}



func (scrapedCities *ScrapedCities) GetScrapedCities(scrapingId string) (scrapedResultCities *ScrapedCities, code int) {
	code = 200
	sql := "select r.city_name from scraping_pieces_index r left join scraping_results t on  " +
		"t.piece_id = r.piece_id where t.scraping_id = '" + scrapingId+ "' group by r.city_name;"


	db := GetDb()
	// perform a db.Query insert
	results, err := db.Query(sql)


	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	var Resutl []ScrapedCity

	for results.Next() {
		var scrapedCity ScrapedCity
		// for each row, scan the result into our scrapedCity composite object
		err = results.Scan(&scrapedCity.CityName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the scrapedCity's Name attribute
		Resutl = append(Resutl, scrapedCity)
	}

	scrapedResultCities = &ScrapedCities{ScrapedCities: Resutl}
	return scrapedResultCities, code
}

