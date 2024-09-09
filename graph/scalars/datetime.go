package scalars

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"time"
)

const dateTimeFormat = time.RFC3339

func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.MarshalString(t.Format(dateTimeFormat))
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("DateTime must be a string")
	}
	return time.Parse(dateTimeFormat, str)
}
