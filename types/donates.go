package types

type GetDonates struct {
	General
	Count     int    `json:"len,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	StartDate int    `json:"start_date,omitempty"`
	EndDate   int    `json:"end_date,omitempty"`
	Sort      string `json:"sort,omitempty"`
	Reverse   bool   `json:"reverse,omitempty"`
}

type GetLastDonates struct {
	General
	OffsetID int `json:"last,omitempty"`
}

type ChangeStatus struct {
	General
	ID     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

type Answer struct {
	General
	ID     int    `json:"id,omitempty"`
	Answer string `json:"answer,omitempty"`
}

type ChangeReward struct {
	General
	ID     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}
