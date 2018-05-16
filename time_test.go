package ktime

import (
	"encoding/json"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func TestGetBSON(t *testing.T) {
	type s1 struct{ Time Time }
	type s2 struct{ Time time.Time }
	d := time.Date(1989, time.October, 9, 0, 0, 0, 0, time.UTC)
	s1v := s1{Time{d}}

	bytes, err := bson.Marshal(&s1v)
	if err != nil {
		t.Fatal(err)
	}

	var s2v s2
	err = bson.Unmarshal(bytes, &s2v)
	if err != nil {
		t.Error(err)
	} else if !s2v.Time.Equal(d) {
		t.Errorf("Expected %s, got %s", d, s2v.Time)
	}
}

func TestSetBSON(t *testing.T) {
	type s1 struct{ Time time.Time }
	type s2 struct{ Time Time }
	d := time.Date(1989, time.October, 9, 0, 0, 0, 0, time.UTC)
	s1v := s1{d}

	bytes, err := bson.Marshal(&s1v)
	if err != nil {
		t.Fatal(err)
	}

	var s2v s2
	err = bson.Unmarshal(bytes, &s2v)
	if err != nil {
		t.Error(err)
	} else if !s2v.Time.Equal(d) {
		t.Errorf("Expected %s, got %s", d, s2v.Time)
	}
}

func TestMarshalJSON(t *testing.T) {
	type s1 struct{ Time Time }
	type s2 struct{ Time int64 }
	d := time.Date(1989, time.October, 9, 0, 0, 0, 0, time.UTC)
	ms := d.UnixNano() / 1000000
	s1v := s1{Time{d}}

	bytes, err := json.Marshal(&s1v)
	if err != nil {
		t.Fatal(err)
	}

	var s2v s2
	err = json.Unmarshal(bytes, &s2v)
	if err != nil {
		t.Error(err)
	} else if ms != s2v.Time {
		t.Errorf("Expected %d, got %d", ms, s2v.Time)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	type s1 struct{ Time int64 }
	type s2 struct{ Time Time }
	type s3 struct{ Time float64 }
	d := time.Date(1989, time.October, 9, 0, 0, 0, 0, time.UTC)
	ms := d.UnixNano() / 1000000
	s1v := s1{ms}

	bytes, err := json.Marshal(&s1v)
	if err != nil {
		t.Fatal(err)
	}

	var s2v s2
	err = json.Unmarshal(bytes, &s2v)
	if err != nil {
		t.Error(err)
	} else if !s2v.Time.Equal(d) {
		t.Errorf("Expected %s, got %s", d, s2v.Time)
	}

	s3v := s3{1989.1009}
	bytes, err = json.Marshal(&s3v)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bytes, &s2v)
	if err == nil {
		t.Errorf("Expected error, got %s", s2v.Time)
	}

	err = json.Unmarshal([]byte(`{"Time":true}`), &s2v)
	if err == nil {
		t.Errorf("Expected error, got %s", s2v.Time)
	}
}

func TestMS(t *testing.T) {
	d := time.Date(1989, time.October, 9, 0, 0, 0, 0, time.UTC)
	dms := d.UnixNano() / 1000000
	tms := Time{d}.MS()

	if tms != dms {
		t.Errorf("Expected %d, got %d", dms, tms)
	}
}

func TestFromMS(t *testing.T) {
	d := time.Date(1989, time.October, 9, 0, 0, 0, 0, time.UTC)
	ms := d.UnixNano() / 1000000
	tm := FromMS(ms)

	if !tm.Equal(d) {
		t.Errorf("Expected %s, got %s", d, tm)
	}
}
