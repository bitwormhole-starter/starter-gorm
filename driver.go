package libgorm

// Driver 驱动
type Driver interface {
	Open(c *Configuration) (Database, error)
	Registration() *DriverRegistration
}

// DriverRegistration 驱动注册信息
type DriverRegistration struct {
	Name   string
	Driver Driver
}

// DriverManager 驱动管理器
type DriverManager interface {
	FindDriver(name string) (Driver, error)
}
