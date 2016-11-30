package augment

import (
	"fmt"
	"time"

	osm "github.com/paulmach/go.osm"
	"github.com/paulmach/go.osm/augment/internal/core"
)

// NoHistoryError is returned if there is no entry in the history
// map for a specific child.
type NoHistoryError struct {
	ElementType osm.ElementType
	ElementID   int64
}

// Error returns a pretty string of the error.
func (e *NoHistoryError) Error() string {
	return fmt.Sprintf("element history not found for %s %d", e.ElementType, e.ElementID)
}

// NoVisibleChildError is returned if there are no visible children
// for a parent at a given time.
type NoVisibleChildError struct {
	ElementType osm.ElementType
	ElementID   int64
	Timestamp   time.Time
}

// Error returns a pretty string of the error.
func (e *NoVisibleChildError) Error() string {
	return fmt.Sprintf("no visible child for %s %d at %v", e.ElementType, e.ElementID, e.Timestamp)
}

// UnsupportedMemberTypeError is returned if a relation member is not a
// node, way or relation.
type UnsupportedMemberTypeError struct {
	RelationID osm.RelationID
	MemberType osm.ElementType
	Index      int
}

// Error returns a pretty string of the error.
func (e *UnsupportedMemberTypeError) Error() string {
	return fmt.Sprintf("unsupported member type %v for relation %d at %d", e.MemberType, e.RelationID, e.Index)
}

func mapErrors(err error) error {
	switch t := err.(type) {
	case *core.NoHistoryError:
		return &NoHistoryError{
			ElementType: core.TypeMapToOSM[t.ChildID.Type],
			ElementID:   t.ChildID.ID,
		}
	case *core.NoVisibleChildError:
		return &NoVisibleChildError{
			ElementType: core.TypeMapToOSM[t.ChildID.Type],
			ElementID:   t.ChildID.ID,
			Timestamp:   t.Timestamp,
		}
	}

	return err
}
