package cmd

import (
	"context"
	"fmt"

	"github.com/DavidHuie/gomigrate"
	"github.com/gocraft/dbr"
	"github.com/ozonebg/gofluence/internal/config"
	"github.com/ozonebg/gofluence/internal/db"
	"github.com/ozonebg/gofluence/pkg/protocol/grpc"
	"github.com/ozonebg/gofluence/pkg/protocol/rest"
	api "github.com/ozonebg/gofluence/pkg/services"
	log "github.com/sirupsen/logrus"
)

var (
	adapter gomigrate.Migratable
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%v:%v@db/%v?sslmode=disable", config.GetDBUser(), config.GetDBPassword(), config.GetDBName())
	log.WithField("dsn", dsn).Info("connection string")
	session := db.CreateDbConnection(dsn)
	defer session.Close()
	defer session.Connection.Close()

	err := performMigration(session)
	if err != nil {
		log.WithError(err).Fatal("failed to run migrations")
	}

	userapi := api.NewUserServiceServer()

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, config.GetGRPCPort(), config.GetHTTPPort())
	}()

	return grpc.RunServer(ctx, userapi, config.GetGRPCPort())
}

func performMigration(sess *dbr.Session) error {
	adapter = gomigrate.Postgres{}
	migrator, err := gomigrate.NewMigratorWithLogger(sess.DB, adapter, "./migrations", log.New())
	if err != nil {
		return err
	}

	err = migrator.Migrate()
	if err != nil {
		return err
	}

	return nil
}
