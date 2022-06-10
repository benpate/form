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

	/*
		var b strings.Builder
		b.WriteString("on load log 'here' send go end on change send go end on go log me send checkFormRules(changed:me as Values) end ")

		if !rule.IsEmpty() {
			b.WriteString("on checkFormRules(changed) from window ")
			b.WriteString("set value to event.detail.changed['" + rule.Path + "'] if value is undefined then exit else if (value " + rule.Operator() + " " + rule.Value + ") then show me else hide me end ")
		}

		return b.String()
	*/
}
