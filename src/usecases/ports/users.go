package ports

type UserOutputPort struct {
	ID       int    `json:"id"`
	NickName string `json:"nickname"`
	Old      int    `json:"old"`
}
