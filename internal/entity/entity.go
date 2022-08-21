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
	Id           string `json:"id"`
	Title        string `json:"title,omitempty"`
	Desc         string `json:"desc,omitempty"`
	Status       string `json:"status,omitempty"`
	Category     string `json:"category,omitempty"`
	ImageLink    string `json:"image,omitempty"`
	RawPrice     int    `json:"raw-price,omitempty"`
	StartPrice   int    `json:"start-price,omitempty"`
	OwnerId      string `json:"owner"`
	CurrentPrice int    `json:"current-price"`
	CreatedAt    string `json:"created-at,omitempty"`
	UpdatedAt    string `json:"updated-at,omitempty"`
}

type UserInfo struct {
	Id        string `json:"id"`
	Username  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	Status    string `json:"user-status,omitempty"`
	Bidded    string `json:"bidded,omitempty"`
	BidsWon   string `json:"bids-won,omitempty"`
	ChatId    string `json:"chat-id,omitempty"`
	CreatedAt string `json:"created-at,omitempty"`
}
