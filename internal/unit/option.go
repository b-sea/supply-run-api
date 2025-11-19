package unit

// Metric, et al. are standard unit systems and base types.
var (
	Metric = SetSystem("metric") //nolint: gochecknoglobals
	US     = SetSystem("us")     //nolint: gochecknoglobals

	Mass   = SetBaseType("mass")   //nolint: gochecknoglobals
	Volume = SetBaseType("volume") //nolint: gochecknoglobals
)

// Option is a Unit creation option.
type Option func(u *Unit)

// SetSystem sets a system on the Unit.
func SetSystem(system string) Option {
	return func(u *Unit) {
		u.system = system
	}
}

// SetBaseType sets a base type of the Unit.
func SetBaseType(base string) Option {
	return func(u *Unit) {
		u.base = base
	}
}

// WithCustomPlural sets the plurality naming for the Unit.
func WithCustomPlural(plural string) Option {
	return func(u *Unit) {
		u.plural = plural
	}
}

// WithNoPlural removes any plurality on the Unit.
func WithNoPlural() Option {
	return func(u *Unit) {
		u.plural = u.name
	}
}
