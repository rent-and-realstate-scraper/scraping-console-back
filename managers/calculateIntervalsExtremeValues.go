package managers

import (
	"math"
	"scraping-console-back/models"
)

type Interval struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type ExtremeValues struct {
	Options map[string]Interval
}

func CalculateExtremeValues(scrapingResults []models.ScrapingResultForCity) (extremeValues ExtremeValues) {
	options := make(map[string]Interval)
	initializeIntervals(options)
	extremeValues = ExtremeValues{Options: options}
	for _, result := range scrapingResults {
		options["number_of_ads_buy"] = Interval{Max: math.Max(options["number_of_ads_buy"].Max, float64(result.NumberOfAdsBuy)),
			Min: math.Min(options["number_of_ads_buy"].Min, float64(result.NumberOfAdsBuy))}
		options["number_of_ads_rent"] = Interval{Max: math.Max(options["number_of_ads_rent"].Max, float64(result.NumberOfAdsRent)),
			Min: math.Min(options["number_of_ads_rent"].Min, float64(result.NumberOfAdsRent))}
		options["average_prize_buy"] = Interval{Max: math.Max(options["average_prize_buy"].Max, float64(result.AveragePrizeBuy)),
			Min: math.Min(options["average_prize_buy"].Min, float64(result.AveragePrizeBuy))}
		options["average_prize_rent"] = Interval{Max: math.Max(options["average_prize_rent"].Max, float64(result.AveragePrizeRent)),
			Min: math.Min(options["average_prize_rent"].Min, float64(result.AveragePrizeRent))}
	}
	return extremeValues
}

func initializeIntervals(options map[string]Interval) {
	options["number_of_ads_buy"] = Interval{0, 999999}
	options["number_of_ads_rent"] = Interval{0, 999999}
	options["average_prize_buy"] = Interval{0, 999999}
	options["average_prize_rent"] = Interval{0, 999999}
}
