package models

type MetricsResponse struct {
	MeanUnitsSold        uint   `json:"mean_units_sold"`
	CheapestBook         string `json:"cheapest_book"`
	BooksWrittenByAuthor uint   `json:"books_written_by_author"`
}

type GetMetricsRequest struct {
	Author string `form:"author"`
}
