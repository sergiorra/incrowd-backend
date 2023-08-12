package article

import (
	"context"
	"fmt"

	"incrowd-backend/domain/models"
	"incrowd-backend/infrastructure/database"
	"incrowd-backend/internal/common"
	"incrowd-backend/internal/context_wrapper"
	"incrowd-backend/internal/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type articleRespositry struct {
	client *mongo.Client
}

func NewArticleRespository(client *mongo.Client) *articleRespositry {
	return &articleRespositry{
		client: client,
	}
}

// GetByTeamIDAndID retrieves an article given a teamID (Collection) and an article ID
func (r *articleRespositry) GetByTeamIDAndID(ctx context.Context, teamID string, id string) (*models.Article, error) {
	collection := r.client.Database(database.IncrowdDB).Collection(teamID)

	filter := bson.M{"_id": id}
	var result Article

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, models.ErrArticleIDNotFound
		}
		log.Errort(context_wrapper.GetCorrelationID(ctx), fmt.Sprintf("could not find article ID %s in collection %s with error %s", id, teamID, err))
		return nil, err
	}

	return mapArticleToDomainModel(&result), nil
}

// ListByTeamID retrieves a list of articles given a teamID (Collection). It allows adding pagination and sorting to the result.
func (r *articleRespositry) ListByTeamID(ctx context.Context, md models.MetaData, teamID string) ([]*models.Article, error) {
	collection := r.client.Database(database.IncrowdDB).Collection(teamID)

	filter := bson.M{}
	skip := int64(md.Page * md.Count)
	opts := options.Find().SetSkip(skip).SetLimit(int64(md.Count))

	order := 1
	if md.Order == "desc" {
		order = -1
	}

	opts.SetSort(bson.M{md.Sort: order})

	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var articles []Article
	err = cur.All(ctx, &articles)
	if err != nil {
		return nil, err
	}

	return mapArticleListToDomainModel(articles), nil
}

// Upsert inserts or updates a list of articles into a collection. It also creates the collection if it does not exist yet.
func (r *articleRespositry) Upsert(ctx context.Context, teamID string, articles []*models.Article) error {
	collectionNames, err := r.GetCollectionNames(ctx)
	if err != nil {
		return fmt.Errorf("failed to get collection names with error %s", err)
	}

	if !common.Contains(collectionNames, teamID) {
		if err = r.client.Database(database.IncrowdDB).CreateCollection(ctx, teamID); err != nil {
			return fmt.Errorf("failed to create collection with error %s", err)
		}
	}

	collection := r.client.Database(database.IncrowdDB).Collection(teamID)

	var writes []mongo.WriteModel

	for _, article := range articles {
		filter := bson.D{{"_id", article.ID}}

		update := bson.D{
			{"$set", mapArticleToMongoDbModel(article)},
		}

		model := mongo.NewUpdateOneModel()
		model.Filter = filter
		model.Update = update
		model.Upsert = common.BoolToPointer(true)

		writes = append(writes, model)
	}

	opts := options.BulkWrite().SetOrdered(false)
	_, err = collection.BulkWrite(ctx, writes, opts)
	if err != nil {
		return fmt.Errorf("bulk write failed due to error %s", err)
	}

	return nil
}

// GetCollectionNames retrieves all existing collection names
func (r *articleRespositry) GetCollectionNames(ctx context.Context) ([]string, error) {
	collectionNames, err := r.client.Database(database.IncrowdDB).ListCollectionNames(ctx, bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf("could not list all collection names from DB %s with error %s", database.IncrowdDB, err)
	}

	return collectionNames, nil
}
