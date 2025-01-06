package auction

import (
	"context"
	"testing"
	"time"

	"fullcycle-auction_go/internal/entity/auction_entity" // Importação correta

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionClosure(t *testing.T) {
	// Cria uma conexão com o MongoDB local ou em container para o teste
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Use o banco de dados de teste
	db := client.Database("testdb")
	repo := NewAuctionRepository(db)

	// Cria um leilão com duração de 2 segundos usando a função CreateAuction do auction_entity
	auction, err := auction_entity.CreateAuction(
		"Test Product",
		"Electronics",
		"A cool product",
		auction_entity.New, // Condição do produto
	)
	assert.Nil(t, err)

	// Adiciona o leilão no banco de dados através do repositório
	err = repo.CreateAuction(context.Background(), auction)
	assert.Nil(t, err)

	// Simula o tempo de duração de 2 segundos para o leilão expirar
	time.Sleep(3 * time.Second)

	// Executa a função que verifica e fecha leilões vencidos
	go repo.MonitorExpiredAuctions(context.Background())

	// Verifica se o status do leilão foi atualizado para "closed"
	var result AuctionEntityMongo
	err = db.Collection("auctions").FindOne(context.Background(), bson.M{"_id": auction.Id}).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, "closed", result.Status)
}
