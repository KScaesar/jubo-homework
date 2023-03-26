package database

type DbConfig struct {
	User        string `configs:"user"`
	Password    string `configs:"password"`
	Host        string `configs:"host"`
	Port        string `configs:"port"`
	Database    string `configs:"database"`
	MaxConn     int    `configs:"max_conn"`
	MaxIdleConn int    `configs:"max_idle_conn"`
	DebugMode   bool   `configs:"debug_mode"`
}

func (c *DbConfig) setDefault() {
	if c.MaxConn <= 0 {
		const default_ = 8
		c.MaxConn = default_
	}

	if c.MaxIdleConn <= 0 {
		const default_ = 4
		c.MaxIdleConn = default_
	}
}
