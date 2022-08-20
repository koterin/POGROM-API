package entity

type Response interface {
	getId() string
}

func (i Item) getId() string {
	return i.Id
}

func (i UserInfo) getId() string {
	return i.Id
}

type ClientRequest struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Redirect string `json:"redirect"`
}

type ServerResponse struct {
	Status   int    `json:"status"`
	Response string `json:"response,omitempty"`
}

type Item struct {
	Id        string `json:"id"`
	Desc      string `json:"desc,omitempty"`
	Title     string `json:"title,omitempty"`
	Status    string `json:"status,omitempty"`
	Category  string `json:"category,omitempty"`
	CreatedAt string `json:"created-at,omitempty"`
}

type UserInfo struct {
	Id        string `json:"id"`
	Username  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	Status    string `json:"user-status,omitempty"`
	Bidded    string `json:"bidded,omitempty"`
	BidsWon   string `json:"bids-won,omitempty"`
	CreatedAt string `json:"created-at,omitempty"`
}
