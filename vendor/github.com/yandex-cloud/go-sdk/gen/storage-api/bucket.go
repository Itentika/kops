// Code generated by sdkgen. DO NOT EDIT.

//nolint
package storage

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	storage "github.com/yandex-cloud/go-genproto/yandex/cloud/storage/v1"
)

//revive:disable

// BucketServiceClient is a storage.BucketServiceClient with
// lazy GRPC connection initialization.
type BucketServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements storage.BucketServiceClient
func (c *BucketServiceClient) Create(ctx context.Context, in *storage.CreateBucketRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements storage.BucketServiceClient
func (c *BucketServiceClient) Delete(ctx context.Context, in *storage.DeleteBucketRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteHTTPSConfig implements storage.BucketServiceClient
func (c *BucketServiceClient) DeleteHTTPSConfig(ctx context.Context, in *storage.DeleteBucketHTTPSConfigRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).DeleteHTTPSConfig(ctx, in, opts...)
}

// Get implements storage.BucketServiceClient
func (c *BucketServiceClient) Get(ctx context.Context, in *storage.GetBucketRequest, opts ...grpc.CallOption) (*storage.Bucket, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).Get(ctx, in, opts...)
}

// GetHTTPSConfig implements storage.BucketServiceClient
func (c *BucketServiceClient) GetHTTPSConfig(ctx context.Context, in *storage.GetBucketHTTPSConfigRequest, opts ...grpc.CallOption) (*storage.HTTPSConfig, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).GetHTTPSConfig(ctx, in, opts...)
}

// GetStats implements storage.BucketServiceClient
func (c *BucketServiceClient) GetStats(ctx context.Context, in *storage.GetBucketStatsRequest, opts ...grpc.CallOption) (*storage.BucketStats, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).GetStats(ctx, in, opts...)
}

// List implements storage.BucketServiceClient
func (c *BucketServiceClient) List(ctx context.Context, in *storage.ListBucketsRequest, opts ...grpc.CallOption) (*storage.ListBucketsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).List(ctx, in, opts...)
}

type BucketIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *BucketServiceClient
	request *storage.ListBucketsRequest

	items []*storage.Bucket
}

func (c *BucketServiceClient) BucketIterator(ctx context.Context, req *storage.ListBucketsRequest, opts ...grpc.CallOption) *BucketIterator {
	var pageSize int64
	const defaultPageSize = 1000

	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &BucketIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *BucketIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Buckets
	return len(it.items) > 0
}

func (it *BucketIterator) Take(size int64) ([]*storage.Bucket, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*storage.Bucket

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *BucketIterator) TakeAll() ([]*storage.Bucket, error) {
	return it.Take(0)
}

func (it *BucketIterator) Value() *storage.Bucket {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *BucketIterator) Error() error {
	return it.err
}

// SetHTTPSConfig implements storage.BucketServiceClient
func (c *BucketServiceClient) SetHTTPSConfig(ctx context.Context, in *storage.SetBucketHTTPSConfigRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).SetHTTPSConfig(ctx, in, opts...)
}

// Update implements storage.BucketServiceClient
func (c *BucketServiceClient) Update(ctx context.Context, in *storage.UpdateBucketRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return storage.NewBucketServiceClient(conn).Update(ctx, in, opts...)
}