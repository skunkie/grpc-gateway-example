package swagger

import "os"

type docs struct{}

func (d *docs) GetAsset() func(name string) ([]byte, error)          { return GetAsset() }
func (d *docs) GetAssetInfo() func(name string) (os.FileInfo, error) { return GetAssetInfo() }
func (d *docs) GetAssetDir() func(name string) ([]string, error)     { return GetAssetDir() }

func GetAsset() func(name string) ([]byte, error)          { return Asset }
func GetAssetInfo() func(name string) (os.FileInfo, error) { return AssetInfo }
func GetAssetDir() func(name string) ([]string, error)     { return AssetDir }
func GetDocs() *docs                                       { return &docs{} }
