// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/tweets.proto

package apiconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	api "github.com/codesmith-dev/twitter/internal/gen/api"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// TweetServiceName is the fully-qualified name of the TweetService service.
	TweetServiceName = "api.TweetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TweetServiceGetTweetProcedure is the fully-qualified name of the TweetService's GetTweet RPC.
	TweetServiceGetTweetProcedure = "/api.TweetService/GetTweet"
	// TweetServiceListTweetsProcedure is the fully-qualified name of the TweetService's ListTweets RPC.
	TweetServiceListTweetsProcedure = "/api.TweetService/ListTweets"
	// TweetServiceCreateTweetProcedure is the fully-qualified name of the TweetService's CreateTweet
	// RPC.
	TweetServiceCreateTweetProcedure = "/api.TweetService/CreateTweet"
	// TweetServiceUpdateTweetProcedure is the fully-qualified name of the TweetService's UpdateTweet
	// RPC.
	TweetServiceUpdateTweetProcedure = "/api.TweetService/UpdateTweet"
	// TweetServiceDeleteTweetProcedure is the fully-qualified name of the TweetService's DeleteTweet
	// RPC.
	TweetServiceDeleteTweetProcedure = "/api.TweetService/DeleteTweet"
)

// TweetServiceClient is a client for the api.TweetService service.
type TweetServiceClient interface {
	// GetTweet gets a tweet by id.
	GetTweet(context.Context, *connect.Request[api.GetTweetRequest]) (*connect.Response[api.Tweet], error)
	// ListTweets lists the current tweets of the system.
	ListTweets(context.Context, *connect.Request[api.ListTweetRequest]) (*connect.Response[api.ListTweetResponse], error)
	// CreateTweet creates a new tweet.
	CreateTweet(context.Context, *connect.Request[api.CreateTweetRequest]) (*connect.Response[api.Tweet], error)
	// UpdateTweet update a tweet.
	UpdateTweet(context.Context, *connect.Request[api.UpdateTweetRequest]) (*connect.Response[api.Tweet], error)
	// DeleteTweet deletes a tweet.
	DeleteTweet(context.Context, *connect.Request[api.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewTweetServiceClient constructs a client for the api.TweetService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTweetServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TweetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tweetServiceClient{
		getTweet: connect.NewClient[api.GetTweetRequest, api.Tweet](
			httpClient,
			baseURL+TweetServiceGetTweetProcedure,
			opts...,
		),
		listTweets: connect.NewClient[api.ListTweetRequest, api.ListTweetResponse](
			httpClient,
			baseURL+TweetServiceListTweetsProcedure,
			opts...,
		),
		createTweet: connect.NewClient[api.CreateTweetRequest, api.Tweet](
			httpClient,
			baseURL+TweetServiceCreateTweetProcedure,
			opts...,
		),
		updateTweet: connect.NewClient[api.UpdateTweetRequest, api.Tweet](
			httpClient,
			baseURL+TweetServiceUpdateTweetProcedure,
			opts...,
		),
		deleteTweet: connect.NewClient[api.DeleteTweetRequest, emptypb.Empty](
			httpClient,
			baseURL+TweetServiceDeleteTweetProcedure,
			opts...,
		),
	}
}

// tweetServiceClient implements TweetServiceClient.
type tweetServiceClient struct {
	getTweet    *connect.Client[api.GetTweetRequest, api.Tweet]
	listTweets  *connect.Client[api.ListTweetRequest, api.ListTweetResponse]
	createTweet *connect.Client[api.CreateTweetRequest, api.Tweet]
	updateTweet *connect.Client[api.UpdateTweetRequest, api.Tweet]
	deleteTweet *connect.Client[api.DeleteTweetRequest, emptypb.Empty]
}

// GetTweet calls api.TweetService.GetTweet.
func (c *tweetServiceClient) GetTweet(ctx context.Context, req *connect.Request[api.GetTweetRequest]) (*connect.Response[api.Tweet], error) {
	return c.getTweet.CallUnary(ctx, req)
}

// ListTweets calls api.TweetService.ListTweets.
func (c *tweetServiceClient) ListTweets(ctx context.Context, req *connect.Request[api.ListTweetRequest]) (*connect.Response[api.ListTweetResponse], error) {
	return c.listTweets.CallUnary(ctx, req)
}

// CreateTweet calls api.TweetService.CreateTweet.
func (c *tweetServiceClient) CreateTweet(ctx context.Context, req *connect.Request[api.CreateTweetRequest]) (*connect.Response[api.Tweet], error) {
	return c.createTweet.CallUnary(ctx, req)
}

// UpdateTweet calls api.TweetService.UpdateTweet.
func (c *tweetServiceClient) UpdateTweet(ctx context.Context, req *connect.Request[api.UpdateTweetRequest]) (*connect.Response[api.Tweet], error) {
	return c.updateTweet.CallUnary(ctx, req)
}

// DeleteTweet calls api.TweetService.DeleteTweet.
func (c *tweetServiceClient) DeleteTweet(ctx context.Context, req *connect.Request[api.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteTweet.CallUnary(ctx, req)
}

// TweetServiceHandler is an implementation of the api.TweetService service.
type TweetServiceHandler interface {
	// GetTweet gets a tweet by id.
	GetTweet(context.Context, *connect.Request[api.GetTweetRequest]) (*connect.Response[api.Tweet], error)
	// ListTweets lists the current tweets of the system.
	ListTweets(context.Context, *connect.Request[api.ListTweetRequest]) (*connect.Response[api.ListTweetResponse], error)
	// CreateTweet creates a new tweet.
	CreateTweet(context.Context, *connect.Request[api.CreateTweetRequest]) (*connect.Response[api.Tweet], error)
	// UpdateTweet update a tweet.
	UpdateTweet(context.Context, *connect.Request[api.UpdateTweetRequest]) (*connect.Response[api.Tweet], error)
	// DeleteTweet deletes a tweet.
	DeleteTweet(context.Context, *connect.Request[api.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewTweetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTweetServiceHandler(svc TweetServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tweetServiceGetTweetHandler := connect.NewUnaryHandler(
		TweetServiceGetTweetProcedure,
		svc.GetTweet,
		opts...,
	)
	tweetServiceListTweetsHandler := connect.NewUnaryHandler(
		TweetServiceListTweetsProcedure,
		svc.ListTweets,
		opts...,
	)
	tweetServiceCreateTweetHandler := connect.NewUnaryHandler(
		TweetServiceCreateTweetProcedure,
		svc.CreateTweet,
		opts...,
	)
	tweetServiceUpdateTweetHandler := connect.NewUnaryHandler(
		TweetServiceUpdateTweetProcedure,
		svc.UpdateTweet,
		opts...,
	)
	tweetServiceDeleteTweetHandler := connect.NewUnaryHandler(
		TweetServiceDeleteTweetProcedure,
		svc.DeleteTweet,
		opts...,
	)
	return "/api.TweetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TweetServiceGetTweetProcedure:
			tweetServiceGetTweetHandler.ServeHTTP(w, r)
		case TweetServiceListTweetsProcedure:
			tweetServiceListTweetsHandler.ServeHTTP(w, r)
		case TweetServiceCreateTweetProcedure:
			tweetServiceCreateTweetHandler.ServeHTTP(w, r)
		case TweetServiceUpdateTweetProcedure:
			tweetServiceUpdateTweetHandler.ServeHTTP(w, r)
		case TweetServiceDeleteTweetProcedure:
			tweetServiceDeleteTweetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTweetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTweetServiceHandler struct{}

func (UnimplementedTweetServiceHandler) GetTweet(context.Context, *connect.Request[api.GetTweetRequest]) (*connect.Response[api.Tweet], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.TweetService.GetTweet is not implemented"))
}

func (UnimplementedTweetServiceHandler) ListTweets(context.Context, *connect.Request[api.ListTweetRequest]) (*connect.Response[api.ListTweetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.TweetService.ListTweets is not implemented"))
}

func (UnimplementedTweetServiceHandler) CreateTweet(context.Context, *connect.Request[api.CreateTweetRequest]) (*connect.Response[api.Tweet], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.TweetService.CreateTweet is not implemented"))
}

func (UnimplementedTweetServiceHandler) UpdateTweet(context.Context, *connect.Request[api.UpdateTweetRequest]) (*connect.Response[api.Tweet], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.TweetService.UpdateTweet is not implemented"))
}

func (UnimplementedTweetServiceHandler) DeleteTweet(context.Context, *connect.Request[api.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.TweetService.DeleteTweet is not implemented"))
}
