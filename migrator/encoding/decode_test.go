package encoding_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/DiamondDrakeVentures/ShowReport/migrator/data/legacy"
	"github.com/DiamondDrakeVentures/ShowReport/migrator/encoding"
	"github.com/stretchr/testify/suite"
)

func TestDecoder(t *testing.T) {
	s := new(DecoderSuite)
	s.PayloadDir = "testdata"
	s.PayloadFile = "clean.csv"

	suite.Run(t, s)
}

type DecoderSuite struct {
	suite.Suite

	PayloadDir  string
	PayloadFile string
}

func (s *DecoderSuite) payload() string {
	return filepath.Join(s.PayloadDir, s.PayloadFile)
}

func (s *DecoderSuite) TestDecodeSingle() {
	f, err := os.Open(s.payload())
	s.Require().NoError(err)

	decoder := encoding.NewDecoder(f)

	var fw legacy.FW
	s.NoError(decoder.Decode(&fw))

	dumpStruct(s.T(), fw)

	s.Require().NoError(f.Close())
}

func (s *DecoderSuite) TestDecodeSlice() {
	f, err := os.Open(s.payload())
	s.Require().NoError(err)

	decoder := encoding.NewDecoder(f)

	var fw []legacy.FW
	s.NoError(decoder.Decode(&fw))

	dumpStruct(s.T(), fw)

	s.Require().NoError(f.Close())
}

func dumpStruct(t testing.TB, v any) {
	dmp := func(t testing.TB, prefix string, rv reflect.Value) {
		t.Helper()

		for i := 0; i < rv.NumField(); i++ {
			field := rv.Type().Field(i)

			k := field.Name
			var v string

			if csvField := field.Tag.Get("csv"); csvField != "" {
				k = csvField
			}

			if field.Type.Kind() == reflect.Struct {
				format := strings.ToLower(field.Tag.Get("time"))

				val := rv.Field(i).Interface()
				if t, ok := val.(time.Time); ok && !t.IsZero() {
					v = t.Format(timeFormat(format))
				}
			} else {
				v = rv.Field(i).String()
			}

			t.Logf("%s%s: %s", prefix, k, v)
		}
	}

	t.Helper()

	rv := reflect.ValueOf(v)

	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			t.Logf("%s: NIL", rv.Type().Name())
			return
		}
		rv = rv.Elem()
	}

	if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
		t.Logf("[]%s:", rv.Type().Elem().Name())

		for i := 0; i < rv.Len(); i++ {
			el := rv.Index(i)
			if el.Kind() == reflect.Pointer {
				if el.IsNil() {
					t.Logf("  %s: NIL", el.Type().Name())
					continue
				}
			}
			t.Logf("  %s:", el.Type().Name())

			dmp(t, "    ", el)
		}
	} else {
		t.Logf("%s:", rv.Type().Name())
		dmp(t, "  ", rv)
	}
}

func timeFormat(format string) string {
	switch strings.ToLower(format) {
	case "date":
		return encoding.FormatDate

	case "datetime", "date_time":
		return encoding.FormatDateTime

	case "timestamp":
		return encoding.FormatTimestamp

	case "time":
		fallthrough

	default:
		return encoding.FormatTime
	}
}
