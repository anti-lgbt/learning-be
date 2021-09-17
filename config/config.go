package config

func InitializeConfig() error {
	NewLoggerService()
	if err := ConnectDatabase(); err != nil {
		return err
	}
	if err := InitSessionStore(); err != nil {
		return err
	}

	return nil
}
