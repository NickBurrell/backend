package app

type Config struct {
	Database struct {
		Address      string
		Username     string
		Password     string
		DatabaseName string
	}

	Redis struct {
		Address  string
		Password string
		Database int
	}
	LogFile string
}
