package main

import (
	"context"
	"log"
	"net/http"
)

type haberdasher struct{}

func (h *haberdasher) ListHats(ctx context.Context, r *ListHatsRequest) (*ListHatsResponse, error) {
	log.Printf("list %+v", r)
	return &ListHatsResponse{}, nil
}

func (h *haberdasher) GetHat(ctx context.Context, r *GetHatRequest) (*GetHatResponse, error) {
	log.Printf("get %+v", r)
	return &GetHatResponse{}, nil
}

func (h *haberdasher) CreateHat(ctx context.Context, r *CreateHatRequest) (*CreateHatResponse, error) {
	log.Printf("create %+v", r)
	return &CreateHatResponse{}, nil
}

func (h *haberdasher) UpdateHat(ctx context.Context, r *UpdateHatRequest) (*UpdateHatResponse, error) {
	log.Printf("update %+v", r)
	return &UpdateHatResponse{}, nil
}

func (h *haberdasher) DeleteHat(ctx context.Context, r *DeleteHatRequest) (*DeleteHatResponse, error) {
	log.Printf("delete %+v", r)
	return &DeleteHatResponse{}, nil
}

func main() {
	gateway := HaberdasherGateway()
	server := NewHaberdasherServer(&haberdasher{}, nil)
	log.Fatal(http.ListenAndServe(":8080", gateway(server)))
}
