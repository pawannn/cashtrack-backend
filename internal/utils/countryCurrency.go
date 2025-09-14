package utils

func GetCurrency(country string) string {
	if currency, exist := countryCurrency[country]; exist {
		return currency
	}
	return "USD"
}

func CheckCountry(county string) bool {
	_, exist := countryCurrency[county]
	return exist
}

var countryCurrency = map[string]string{
	"AE": "AED", // United Arab Emirates
	"AU": "AUD", // Australia
	"BR": "BRL", // Brazil
	"CA": "CAD", // Canada
	"CN": "CNY", // China
	"DE": "EUR", // Germany
	"ES": "EUR", // Spain
	"FR": "EUR", // France
	"GB": "GBP", // United Kingdom
	"IN": "INR", // India
	"IT": "EUR", // Italy
	"JP": "JPY", // Japan
	"KR": "KRW", // South Korea
	"MX": "MXN", // Mexico
	"RU": "RUB", // Russia
	"SA": "SAR", // Saudi Arabia
	"US": "USD", // United States
	"ZA": "ZAR", // South Africa
}
