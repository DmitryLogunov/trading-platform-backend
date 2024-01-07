package timeframes

import "errors"

const (
	OneSec = iota
	OneMin
	ThreeMin
	FiveMin
	FifteenMin
	ThirtyMin
	OneHour
	FourHours
	TwelveHours
	OneDay
	OneWeek
	OneMonth
)

func GetTimeframe(t uint) (s string, err error) {
	if t == OneSec {
		return "1s", nil
	}

	if t == OneMin {
		return "1m", nil
	}

	if t == ThreeMin {
		return "3m", nil
	}

	if t == FiveMin {
		return "5m", nil
	}

	if t == FifteenMin {
		return "15m", nil
	}

	if t == ThirtyMin {
		return "30m", nil
	}

	if t == OneHour {
		return "1h", nil
	}

	if t == FourHours {
		return "4h", nil
	}

	if t == TwelveHours {
		return "12h", nil
	}

	if t == OneDay {
		return "1d", nil
	}

	if t == OneWeek {
		return "1w", nil
	}

	if t == OneMonth {
		return "1M", nil
	}

	return "", errors.New("unknown timeframe")
}
