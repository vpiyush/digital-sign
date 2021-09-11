package endpoints

import (
	"context"
	"errors"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/vpiyush/digital-sign/internal"
	"github.com/vpiyush/digital-sign/pkg/watermark"
)

type EP struct {
	GetDocumentEP   endpoint.Endpoint
	AddDocumentEP   endpoint.Endpoint
	StatusEP        endpoint.Endpoint
	ServiceStatusEP endpoint.Endpoint
	WatermarkEP     endpoint.Endpoint
}

func NewEndpointSet(svc watermark.Service) EP {
	return EP{
		GetDocumentEP:   MakeGetDocumentsEndpoint(svc),
		AddDocumentEP:   MakeAddDocumentEndpoint(svc),
		StatusEP:        MakeStatusEndpoint(svc),
		ServiceStatusEP: MakeServiceStatusEndpoint(svc),
		WatermarkEP:     MakeWatermarkEndpoint(svc),
	}
}

func MakeGetDocumentsEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDocumentRequest)
		docs, err := svc.GetDocuments(ctx, req.Filters...)
		if err != nil {
			return GetDocumentResponse{docs, err.Error()}, nil
		}
		return GetDocumentResponse{docs, ""}, nil
	}
}
func MakeStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StatusRequest)
		status, err := svc.Status(ctx, req.TicketID)
		if err != nil {
			return StatusResponse{Status: status, Err: err.Error()}, nil
		}
		return StatusResponse{Status: status, Err: ""}, nil
	}
}

func MakeAddDocumentEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddDocumentRequest)
		ticketID, err := svc.AddDocument(ctx, req.Document)
		if err != nil {
			return AddDocumentResponse{TicketID: ticketID, Err: err.Error()}, nil
		}
		return AddDocumentResponse{TicketID: ticketID, Err: ""}, nil
	}
}

func MakeWatermarkEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WatermarkRequest)
		code, err := svc.Watermark(ctx, req.TicketID, req.Mark)
		if err != nil {
			return WatermarkResponse{Code: code, Err: err.Error()}, nil
		}
		return WatermarkResponse{Code: code, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *EP) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetDocumentEP(ctx, GetDocumentRequest{Filters: filters})
	if err != nil {
		return []internal.Document{}, err
	}
	getResp := resp.(GetDocumentResponse)
	if getResp.Err != "" {
		return []internal.Document{}, errors.New(getResp.Err)
	}
	return getResp.Documents, nil
}

func (s *EP) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEP(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

func (s *EP) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddDocumentEP(ctx, AddDocumentRequest{Document: doc})
	if err != nil {
		return "", err
	}
	adResp := resp.(AddDocumentResponse)
	if adResp.Err != "" {
		return "", errors.New(adResp.Err)
	}
	return adResp.TicketID, nil
}

func (s *EP) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	resp, err := s.StatusEP(ctx, StatusRequest{TicketID: ticketID})
	if err != nil {
		return internal.Failed, err
	}
	stsResp := resp.(StatusResponse)
	if stsResp.Err != "" {
		return internal.Failed, errors.New(stsResp.Err)
	}
	return stsResp.Status, nil
}

func (s *EP) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	resp, err := s.WatermarkEP(ctx, WatermarkRequest{TicketID: ticketID, Mark: mark})
	wmResp := resp.(WatermarkResponse)
	if err != nil {
		return wmResp.Code, err
	}
	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}
	return wmResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
