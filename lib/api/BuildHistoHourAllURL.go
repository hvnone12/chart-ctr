package api

import "strings"

// BuildHistoHourAllURL :: Construct URL for API call
func BuildHistoHourAllURL(coin, base string) string {
	scheme := "https://"
	host := "min-api.cryptocompare.com"
	path := "/data/histohour"
	link := scheme + host + path

	fsym := "?fsym=" + strings.ToUpper(coin)
	tsym := "&tsym=" + strings.ToUpper(base)
	limit := "&limit=2000"
	aggregate := "&aggregate=6"
	query := fsym + tsym + limit + aggregate

	return link + query
}
