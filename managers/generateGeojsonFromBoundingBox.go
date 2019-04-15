package managers

import (
	"scraping-console-back/models"
)

type Geojson struct {
	Type     string    `json:"type"` //FeatureCollection
	Features []Feature `json:"features"`
}

type Feature struct {
	Type        string     `json:"type"` //Feature
	Properties  Properties `json:"properties"`
	BoundingBox []float64  `json:"bbox"`
	Geometry    Geometry   `json:"geometry"`
}

type Properties struct {
	AveragePrizeRent float64 `json:"average_prize_rent"`
	AveragePrizeBuy  float64 `json:"average_prize_buy"`
	NumberOfAdsRent  int     `json:"number_of_ads_rent"`
	NumberOfAdsBuy   int     `json:"number_of_ads_buy"`
}

type Geometry struct {
	Type        string        `json:"type"` //Polygon
	Coordinates [][][]float64 `json:"coordinates"`
}

func GenerateGeoJsonFromResult(scrapingResults []models.ScrapingResultForCity) (geojson Geojson) {
	var features []Feature
	geojson = Geojson{"FeatureCollection", features}
	for _, result := range scrapingResults {
		feature := generateFeature(result)
		features = append(features, feature)
		geojson.Features = features
	}

	return geojson
}
func generateFeature(scrapingResult models.ScrapingResultForCity) (feature Feature) {

	feature = Feature{Type: "Feature", Geometry: Geometry{Type: "Polygon"}}

	feature.Properties = Properties{AveragePrizeBuy: scrapingResult.AveragePrizeBuy,
		AveragePrizeRent: scrapingResult.AveragePrizeRent,
		NumberOfAdsBuy:   scrapingResult.NumberOfAdsBuy,
		NumberOfAdsRent:  scrapingResult.NumberOfAdsRent}

	bbox := []float64{generateBoundingBox(scrapingResult)[1][0], generateBoundingBox(scrapingResult)[1][1], generateBoundingBox(scrapingResult)[0][0], generateBoundingBox(scrapingResult)[0][1]}
	feature.BoundingBox = bbox
	coordinates := make([][][]float64, 1)
	coordinates[0] = [][]float64{
		[]float64{bbox[0], bbox[3]},
		[]float64{bbox[2], bbox[3]},
		[]float64{bbox[2], bbox[1]},
		[]float64{bbox[0], bbox[1]},
		[]float64{bbox[0], bbox[3]}}

	feature.Geometry.Coordinates = coordinates

	return feature
}

func generateBoundingBox(scrapingResult models.ScrapingResultForCity) (BoundingBox [][]float64) {
	firstPoint := []float64{scrapingResult.BoundingBox1X, scrapingResult.BoundingBox1Y}
	secondPoint := []float64{scrapingResult.BoundingBox2X, scrapingResult.BoundingBox2Y}
	return [][]float64{firstPoint, secondPoint}
}
