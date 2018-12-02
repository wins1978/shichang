package model

// Conf 配置信息
type Conf struct {
	UploadFolder string            `yaml:"UploadFolder"`
	LogFolder    string            `yaml:"LogFolder"`
	ColumnMap    map[string]string `yaml:"ColumnMap"`
}
