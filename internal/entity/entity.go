package entity

type ClientRequest struct {
    Email    string  `json:"email"`
    Code     string  `json:"code"`
    Redirect string  `json:"redirect"`
}

type ServerResponse struct {
    Status          int     `json:"status"`
    Response        string  `json:"response,omitempty"`
    ChatId          string  `json:"chat-id,omitempty"`
    Username        string  `json:"username,omitempty"`
    UserId          string  `json:"user-id,omitempty"`
    ItemId          string  `json:"item-id,omitempty"`
    ItemDesc        string  `json:"description,omitempty"`
    ItemTitle       string  `json:"title,omitempty"`
    ItemStatus      string  `json:"status,omitempty"`
    ItemCategory    string  `json:"category,omitempty"`
    ItemCreatedAt   string  `json:"created-at,omitempty"`
}

type Items struct {
    Id          string
    Desc        string
    Title       string
    Status      string
    Category    string
    CreatedAt   string
}
