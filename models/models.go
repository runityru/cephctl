package models

type CephConfig map[string]map[string]string

type CephConfigDifferenceKind string

const (
	CephConfigDifferenceKindAdd    CephConfigDifferenceKind = "add"
	CephConfigDifferenceKindChange CephConfigDifferenceKind = "change"
	CephConfigDifferenceKindRemove CephConfigDifferenceKind = "remove"
)

type CephConfigDifference struct {
	Kind     CephConfigDifferenceKind
	Section  string
	Key      string
	OldValue *string
	Value    *string
}
