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
	Properties  []Property `json:"properties"`
	BoundingBox []string   `json:"bbox"`
	Geometry    Geometry   `json:"geometry"`
}

type Property struct {
	AveragePrizeRent string `json:"average_prize_rent"`
	AveragePrizeBuy  string `json:"average_prize_buy"`
	NumberOfAdsRent  string `json:"number_of_ads_rent"`
	NumberOfAdsBuy   string `json:"number_of_ads_buy"`
}

type Geometry struct {
	Type        string       `json:"type"` //Polygon
	Coordinates [][][]string `json:"coordinates"`
}

func GenerateGeoJsonFromResult(scrapingResults []models.ScrapingResultForCity) (geojson Geojson) {
	var features []Feature
	geojson = Geojson{"FeatureCollection", features}
	for _, result := range scrapingResults {
		feature := generateFeature(result)
		features = append(features, feature)
	}

	return geojson
}
func generateFeature(scrapingResult models.ScrapingResultForCity) (feature Feature) {

	feature = Feature{Type: "Feature", Properties: []Property{}, Geometry: Geometry{Type: "Polygon"}}

	return feature
}
