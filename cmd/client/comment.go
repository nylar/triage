package main

import (
	"context"
	"os"
	"time"

	"github.com/nylar/triage/comment/commentpb"
	"github.com/nylar/triage/pkg/timeutil"
	"github.com/olekukonko/tablewriter"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func (cl *client) listComment() func(*cli.Context) error {
	return func(c *cli.Context) error {
		ticketID := c.String("ticket_id")
		if ticketID == "" {
			log.Fatal().Msg("Ticket ID can't be empty")
		}

		resp, err := cl.commentService.List(context.Background(), &commentpb.ListRequest{
			TicketId: ticketID,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Couldn't list comments")
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers(&commentpb.Comment{}))
		table.SetAutoMergeCells(true)

		for _, comment := range resp.Comments {
			table.Append([]string{
				comment.Id,
				comment.TicketId,
				comment.Content,
				timeutil.TimestampToTime(comment.CreatedAt).Format(time.RFC3339),
				timeutil.TimestampToTime(comment.UpdatedAt).Format(time.RFC3339),
			})
		}
		table.Render()

		return nil
	}
}

func (cl *client) createComment() func(*cli.Context) error {
	return func(c *cli.Context) error {
		ticketID := c.String("ticket_id")
		if ticketID == "" {
			log.Fatal().Msg("Ticket ID can't be empty")
		}

		content := c.String("content")
		if content == "" {
			log.Fatal().Msg("Content can't be empty")
		}

		resp, err := cl.commentService.Create(context.Background(), &commentpb.CreateRequest{
			TicketId: ticketID,
			Content:  content,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Couldn't create comment")
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers(&commentpb.Comment{}))
		table.SetAutoMergeCells(true)

		table.Append([]string{
			resp.Comment.Id,
			resp.Comment.TicketId,
			resp.Comment.Content,
			timeutil.TimestampToTime(resp.Comment.CreatedAt).Format(time.RFC3339),
			timeutil.TimestampToTime(resp.Comment.UpdatedAt).Format(time.RFC3339),
		})
		table.Render()

		return nil
	}
}
