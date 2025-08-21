package form

type Rule struct {
	Path  string `json:"path"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

func (rule Rule) Operator() string {
	if rule.Op == "" {
		return "=="
	}

	return rule.Op
}

func (rule Rule) IsEmpty() bool {
	return (rule.Path == "")
}

func (rule Rule) HyperscriptRules() string {
	return ""

}
