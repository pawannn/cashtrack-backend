package utils

func GetCurrency(country string) string {
	if currency, exist := countryCurrency[country]; exist {
		return currency
	}
	return "USD"
}

var countryCurrency = map[string]string{
	"ARE": "AED",
	"AUS": "AUD",
	"BRA": "BRL",
	"CAN": "CAD",
	"CHN": "CNY",
	"DEU": "EUR",
	"ESP": "EUR",
	"FRA": "EUR",
	"GBR": "GBP",
	"IND": "INR",
	"ITA": "EUR",
	"JPN": "JPY",
	"KOR": "KRW",
	"MEX": "MXN",
	"RUS": "RUB",
	"SAU": "SAR",
	"USA": "USD",
	"ZAF": "ZAR",
}
