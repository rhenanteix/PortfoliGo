package insights

type Engine struct {
	rules []Rule
}

type Config struct {
	SinglePayerConcentration float64
	OverduePercentage        float64
	ShortTermPercentage      float64
	DataQualityPercentage    float64
}

func DefaultConfig() Config {
	return Config{
		SinglePayerConcentration: 30,
		OverduePercentage:        5,
		ShortTermPercentage:      60,
		DataQualityPercentage:    3,
	}
}

func NewEngine(
	config Config,
) *Engine {

	return &Engine{
		rules: []Rule{
			NewConcentrationRule(
				config.SinglePayerConcentration,
			),
			NewOverdueRule(
				config.OverduePercentage,
			),
			NewShortTermRule(
				config.ShortTermPercentage,
			),
			NewDataQualityRule(
				config.DataQualityPercentage,
			),
		},
	}
}

func (e *Engine) Evaluate(
	ctx Context,
) []Insight {

	var result []Insight

	for _, rule := range e.rules {

		insight := rule.Evaluate(ctx)

		if insight == nil {
			continue
		}

		result = append(
			result,
			*insight,
		)
	}

	return result
}
