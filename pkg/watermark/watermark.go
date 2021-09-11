package watermark

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/lithammer/shortuuid/v3"
	"github.com/vpiyush/digital-sign/internal"
)

type watermark struct{}

func (w *watermark) GetDocuments(ctx context.Context, filter ...internal.Filter) ([]internal.Document, error) {
	// TODO: Query database using filter and return the list of documents
	doc := internal.Document{
		Content: "book",
		Title:   "the code",
		Author:  "bruce",
		Topic:   "Science",
	}
	return []internal.Document{doc}, nil
}

func (w *watermark) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	// TODO: Query database to get current status
	return internal.InProgress, nil
}

func (w *watermark) Watermark(ctx context.Context, ticketID string, mark string) (int, error) {
	// TODO: fetch the database entry using ticket id and update
	// the water mark field
	// Check if the watermarking status is already in progress or started or finished
	// return error in such case
	// return error if no entry found using the ticket ID
	return http.StatusOK, nil
}

func (w *watermark) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	// TODO: Add the document in database, generated ticket ID with the document
	newTicketID := shortuuid.New()
	return newTicketID, nil
}

func (w *watermark) ServicStatus(ctx context.Context) (int, error) {
	logger.Log("Checking the service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
