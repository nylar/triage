package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/nylar/triage/base"
	"github.com/nylar/triage/config"
	"github.com/nylar/triage/pkg/clock"
	"github.com/nylar/triage/ticket"
	"github.com/nylar/triage/ticket/ticketpb"
	toml "github.com/pelletier/go-toml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	major = 0
	minor = 1
	patch = 0
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	log.Info().Str("version", version()).Msg("Triage")

	configFile, err := os.Open(os.Getenv("TRIAGE_CONFIG_PATH"))
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't open config file")
	}
	defer configFile.Close()

	conf := &config.Config{}
	if err := toml.NewDecoder(configFile).Decode(conf); err != nil {
		log.Fatal().Err(err).Msg("Couldn't parse config")
	}

	idGenerator := uuid.NewRandom
	realClock := clock.Real{}

	var ticketService ticket.Service
	if conf.SQL != nil {
		log.Info().Str("vendor", conf.SQL.Vendor).Msg("Using SQL database")

		dsn, err := conf.SQL.DSN()
		if err != nil {
			log.Fatal().Err(err).
				Msg("Couldn't build database connection string")
		}

		db, err := sql.Open(conf.SQL.Vendor, dsn)
		if err != nil {
			log.Fatal().Err(err).Msg("Couldn't open database")
		}

		if err := db.Ping(); err != nil {
			log.Fatal().Err(err).Msg("Couldn't ping database")
		}

		baseSQL := base.SQL{
			DB:          db,
			Placeholder: conf.SQL.Placeholder(),
			IDGenerator: idGenerator,
			Clock:       realClock,
		}

		ticketService = &ticket.SQL{
			SQL: baseSQL,
		}
	} else if conf.Bolt != nil {
		log.Info().Msg("Using bolt database")

		db, err := bolt.Open(conf.Bolt.Path, 0777, nil)
		if err != nil {
			log.Fatal().Err(err).Msg("Couldn't open database")
		}
		defer db.Close()

		baseBolt := base.Bolt{
			DB:          db,
			IDGenerator: idGenerator,
			Clock:       realClock,
		}

		if err := baseBolt.Bootstrap(); err != nil {
			log.Fatal().Err(err).Msg("Couldn't bootstrap database")
		}

		ticketService = &ticket.Bolt{
			Bolt: baseBolt,
		}

	} else {
		log.Fatal().Msg("A datastore has not been configured")
	}

	server := grpc.NewServer()

	ticketpb.RegisterTicketServiceServer(server, ticketService)

	listener, err := net.Listen("tcp", conf.Server.Address())
	if err != nil {
		log.Fatal().Err(err).Str("address", conf.Server.Address()).
			Msg("Couldn't listen to TCP server")
	}

	log.Info().Str("address", conf.Server.Address()).Msg("Started serving")
	if err := server.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("Couldn't start gRPC server")
	}
}

func version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
