package voteparam

type Vote struct {
	Target   string `json:"target"`
	Rate     int    `json:"rate"`
	OptionID int    `json:"option_id"`
	TargetID int    `json:"target_id"`
}

type Option struct {
	OptionID int    `json:"option_id"`
	Name     string `json:"name"`
	Weight   int    `json:"weight"`
	Owner    string `json:"owner"`
}
