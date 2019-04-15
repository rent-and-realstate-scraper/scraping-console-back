package managers

import (
	"fmt"
	"scraping-console-back/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateGeojsonFromBoundingBox(t *testing.T) {

	result := models.ScrapingResultForCity{}
	result.DeviceId = "testdevice"
	result.ScrapingId = "testScrapingId"
	result.NumberOfAdsBuy = 2
	result.NumberOfAdsRent = 3
	result.AveragePrizeBuy = 1000.0
	result.AveragePrizeRent = 12.0

	result.BoundingBox1X = 1.1
	result.BoundingBox1Y = 1.2
	result.BoundingBox2X = 1.3
	result.BoundingBox2Y = 1.4

	results := []models.ScrapingResultForCity{result}
	geojson := GenerateGeoJsonFromResult(results)
	fmt.Println(geojson)
	assert.NotNil(t, geojson, "is not nil")
	assert.Equal(t, geojson.Features[0].Properties.NumberOfAdsBuy, result.NumberOfAdsBuy, "fields are filled")

}
