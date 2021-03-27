package tz_converter

import "time"

// TimeIn returns a time for a particular time zone based on the IANA name given
func TimeIn(t time.Time, name string) (time.Time, error) {
	location, err := time.LoadLocation(name)
	if err != nil {
		return time.Time{}, err
	}
	t = t.In(location)
	return t, nil
}
