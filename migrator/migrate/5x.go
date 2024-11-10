package migrate

import (
	"os"
	"time"

	"github.com/DiamondDrakeVentures/ShowReport/migrator/data"
	"github.com/DiamondDrakeVentures/ShowReport/migrator/data/legacy"
	"github.com/DiamondDrakeVentures/ShowReport/migrator/encoding"
)

// migrateTo5x migrates reports to the 5x format.
type migrateTo5x struct {
	srcFile    string
	destFile   string
	srcFormat  string
	updateUser string
}

// To5x returns a new Migrator that migrate reports to the 5x format.
func To5x(srcFile, destFile, srcFormat, updateUser string) Migrator {
	return &migrateTo5x{
		srcFile:    srcFile,
		destFile:   destFile,
		srcFormat:  srcFormat,
		updateUser: updateUser,
	}
}

func (m migrateTo5x) Source() string {
	return m.srcFile
}

func (m migrateTo5x) SourceFormat() string {
	return m.srcFormat
}

func (m migrateTo5x) Dest() string {
	return m.destFile
}

func (m migrateTo5x) DestFormat() string {
	return "5x"
}

func (m migrateTo5x) User() string {
	return m.updateUser
}

func (m migrateTo5x) Execute() error {
	src, err := os.Open(m.srcFile)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(m.destFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	decoder := encoding.NewDecoder(src)
	encoder := encoding.NewEncoder(dst)

	var recToMigrate int
	var reports []data.ShowReport5X

	if m.srcFormat == "fw" {
		var fw []legacy.FW
		err = decoder.Decode(&fw)
		if err != nil {
			return err
		}

		recToMigrate = len(fw)
		reports = make([]data.ShowReport5X, recToMigrate)

		m.fromFW(&fw, &reports)
	} else if m.srcFormat == "iv" {
		var iv []legacy.IV
		err = decoder.Decode(&iv)
		if err != nil {
			return err
		}

		recToMigrate = len(iv)
		reports = make([]data.ShowReport5X, recToMigrate)

		m.fromIV(&iv, &reports)
	}

	err = encoder.Encode(&reports)
	if err != nil {
		return err
	}

	return nil
}

// fromFW migrates from FW to 5x.
func (m migrateTo5x) fromFW(fw *[]legacy.FW, reports *[]data.ShowReport5X) error {
	for i, report := range *fw {
		(*reports)[i] = data.ShowReport5X{
			CreateTime: report.Timestamp,
			CreateUser: report.Email,
			UpdateTime: time.Now(),
			UpdateUser: m.updateUser,
			SubmitTime: report.Timestamp,
			SubmitUser: report.Email,

			ProjectID:     "20240530",
			EpisodeID:     report.EpisodeID,
			EpisodeNumber: report.EpisodeNumber,
			Status:        report.Status,
			Date:          report.Date,

			Setup:       report.Setup,
			SoundCheck:  report.SoundCheck,
			MicCheck:    report.MicCheck,
			StreamStart: report.StreamStart,
			ShowStart:   report.ShowStart,

			Intermission1Start: report.Intermission1Start,
			Intermission1End:   report.Intermission1End,
			Intermission2Start: report.Intermission2Start,
			Intermission2End:   report.Intermission2End,
			Intermission3Start: report.Intermission3Start,
			Intermission3End:   report.Intermission3End,

			ShowStop:  report.ShowStop,
			StreamEnd: report.StreamEnd,
			Teardown:  report.Teardown,

			Notes: report.Notes,
		}
	}

	return nil
}

// fromIV migrates from IV to 5x.
func (m migrateTo5x) fromIV(iv *[]legacy.IV, reports *[]data.ShowReport5X) error {
	for i, report := range *iv {
		(*reports)[i] = data.ShowReport5X{
			CreateTime: report.Timestamp,
			CreateUser: report.Email,
			UpdateTime: time.Now(),
			UpdateUser: m.updateUser,
			SubmitTime: report.Timestamp,
			SubmitUser: report.Email,

			ProjectID:     "20221215",
			EpisodeID:     report.EpisodeID,
			EpisodeNumber: report.EpisodeNumber,
			Status:        report.Status,
			Date:          report.Date,

			Setup:       report.Setup,
			SoundCheck:  report.SoundCheck,
			MicCheck:    report.MicCheck,
			StreamStart: report.StreamStart,
			ShowStart:   report.ShowStart,

			Intermission1Start: report.Intermission1Start,
			Intermission1End:   report.Intermission1End,
			Intermission2Start: report.Intermission2Start,
			Intermission2End:   report.Intermission2End,
			Intermission3Start: report.Intermission3Start,
			Intermission3End:   report.Intermission3End,

			ShowStop:  report.ShowStop,
			StreamEnd: report.StreamEnd,
			Teardown:  report.Teardown,

			Notes: report.Notes,
		}
	}

	return nil
}
