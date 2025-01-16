package query

type Limit struct {
	Offset int
	Limit  int
}

func (l Limit) Increase(i int) Limit {
	return Limit{
		Offset: l.Offset,
		Limit:  l.Limit + 1,
	}
}

func NewLimit(from int, to int) Limit { return Limit{Offset: from, Limit: to} }
