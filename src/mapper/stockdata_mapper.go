package mapper

import (
	"StockData/src/entity"
	muzwang_stock_data "StockData/src/idl"
	tdxreader "StockData/src/idl/tdx"
)

func MapStockDataFromProtoToEntity(tdxFileData *tdxreader.TdxFileData) *entity.StockData {
	stockData := &entity.StockData{
		Meta: entity.MetaData{},
		Rows: []entity.RowData{},
	}

	if tdxFileData == nil || tdxFileData.Data == nil {
		return stockData
	}

	for _, data := range tdxFileData.Data {
		stockData.Rows = append(stockData.Rows, entity.RowData{
			Amount: data.Amount,
			DateTime: data.DateTime,
			High: data.High,
			Low: data.Low,
			Open: data.Open,
			Close: data.Close,
		})
	}

	return stockData
}

func MapMetricFromEntityToProto(metric entity.Metric) tdxreader.DateMetric {
	switch metric {
	case entity.MetricDaily:
		return tdxreader.ONE_DAY
	case entity.MetricOneMinute:
		return tdxreader.ONE_MIN
	case entity.MetricFiveMinutes:
		return tdxreader.FIVE_MIN
	default:
		return tdxreader.UNDEFINED
	}
}

func MapMetricFromProtoToEntity(metric muzwang_stock_data.Metric) entity.Metric {
	switch metric {
	case muzwang_stock_data.MetricOneMinute:
		return entity.MetricOneMinute
	case muzwang_stock_data.MetricFiveMinutes:
		return entity.MetricFiveMinutes
	case muzwang_stock_data.MetricOneDay:
		return entity.MetricDaily
	default:
		return entity.MetricInvalid
	}
}

func MapStockDataFromEntityToProto(data *entity.StockData) *muzwang_stock_data.GetStockDataResponse {
	stockData := &muzwang_stock_data.GetStockDataResponse{
		Meta: &muzwang_stock_data.MetaData{
		},
		Data: []*muzwang_stock_data.RowData{},
	}

	if data == nil || len(data.Rows) == 0 {
		return stockData
	}

	for _, row := range data.Rows {
		stockData.Data = append(stockData.Data, &muzwang_stock_data.RowData{
			Amount: row.Amount,
			Volume: row.Volume,
			Open: row.Open,
			Close: row.Close,
			High: row.High,
			Low: row.Low,
		})
	}

	return stockData
}