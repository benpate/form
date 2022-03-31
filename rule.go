package form

import "strings"

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

// form.Rule{Path:"your.mom", In:"one, two, three"}

func (rule Rule) IsEmpty() bool {
	return (rule.Path == "")
}

func (rule Rule) HyperscriptRules() string {

	var b strings.Builder

	b.WriteString("on change send checkFormRules(changed:me as Values) end ")

	if !rule.IsEmpty() {
		b.WriteString("on checkFormRules(changed) from window ")
		b.WriteString("set value to event.detail.changed['" + rule.Path + "'] if value is undefined then exit else if (value " + rule.Operator() + " " + rule.Value + ") then show me else hide me end ")
	}

	return b.String()
}
