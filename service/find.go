package service

import (
	"chat-demo/conf"
	"chat-demo/model/ws"
	"context"
	"time"
)

func InsertMsg(database, id, content string, read uint, expire int64) error {
	collection := conf.MongoDBClient.Database(database).Collection(id)
	comment := ws.Trainer{
		content,
		time.Now().Unix(),
		time.Now().Unix() + expire,
		read,
	}
	_, err := collection.InsertOne(context.TODO(), comment)
	return err
}
