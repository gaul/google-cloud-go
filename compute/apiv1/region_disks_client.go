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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRegionDisksClientHook clientHook

// RegionDisksCallOptions contains the retry settings for each method of RegionDisksClient.
type RegionDisksCallOptions struct {
	AddResourcePolicies    []gax.CallOption
	CreateSnapshot         []gax.CallOption
	Delete                 []gax.CallOption
	Get                    []gax.CallOption
	GetIamPolicy           []gax.CallOption
	Insert                 []gax.CallOption
	List                   []gax.CallOption
	RemoveResourcePolicies []gax.CallOption
	Resize                 []gax.CallOption
	SetIamPolicy           []gax.CallOption
	SetLabels              []gax.CallOption
	TestIamPermissions     []gax.CallOption
}

// internalRegionDisksClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalRegionDisksClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddResourcePolicies(context.Context, *computepb.AddResourcePoliciesRegionDiskRequest, ...gax.CallOption) (*Operation, error)
	CreateSnapshot(context.Context, *computepb.CreateSnapshotRegionDiskRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeleteRegionDiskRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetRegionDiskRequest, ...gax.CallOption) (*computepb.Disk, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyRegionDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertRegionDiskRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListRegionDisksRequest, ...gax.CallOption) *DiskIterator
	RemoveResourcePolicies(context.Context, *computepb.RemoveResourcePoliciesRegionDiskRequest, ...gax.CallOption) (*Operation, error)
	Resize(context.Context, *computepb.ResizeRegionDiskRequest, ...gax.CallOption) (*Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyRegionDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	SetLabels(context.Context, *computepb.SetLabelsRegionDiskRequest, ...gax.CallOption) (*Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsRegionDiskRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// RegionDisksClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The RegionDisks API.
type RegionDisksClient struct {
	// The internal transport-dependent client.
	internalClient internalRegionDisksClient

	// The call options for this service.
	CallOptions *RegionDisksCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RegionDisksClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RegionDisksClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *RegionDisksClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddResourcePolicies adds existing resource policies to a regional disk. You can only add one policy which will be applied to this disk for scheduling snapshot creation.
func (c *RegionDisksClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AddResourcePolicies(ctx, req, opts...)
}

// CreateSnapshot creates a snapshot of a specified persistent disk. For regular snapshot creation, consider using snapshots.insert instead, as that method supports more features, such as creating snapshots in a project different from the source disk project.
func (c *RegionDisksClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.CreateSnapshot(ctx, req, opts...)
}

// Delete deletes the specified regional persistent disk. Deleting a regional disk removes all the replicas of its data permanently and is irreversible. However, deleting a disk does not delete any snapshots previously made from the disk. You must separately delete snapshots.
func (c *RegionDisksClient) Delete(ctx context.Context, req *computepb.DeleteRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns a specified regional persistent disk.
func (c *RegionDisksClient) Get(ctx context.Context, req *computepb.GetRegionDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *RegionDisksClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert creates a persistent regional disk in the specified project using the data included in the request.
func (c *RegionDisksClient) Insert(ctx context.Context, req *computepb.InsertRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of persistent disks contained within the specified region.
func (c *RegionDisksClient) List(ctx context.Context, req *computepb.ListRegionDisksRequest, opts ...gax.CallOption) *DiskIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// RemoveResourcePolicies removes resource policies from a regional disk.
func (c *RegionDisksClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.RemoveResourcePolicies(ctx, req, opts...)
}

// Resize resizes the specified regional persistent disk.
func (c *RegionDisksClient) Resize(ctx context.Context, req *computepb.ResizeRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Resize(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *RegionDisksClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// SetLabels sets the labels on the target regional disk.
func (c *RegionDisksClient) SetLabels(ctx context.Context, req *computepb.SetLabelsRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *RegionDisksClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsRegionDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type regionDisksRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// operationClient is used to call the operation-specific management service.
	operationClient *RegionOperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewRegionDisksRESTClient creates a new region disks rest client.
//
// The RegionDisks API.
func NewRegionDisksRESTClient(ctx context.Context, opts ...option.ClientOption) (*RegionDisksClient, error) {
	clientOpts := append(defaultRegionDisksRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &regionDisksRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	o := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opC, err := NewRegionOperationsRESTClient(ctx, o...)
	if err != nil {
		return nil, err
	}
	c.operationClient = opC

	return &RegionDisksClient{internalClient: c, CallOptions: &RegionDisksCallOptions{}}, nil
}

func defaultRegionDisksRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *regionDisksRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *regionDisksRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	if err := c.operationClient.Close(); err != nil {
		return err
	}
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *regionDisksRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddResourcePolicies adds existing resource policies to a regional disk. You can only add one policy which will be applied to this disk for scheduling snapshot creation.
func (c *regionDisksRESTClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetRegionDisksAddResourcePoliciesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/addResourcePolicies", req.GetProject(), req.GetRegion(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// CreateSnapshot creates a snapshot of a specified persistent disk. For regular snapshot creation, consider using snapshots.insert instead, as that method supports more features, such as creating snapshots in a project different from the source disk project.
func (c *regionDisksRESTClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetSnapshotResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/createSnapshot", req.GetProject(), req.GetRegion(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Delete deletes the specified regional persistent disk. Deleting a regional disk removes all the replicas of its data permanently and is irreversible. However, deleting a disk does not delete any snapshots previously made from the disk. You must separately delete snapshots.
func (c *regionDisksRESTClient) Delete(ctx context.Context, req *computepb.DeleteRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v", req.GetProject(), req.GetRegion(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Get returns a specified regional persistent disk.
func (c *regionDisksRESTClient) Get(ctx context.Context, req *computepb.GetRegionDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v", req.GetProject(), req.GetRegion(), req.GetDisk())

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Disk{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *regionDisksRESTClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/getIamPolicy", req.GetProject(), req.GetRegion(), req.GetResource())

	params := url.Values{}
	if req != nil && req.OptionsRequestedPolicyVersion != nil {
		params.Add("optionsRequestedPolicyVersion", fmt.Sprintf("%v", req.GetOptionsRequestedPolicyVersion()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Policy{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// Insert creates a persistent regional disk in the specified project using the data included in the request.
func (c *regionDisksRESTClient) Insert(ctx context.Context, req *computepb.InsertRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetDiskResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks", req.GetProject(), req.GetRegion())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}
	if req != nil && req.SourceImage != nil {
		params.Add("sourceImage", fmt.Sprintf("%v", req.GetSourceImage()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// List retrieves the list of persistent disks contained within the specified region.
func (c *regionDisksRESTClient) List(ctx context.Context, req *computepb.ListRegionDisksRequest, opts ...gax.CallOption) *DiskIterator {
	it := &DiskIterator{}
	req = proto.Clone(req).(*computepb.ListRegionDisksRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.Disk, string, error) {
		resp := &computepb.DiskList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks", req.GetProject(), req.GetRegion())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// RemoveResourcePolicies removes resource policies from a regional disk.
func (c *regionDisksRESTClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetRegionDisksRemoveResourcePoliciesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/removeResourcePolicies", req.GetProject(), req.GetRegion(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// Resize resizes the specified regional persistent disk.
func (c *regionDisksRESTClient) Resize(ctx context.Context, req *computepb.ResizeRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetRegionDisksResizeRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/resize", req.GetProject(), req.GetRegion(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *regionDisksRESTClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyRegionDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetRegionSetPolicyRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/setIamPolicy", req.GetProject(), req.GetRegion(), req.GetResource())

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Policy{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// SetLabels sets the labels on the target regional disk.
func (c *regionDisksRESTClient) SetLabels(ctx context.Context, req *computepb.SetLabelsRegionDiskRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetRegionSetLabelsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/setLabels", req.GetProject(), req.GetRegion(), req.GetResource())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	op := &Operation{
		&regionOperationsHandle{
			c:       c.operationClient,
			proto:   resp,
			project: req.GetProject(),
			region:  req.GetRegion(),
		},
	}
	return op, nil
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *regionDisksRESTClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsRegionDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTestPermissionsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/disks/%v/testIamPermissions", req.GetProject(), req.GetRegion(), req.GetResource())

	// Build HTTP headers from client and context metadata.
	headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.TestPermissionsResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
