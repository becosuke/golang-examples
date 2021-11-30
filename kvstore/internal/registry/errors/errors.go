package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KVStoreError interface {
	error
	GRPCStatus() *status.Status
}

type kvstoreError struct {
	codes.Code
	msg string
}

func NewKVStoreError(code codes.Code, msg string) *kvstoreError {
	return &kvstoreError{Code: code, msg: msg}
}

func (e *kvstoreError) GRPCStatus() *status.Status {
	return status.New(e.Code, e.msg)
}

func (e *kvstoreError) Error() string {
	return e.String()
}

var (
	ErrSyncmapNotFound        = NewKVStoreError(codes.NotFound, "not exists")
	ErrSyncmapInvalidData     = NewKVStoreError(codes.Internal, "returns invalid data")
	ErrSyncmapInvalidArgument = NewKVStoreError(codes.InvalidArgument, "invalid argument")
)
