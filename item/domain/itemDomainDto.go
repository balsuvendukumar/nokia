package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemDetail struct {
	ID                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name              string             `json:"Name,omitempty" bson:"item,omitempty"`
	Manufacturer      string             `json:"Manufacturer" bson:"manufacturer,omitempty"`
	Price             int                `json:"Price,omitempty" bson:"price,omitempty"`
	ManufacturingDate string             `json:"ManufacturingDate" bson:"manufacturing_date"`
	ExpiryDate        string             `json:"ExpiryDate" bson:"expiry_date"`
	ItemID            int                `json:"item_id,omitempty" bson:"item_id"`
}
