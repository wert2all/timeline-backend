// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AddTimeline struct {
	Name *string `json:"name,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type ShortUserTimeline struct {
	ID   int     `json:"id"`
	Name *string `json:"name,omitempty"`
}

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name"`
	Email     string               `json:"email"`
	Avatar    *string              `json:"avatar,omitempty"`
	IsNew     bool                 `json:"isNew"`
	Timelines []*ShortUserTimeline `json:"timelines"`
}
