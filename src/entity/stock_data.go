package entity

type Metric int

const (
	MetricInvalid Metric = iota
	MetricDaily
	MetricOneMinute
	MetricFiveMinutes
)

type ReadLocalStockDataRequest struct {
	LocalPath string
	DataMetric Metric
}

type StockData struct {
	Meta MetaData
	Rows []RowData
}

type MetaData struct {
	Metric Metric
}

type RowData struct {
	DateTime string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Amount   float64
	Volume   float64
}
