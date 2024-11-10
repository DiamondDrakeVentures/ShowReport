package legacy

import "time"

// IV defines the data structure for the legacy Show Report format for Into Volthune.
type IV struct {
	Timestamp     time.Time `csv:"Timestamp" time:"timestamp"`
	Email         string    `csv:"Email"`
	Status        string    `csv:"Status"`
	Date          time.Time `csv:"Date" time:"date"`
	EpisodeNumber string    `csv:"Episode Number"`
	EpisodeID     string    `csv:"Episode ID"`

	Setup       time.Time `csv:"Setup" time:"datetimesec"`
	SoundCheck  time.Time `csv:"Sound Check"`
	MicCheck    time.Time `csv:"Mic Check"`
	StreamStart time.Time `csv:"Stream Start"`
	ShowStart   time.Time `csv:"Show Start"`

	Intermission1Start time.Time `csv:"Intermission 1 Start"`
	Intermission1End   time.Time `csv:"Intermission 1 End"`
	Intermission2Start time.Time `csv:"Intermission 2 Start"`
	Intermission2End   time.Time `csv:"Intermission 2 End"`
	Intermission3Start time.Time `csv:"Intermission 3 Start"`
	Intermission3End   time.Time `csv:"Intermission 3 End"`

	ShowStop  time.Time `csv:"Show Stop"`
	StreamEnd time.Time `csv:"Stream End"`
	Teardown  time.Time `csv:"Teardown" time:"datetime"`

	Notes string `csv:"Notes"`
}
