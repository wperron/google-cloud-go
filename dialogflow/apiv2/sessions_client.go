// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package dialogflow

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newSessionsClientHook clientHook

// SessionsCallOptions contains the retry settings for each method of SessionsClient.
type SessionsCallOptions struct {
	DetectIntent          []gax.CallOption
	StreamingDetectIntent []gax.CallOption
}

func defaultSessionsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSessionsCallOptions() *SessionsCallOptions {
	return &SessionsCallOptions{
		DetectIntent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		StreamingDetectIntent: []gax.CallOption{},
	}
}

// internalSessionsClient is an interface that defines the methods availaible from Dialogflow API.
type internalSessionsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	DetectIntent(context.Context, *dialogflowpb.DetectIntentRequest, ...gax.CallOption) (*dialogflowpb.DetectIntentResponse, error)
	StreamingDetectIntent(context.Context, ...gax.CallOption) (dialogflowpb.Sessions_StreamingDetectIntentClient, error)
}

// SessionsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service used for session interactions.
//
// For more information, see the API interactions
// guide (at https://cloud.google.com/dialogflow/docs/api-overview).
type SessionsClient struct {
	// The internal transport-dependent client.
	internalClient internalSessionsClient

	// The call options for this service.
	CallOptions *SessionsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SessionsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SessionsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SessionsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// DetectIntent processes a natural language query and returns structured, actionable data
// as a result. This method is not idempotent, because it may cause contexts
// and session entity types to be updated, which in turn might affect
// results of future queries.
//
// If you might use
// Agent Assist (at https://cloud.google.com/dialogflow/docs/#aa)
// or other CCAI products now or in the future, consider using
// AnalyzeContent
// instead of DetectIntent. AnalyzeContent has additional
// functionality for Agent Assist and other CCAI products.
//
// Note: Always use agent versions for production traffic.
// See Versions and
// environments (at https://cloud.google.com/dialogflow/es/docs/agents-versions).
func (c *SessionsClient) DetectIntent(ctx context.Context, req *dialogflowpb.DetectIntentRequest, opts ...gax.CallOption) (*dialogflowpb.DetectIntentResponse, error) {
	return c.internalClient.DetectIntent(ctx, req, opts...)
}

// StreamingDetectIntent processes a natural language query in audio format in a streaming fashion
// and returns structured, actionable data as a result. This method is only
// available via the gRPC API (not REST).
//
// If you might use
// Agent Assist (at https://cloud.google.com/dialogflow/docs/#aa)
// or other CCAI products now or in the future, consider using
// StreamingAnalyzeContent
// instead of StreamingDetectIntent. StreamingAnalyzeContent has
// additional functionality for Agent Assist and other CCAI products.
//
// Note: Always use agent versions for production traffic.
// See Versions and
// environments (at https://cloud.google.com/dialogflow/es/docs/agents-versions).
func (c *SessionsClient) StreamingDetectIntent(ctx context.Context, opts ...gax.CallOption) (dialogflowpb.Sessions_StreamingDetectIntentClient, error) {
	return c.internalClient.StreamingDetectIntent(ctx, opts...)
}

// sessionsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type sessionsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SessionsClient
	CallOptions **SessionsCallOptions

	// The gRPC API client.
	sessionsClient dialogflowpb.SessionsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSessionsClient creates a new sessions client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service used for session interactions.
//
// For more information, see the API interactions
// guide (at https://cloud.google.com/dialogflow/docs/api-overview).
func NewSessionsClient(ctx context.Context, opts ...option.ClientOption) (*SessionsClient, error) {
	clientOpts := defaultSessionsGRPCClientOptions()
	if newSessionsClientHook != nil {
		hookOpts, err := newSessionsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SessionsClient{CallOptions: defaultSessionsCallOptions()}

	c := &sessionsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		sessionsClient:   dialogflowpb.NewSessionsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *sessionsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *sessionsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *sessionsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *sessionsGRPCClient) DetectIntent(ctx context.Context, req *dialogflowpb.DetectIntentRequest, opts ...gax.CallOption) (*dialogflowpb.DetectIntentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 220000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session", url.QueryEscape(req.GetSession())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DetectIntent[0:len((*c.CallOptions).DetectIntent):len((*c.CallOptions).DetectIntent)], opts...)
	var resp *dialogflowpb.DetectIntentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionsClient.DetectIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionsGRPCClient) StreamingDetectIntent(ctx context.Context, opts ...gax.CallOption) (dialogflowpb.Sessions_StreamingDetectIntentClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp dialogflowpb.Sessions_StreamingDetectIntentClient
	opts = append((*c.CallOptions).StreamingDetectIntent[0:len((*c.CallOptions).StreamingDetectIntent):len((*c.CallOptions).StreamingDetectIntent)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionsClient.StreamingDetectIntent(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
