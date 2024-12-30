// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Account struct {
	ID       int                `json:"id"`
	Name     *string            `json:"name,omitempty"`
	Avatar   *string            `json:"avatar,omitempty"`
	Settings []*AccountSettings `json:"settings"`
}

type AccountSettingInput struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}

type AccountSettings struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AddTimeline struct {
	Name      *string `json:"name,omitempty"`
	AccountID int     `json:"accountId"`
}

type ExistTimelineEventInput struct {
	ID               int           `json:"id"`
	TimelineID       int           `json:"timelineId"`
	Date             time.Time     `json:"date"`
	Type             *TimelineType `json:"type,omitempty"`
	Title            *string       `json:"title,omitempty"`
	Description      *string       `json:"description,omitempty"`
	ShowTime         *bool         `json:"showTime,omitempty"`
	URL              *string       `json:"url,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	PreviewlyImageID *int          `json:"previewlyImageId,omitempty"`
}

type Limit struct {
	From *int `json:"from,omitempty"`
	To   *int `json:"to,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type ShortAccount struct {
	ID             int                `json:"id"`
	Name           *string            `json:"name,omitempty"`
	PreviewlyToken string             `json:"previewlyToken"`
	Avatar         *string            `json:"avatar,omitempty"`
	Settings       []*AccountSettings `json:"settings"`
}

type ShortTimeline struct {
	ID   int     `json:"id"`
	Name *string `json:"name,omitempty"`
}

type TimelineEvent struct {
	ID          int          `json:"id"`
	Date        time.Time    `json:"date"`
	Type        TimelineType `json:"type"`
	Title       *string      `json:"title,omitempty"`
	Description *string      `json:"description,omitempty"`
	ShowTime    *bool        `json:"showTime,omitempty"`
	URL         *string      `json:"url,omitempty"`
	Tags        []string     `json:"tags"`
}

type TimelineEventInput struct {
	ID               *int          `json:"id,omitempty"`
	TimelineID       int           `json:"timelineId"`
	Date             time.Time     `json:"date"`
	Type             *TimelineType `json:"type,omitempty"`
	Title            *string       `json:"title,omitempty"`
	Description      *string       `json:"description,omitempty"`
	ShowTime         *bool         `json:"showTime,omitempty"`
	URL              *string       `json:"url,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	PreviewlyImageID *int          `json:"previewlyImageId,omitempty"`
}

type User struct {
	ID       int             `json:"id"`
	Name     *string         `json:"name,omitempty"`
	Email    string          `json:"email"`
	Avatar   *string         `json:"avatar,omitempty"`
	IsNew    bool            `json:"isNew"`
	Accounts []*ShortAccount `json:"accounts"`
}

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

var AllStatus = []Status{
	StatusSuccess,
	StatusError,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusSuccess, StatusError:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TimelineType string

const (
	TimelineTypeDefault   TimelineType = "default"
	TimelineTypeSelebrate TimelineType = "selebrate"
)

var AllTimelineType = []TimelineType{
	TimelineTypeDefault,
	TimelineTypeSelebrate,
}

func (e TimelineType) IsValid() bool {
	switch e {
	case TimelineTypeDefault, TimelineTypeSelebrate:
		return true
	}
	return false
}

func (e TimelineType) String() string {
	return string(e)
}

func (e *TimelineType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TimelineType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TimelineType", str)
	}
	return nil
}

func (e TimelineType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
