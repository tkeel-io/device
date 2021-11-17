// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	json "encoding/json"
	go_restful "github.com/emicklei/go-restful"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.go_restful.json.

type GroupHTTPServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CommonResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*CommonResponse, error)
	GetGroup(context.Context, *GetGroupRequest) (*CommonResponse, error)
	ListGroup(context.Context, *ListGroupRequest) (*CommonResponse, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*CommonResponse, error)
}

type GroupHTTPHandler struct {
	srv GroupHTTPServer
}

func newGroupHTTPHandler(s GroupHTTPServer) *GroupHTTPHandler {
	return &GroupHTTPHandler{srv: s}
}

func (h *GroupHTTPHandler) CreateGroup(req *go_restful.Request, resp *go_restful.Response) {
	in := CreateGroupRequest{}
	if err := transportHTTP.GetBody(req, &in.Group); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.CreateGroup(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
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

func (h *GroupHTTPHandler) DeleteGroup(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteGroupRequest{}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.DeleteGroup(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
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

func (h *GroupHTTPHandler) GetGroup(req *go_restful.Request, resp *go_restful.Response) {
	in := GetGroupRequest{}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.GetGroup(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
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

func (h *GroupHTTPHandler) ListGroup(req *go_restful.Request, resp *go_restful.Response) {
	in := ListGroupRequest{}

	out, err := h.srv.ListGroup(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
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

func (h *GroupHTTPHandler) UpdateGroup(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateGroupRequest{}
	if err := transportHTTP.GetBody(req, &in.Group); err != nil {
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

	out, err := h.srv.UpdateGroup(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
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

func RegisterGroupHTTPServer(container *go_restful.Container, srv GroupHTTPServer) {
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

	handler := newGroupHTTPHandler(srv)
	ws.Route(ws.POST("/group").
		To(handler.CreateGroup))
	ws.Route(ws.PUT("/group/{id}").
		To(handler.UpdateGroup))
	ws.Route(ws.DELETE("/group/{id}").
		To(handler.DeleteGroup))
	ws.Route(ws.GET("/group/{id}").
		To(handler.GetGroup))
	ws.Route(ws.GET("/group").
		To(handler.ListGroup))
}
