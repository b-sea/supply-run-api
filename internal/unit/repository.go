package unit

import "context"

// Repository defines all data interactions required for units.
type Repository interface {
	CreateUnit(ctx context.Context, unit *Unit) error
	CreateConversion(ctx context.Context, conversion *Conversion) error
}
