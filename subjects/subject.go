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

const (
	// Zone
	UserCreated Subject = "zone:created"
	UserUpdated Subject = "zone:updated"
	UserDeleted Subject = "zone:deleted"
)

func GetUserSubjects() []string {
	return []string{
		string(UserCreated),
		string(UserUpdated),
		string(UserDeleted),
	}
}