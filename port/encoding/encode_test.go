package encoding_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DiamondDrakeVentures/ShowReport/port/data/legacy"
	"github.com/DiamondDrakeVentures/ShowReport/port/encoding"
	"github.com/stretchr/testify/suite"
)

func TestEncoder(t *testing.T) {
	s := new(EncoderSuite)
	s.PayloadDir = "testdata"
	s.PayloadFile = "clean.csv"
	s.TargetFile = "target.csv"

	suite.Run(t, s)
}

type EncoderSuite struct {
	suite.Suite

	PayloadDir  string
	PayloadFile string
	TargetFile  string

	PayloadData []legacy.FW
}

func (s *EncoderSuite) payload() string {
	return filepath.Join(s.PayloadDir, s.PayloadFile)
}

func (s *EncoderSuite) target() string {
	dir := s.T().TempDir()

	return filepath.Join(dir, s.TargetFile)
}

func (s *EncoderSuite) TestEncodeSingle() {
	f, err := os.Open(s.payload())
	s.Require().NoError(err)

	decoder := encoding.NewDecoder(f)

	var fw legacy.FW
	err = decoder.Decode(&fw)
	s.Require().NoError(err)

	s.Require().NoError(f.Close())

	f, err = os.Create(s.target())
	s.Require().NoError(err)

	encoder := encoding.NewEncoder(f)

	s.NoError(encoder.Encode(&fw))

	s.Require().NoError(f.Close())
}

func (s *EncoderSuite) TestEncodeSlice() {
	f, err := os.Open(s.payload())
	s.Require().NoError(err)

	decoder := encoding.NewDecoder(f)

	var fw []legacy.FW
	err = decoder.Decode(&fw)
	s.Require().NoError(err)

	s.Require().NoError(f.Close())

	f, err = os.Create(s.target())
	s.Require().NoError(err)

	encoder := encoding.NewEncoder(f)

	s.NoError(encoder.Encode(&fw))

	s.Require().NoError(f.Close())
}

func (s *EncoderSuite) TestEncodeArray() {
	f, err := os.Open(s.payload())
	s.Require().NoError(err)

	decoder := encoding.NewDecoder(f)

	var fw []legacy.FW
	err = decoder.Decode(&fw)
	s.Require().NoError(err)

	s.Require().NoError(f.Close())

	f, err = os.Create(s.target())
	s.Require().NoError(err)

	encoder := encoding.NewEncoder(f)

	farr := [3]legacy.FW(fw)

	s.NoError(encoder.Encode(&farr))

	s.Require().NoError(f.Close())
}
