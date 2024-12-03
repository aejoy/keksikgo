package types

type ChangeCampaign struct {
	General
	ID            int    `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Status        string `json:"status,omitempty"`
	End           int    `json:"end,omitempty"`
	Point         int    `json:"point,omitempty"`
	StartReceived int    `json:"start_received,omitempty"`
	StartBackers  int    `json:"start_backers,omitempty"`
}

type ChangeCampaignReward struct {
	General
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"desc,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Limits      int    `json:"limits,omitempty"`
	Status      string `json:"status,omitempty"`
}
