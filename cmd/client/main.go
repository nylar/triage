package main

import (
	"os"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/nylar/triage/comment/commentpb"
	"github.com/nylar/triage/ticket/ticketpb"
	toml "github.com/pelletier/go-toml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	strcase "github.com/stoewer/go-strcase"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

type client struct {
	ticketService  ticketpb.TicketServiceClient
	commentService commentpb.CommentServiceClient
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	configFile, err := os.Open(os.Getenv("TRIAGE_CLIENT_CONFIG_PATH"))
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't open config file")
	}
	defer configFile.Close()

	conf := &Config{}
	if err := toml.NewDecoder(configFile).Decode(conf); err != nil {
		log.Fatal().Err(err).Msg("Couldn't parse config")
	}

	rpc, err := grpc.Dial(conf.Server.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", conf.Server.Address).
			Msg("Couldn't connect to server")
	}

	cl := client{
		ticketService:  ticketpb.NewTicketServiceClient(rpc),
		commentService: commentpb.NewCommentServiceClient(rpc),
	}

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "ticket",
			Aliases: []string{"t"},
			Usage:   "ticket management",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "Add a new ticket",
					Action: cl.createTicket(),
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "subject",
						},
					},
				},
				{
					Name:   "list",
					Usage:  "List a new ticket",
					Action: cl.listTicket(),
				},
			},
		},
		{
			Name:    "comment",
			Aliases: []string{"c"},
			Usage:   "comment management",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "Add a new comment to a ticket",
					Action: cl.createComment(),
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "content",
						},
						cli.StringFlag{
							Name: "ticket_id",
						},
					},
				},
				{
					Name:   "list",
					Usage:  "List a ticket's comments",
					Action: cl.listComment(),
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "ticket_id",
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Couldn't run app")
	}
}

func headers(msg proto.Message) []string {
	value := reflect.Indirect(reflect.ValueOf(msg)).Type()

	var headers []string

	for i := 0; i < value.NumField(); i++ {
		headers = append(headers, strcase.SnakeCase(value.Field(i).Name))
	}

	return headers
}
