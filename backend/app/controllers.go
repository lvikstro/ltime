// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "LTime API": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/lvikstro/ltime/backend/design
// --out=$(GOPATH)/src/github.com/lvikstro/ltime/backend
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthenticationController is the controller interface for the Authentication actions.
type AuthenticationController interface {
	goa.Muxer
	Login(*LoginAuthenticationContext) error
	Register(*RegisterAuthenticationContext) error
}

// MountAuthenticationController "mounts" a Authentication resource controller on the given service.
func MountAuthenticationController(service *goa.Service, ctrl AuthenticationController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/auth/login", ctrl.MuxHandler("preflight", handleAuthenticationOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/auth/register", ctrl.MuxHandler("preflight", handleAuthenticationOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginAuthenticationContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*LoginPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Login(rctx)
	}
	h = handleAuthenticationOrigin(h)
	service.Mux.Handle("POST", "/api/auth/login", ctrl.MuxHandler("login", h, unmarshalLoginAuthenticationPayload))
	service.LogInfo("mount", "ctrl", "Authentication", "action", "Login", "route", "POST /api/auth/login")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRegisterAuthenticationContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RegisterPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Register(rctx)
	}
	h = handleAuthenticationOrigin(h)
	service.Mux.Handle("POST", "/api/auth/register", ctrl.MuxHandler("register", h, unmarshalRegisterAuthenticationPayload))
	service.LogInfo("mount", "ctrl", "Authentication", "action", "Register", "route", "POST /api/auth/register")
}

// handleAuthenticationOrigin applies the CORS response headers corresponding to the origin.
func handleAuthenticationOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT, OPTION")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalLoginAuthenticationPayload unmarshals the request body into the context request data Payload field.
func unmarshalLoginAuthenticationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &loginPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalRegisterAuthenticationPayload unmarshals the request body into the context request data Payload field.
func unmarshalRegisterAuthenticationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &registerPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT, OPTION")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
