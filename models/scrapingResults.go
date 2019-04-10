package models

import (
	"time"
)

type ScrapingResults struct {
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
