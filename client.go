// Solø API client

package gosolo

import (
	"context"
	"crypto/tls"
	"net/url"
	"time"

	//
	"github.com/firestuff/patchy/patchyc"
	"github.com/gopatchy/metadata"
)

type (
	Filter     = patchyc.Filter
	GetOpts    = patchyc.GetOpts
	ListOpts   = patchyc.ListOpts
	UpdateOpts = patchyc.UpdateOpts
)

type TaskResponse struct {
	metadata.Metadata
	UserID   *string    `json:"userID,omitempty"`
	Name     *string    `json:"name,omitempty"`
	Complete *bool      `json:"complete,omitempty"`
	After    *time.Time `json:"after,omitempty"`
}

type TaskRequest struct {
	UserID   *string    `json:"userID,omitempty"`
	Name     *string    `json:"name,omitempty"`
	Complete *bool      `json:"complete,omitempty"`
	After    *time.Time `json:"after,omitempty"`
}

type TokenResponse struct {
	metadata.Metadata
	UserID *string `json:"userID,omitempty"`
	Token  *string `json:"token,omitempty"`
}

type TokenRequest struct {
	UserID *string `json:"userID,omitempty"`
	Token  *string `json:"token,omitempty"`
}

type UserResponse struct {
	metadata.Metadata
	Name         *string `json:"name,omitempty"`
	Email        *string `json:"email,omitempty"`
	Password     *string `json:"password,omitempty"`
	ServiceAdmin *bool   `json:"serviceAdmin,omitempty"`
}

type UserRequest struct {
	Name         *string `json:"name,omitempty"`
	Email        *string `json:"email,omitempty"`
	Password     *string `json:"password,omitempty"`
	ServiceAdmin *bool   `json:"serviceAdmin,omitempty"`
}

type Client struct {
	patchyClient *patchyc.Client
}

func NewClient(baseURL string) *Client {
	baseURL, err := url.JoinPath(baseURL, "/v1")
	if err != nil {
		panic(err)
	}

	return &Client{
		patchyClient: patchyc.NewClient(baseURL),
	}
}

func (c *Client) SetTLSClientConfig(cfg *tls.Config) *Client {
	c.patchyClient.SetTLSClientConfig(cfg)
	return c
}

func (c *Client) SetDebug(debug bool) *Client {
	c.patchyClient.SetDebug(debug)
	return c
}

func (c *Client) SetBasicAuth(user, pass string) *Client {
	c.patchyClient.SetBasicAuth(user, pass)
	return c
}

func (c *Client) SetAuthToken(token string) *Client {
	c.patchyClient.SetAuthToken(token)
	return c
}

//// Task

func (c *Client) CreateTask(ctx context.Context, obj *TaskRequest) (*TaskResponse, error) {
	return CreateName[TaskResponse, TaskRequest](ctx, c, "task", obj)
}

func (c *Client) DeleteTask(ctx context.Context, id string, opts *UpdateOpts) error {
	return DeleteName[TaskResponse](ctx, c, "task", id, opts)
}

func (c *Client) FindTask(ctx context.Context, shortID string) (*TaskResponse, error) {
	return FindName[TaskResponse](ctx, c, "task", shortID)
}

func (c *Client) GetTask(ctx context.Context, id string, opts *GetOpts) (*TaskResponse, error) {
	return GetName[TaskResponse](ctx, c, "task", id, opts)
}

func (c *Client) ListTask(ctx context.Context, opts *ListOpts) ([]*TaskResponse, error) {
	return ListName[TaskResponse](ctx, c, "task", opts)
}

func (c *Client) ReplaceTask(ctx context.Context, id string, obj *TaskRequest, opts *UpdateOpts) (*TaskResponse, error) {
	return ReplaceName[TaskResponse, TaskRequest](ctx, c, "task", id, obj, opts)
}

func (c *Client) UpdateTask(ctx context.Context, id string, obj *TaskRequest, opts *UpdateOpts) (*TaskResponse, error) {
	return UpdateName[TaskResponse, TaskRequest](ctx, c, "task", id, obj, opts)
}

func (c *Client) StreamGetTask(ctx context.Context, id string, opts *GetOpts) (*patchyc.GetStream[TaskResponse], error) {
	return StreamGetName[TaskResponse](ctx, c, "task", id, opts)
}

func (c *Client) StreamListTask(ctx context.Context, opts *ListOpts) (*patchyc.ListStream[TaskResponse], error) {
	return StreamListName[TaskResponse](ctx, c, "task", opts)
}

//// Token

func (c *Client) CreateToken(ctx context.Context, obj *TokenRequest) (*TokenResponse, error) {
	return CreateName[TokenResponse, TokenRequest](ctx, c, "token", obj)
}

func (c *Client) DeleteToken(ctx context.Context, id string, opts *UpdateOpts) error {
	return DeleteName[TokenResponse](ctx, c, "token", id, opts)
}

func (c *Client) FindToken(ctx context.Context, shortID string) (*TokenResponse, error) {
	return FindName[TokenResponse](ctx, c, "token", shortID)
}

func (c *Client) GetToken(ctx context.Context, id string, opts *GetOpts) (*TokenResponse, error) {
	return GetName[TokenResponse](ctx, c, "token", id, opts)
}

func (c *Client) ListToken(ctx context.Context, opts *ListOpts) ([]*TokenResponse, error) {
	return ListName[TokenResponse](ctx, c, "token", opts)
}

func (c *Client) ReplaceToken(ctx context.Context, id string, obj *TokenRequest, opts *UpdateOpts) (*TokenResponse, error) {
	return ReplaceName[TokenResponse, TokenRequest](ctx, c, "token", id, obj, opts)
}

func (c *Client) UpdateToken(ctx context.Context, id string, obj *TokenRequest, opts *UpdateOpts) (*TokenResponse, error) {
	return UpdateName[TokenResponse, TokenRequest](ctx, c, "token", id, obj, opts)
}

func (c *Client) StreamGetToken(ctx context.Context, id string, opts *GetOpts) (*patchyc.GetStream[TokenResponse], error) {
	return StreamGetName[TokenResponse](ctx, c, "token", id, opts)
}

func (c *Client) StreamListToken(ctx context.Context, opts *ListOpts) (*patchyc.ListStream[TokenResponse], error) {
	return StreamListName[TokenResponse](ctx, c, "token", opts)
}

//// User

func (c *Client) CreateUser(ctx context.Context, obj *UserRequest) (*UserResponse, error) {
	return CreateName[UserResponse, UserRequest](ctx, c, "user", obj)
}

func (c *Client) DeleteUser(ctx context.Context, id string, opts *UpdateOpts) error {
	return DeleteName[UserResponse](ctx, c, "user", id, opts)
}

func (c *Client) FindUser(ctx context.Context, shortID string) (*UserResponse, error) {
	return FindName[UserResponse](ctx, c, "user", shortID)
}

func (c *Client) GetUser(ctx context.Context, id string, opts *GetOpts) (*UserResponse, error) {
	return GetName[UserResponse](ctx, c, "user", id, opts)
}

func (c *Client) ListUser(ctx context.Context, opts *ListOpts) ([]*UserResponse, error) {
	return ListName[UserResponse](ctx, c, "user", opts)
}

func (c *Client) ReplaceUser(ctx context.Context, id string, obj *UserRequest, opts *UpdateOpts) (*UserResponse, error) {
	return ReplaceName[UserResponse, UserRequest](ctx, c, "user", id, obj, opts)
}

func (c *Client) UpdateUser(ctx context.Context, id string, obj *UserRequest, opts *UpdateOpts) (*UserResponse, error) {
	return UpdateName[UserResponse, UserRequest](ctx, c, "user", id, obj, opts)
}

func (c *Client) StreamGetUser(ctx context.Context, id string, opts *GetOpts) (*patchyc.GetStream[UserResponse], error) {
	return StreamGetName[UserResponse](ctx, c, "user", id, opts)
}

func (c *Client) StreamListUser(ctx context.Context, opts *ListOpts) (*patchyc.ListStream[UserResponse], error) {
	return StreamListName[UserResponse](ctx, c, "user", opts)
}

//// Generic

func CreateName[TOut, TIn any](ctx context.Context, c *Client, name string, obj *TIn) (*TOut, error) {
	return patchyc.CreateName[TOut, TIn](ctx, c.patchyClient, name, obj)
}

func DeleteName[TOut any](ctx context.Context, c *Client, name, id string, opts *UpdateOpts) error {
	return patchyc.DeleteName[TOut](ctx, c.patchyClient, name, id, opts)
}

func FindName[TOut any](ctx context.Context, c *Client, name, shortID string) (*TOut, error) {
	return patchyc.FindName[TOut](ctx, c.patchyClient, name, shortID)
}

func GetName[TOut any](ctx context.Context, c *Client, name, id string, opts *GetOpts) (*TOut, error) {
	return patchyc.GetName[TOut](ctx, c.patchyClient, name, id, opts)
}

func ListName[TOut any](ctx context.Context, c *Client, name string, opts *ListOpts) ([]*TOut, error) {
	return patchyc.ListName[TOut](ctx, c.patchyClient, name, opts)
}

func ReplaceName[TOut, TIn any](ctx context.Context, c *Client, name, id string, obj *TIn, opts *UpdateOpts) (*TOut, error) {
	return patchyc.ReplaceName[TOut, TIn](ctx, c.patchyClient, name, id, obj, opts)
}

func UpdateName[TOut, TIn any](ctx context.Context, c *Client, name, id string, obj *TIn, opts *UpdateOpts) (*TOut, error) {
	return patchyc.UpdateName[TOut, TIn](ctx, c.patchyClient, name, id, obj, opts)
}

func StreamGetName[TOut any](ctx context.Context, c *Client, name, id string, opts *GetOpts) (*patchyc.GetStream[TOut], error) {
	return patchyc.StreamGetName[TOut](ctx, c.patchyClient, name, id, opts)
}

func StreamListName[TOut any](ctx context.Context, c *Client, name string, opts *ListOpts) (*patchyc.ListStream[TOut], error) {
	return patchyc.StreamListName[TOut](ctx, c.patchyClient, name, opts)
}

//// Utility generic

func P[T any](v T) *T {
	return patchyc.P(v)
}

// vim: set filetype=go:
