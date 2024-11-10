package migrate_test

import (
	"path/filepath"
	"testing"

	"github.com/DiamondDrakeVentures/ShowReport/migrator/migrate"
	"github.com/stretchr/testify/suite"
)

func TestMigrateTo5x(t *testing.T) {
	s := new(MigrateTo5xSuite)
	s.srcDir = "testdata"
	s.srcFile = "iv.csv"
	s.dstFile = "migrated.csv"

	suite.Run(t, s)
}

type MigrateTo5xSuite struct {
	suite.Suite

	srcDir  string
	srcFile string
	dstFile string
}

func (s *MigrateTo5xSuite) source() string {
	s.T().Helper()

	return filepath.Join(s.srcDir, s.srcFile)
}

func (s *MigrateTo5xSuite) dest() string {
	s.T().Helper()

	return filepath.Join(s.T().TempDir(), s.dstFile)
}

func (s *MigrateTo5xSuite) TestSource() {
	m := migrate.To5x(s.source(), s.dest(), "iv", "test@test.local")
	s.Equal(s.source(), m.Source())
}

func (s *MigrateTo5xSuite) TestSourceFormat() {
	m := migrate.To5x(s.source(), s.dest(), "iv", "test@test.local")
	s.Equal("iv", m.SourceFormat())
}

func (s *MigrateTo5xSuite) TestDest() {
	d := s.dest()
	m := migrate.To5x(s.source(), d, "iv", "test@test.local")
	s.Equal(d, m.Dest())
}

func (s *MigrateTo5xSuite) TestDestFormat() {
	m := migrate.To5x(s.source(), s.dest(), "iv", "test@test.local")
	s.Equal("5x", m.DestFormat())
}

func (s *MigrateTo5xSuite) TestUser() {
	m := migrate.To5x(s.source(), s.dest(), "iv", "test@test.local")
	s.Equal("test@test.local", m.User())
}

func (s *MigrateTo5xSuite) TestExecute() {
	m := migrate.To5x(s.source(), s.dest(), "iv", "test@test.local")

	s.NoError(m.Execute())
}
