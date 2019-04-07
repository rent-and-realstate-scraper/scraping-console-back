package models

import "fmt"

type ScrapedCities struct {
	ScrapedCities []string `json:"scraped_cities"`
}


func (scrapedCities *ScrapedCities) GetScrapedCities(scrapingId string) (scrapedResultCities *ScrapedCities) {

	sql := fmt.Sprintf("select r.city_name from scraping_pieces_index r left join scraping_results t on  " +
		"t.piece_id = r.piece_id where t.scraping_id = '%s' group by r.city_name;",scrapingId)

	_, err := db.Query("SET_SQL_MODE = ''")

	db := GetDb()
	results, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	var result []string
	for results.Next() {
		var scrapedCity string
		err = results.Scan(&scrapedCity)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, scrapedCity)
	}

	scrapedResultCities = &ScrapedCities{ScrapedCities: result}
	return scrapedResultCities
}

