package types

// Для обхода рефлексии

type General struct {
	Version int    `json:"v,omitempty"`
	Group   int    `json:"group,omitempty"`
	Token   string `json:"token,omitempty"`
}

func (g *General) SetRequiredFields(version, group int, token string) {
	g.Version = version
	g.Group = group
	g.Token = token
}
