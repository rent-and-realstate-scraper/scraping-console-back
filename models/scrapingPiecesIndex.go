package models

import (
	"fmt"
)

type ScrapingPiecesIndex struct {
	PieceId            string  `json:"piece_id" db:"piece_id"`
	PieceName          string  `json:"piece_name" db:"piece_name"`
	CityName           string  `json:"city_name" db:"city_name"`
	DeviceId           string  `json:"device_id" db:"device_id"`
	Scraped            bool    `json:"scraped" db:"scraped"`
	BoundingBox1X      float64 `json:"bounding_box1_x" db:"bounding_box1_x"`
	BoundingBox1Y      float64 `json:"bounding_box1_y" db:"bounding_box1_y"`
	BoundingBox2X      float64 `json:"bounding_box2_x" db:"bounding_box2_x"`
	BoundingBox2Y      float64 `json:"bounding_box2_y" db:"bounding_box2_y"`
	CenterPointX       float64 `json:"center_point_x" db:"center_point_x"`
	CenterPointY       float64 `json:"center_point_y" db:"center_point_y"`
	GeojsonCoordinates string  `json:"geojson_coordinates" db:"geojson_coordinates"`
	Method             string  `json:"method" db:"method"`
}

func IndexGetScrapingCount(scraped bool, deviceId string) (count int) {
	sql := fmt.Sprint("select count(*) from scraping_pieces_index where scraped=%b and device_id = \"%s\";", scraped, deviceId)

	db := GetDb()
	rows := db.QueryRow(sql)
	err := rows.Scan(&count)

	if err != nil {
		panic(err.Error())
	}

	return count

}

func ListDevices() (devices []string) {
	sql := "select device_id from scraping_pieces_index group by device_id"
	db := GetDb()
	rows, _ := db.Query(sql)
	for rows.Next() {
		var device string
		err := rows.Scan(device)
		if err != nil {
			fmt.Println(err)
		}
		devices = append(devices, device)
	}

	return devices

}
