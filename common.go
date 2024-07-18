package natsjs

import (
	"github.com/practionnetworks/natsjs/stream"
	"github.com/practionnetworks/natsjs/subjects"
)

const (
	Zone = stream.ZoneStream
)

const (
	ZoneCreate = subjects.ZoneCreatedSubject
	ZoneUpdate = subjects.ZoneUpdatedSubject
	ZoneDelete = subjects.ZoneDeletedSubject
)
