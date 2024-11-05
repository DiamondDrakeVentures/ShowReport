package data

import "time"

// ShowReport5X defines the data structure for Show Report v5.x.
type ShowReport5X struct {
	CreateTime time.Time `csv:"Create Time" time:"datetimesec"`
	CreateUser string    `csv:"Create User"`
	UpdateTime time.Time `csv:"Update Time" time:"datetimesec"`
	UpdateUser string    `csv:"Update User"`
	SubmitTime time.Time `csv:"Submit Time" time:"datetimesec"`
	SubmitUser string    `csv:"Submit User"`

	ProjectID     string    `csv:"Project ID"`
	EpisodeID     string    `csv:"Episode ID"`
	EpisodeNumber string    `csv:"Episode Number"`
	Status        string    `csv:"Status"`
	Date          time.Time `csv:"Date" time:"date"`

	Setup       time.Time `csv:"Setup" time:"datetime"`
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
