package model

type ShortAccount struct {
	ID             int                `json:"id"`
	Name           *string            `json:"name,omitempty"`
	PreviewlyToken string             `json:"previewlyToken"`
	AvatarID       *int               `json:"avatarId,omitempty"`
	About          *string            `json:"about,omitempty"`
	Settings       []*AccountSettings `json:"settings"`
}
