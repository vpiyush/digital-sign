package watermark

import (
	"context"

	"github.com/vpiyush/digital-sign/internal"
)

type Service interface {
	GetDocuments(ctx context.Context, filter ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
