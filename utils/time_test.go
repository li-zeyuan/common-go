package utils

import "testing"

func TestTimeStamp2Time(t *testing.T) {
	t.Log(TimeStamp2Time(-29926704211))
}

func TestParseDay(t *testing.T) {
	date, err := ParseDay("1806-01-02")

	t.Log(date)
	t.Log(err)
}

func TestTodayStartUTC(t *testing.T) {
	t.Log(TodayStartUTC())
}

func TestTimeStamp2TimeString(t *testing.T) {
	t.Log(TimeStamp2UTCTimeString(1696737535))
}

func TestCur2TodayEndDuration(t *testing.T) {
	t.Log(Cur2TodayEndDuration())
}
