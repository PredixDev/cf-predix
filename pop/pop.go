package pop

type POP struct {
	Label       string `json:"label"`
	ApiEndpoint string `json:"apiEndpoint"`
}

func NewPOP(label, apiEndpoint string) POP {
	return POP{label, apiEndpoint}
}
