package service

import (
	"chat-demo/conf"
	"chat-demo/model/ws"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type SendSortMsg struct {
	Contend  string `json:"contend"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

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

func FindMany(database, sendID, id string, time int64, pageSize int64) (results []ws.Result, err error) {
	var resultMe []ws.Trainer  //id
	var resultYou []ws.Trainer //sendID
	sendIDCollection := conf.MongoDBClient.Database(database).Collection(sendID)
	idCollection := conf.MongoDBClient.Database(database).Collection(id)
	sendIDTimeCurcor, err := sendIDCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(pageSize))
	idTimeCurcor, err := idCollection.Find(context.TODO(),
		options.Find().SetSort(bson.D{{"startTime", -1}}),
		options.Find().SetLimit(pageSize))
	err = sendIDTimeCurcor.All(context.TODO(), &resultYou)
	err = idTimeCurcor.All(context.TODO(), &resultMe)
	results, _ = AppendAndSort(resultMe, resultYou)
	return
}

func AppendAndSort(resultMe, resultYou []ws.Trainer) (results []ws.Result, err error) {
	for _, r := range resultMe {
		sendSort := SendSortMsg{ //构造放回的msg
			Contend:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{ //构造返回所有内容，包括传送者
			StartTime: r.StartTime,
			Msg:       fmt.Sprint("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}

	for _, r := range resultYou {
		sendSort := SendSortMsg{ //构造放回的msg
			Contend:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{ //构造返回所有内容，包括传送者
			StartTime: r.StartTime,
			Msg:       fmt.Sprint("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}
	return
}
