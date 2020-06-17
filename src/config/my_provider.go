package config

import (
	"fmt"
	"go.uber.org/config"
	"path/filepath"
)

var Module = New


func New() (config.Provider, error ) {
	absFp, _ := filepath.Abs(filepath.FromSlash("./src/config/base.yaml"))
	fmt.Println(absFp)
	return config.NewYAMLProviderFromFiles(absFp)
}