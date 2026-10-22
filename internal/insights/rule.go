package insights

type Rule interface {
	Evaluate(ctx Context) *Insight
}
