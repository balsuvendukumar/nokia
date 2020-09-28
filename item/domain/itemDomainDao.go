package domain

import (
	"context"
	"net/http"
	"time"

	"github.com/balsuvendukumar/item/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type increamenter struct {
	ID     string `bson:"_id"`
	SeqVal int    `bson:"seqVal"`
}

func (i increamenter) getNextItemID() int {
	collection := client.Database("item").Collection("increamenter")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	filter := bson.M{"_id": "items"}
	update := bson.M{"$inc": bson.M{"seqVal": 1}}
	_ = collection.FindOneAndUpdate(ctx, filter, update).Decode(&i)
	return i.SeqVal + 1
}

var client *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb")
	client, _ = mongo.Connect(ctx, clientOptions)
	//fmt.Println(client)

}

func InsertItem(r ItemDetail) (*ItemDetail, *util.ItemError) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := client.Database("item").Collection("itemDetails")
	var i increamenter
	r.ItemID = i.getNextItemID()
	res, err := collection.InsertOne(ctx, r)

	if err != nil {
		er := &util.ItemError{
			http.StatusInternalServerError,
			"data not inserted",
		}
		return nil, er

	}

	filter := bson.M{"_id": bson.M{"$eq": res.InsertedID}}
	collection.FindOne(ctx, filter).Decode(&r)

	return &r, nil
}
func FetchOneItem(key int) (*ItemDetail, *util.ItemError) {
	var gItem ItemDetail
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := client.Database("item").Collection("itemDetails")
	filter := bson.M{"item_id": bson.M{"$eq": key}}

	err := collection.FindOne(ctx, filter).Decode(&gItem)
	if err != nil {
		er := &util.ItemError{
			http.StatusNotFound, "the entered key is not found on DB",
		}
		return nil, er
	}
	return &gItem, nil

}
