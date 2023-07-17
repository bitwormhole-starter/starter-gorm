package libgorm

// Configuration 数据源配置信息
type Configuration struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Enabled  bool
}
