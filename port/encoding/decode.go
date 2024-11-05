package encoding

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"reflect"
	"time"
)

// A Decoder reads records from a CSV-encoded file.
//
// Internally, a Decoder uses encoding/csv to read the raw CSV file.
// This file is BOM-adjusted.
type Decoder struct {
	r          *bufio.Reader
	c          *csv.Reader
	init       bool
	headers    map[string]int
	revHeaders map[int]string
}

// NewDecoder returns a new Decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:          bufio.NewReader(r),
		headers:    make(map[string]int),
		revHeaders: make(map[int]string),
	}
}

// checkBOM reads the first rune and checks whether it is a BOM rune.
// If it is not a BOM rune, checkBOM unread the rune.
func (d *Decoder) checkBOM() error {
	rr, _, err := d.r.ReadRune()
	if err != nil {
		return err
	}
	if rr != '\uFEFF' {
		d.r.UnreadRune()
	}

	return nil
}

// Decode decodes a record or records from d to v.
//
// If v is a pointer to a struct, Decode will read and decode one record.
// If v is a pointer to a slice, Decode will read and decode all records.
func (d *Decoder) Decode(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return fmt.Errorf("invalid type %s", reflect.TypeOf(v))
	} else if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}

	// Initialize d
	if !d.init {
		err := d.checkBOM()
		if err != nil {
			return err
		}
		d.c = csv.NewReader(d.r)

		data, err := d.c.Read()
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}

		// Populate headers mapping and reverse mapping
		for i, field := range data {
			d.headers[field] = i
			d.revHeaders[i] = field
		}

		d.init = true
	}

	if rv.Kind() == reflect.Slice {
		// v is a slice, parse in a loop into io.EOF
		for {
			val := reflect.New(rv.Type().Elem()).Elem()

			err := d.parse(val)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					return err
				}
			}

			rv.Set(reflect.Append(rv, val))
		}
	} else {
		err := d.parse(rv)
		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}
	}

	return nil
}

// parse parses a record from d to rv.
// Call in a loop to parse *all* records.
func (d *Decoder) parse(rv reflect.Value) error {
	data, err := d.c.Read()
	if err != nil {
		return err
	}

	for i := 0; i < rv.Type().NumField(); i++ {
		field := rv.Type().Field(i)

		var fieldName string
		if csvField := field.Tag.Get("csv"); csvField != "" {
			fieldName = csvField
		} else {
			fieldName = field.Name
		}

		value := data[d.headers[fieldName]]

		if field.Type.Kind() == reflect.String {
			rv.Field(i).SetString(value)
		} else if field.Type.Kind() == reflect.Struct {
			timeLayout := timeFormat(field.Tag.Get("time"))
			if value != "" {
				t, err := time.Parse(timeLayout, value)
				if err != nil {
					return err
				}

				rv.Field(i).Set(reflect.ValueOf(&t).Elem())
			} else {
				rv.Field(i).Set(reflect.ValueOf(time.Time{}))
			}
		}
	}

	return nil
}
