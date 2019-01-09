package model_menu

type Meta struct {
	Title      string
	Icon       string
	Roles      []string
	NoCache    bool `json:"noCache" bson:"noCache"`
	Breadcrumb bool
}
