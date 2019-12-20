package model

type StorageType int8

const (
	Image = iota + 1
	Video
	Text
	Others
)

var StorageTypes = []StorageType{
	Image,
	Video,
	Text,
	Others,
}

func (s StorageType) String() string {
	switch s {
	case Image:
		return "Image"
	case Video:
		return "Video"
	case Text:
		return "Text"
	default:
		return "Others"
	}
}

// Company represents company model
type Storage struct {
	Uuid           string      `json:"uuid"`
	Url            string      `json:"url"`
	Path           string      `json:"path"`
	Type           StorageType `json:"type"`
	OwnerID        string      `json:"-"`
	CustomerNumber string      `json:"-"`
	IsUsed         bool        `json:"-"`
}
