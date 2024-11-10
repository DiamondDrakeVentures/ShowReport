package encoding

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
	"time"
)

// An Encoder writes records to a CSV-encoded file.
//
// Internally, an Encoder uses encoding/csv to write the raw CSV file.
type Encoder struct {
	w          *bufio.Writer
	c          *csv.Writer
	init       bool
	headers    map[string]int
	revHeaders map[int]string
}

// NewEncoder returns a new Encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:          bufio.NewWriter(w),
		headers:    make(map[string]int),
		revHeaders: make(map[int]string),
	}
}

// Encode encodes a record or records from v to e.
//
// If v is a struct or pointer to a struct, Encode will encode and write one record.
// If v is either an array, slice, or a pointer to an array or slice, Encode will encode and write
// all records.
func (e *Encoder) Encode(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return fmt.Errorf("invalid type %s", reflect.TypeOf(v))
		}
		rv = rv.Elem()
	}

	// Initialize e
	if !e.init {
		e.c = csv.NewWriter(e.w)

		var numField int
		if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
			numField = rv.Type().Elem().NumField()
		} else {
			numField = rv.Type().NumField()
		}
		rec := make([]string, numField)

		// Populate headers mapping and reverse mapping, and write header line
		for i := 0; i < numField; i++ {
			var field reflect.StructField
			if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
				field = rv.Type().Elem().Field(i)
			} else {
				field = rv.Type().Field(i)
			}

			var fieldName string
			if csvField := field.Tag.Get("csv"); csvField != "" {
				fieldName = csvField
			} else {
				fieldName = field.Name
			}

			e.headers[fieldName] = i
			e.revHeaders[i] = fieldName
			rec[i] = fieldName
		}

		err := e.c.Write(rec)
		if err != nil {
			return err
		}

		e.c.Flush()
		e.init = true
	}

	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		// v is a slice, array, or pointer to slice or array, write in a loop
		for i := 0; i < rv.Len(); i++ {
			err := e.write(rv.Index(i))
			if err != nil {
				return err
			}
		}
	} else {
		err := e.write(rv)
		if err != nil {
			return err
		}
	}

	e.c.Flush()

	return nil
}

// write writes a record from rv to e.
// Call in a loop to write *all* records.
func (e *Encoder) write(rv reflect.Value) error {
	rec := make([]string, rv.Type().NumField())
	for i := 0; i < rv.Type().NumField(); i++ {
		field := rv.Type().Field(i)

		key := field.Name
		if csv := field.Tag.Get("csv"); csv != "" {
			key = csv
		}

		var val string
		if field.Type.Kind() == reflect.String {
			val = rv.Field(i).String()
		} else if field.Type.Kind() == reflect.Struct {
			if t, ok := rv.Field(i).Interface().(time.Time); ok && !t.IsZero() {
				format := timeFormat(field.Tag.Get("time"))
				val = t.Format(format)
			}
		}

		rec[e.headers[key]] = val
	}

	err := e.c.Write(rec)
	if err != nil {
		return err
	}

	return nil
}
