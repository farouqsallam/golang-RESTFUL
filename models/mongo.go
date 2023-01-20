package models

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	
)
type Video struct{
	Name string
	Desc string
	Size uint32
}

func GetAllVids() []Video{
	coll := Client.Database("Videos").Collection("video")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil{
		log.Fatal("error while getting the document from DB")
	}

	var results []Video

	if err := cursor.All(context.TODO(), &results); err != nil{
		log.Fatal("error monog.go GetAllVids func")
	}

	return results
}

func GetSpecificVid(vidName string) []Video{
	coll := Client.Database("Videos").Collection("video")
	filter := bson.D{{"name", vidName}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil{
		log.Fatal("error while getting the document from DB")
	}

	var results []Video

	if err := cursor.All(context.TODO(), &results); err != nil{
		log.Fatal("error monog.go GetAllVids func")
	}

	return results
}

func InsertAVid(v *Video){
	coll := Client.Database("Videos").Collection("video")

	_, err := coll.InsertOne(context.TODO(), *v)
	if err != nil {
		log.Fatal("error while inserting a video")
	}

}

func DeleteAVid(vidName string) {
	coll := Client.Database("Videos").Collection("video")
	filter := bson.D{{"name", vidName}}

	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil{
		log.Fatal("error deleting")
	}
	
}

func UpdateVid(name string, desc string, size uint32){
	coll := Client.Database("Videos").Collection("video")
	filter := bson.D{{"name", name}}
	update := bson.D{{"$set", bson.D{{"name", name}, {"desc", desc}, {"size", size}}}}

	coll.UpdateMany(context.TODO(), filter, update)
	w.Write([]byte("Video updated"))
}