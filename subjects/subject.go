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
	UserCreated Subject = "User:created"
	UserUpdated Subject = "User:updated"
	UserDeleted Subject = "User:deleted"
)

func GetUserSubjects() []string {
	return []string{
		string(UserCreated),
		string(UserUpdated),
		string(UserDeleted),
	}
}