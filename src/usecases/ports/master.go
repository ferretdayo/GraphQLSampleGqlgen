package ports

type MasterOutputPort struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsDelete bool   `json:"is_delete"`
}
