package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"microservice/app/user/service/internal/conf"
	"microservice/app/user/service/internal/data/ent"
	"microservice/app/user/service/internal/data/ent/migrate"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEntClient,
	NewUserRepo,
)

// Data .
type Data struct {
	log *log.Helper
	db  *ent.Client
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log1 := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log1.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log1.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// NewData .
func NewData(
	logger log.Logger,
	entClient *ent.Client,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		log: l,
		db:  entClient,
	}

	return d, func() {
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}
