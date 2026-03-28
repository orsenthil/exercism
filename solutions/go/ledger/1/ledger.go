package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Entry struct {
	Date        string
	Description string
	Change      int // in cents
}

type Locale string

const (
	LocaleUS    Locale = "en-US"
	LocaleDutch Locale = "nl-NL"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type Formatter interface {
	FormatDate(date time.Time) string
	FormatMoney(amount int, currency Currency) string
	GetDateLayout() string
	GetDecimalSeparator() string
	GetThousandSeparator() string
	GetCurrencySymbol(currency Currency) string
}

type USFormatter struct{}
type DutchFormatter struct{}

func (f USFormatter) FormatDate(date time.Time) string {
	return date.Format(f.GetDateLayout())
}

func (f USFormatter) FormatMoney(amount int, currency Currency) string {
	dollars := amount / 100
	cents := amount % 100
	return f.GetCurrencySymbol(currency) + strconv.Itoa(dollars) + "." + f.GetDecimalSeparator() + strconv.Itoa(cents)
}

func (f USFormatter) GetDateLayout() string {
	return "01/02/2006"
}

func (f USFormatter) GetDecimalSeparator() string {
	return "."
}

func (f USFormatter) GetThousandSeparator() string {
	return ","
}

func (f USFormatter) GetCurrencySymbol(currency Currency) string {
	switch currency {
	case USD:
		return "$"
	case EUR:
		return "€"
	default:
		return ""

	}
}

func (f DutchFormatter) FormatDate(date time.Time) string {
	return date.Format(f.GetDateLayout())
}

func (f DutchFormatter) FormatMoney(amount int, currency Currency) string {
	dollars := amount / 100
	cents := amount % 100
	return f.GetCurrencySymbol(currency) + strconv.Itoa(dollars) + f.GetThousandSeparator() + f.GetDecimalSeparator() + strconv.Itoa(cents)
}

func (f DutchFormatter) GetDateLayout() string {
	return "02-01-2006"
}

func (f DutchFormatter) GetDecimalSeparator() string {
	return ","
}

func (f DutchFormatter) GetThousandSeparator() string {
	return "."
}

func (f DutchFormatter) GetCurrencySymbol(currency Currency) string {
	switch currency {
	case USD:
		return "$"
	case EUR:
		return "€"
	default:
		return ""

	}
}

func formatMoney(amount int, currency Currency, f Formatter) string {
	negative := amount < 0
	if negative {
		amount = -amount
	}

	dollars := amount / 100
	cents := amount % 100

	// Format with proper separators
	moneyStr := formatNumber(dollars, f.GetThousandSeparator())
	result := fmt.Sprintf("%s%s%02d", moneyStr, f.GetDecimalSeparator(), cents)

	if negative {
		return fmt.Sprintf("(%s%s)", f.GetCurrencySymbol(currency), result)
	}

	return fmt.Sprintf("%s%s ", f.GetCurrencySymbol(currency), result)
}

func formatNumber(n int, thousandSep string) string {
	str := strconv.Itoa(n)
	result := ""
	for i, j := len(str)-1, 0; i >= 0; i-- {
		if j > 0 && j%3 == 0 {
			result = thousandSep + result
		}
		result = string(str[i]) + result
		j++
	}

	return result
}

func parseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// Get appropriate formatter based on locale
	var formatter Formatter
	switch locale {
	case "en-US":
		formatter = USFormatter{}
	case "nl-NL":
		formatter = DutchFormatter{}
	default:
		return "", errors.New("unsupported locale")
	}

	sort.SliceStable(entries, func(i, j int) bool {
		date1, _ := parseDate(entries[i].Date)
		date2, _ := parseDate(entries[j].Date)

		if date1.Equal(date2) {
			if entries[i].Description == entries[j].Description {
				return entries[i].Change < entries[j].Change
			}
			return entries[i].Description < entries[j].Description
		}
		return date1.Before(date2)

	})

	// Format the header
	result := "Date       | Description               | Change\n"

	// Format each entry
	for _, entry := range entries {
		date, _ := parseDate(entry.Date)
		desc := formatDescription(entry.Description)
		money := formatMoney(entry.Change, Currency(currency), formatter)

		result += fmt.Sprintf("%s | %-25s | %13s\n", formatter.FormatDate(date), desc, money)
	}

	return result, nil
}

func formatDescription(desc string) string {
	if len(desc) > 25 {
		return desc[:22] + "..."
	}
	return desc
}
