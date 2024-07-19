package natsjs

import (
	"github.com/practionnetworks/natsjs/stream"
	"github.com/practionnetworks/natsjs/subjects"
)

var ZoneStream = stream.ZoneStream

var ZoneSubject = subjects.GetZoneSubjects()

var UserStream = stream.UserStream

var UserSubject = subjects.GetUserSubjects()
