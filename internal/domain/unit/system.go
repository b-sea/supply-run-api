// Package unit defines everything to manage the units of measurement domain.
package unit

// System is a measurement system.
type System int

// NoSystem et all are the different measurement systems.
const (
	NoSystem System = iota
	ImperialSystem
	MetricSystem
)

func (s System) String() string {
	switch s {
	case ImperialSystem:
		return "IMPERIAL"
	case MetricSystem:
		return "METRIC"
	case NoSystem:
		fallthrough
	default:
		return ""
	}
}

// SystemFromString converts a string to a measurement system.
func SystemFromString(s string) System {
	switch s {
	case "IMPERIAL":
		return ImperialSystem
	case "METRIC":
		return MetricSystem
	default:
		return NoSystem
	}
}
