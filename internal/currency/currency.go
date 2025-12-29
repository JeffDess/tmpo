package currency

import (
	"fmt"
	"strings"
)

const DefaultCurrency = "USD"

var currencySymbols = map[string]string{
	// Americas
	"USD": "$",   // United States Dollar
	"CAD": "CA$", // Canadian Dollar
	"BRL": "R$",  // Brazilian Real
	"MXN": "MX$", // Mexican Peso
	"ARS": "AR$", // Argentine Peso

	// Europe
	"EUR": "€",   // Euro
	"GBP": "£",   // British Pound Sterling
	"CHF": "Fr",  // Swiss Franc
	"SEK": "kr",  // Swedish Krona
	"NOK": "kr",  // Norwegian Krone
	"DKK": "kr",  // Danish Krone
	"PLN": "zł",  // Polish Zloty
	"CZK": "Kč",  // Czech Koruna

	// Asia
	"JPY": "¥",   // Japanese Yen
	"CNY": "¥",   // Chinese Yuan
	"INR": "₹",   // Indian Rupee
	"KRW": "₩",   // South Korean Won
	"SGD": "S$",  // Singapore Dollar
	"HKD": "HK$", // Hong Kong Dollar
	"THB": "฿",   // Thai Baht
	"IDR": "Rp",  // Indonesian Rupiah
	"MYR": "RM",  // Malaysian Ringgit
	"PHP": "₱",   // Philippine Peso
	"VND": "₫",   // Vietnamese Dong

	// Oceania
	"AUD": "A$",  // Australian Dollar
	"NZD": "NZ$", // New Zealand Dollar

	// Middle East & Africa
	"AED": "د.إ", // UAE Dirham
	"SAR": "﷼",   // Saudi Riyal
	"ILS": "₪",   // Israeli Shekel
	"ZAR": "R",   // South African Rand
	"EGP": "E£",  // Egyptian Pound
	"TRY": "₺",   // Turkish Lira
}

func FormatCurrency(amount float64, currencyCode string) string {
	currencyCode = strings.ToUpper(strings.TrimSpace(currencyCode))

	if currencyCode == "" || !IsSupported(currencyCode) {
		currencyCode = DefaultCurrency
	}

	symbol := GetSymbol(currencyCode)
	return fmt.Sprintf("%s%.2f", symbol, amount)
}

func GetSymbol(currencyCode string) string {
	currencyCode = strings.ToUpper(strings.TrimSpace(currencyCode))

	if symbol, exists := currencySymbols[currencyCode]; exists {
		return symbol
	}

	return currencyCode
}

func IsSupported(currencyCode string) bool {
	currencyCode = strings.ToUpper(strings.TrimSpace(currencyCode))
	_, exists := currencySymbols[currencyCode]
	return exists
}

func GetSupportedCurrencies() []string {
	currencies := make([]string, 0, len(currencySymbols))
	for code := range currencySymbols {
		currencies = append(currencies, code)
	}
	return currencies
}
