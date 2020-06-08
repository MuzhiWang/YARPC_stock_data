package config

import "go.uber.org/config"

var Module = New


func New() (config.Provider, error ) {
	return config.NewYAMLProviderFromFiles("C:\\Users\\wmz66\\go\\src\\StockData\\src\\config\\base.yaml")
}