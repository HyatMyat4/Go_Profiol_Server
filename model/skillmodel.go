package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type skills struct {
	Id         primitive.ObjectID `bson:"_id"`
	Skill_name *string            `json:"skill_name" validate:"required, min=2, max =100"`
	Image      *string            `json:"image"`
	Width      *uint8             `json:"width"`
	Hight      *uint8             `json:"height"`
}
