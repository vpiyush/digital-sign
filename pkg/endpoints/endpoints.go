package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/vpiyush/digital-sign/pkg/watermark"
)

type ep struct {
	GetDocument   endpoint.Endpoint
	AddDocument   endpoint.Endpoint
	Status        endpoint.Endpoint
	ServiceStatus endpoint.Endpoint
	Watermark     endpoint.Endpoint
}

func MakeGetDocumentsEP(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDocumentRequest)
		docs, err := svc.GetDocuments(ctx, req.Filters...)
		if err != nil {
			return GetDocumentResponse{docs, err.Error()}, nil
		}
		return GetDocumentResponse{docs, ""}, nil
	}
}


