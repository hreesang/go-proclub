package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/hreesang/go-proclub/bot/database"
)

type MatchListener struct {
	ID        bson.ObjectID 	`bson:"_id,omitempty" json:"id"`
	ClubID    string             `bson:"club_id" json:"clubId"` 
	ChannelID string             `bson:"channel_id" json:"channelId"`
}

const (
	collectionName	= "match_listener"
)

func GetMatchListenerFromChannelID(channelID string) (*MatchListener, error) {
	var matchListener *MatchListener

	db, err := database.Setup()
	if err != nil {
		return matchListener, err
	}

	coll := db.Collection(collectionName)
	filter := MatchListener{
		ChannelID: channelID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	if err := coll.FindOne(ctx, filter).Decode(&matchListener); err != nil {
		return matchListener, err
	}

	return matchListener, nil
}
