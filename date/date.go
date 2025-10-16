package date

import "time"

const ISO8601 = "2006-01-02"

// t.Format(dates.ISO8601)
// t.Format("2006-01-02")
// t.Format(ISO8601)
// dates.ISO8601(t)
// dates.Fmt(t)
// fmtdate(t)

func Fmt(t time.Time) string {
	return t.Format(ISO8601)
}

// utils.Date
// dates.New

func New(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalCSV(data []byte) error {
	t, err := time.Parse(ISO8601, string(data))
	if err != nil {
		return err
	}

	*d = Date{Time: t}
	return nil
}
