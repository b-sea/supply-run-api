// Package unit defines everything to manage the units of measurement domain.
package unit

// Type is a base or derived SI measurment types.
type Type int

// NoType et all are the different SI types.
const (
	NoType Type = iota
	MassType
	VolumeType
)

func (t Type) String() string {
	switch t {
	case MassType:
		return "MASS"
	case VolumeType:
		return "VOLUME"
	case NoType:
		fallthrough
	default:
		return ""
	}
}

// TypeFromString converts a string to a SI type.
func TypeFromString(s string) Type {
	switch s {
	case "MASS":
		return MassType
	case "VOLUME":
		return VolumeType
	default:
		return NoType
	}
}
