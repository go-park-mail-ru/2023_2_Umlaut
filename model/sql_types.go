package model

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullInt64 sql.NullInt64

func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Int64)
	}
	return []byte(`null`), nil
}

func (ni *NullInt64) Scan(value interface{}) error {
	var s sql.NullInt64
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{s.Int64, false}
	} else {
		*ni = NullInt64{s.Int64, true}
	}

	return nil
}

type NullString sql.NullString

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return []byte(`null`), nil
}

func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}
