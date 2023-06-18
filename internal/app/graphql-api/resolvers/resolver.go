package resolvers

import (
	"github.com/DmitryLogunov/trading-platform/internal/app/graphql-api/gql-services"
	"github.com/DmitryLogunov/trading-platform/internal/core/scheduler"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your cmd, add any dependencies you require here.

type Resolver struct {
	MongoDB     *mongo.Database
	Scheduler   *scheduler.JobsManager
	GqlServices *gqlServices.GqlServices
}
