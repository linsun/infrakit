package group

import (
	"encoding/json"
	"github.com/docker/infrakit/spi/instance"
)

// Plugin defines the functions for a Group plugin.
type Plugin interface {
	WatchGroup(grp Spec) error

	UnwatchGroup(id ID) error

	InspectGroup(id ID) (Description, error)

	DescribeUpdate(updated Spec) (string, error)

	UpdateGroup(updated Spec) error

	StopUpdate(id ID) error

	DestroyGroup(id ID) error
}

// ID is the unique identifier for a Group.
type ID string

// Spec is the specification for a Group.  The full schema for a Group is defined by the plugin.
type Spec struct {
	// ID is the unique identifier for the group.
	ID ID

	// Properties is the configuration for the group.
	Properties *json.RawMessage
}

// Description is a placeholder for the reported state of a Group.
type Description struct {
	Instances []instance.Description
}
