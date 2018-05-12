package main

import (
	"context"
	"os"
	"time"

	"github.com/nylar/triage/pkg/timeutil"
	"github.com/nylar/triage/ticket/ticketpb"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func (cl *client) listTicket() func(*cli.Context) error {
	return func(c *cli.Context) error {
		resp, err := cl.ticketService.List(context.Background(), &ticketpb.ListRequest{})
		if err != nil {
			log.Fatal().Err(err).Msg("Couldn't list tickets")
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers(&ticketpb.Ticket{}))

		for _, ticket := range resp.Tickets {
			table.Append([]string{
				ticket.Id,
				ticket.Subject,
				timeutil.TimestampToTime(ticket.CreatedAt).Format(time.RFC3339),
				timeutil.TimestampToTime(ticket.UpdatedAt).Format(time.RFC3339),
			})
		}
		table.Render()

		return nil
	}
}

func (cl *client) createTicket() func(*cli.Context) error {
	return func(c *cli.Context) error {
		subject := c.String("subject")
		if subject == "" {
			log.Fatal().Msg("Subject can't be empty")
		}

		resp, err := cl.ticketService.Create(context.Background(), &ticketpb.CreateRequest{
			Subject: c.String("subject"),
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Couldn't create ticket")
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers(&ticketpb.Ticket{}))

		table.Append([]string{
			resp.Ticket.Id,
			resp.Ticket.Subject,
			timeutil.TimestampToTime(resp.Ticket.CreatedAt).Format(time.RFC3339),
			timeutil.TimestampToTime(resp.Ticket.UpdatedAt).Format(time.RFC3339),
		})
		table.Render()

		return nil
	}
}
