package asset

import (
	"fmt"
	"time"

	"github.com/monkeydioude/hako-back/pkg/mongo"
)

type Image struct {
	ID           string `bson:"id,omitempty" json:"id"`
	Type         string `bson:"type,omitempty" json:"type"`
	URL          string `bson:"url,omitempty" json:"url"`
	DateCreation int64  `bson:"date_creation,omitempty" json:"date_creation"`
	UserID       string `bson:"user_id,omitempty" json:"user_id"`
}

func NewImage(name, uid, url string) *Image {
	now := time.Now().Unix()
	return &Image{
		ID:           fmt.Sprintf("%s%s%d", uid, name, now),
		Type:         "image",
		URL:          url,
		DateCreation: now,
		UserID:       uid,
	}
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
