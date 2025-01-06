package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
	Expiration  int64                           `bson:"expiration"`
}
type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func getAuctionDuration() (time.Duration, error) {
	durationStr := os.Getenv("AUCTION_DURATION")
	if durationStr == "" {
		return 0, fmt.Errorf("AUCTION_DURATION environment variable not set")
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return 0, fmt.Errorf("invalid AUCTION_DURATION value: %v", err)
	}

	return time.Duration(duration) * time.Second, nil
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionDuration, err := getAuctionDuration()
	if err != nil {
		logger.Error("Error retrieving auction duration", err)
		return internal_error.NewInternalServerError("Error retrieving auction duration")
	}

	expirationTime := time.Now().Add(auctionDuration).Unix()

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
		Expiration:  expirationTime,
	}

	_, err = ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	return nil
}

func (ar *AuctionRepository) MonitorExpiredAuctions(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute) // A cada minuto
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now().Unix()
			filter := bson.M{"expiration": bson.M{"$lt": now}, "status": bson.M{"$ne": "closed"}}
			update := bson.M{"$set": bson.M{"status": "closed"}}

			_, err := ar.Collection.UpdateMany(ctx, filter, update)
			if err != nil {
				logger.Error("Error updating expired auctions", err)
			}
		}
	}
}
