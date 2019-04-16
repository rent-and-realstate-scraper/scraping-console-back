package managers

import (
	"fmt"
	"scraping-console-back/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateExtremeValues(t *testing.T) {

	result1 := models.ScrapingResultForCity{}
	result1.DeviceId = "testdevice"
	result1.ScrapingId = "testScrapingId"
	result1.NumberOfAdsBuy = 2
	result1.NumberOfAdsRent = 3
	result1.AveragePrizeBuy = 1000.0
	result1.AveragePrizeRent = 12.0

	result2 := models.ScrapingResultForCity{}
	result2.DeviceId = "testdevice2"
	result2.ScrapingId = "testScrapingId2"
	result2.NumberOfAdsBuy = 24
	result2.NumberOfAdsRent = 0
	result2.AveragePrizeBuy = 10077.0
	result2.AveragePrizeRent = 0.01

	results := []models.ScrapingResultForCity{result1, result2}
	intervals := CalculateExtremeValues(results)
	fmt.Println(intervals)
	assert.NotNil(t, intervals, "is not nil")
	// assert.Equal(t, geojson.Features[0].Properties.NumberOfAdsBuy, result1.NumberOfAdsBuy, "fields are filled")

}
