package chrono

import "time"

type Zone string
type Format string

const (
	Zone_UTC     Zone = "UTC"
	Zone_Bangkok Zone = "Asia/Bangkok"
)

const (
	Format_ISO8601  Format = "2006-01-02T15:04:05.000-0700"
	Format_TopValue Format = "2006-01-02 15:04:05"
)

var defaultZone Zone
var defaultLocation *time.Location

func SetDefaultZone(zone Zone) {
	defaultZone = zone
	defaultLocation, _ = time.LoadLocation(string(zone))
}
