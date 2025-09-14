package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Env             = "dev"
	AppPort         = "8080"
	DBName          = ""
	DBHost          = ""
	DBPort          = "5432"
	DBUser          = ""
	DBPass          = ""
	DBSsl           = "disable"
	CacheHost       = "localhost"
	CachePort       = "6379"
	CachePass       = ""
	CacheDB         = "0"
	AuthTokenSecret = ""
	SMSServiceID    = ""
	SMSAccountSID   = ""
	SMSServiceToken = ""
)

type CashTrackCfg struct {
	ENV             string `mapstructure:"ENV"`
	Port            int    `mapstructure:"APP_PORT"`
	DBName          string `mapstructure:"DB_NAME"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          int    `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPass          string `mapstructure:"DB_PASS"`
	DBSsl           string `mapstructure:"DB_SSL"`
	CacheHost       string `mapstructure:"CACHEDB_HOST"`
	CachePort       int    `mapstructure:"CACHEDB_PORT"`
	CachePass       string `mapstructure:"CACHEDB_PASS"`
	CacheDB         int    `mapstructure:"CACHEDB_DB"`
	AuthTokenSecret string `mapstructure:"TOKEN_SECRET"`
	SMSServiceID    string `mapstructure:"SMS_SERVICE_ID"`
	SMSAccountSID   string `mapstructure:"SMS_ACCOUNT_SID"`
	SMSServiceToken string `mapstructure:"SMS_SERVICE_TOKEN"`
}

func LoadConfig() (*CashTrackCfg, error) {
	// If running in production, use embedded values
	if Env == "prod" || Env == "production" {
		fmt.Println("Using embedded production configuration")
		return &CashTrackCfg{
			ENV:             Env,
			Port:            mustAtoi(AppPort),
			DBName:          DBName,
			DBHost:          DBHost,
			DBPort:          mustAtoi(DBPort),
			DBUser:          DBUser,
			DBPass:          DBPass,
			DBSsl:           DBSsl,
			CacheHost:       CacheHost,
			CachePort:       mustAtoi(CachePort),
			CachePass:       CachePass,
			CacheDB:         mustAtoi(CacheDB),
			AuthTokenSecret: AuthTokenSecret,
			SMSServiceID:    SMSServiceID,
			SMSAccountSID:   SMSAccountSID,
			SMSServiceToken: SMSServiceToken,
		}, nil
	}

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var newCfg CashTrackCfg
	if err := viper.Unmarshal(&newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}

func mustAtoi(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}
