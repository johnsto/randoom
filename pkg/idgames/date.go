package idgames

import "time"

// Date is a holder for a Time value that can be parsed from a date-only string within a JSON document.
type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(data))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
