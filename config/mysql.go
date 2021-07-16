package config

type Mysql struct {
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// func (m *Mysql) Dsn() string {
// 	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" +
// 		m.DBName + "?charset=" + m.Charset + "&parseTime=" + m.ParseTime + "&loc=Local"
// }
