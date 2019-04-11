package models

import "fmt"

type ScrapingPiecesIndex struct {
	PieceId            string  `json:"piece_id"`
	PieceName          string  `json:"piece_name"`
	CityName           string  `json:"city_name"`
	DeviceId           string  `json:"device_id"`
	Scraped            bool    `json:"scraped"`
	BoundingBox1X      float64 `json:"bounding_box1_x"`
	BoundingBox1Y      float64 `json:"bounding_box1_y"`
	BoundingBox2X      float64 `json:"bounding_box2_x"`
	BoundingBox2Y      float64 `json:"bounding_box2_y"`
	CenterPointX       float64 `json:"center_point_x"`
	CenterPointY       float64 `json:"center_point_y"`
	GeojsonCoordinates string  `json:"geojson_coordinates"`
	Method             string  `json:"method"`
}

func IndexGetScrapingCount(scraped bool, deviceId string) (count int) {
	sql := fmt.Sprint("select count(*) from scraping_pieces_index where scraped=%b and device_id = \"%s\";", scraped, deviceId)

	db := GetDb()
	rows, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
	}

	return count

}
