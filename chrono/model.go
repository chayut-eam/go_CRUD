package chrono

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Unix struct {
	time.Time
}

type ISO8601 struct {
	time.Time
}

type TopValue struct {
	time.Time
}

var nilTime time.Time = time.Time{}.UTC()

func (u Unix) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", (u.Time.UTC().Unix()))), nil
}

func (u *Unix) UnmarshalJSON(b []byte) error {
	var epoch int64
	err := json.Unmarshal(b, &epoch)
	if err != nil {
		return err
	}
	u.Time = time.Unix(epoch, 0).UTC()
	return nil
}

func (t ISO8601) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.In(defaultLocation).Format(string(Format_ISO8601)) + `"`), nil
}

func (t *ISO8601) UnmarshalJSON(b []byte) error {
	// trim leading and trailing double quotes
	s := strings.Trim(string(b), "\"")

	// parsed time is always in UTC
	var err error
	t.Time, err = time.Parse(string(Format_ISO8601), s)
	if err != nil {
		return err
	}
	return nil
}

func (t ISO8601) ToUnix() Unix {
	epoch := t.Time.UTC().Unix()
	return Unix{
		Time: time.Unix(epoch, 0),
	}
}

func (t TopValue) MarshalJSON() ([]byte, error) {
	if nilTime.UTC().Unix() == t.Time.UTC().Unix() {
		return []byte("null"), nil
	}
	// return []byte(`"` + t.Time.In(defaultLocation).Format(string(Format_TopValue)) + `"`), nil
	return []byte(`"` + t.Time.In(defaultLocation).Format(string(Format_ISO8601)) + `"`), nil
}

func (t *TopValue) UnmarshalJSON(b []byte) error {
	// trim leading and trailing double quotes
	s := strings.Trim(string(b), "\"")
	if len(s) == 0 {
		t.Time = nilTime
		return nil
	}

	// parsed time is always in UTC
	var err error
	t.Time, err = time.Parse(string(Format_TopValue), s)
	if err != nil {
		return err
	}

	// if the parsing time is sent in Asia/Bangkok, we have to adjust offset by -7 hours
	if defaultZone == Zone_Bangkok {
		t.Time = t.Time.Add(time.Hour * -7)
	}

	return nil
}

func Now() ISO8601 {
	return ISO8601{
		Time: time.Now().UTC(),
	}
}

func Timestamp() *int64 {
	timestamp := time.Now().UTC().Unix()
	return &timestamp
}
