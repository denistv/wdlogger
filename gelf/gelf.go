package gelf

const (
	VersionField      = "version"
	HostField         = "host"
	ShortMessageField = "short_message"
	FullMessageField  = "full_message"
	TimestampField    = "timestamp"
	LevelField        = "level"
	FacilityField     = "facility"
	LineField         = "line"
)

const additionalFieldPrefix = "_"

func AdditionalField(name string) string {
	return additionalFieldPrefix + name
}
