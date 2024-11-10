package migrate

// Migrator migrates report from a given format to another.
type Migrator interface {
	// Execute executes the migration process.
	Execute() error

	// Source returns the filename of the source file.
	Source() string
	// SourceFormat returns the source file's format ID.
	SourceFormat() string
	// Dest returns the filename of the target file.
	Dest() string
	// DestFormat returns the target file's format ID.
	DestFormat() string
	// User returns the user attributed for the migrated data.
	User() string
}
