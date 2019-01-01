package asset

import (
	"fmt"
	"time"

	"github.com/monkeydioude/hako-back/pkg/mongo"
)

type Image struct {
	ID           string `bson:"id,omitempty" json:"id"`
	Type         string `bson:"type,omitempty" json:"type"`
	URL          string `json:"url,omitempty"`
	DateCreation int64  `bson:"date_creation,omitempty" json:"date_creation"`
	UserID       string `bson:"user_id,omitempty" json:"user_id"`
}

func NewImage(UID, ID string) *Image {
	return &Image{
		ID:           ID,
		Type:         "image",
		DateCreation: time.Now().Unix(),
		UserID:       UID,
	}
}

func GenerateImageURL(imgBaseURL, userID, ID string) string {
	return fmt.Sprintf("%s%s/%s", imgBaseURL, userID, ID)
}

func (i *Image) GetType() string {
	return i.Type
}

func (i *Image) GetURL() string {
	return i.URL
}

func (i *Image) GetID() string {
	return i.ID
}

func (i *Image) GetUserID() string {
	return i.UserID
}

func (i *Image) GetDateCreation() int64 {
	return i.DateCreation
}

func (i *Image) Store(db *mongo.DB) (interface{}, error) {
	return db.Collection("asset").InsertOne(i)
}

func (i *Image) Spawn() mongo.Spawnable {
	return &Image{}
}

func (i *Image) GenerateUrl(imgBaseURL string) string {
	return fmt.Sprintf("%s%s/%s", imgBaseURL, i.UserID, i.ID)
}
