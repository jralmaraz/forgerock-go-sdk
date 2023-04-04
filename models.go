package accessmanagerclient

type Realm struct {
	ID         string   `json:"_id"`
	Rev        string   `json:"_rev"`
	ParentPath string   `json:"parentPath"`
	Active     bool     `json:"active"`
	Name       string   `json:"name"`
	Aliases    []string `json:"aliases"`
}
