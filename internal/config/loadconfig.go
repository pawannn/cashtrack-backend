package config

import "github.com/spf13/viper"

type CashTrackCfg struct {
	Port            int    `mapstructure:"APP_PORT"`
	DBName          string `mapstructure:"DB_NAME"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          int    `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPass          string `mapstructure:"DB_PASS"`
	DBSsl           string `mapstructure:"DB_SSL"`
	CacheHost       string `mapstructure:"CACHEDB_HOST"`
	CachePort       string `mapstructure:"CACHEDB_PORT"`
	CachePass       string `mapstructure:"CACHEDB_PASS"`
	CacheDB         string `mapstructure:"CACHEDB_DB"`
	AuthTokenSecret string `mapstructure:"TOKEN_SECRET"`
	SMSServiceID    string `mapstructure:"SMS_SERVICE_ID"`
	SMSAccountSID   string `mapstructure:"SMS_ACCOUNT_SID"`
	SMSServiceToken string `mapstructure:"SMS_SERVICE_TOKEN"`
}

func LoadConfig() (*CashTrackCfg, error) {
	var newCfg CashTrackCfg
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}
