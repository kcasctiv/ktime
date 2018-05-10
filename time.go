package ktime

import (
	"encoding/json"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Time is milliseconds in JSON, otherwise - time.Time
type Time time.Time

// GetBSON bson Getter interface implementation
func (t Time) GetBSON() (interface{}, error) {
	return time.Time(t), nil
}

// SetBSON bson Setter interface implementation
func (t *Time) SetBSON(raw bson.Raw) error {
	var tt time.Time
	err := raw.Unmarshal(&tt)
	if err != nil {
		return err
	}

	*t = Time(tt)
	return nil
}

// MarshalJSON json Marshaler interface implementation
func (t Time) MarshalJSON() ([]byte, error) {
	stamp := strconv.FormatInt(t.MS(), 10)
	return []byte(stamp), nil
}

// UnmarshalJSON json Unmarshaler interface implementation
func (t *Time) UnmarshalJSON(data []byte) error {
	var n json.Number
	err := json.Unmarshal(data, &n)
	if err != nil {
		return err
	}

	ms, err := n.Int64()
	if err != nil {
		return err
	}

	*t = FromMS(ms)
	return nil
}

// MS representation in milliseconds
func (t Time) MS() int64 {
	return time.Time(t).UnixNano() / 1000000
}

// FromMS creates new Time with initial value,
// that equals milliseconds from ms param
func FromMS(ms int64) Time {
	sec := int64(ms / 1000)
	nsec := int64(int64(ms)-sec*1000) * 1000000
	return Time(time.Unix(sec, nsec))
}
