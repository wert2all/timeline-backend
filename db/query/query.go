package query

type Limit struct {
	Offset int
	Limit  int
}

func NewLimit(from int, to int) Limit { return Limit{Offset: from, Limit: to} }
