package subjects

type Subject string

const (
	// Zone
	ZoneCreated Subject = "zone:created"
	ZoneUpdated Subject = "zone:updated"
	ZoneDeleted Subject = "zone:deleted"
)

func GetZoneSubjects() []string {
	return []string{
		string(ZoneCreated),
		string(ZoneUpdated),
		string(ZoneDeleted),
	}
}
