package gosolo

import (
	"context"
	"crypto/tls"
	"net/url"
	"time"

	// External
	"github.com/firestuff/patchy/patchyc"
)

type (
	Filter     = patchyc.Filter
	GetOpts    = patchyc.GetOpts
	ListOpts   = patchyc.ListOpts
	UpdateOpts = patchyc.UpdateOpts
)

type Task struct {
	ID         string     `json:"id"`
	ETag       string     `json:"etag"`
	Generation int64      `json:"generation"`
	UserID     *string    `json:"userID"`
	Name       *string    `json:"name"`
	Complete   *bool      `json:"complete"`
	After      *time.Time `json:"after"`
}

type Token struct {
	ID         string  `json:"id"`
	ETag       string  `json:"etag"`
	Generation int64   `json:"generation"`
	UserID     *string `json:"userID"`
	Token      *string `json:"token"`
}

type User struct {
	ID           string  `json:"id"`
	ETag         string  `json:"etag"`
	Generation   int64   `json:"generation"`
	Name         *string `json:"name"`
	Email        *string `json:"email"`
	Password     *string `json:"password"`
	ServiceAdmin *bool   `json:"serviceAdmin"`
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

func (c *Client) CreateTask(ctx context.Context, obj *Task) (*Task, error) {
	return CreateName[Task](ctx, c, "task", obj)
}

func (c *Client) DeleteTask(ctx context.Context, id string, opts *UpdateOpts) error {
	return DeleteName[Task](ctx, c, "task", id, opts)
}

func (c *Client) FindTask(ctx context.Context, shortID string) (*Task, error) {
	return FindName[Task](ctx, c, "task", shortID)
}

func (c *Client) GetTask(ctx context.Context, id string, opts *GetOpts) (*Task, error) {
	return GetName[Task](ctx, c, "task", id, opts)
}

func (c *Client) ListTask(ctx context.Context, opts *ListOpts) ([]*Task, error) {
	return ListName[Task](ctx, c, "task", opts)
}

func (c *Client) ReplaceTask(ctx context.Context, id string, obj *Task, opts *UpdateOpts) (*Task, error) {
	return ReplaceName[Task](ctx, c, "task", id, obj, opts)
}

func (c *Client) UpdateTask(ctx context.Context, id string, obj *Task, opts *UpdateOpts) (*Task, error) {
	return UpdateName[Task](ctx, c, "task", id, obj, opts)
}

func (c *Client) StreamGetTask(ctx context.Context, id string, opts *GetOpts) (*patchyc.GetStream[Task], error) {
	return StreamGetName[Task](ctx, c, "task", id, opts)
}

func (c *Client) StreamListTask(ctx context.Context, opts *ListOpts) (*patchyc.ListStream[Task], error) {
	return StreamListName[Task](ctx, c, "task", opts)
}

//// Token

func (c *Client) CreateToken(ctx context.Context, obj *Token) (*Token, error) {
	return CreateName[Token](ctx, c, "token", obj)
}

func (c *Client) DeleteToken(ctx context.Context, id string, opts *UpdateOpts) error {
	return DeleteName[Token](ctx, c, "token", id, opts)
}

func (c *Client) FindToken(ctx context.Context, shortID string) (*Token, error) {
	return FindName[Token](ctx, c, "token", shortID)
}

func (c *Client) GetToken(ctx context.Context, id string, opts *GetOpts) (*Token, error) {
	return GetName[Token](ctx, c, "token", id, opts)
}

func (c *Client) ListToken(ctx context.Context, opts *ListOpts) ([]*Token, error) {
	return ListName[Token](ctx, c, "token", opts)
}

func (c *Client) ReplaceToken(ctx context.Context, id string, obj *Token, opts *UpdateOpts) (*Token, error) {
	return ReplaceName[Token](ctx, c, "token", id, obj, opts)
}

func (c *Client) UpdateToken(ctx context.Context, id string, obj *Token, opts *UpdateOpts) (*Token, error) {
	return UpdateName[Token](ctx, c, "token", id, obj, opts)
}

func (c *Client) StreamGetToken(ctx context.Context, id string, opts *GetOpts) (*patchyc.GetStream[Token], error) {
	return StreamGetName[Token](ctx, c, "token", id, opts)
}

func (c *Client) StreamListToken(ctx context.Context, opts *ListOpts) (*patchyc.ListStream[Token], error) {
	return StreamListName[Token](ctx, c, "token", opts)
}

//// User

func (c *Client) CreateUser(ctx context.Context, obj *User) (*User, error) {
	return CreateName[User](ctx, c, "user", obj)
}

func (c *Client) DeleteUser(ctx context.Context, id string, opts *UpdateOpts) error {
	return DeleteName[User](ctx, c, "user", id, opts)
}

func (c *Client) FindUser(ctx context.Context, shortID string) (*User, error) {
	return FindName[User](ctx, c, "user", shortID)
}

func (c *Client) GetUser(ctx context.Context, id string, opts *GetOpts) (*User, error) {
	return GetName[User](ctx, c, "user", id, opts)
}

func (c *Client) ListUser(ctx context.Context, opts *ListOpts) ([]*User, error) {
	return ListName[User](ctx, c, "user", opts)
}

func (c *Client) ReplaceUser(ctx context.Context, id string, obj *User, opts *UpdateOpts) (*User, error) {
	return ReplaceName[User](ctx, c, "user", id, obj, opts)
}

func (c *Client) UpdateUser(ctx context.Context, id string, obj *User, opts *UpdateOpts) (*User, error) {
	return UpdateName[User](ctx, c, "user", id, obj, opts)
}

func (c *Client) StreamGetUser(ctx context.Context, id string, opts *GetOpts) (*patchyc.GetStream[User], error) {
	return StreamGetName[User](ctx, c, "user", id, opts)
}

func (c *Client) StreamListUser(ctx context.Context, opts *ListOpts) (*patchyc.ListStream[User], error) {
	return StreamListName[User](ctx, c, "user", opts)
}

//// Generic

func CreateName[T any](ctx context.Context, c *Client, name string, obj *T) (*T, error) {
	return patchyc.CreateName[T](ctx, c.patchyClient, name, obj)
}

func DeleteName[T any](ctx context.Context, c *Client, name, id string, opts *UpdateOpts) error {
	return patchyc.DeleteName[T](ctx, c.patchyClient, name, id, opts)
}

func FindName[T any](ctx context.Context, c *Client, name, shortID string) (*T, error) {
	return patchyc.FindName[T](ctx, c.patchyClient, name, shortID)
}

func GetName[T any](ctx context.Context, c *Client, name, id string, opts *GetOpts) (*T, error) {
	return patchyc.GetName[T](ctx, c.patchyClient, name, id, opts)
}

func ListName[T any](ctx context.Context, c *Client, name string, opts *ListOpts) ([]*T, error) {
	return patchyc.ListName[T](ctx, c.patchyClient, name, opts)
}

func ReplaceName[T any](ctx context.Context, c *Client, name, id string, obj *T, opts *UpdateOpts) (*T, error) {
	return patchyc.ReplaceName[T](ctx, c.patchyClient, name, id, obj, opts)
}

func UpdateName[T any](ctx context.Context, c *Client, name, id string, obj *T, opts *UpdateOpts) (*T, error) {
	return patchyc.UpdateName[T](ctx, c.patchyClient, name, id, obj, opts)
}

func StreamGetName[T any](ctx context.Context, c *Client, name, id string, opts *GetOpts) (*patchyc.GetStream[T], error) {
	return patchyc.StreamGetName[T](ctx, c.patchyClient, name, id, opts)
}

func StreamListName[T any](ctx context.Context, c *Client, name string, opts *ListOpts) (*patchyc.ListStream[T], error) {
	return patchyc.StreamListName[T](ctx, c.patchyClient, name, opts)
}

//// Utility generic

func P[T any](v T) *T {
	return patchyc.P(v)
}
