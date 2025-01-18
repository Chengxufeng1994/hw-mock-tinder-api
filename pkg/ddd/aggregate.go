package ddd

const (
	AggregateNameKey    = "aggregate-name"
	AggregateIDKey      = "aggregate-id"
	AggregateVersionKey = "aggregate-version"
)

// Aggregate
type (
	AggregateNamer interface {
		AggregateName() string
	}

	Aggregate interface {
		IDer
		AggregateNamer
		IDSetter
		NameSetter
	}

	aggregate struct {
		Entity
	}
)

var _ Aggregate = (*aggregate)(nil)

func NewAggregate(id, name string) *aggregate {
	return &aggregate{
		Entity: NewEntity(id, name),
	}
}

func (agg aggregate) AggregateName() string { return agg.EntityName() }
