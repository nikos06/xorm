package relation

const (
	None = iota
	HasOne
	BelongsTo
	HasMany
)

type Relation struct {
	Type       int
	Conditions string
	Params     []interface{}
	ForeignKey string
	OwnerKey   string
	TableName  string
}
