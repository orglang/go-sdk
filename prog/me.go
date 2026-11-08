package prog

type Spec struct {
	Exps []ProgExp `parser:"@@*"`
}
