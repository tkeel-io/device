// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	json "encoding/json"
	go_restful "github.com/emicklei/go-restful"
	errors "github.com/tkeel-io/kit/errors"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.go_restful.json.errors.

type DeviceHTTPServer interface {
	AddDeviceExt(context.Context, *AddDeviceExtRequest) (*AddDeviceExtResponse, error)
	CreateDevice(context.Context, *CreateDeviceRequest) (*CreateDeviceResponse, error)
	DeleteDevice(context.Context, *DeleteDeviceRequest) (*DeleteDeviceResponse, error)
	DeleteDeviceExt(context.Context, *DeleteDeviceExtRequest) (*DeleteDeviceExtResponse, error)
	EnableDevice(context.Context, *EnableDeviceRequest) (*EnableDeviceResponse, error)
	GetDevice(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	ListDevice(context.Context, *ListDeviceRequest) (*ListDeviceResponse, error)
	UpdateDevice(context.Context, *UpdateDeviceRequest) (*UpdateDeviceResponse, error)
	UpdateDeviceExt(context.Context, *UpdateDeviceExtRequest) (*UpdateDeviceExtResponse, error)
}

type DeviceHTTPHandler struct {
	srv DeviceHTTPServer
}

func newDeviceHTTPHandler(s DeviceHTTPServer) *DeviceHTTPHandler {
	return &DeviceHTTPHandler{srv: s}
}

func (h *DeviceHTTPHandler) AddDeviceExt(req *go_restful.Request, resp *go_restful.Response) {
	in := AddDeviceExtRequest{}
	if err := transportHTTP.GetBody(req, &in.Ext); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.AddDeviceExt(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) CreateDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := CreateDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.Dev); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.CreateDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) DeleteDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.Ids); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.DeleteDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) DeleteDeviceExt(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteDeviceExtRequest{}
	if err := transportHTTP.GetBody(req, &in.Keys); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.DeleteDeviceExt(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) EnableDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := EnableDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.Enable); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.EnableDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) GetDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := GetDeviceRequest{}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.GetDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) ListDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := ListDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.Filter); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.ListDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) UpdateDevice(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateDeviceRequest{}
	if err := transportHTTP.GetBody(req, &in.Dev); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.UpdateDevice(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *DeviceHTTPHandler) UpdateDeviceExt(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateDeviceExtRequest{}
	if err := transportHTTP.GetBody(req, &in.Ext); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	ctx := transportHTTP.ContextWithHeader(req.Request.Context(), req.Request.Header)

	out, err := h.srv.UpdateDeviceExt(ctx, &in)
	if err != nil {
		tErr := errors.FromError(err)
		httpCode := errors.GRPCToHTTPStatusCode(tErr.GRPCStatus().Code())
		resp.WriteErrorString(httpCode, tErr.Message)
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func RegisterDeviceHTTPServer(container *go_restful.Container, srv DeviceHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newDeviceHTTPHandler(srv)
	ws.Route(ws.POST("/devices").
		To(handler.CreateDevice))
	ws.Route(ws.PUT("/devices/{id}").
		To(handler.UpdateDevice))
	ws.Route(ws.POST("/devices/delete").
		To(handler.DeleteDevice))
	ws.Route(ws.GET("/devices/{id}").
		To(handler.GetDevice))
	ws.Route(ws.POST("/devices/search").
		To(handler.ListDevice))
	ws.Route(ws.PUT("/devices/{id}/enable").
		To(handler.EnableDevice))
	ws.Route(ws.POST("/devices/{id}/ext").
		To(handler.AddDeviceExt))
	ws.Route(ws.POST("/devices/{id}/ext/delete").
		To(handler.DeleteDeviceExt))
	ws.Route(ws.PUT("/devices/{id}/ext").
		To(handler.UpdateDeviceExt))
}
