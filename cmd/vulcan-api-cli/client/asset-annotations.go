// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": asset-annotations Resource Client
//
// Command:
// $ main

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateAssetAnnotationsPath computes a request path to the create action of asset-annotations.
func CreateAssetAnnotationsPath(teamID string, assetID string) string {
	param0 := teamID
	param1 := assetID

	return fmt.Sprintf("/api/v1/teams/%s/assets/%s/annotations", param0, param1)
}

// Create one or more annotation for a given asset.
func (c *Client) CreateAssetAnnotations(ctx context.Context, path string, payload *AssetAnnotationRequest) (*http.Response, error) {
	req, err := c.NewCreateAssetAnnotationsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAssetAnnotationsRequest create the request corresponding to the create action endpoint of the asset-annotations resource.
func (c *Client) NewCreateAssetAnnotationsRequest(ctx context.Context, path string, payload *AssetAnnotationRequest) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteAssetAnnotationsPath computes a request path to the delete action of asset-annotations.
func DeleteAssetAnnotationsPath(teamID string, assetID string) string {
	param0 := teamID
	param1 := assetID

	return fmt.Sprintf("/api/v1/teams/%s/assets/%s/annotations", param0, param1)
}

// Delete one or more annotation for a given asset.
func (c *Client) DeleteAssetAnnotations(ctx context.Context, path string, payload *AssetAnnotationDeleteRequest) (*http.Response, error) {
	req, err := c.NewDeleteAssetAnnotationsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAssetAnnotationsRequest create the request corresponding to the delete action endpoint of the asset-annotations resource.
func (c *Client) NewDeleteAssetAnnotationsRequest(ctx context.Context, path string, payload *AssetAnnotationDeleteRequest) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListAssetAnnotationsPath computes a request path to the list action of asset-annotations.
func ListAssetAnnotationsPath(teamID string, assetID string) string {
	param0 := teamID
	param1 := assetID

	return fmt.Sprintf("/api/v1/teams/%s/assets/%s/annotations", param0, param1)
}

// List annotations of a given asset.
func (c *Client) ListAssetAnnotations(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAssetAnnotationsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAssetAnnotationsRequest create the request corresponding to the list action endpoint of the asset-annotations resource.
func (c *Client) NewListAssetAnnotationsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// PutAssetAnnotationsPath computes a request path to the put action of asset-annotations.
func PutAssetAnnotationsPath(teamID string, assetID string) string {
	param0 := teamID
	param1 := assetID

	return fmt.Sprintf("/api/v1/teams/%s/assets/%s/annotations", param0, param1)
}

// Override all annotations with a new list
func (c *Client) PutAssetAnnotations(ctx context.Context, path string, payload *AssetAnnotationRequest) (*http.Response, error) {
	req, err := c.NewPutAssetAnnotationsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPutAssetAnnotationsRequest create the request corresponding to the put action endpoint of the asset-annotations resource.
func (c *Client) NewPutAssetAnnotationsRequest(ctx context.Context, path string, payload *AssetAnnotationRequest) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateAssetAnnotationsPath computes a request path to the update action of asset-annotations.
func UpdateAssetAnnotationsPath(teamID string, assetID string) string {
	param0 := teamID
	param1 := assetID

	return fmt.Sprintf("/api/v1/teams/%s/assets/%s/annotations", param0, param1)
}

// Update one or more annotation for a given asset.
func (c *Client) UpdateAssetAnnotations(ctx context.Context, path string, payload *AssetAnnotationRequest) (*http.Response, error) {
	req, err := c.NewUpdateAssetAnnotationsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAssetAnnotationsRequest create the request corresponding to the update action endpoint of the asset-annotations resource.
func (c *Client) NewUpdateAssetAnnotationsRequest(ctx context.Context, path string, payload *AssetAnnotationRequest) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
