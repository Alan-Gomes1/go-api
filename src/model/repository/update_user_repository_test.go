package repository

import (
	"os"
	"testing"

	"github.com/Alan-Gomes1/go-api/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	os.Setenv(MONGODB_USER_COLLECTION, collectionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDB.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain("test@example.com", "test", "jhon", 18)
		domain.SetID(primitive.NewObjectID().Hex())
		err := repo.UpdateUser(domain.GetID(), domain)
		assert.Nil(t, err)
	})
}
