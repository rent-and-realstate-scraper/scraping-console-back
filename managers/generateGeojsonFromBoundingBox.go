type struct Geojson {
	Type string `json:"type"` //FeatureCollection
	Features []Feature `json:"features"`
}

type struct Feature {
	Type string `json:"type"` //Feature
	Properties []Property `json:"properties"`
	BoundingBox []string `json:"bbox"`
	Geometry Geometry `json:"geometry"`
}

type struct Property {
	AveragePrizeRent string `json:"average_prize_rent"`
	AveragePrizeBuy string `json:"average_prize_buy"`
	NumberOfAdsRent string `json:"number_of_ads_rent"`
	NumberOfAdsBuy string `json:"number_of_ads_buy"`
}

type struct Geometry {
	Type string `json:"type"` //Polygon
	Coordinates [][][]string `json:"coordinates"`
}

