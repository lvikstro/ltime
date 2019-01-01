// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "LTime API": authentication Resource Client
//
// Command:
// $ goagen
// --design=github.com/lvikstro/ltime/backend/design
// --out=$(GOPATH)/src/github.com/lvikstro/ltime/backend
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// LoginAuthenticationPath computes a request path to the login action of authentication.
func LoginAuthenticationPath() string {

	return fmt.Sprintf("/api/auth/login")
}

// Sign a user in
func (c *Client) LoginAuthentication(ctx context.Context, path string, payload *LoginPayload) (*http.Response, error) {
	req, err := c.NewLoginAuthenticationRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginAuthenticationRequest create the request corresponding to the login action endpoint of the authentication resource.
func (c *Client) NewLoginAuthenticationRequest(ctx context.Context, path string, payload *LoginPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// RegisterAuthenticationPath computes a request path to the register action of authentication.
func RegisterAuthenticationPath() string {

	return fmt.Sprintf("/api/auth/register")
}

// Create a new user
func (c *Client) RegisterAuthentication(ctx context.Context, path string, payload *RegisterPayload) (*http.Response, error) {
	req, err := c.NewRegisterAuthenticationRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegisterAuthenticationRequest create the request corresponding to the register action endpoint of the authentication resource.
func (c *Client) NewRegisterAuthenticationRequest(ctx context.Context, path string, payload *RegisterPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}
