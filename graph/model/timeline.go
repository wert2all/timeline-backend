package model

type Timeline struct {
	ID        int           `json:"id"`
	Name      *string       `json:"name,omitempty"`
	AccountID int           `json:"accountId"`
	Account   *ShortAccount `json:"account"`
}
