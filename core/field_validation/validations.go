package field_validation

type StringPlain struct {
	Type string `json:"type"`
}

func (t *StringPlain) setType(s string) {
	t.Type = s
}

type StringEnum struct {
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

func (t *StringEnum) setType(s string) {
	t.Type = s
}

type IntegerBetween struct {
	Type string `json:"type"`
	From int    `json:"from"`
	To   int    `json:"to"`
}

func (t *IntegerBetween) setType(s string) {
	t.Type = s
}

type IntegerEnum struct {
	Type    string `json:"type"`
	Options []int  `json:"options"`
}

func (t *IntegerEnum) setType(s string) {
	t.Type = s
}

type FloatBetween struct {
	Type string  `json:"type"`
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

func (t *FloatBetween) setType(s string) {
	t.Type = s
}

type FloatEnum struct {
	Type    string    `json:"type"`
	Options []float64 `json:"options"`
}

func (t *FloatEnum) setType(s string) {
	t.Type = s
}

type DateEnum struct {
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

func (t *DateEnum) setType(s string) {
	t.Type = s
}

type RelationshipOneToMany struct {
	Type string `json:"type"`
	To   string `json:"to"`
}

func (t *RelationshipOneToMany) setType(s string) {
	t.Type = s
}

type RelationshipOneToOne struct {
	Type string `json:"type"`
	To   string `json:"to"`
}

func (t *RelationshipOneToOne) setType(s string) {
	t.Type = s
}
