package entity

type ClientRequest struct {
    Email    string  `json:"email"`
    Code     string  `json:"code"`
    Redirect string  `json:"redirect"`
}

type ServerResponse struct {
    Status      int     `json:"status"`
    Response    string  `json:"response"`
    Flow        string  `json:"flow,omitempty"`
}
