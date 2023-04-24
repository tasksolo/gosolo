// Solø API client

package gosolo

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	//
	"github.com/go-resty/resty/v2"
	"github.com/gopatchy/jsrest"
	"github.com/gopatchy/metadata"
	"golang.org/x/exp/slices"
)

type Task struct {
	metadata.Metadata

	ListETag string    `json:"-"`
	UserID   string    `json:"userID,omitempty"`
	Name     string    `json:"name,omitempty"`
	Complete bool      `json:"complete,omitempty"`
	After    time.Time `json:"after,omitempty"`
}

type Token struct {
	metadata.Metadata

	ListETag string `json:"-"`
	UserID   string `json:"userID,omitempty"`
	Token    string `json:"token,omitempty"`
}

type User struct {
	metadata.Metadata

	ListETag     string `json:"-"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	ServiceAdmin bool   `json:"serviceAdmin,omitempty"`
}

type GetOpts[T any] struct {
	Prev *T
}

type ListOpts[T any] struct {
	Stream  string
	Limit   int64
	Offset  int64
	After   string
	Sorts   []string
	Filters []Filter

	Prev []*T
}

type Filter struct {
	Path  string
	Op    string
	Value string
}

type UpdateOpts[T any] struct {
	Prev *T
}

type Client struct {
	rst *resty.Client
}

var (
	ErrNotFound            = fmt.Errorf("not found")
	ErrMultipleFound       = fmt.Errorf("multiple found")
	ErrInvalidStreamEvent  = fmt.Errorf("invalid stream event")
	ErrInvalidStreamFormat = fmt.Errorf("invalid stream format")
)

func NewClient(baseURL string) *Client {
	baseURL, err := url.JoinPath(baseURL, "/v1")
	if err != nil {
		panic(err)
	}

	rst := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Accept", "application/json").
		SetJSONEscapeHTML(false)

	// TODO: SetTimeout()
	// TODO: SetRetry*() or roll our own
	// TODO: Add Idempotency-Key support

	return &Client{
		rst: rst,
	}
}

func (c *Client) SetTLSClientConfig(cfg *tls.Config) *Client {
	c.rst.SetTLSClientConfig(cfg)
	return c
}

func (c *Client) SetDebug(debug bool) *Client {
	c.rst.SetDebug(debug)
	return c
}

func (c *Client) SetHeader(name, value string) *Client {
	c.rst.SetHeader(name, value)
	return c
}

func (c *Client) ResetAuth() *Client {
	c.rst.Token = ""
	c.rst.UserInfo = nil

	return c
}

func (c *Client) SetBasicAuth(user, pass string) *Client {
	c.ResetAuth()
	c.rst.SetBasicAuth(user, pass)

	return c
}

func (c *Client) SetAuthToken(token string) *Client {
	c.ResetAuth()
	c.rst.SetAuthToken(token)

	return c
}

func (c *Client) DebugInfo(ctx context.Context) (map[string]any, error) {
	return c.fetchMap(ctx, "_debug")
}

func (c *Client) OpenAPI(ctx context.Context) (map[string]any, error) {
	return c.fetchMap(ctx, "_openapi")
}

func (c *Client) GoClient(ctx context.Context) (string, error) {
	return c.fetchString(ctx, "_client.go")
}

func (c *Client) TSClient(ctx context.Context) (string, error) {
	return c.fetchString(ctx, "_client.ts")
}

//// Task

func (c *Client) CreateTask(ctx context.Context, obj *Task) (*Task, error) {
	return CreateName[Task](ctx, c, "task", obj)
}

func (c *Client) DeleteTask(ctx context.Context, id string, opts *UpdateOpts[Task]) error {
	return DeleteName[Task](ctx, c, "task", id, opts)
}

func (c *Client) FindTask(ctx context.Context, shortID string) (*Task, error) {
	return FindName[Task](ctx, c, "task", shortID)
}

func (c *Client) GetTask(ctx context.Context, id string, opts *GetOpts[Task]) (*Task, error) {
	return GetName[Task](ctx, c, "task", id, opts)
}

func (c *Client) ListTask(ctx context.Context, opts *ListOpts[Task]) ([]*Task, error) {
	return ListName[Task](ctx, c, "task", opts)
}

func (c *Client) ReplaceTask(ctx context.Context, id string, obj *Task, opts *UpdateOpts[Task]) (*Task, error) {
	return ReplaceName[Task](ctx, c, "task", id, obj, opts)
}

func (c *Client) UpdateTask(ctx context.Context, id string, obj *Task, opts *UpdateOpts[Task]) (*Task, error) {
	return UpdateName[Task](ctx, c, "task", id, obj, opts)
}

func (c *Client) StreamGetTask(ctx context.Context, id string, opts *GetOpts[Task]) (*GetStream[Task], error) {
	return StreamGetName[Task](ctx, c, "task", id, opts)
}

func (c *Client) StreamListTask(ctx context.Context, opts *ListOpts[Task]) (*ListStream[Task], error) {
	return StreamListName[Task](ctx, c, "task", opts)
}

//// Token

func (c *Client) CreateToken(ctx context.Context, obj *Token) (*Token, error) {
	return CreateName[Token](ctx, c, "token", obj)
}

func (c *Client) DeleteToken(ctx context.Context, id string, opts *UpdateOpts[Token]) error {
	return DeleteName[Token](ctx, c, "token", id, opts)
}

func (c *Client) FindToken(ctx context.Context, shortID string) (*Token, error) {
	return FindName[Token](ctx, c, "token", shortID)
}

func (c *Client) GetToken(ctx context.Context, id string, opts *GetOpts[Token]) (*Token, error) {
	return GetName[Token](ctx, c, "token", id, opts)
}

func (c *Client) ListToken(ctx context.Context, opts *ListOpts[Token]) ([]*Token, error) {
	return ListName[Token](ctx, c, "token", opts)
}

func (c *Client) ReplaceToken(ctx context.Context, id string, obj *Token, opts *UpdateOpts[Token]) (*Token, error) {
	return ReplaceName[Token](ctx, c, "token", id, obj, opts)
}

func (c *Client) UpdateToken(ctx context.Context, id string, obj *Token, opts *UpdateOpts[Token]) (*Token, error) {
	return UpdateName[Token](ctx, c, "token", id, obj, opts)
}

func (c *Client) StreamGetToken(ctx context.Context, id string, opts *GetOpts[Token]) (*GetStream[Token], error) {
	return StreamGetName[Token](ctx, c, "token", id, opts)
}

func (c *Client) StreamListToken(ctx context.Context, opts *ListOpts[Token]) (*ListStream[Token], error) {
	return StreamListName[Token](ctx, c, "token", opts)
}

//// User

func (c *Client) CreateUser(ctx context.Context, obj *User) (*User, error) {
	return CreateName[User](ctx, c, "user", obj)
}

func (c *Client) DeleteUser(ctx context.Context, id string, opts *UpdateOpts[User]) error {
	return DeleteName[User](ctx, c, "user", id, opts)
}

func (c *Client) FindUser(ctx context.Context, shortID string) (*User, error) {
	return FindName[User](ctx, c, "user", shortID)
}

func (c *Client) GetUser(ctx context.Context, id string, opts *GetOpts[User]) (*User, error) {
	return GetName[User](ctx, c, "user", id, opts)
}

func (c *Client) ListUser(ctx context.Context, opts *ListOpts[User]) ([]*User, error) {
	return ListName[User](ctx, c, "user", opts)
}

func (c *Client) ReplaceUser(ctx context.Context, id string, obj *User, opts *UpdateOpts[User]) (*User, error) {
	return ReplaceName[User](ctx, c, "user", id, obj, opts)
}

func (c *Client) UpdateUser(ctx context.Context, id string, obj *User, opts *UpdateOpts[User]) (*User, error) {
	return UpdateName[User](ctx, c, "user", id, obj, opts)
}

func (c *Client) StreamGetUser(ctx context.Context, id string, opts *GetOpts[User]) (*GetStream[User], error) {
	return StreamGetName[User](ctx, c, "user", id, opts)
}

func (c *Client) StreamListUser(ctx context.Context, opts *ListOpts[User]) (*ListStream[User], error) {
	return StreamListName[User](ctx, c, "user", opts)
}

//// Generic

func CreateName[T any](ctx context.Context, c *Client, name string, obj *T) (*T, error) {
	created := new(T)

	resp, err := c.rst.R().
		SetContext(ctx).
		SetPathParam("name", name).
		SetBody(obj).
		SetResult(created).
		Post("{name}")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	return created, nil
}

func DeleteName[T any](ctx context.Context, c *Client, name, id string, opts *UpdateOpts[T]) error {
	r := c.rst.R().
		SetContext(ctx).
		SetPathParam("name", name).
		SetPathParam("id", id)

	opts.apply(r)

	resp, err := r.Delete("{name}/{id}")
	if err != nil {
		return err
	}

	if resp.IsError() {
		return jsrest.ReadError(resp.Body())
	}

	return nil
}

func FindName[T any](ctx context.Context, c *Client, name, shortID string) (*T, error) {
	listOpts := &ListOpts[T]{
		Filters: []Filter{
			{
				Path:  "id",
				Op:    "hp",
				Value: shortID,
			},
		},
	}

	objs, err := ListName[T](ctx, c, name, listOpts)
	if err != nil {
		return nil, err
	}

	if len(objs) == 0 {
		return nil, fmt.Errorf("%s (%w)", shortID, ErrNotFound)
	}

	if len(objs) > 1 {
		return nil, fmt.Errorf("%s (%w)", shortID, ErrMultipleFound)
	}

	return objs[0], nil
}

func GetName[T any](ctx context.Context, c *Client, name, id string, opts *GetOpts[T]) (*T, error) {
	obj := new(T)

	r := c.rst.R().
		SetContext(ctx).
		SetPathParam("name", name).
		SetPathParam("id", id).
		SetResult(obj)

	opts.apply(r)

	resp, err := r.Get("{name}/{id}")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, nil
	}

	if opts != nil && opts.Prev != nil && resp.StatusCode() == http.StatusNotModified {
		return opts.Prev, nil
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	return obj, nil
}

func ListName[T any](ctx context.Context, c *Client, name string, opts *ListOpts[T]) ([]*T, error) {
	objs := []*T{}

	r := c.rst.R().
		SetContext(ctx).
		SetPathParam("name", name).
		SetResult(&objs)

	err := opts.apply(r)
	if err != nil {
		return nil, err
	}

	resp, err := r.Get("{name}")
	if err != nil {
		return nil, err
	}

	if opts != nil && opts.Prev != nil && resp.StatusCode() == http.StatusNotModified {
		return opts.Prev, nil
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	setListETag(objs, resp.Header().Get("ETag"))

	return objs, nil
}

func ReplaceName[T any](ctx context.Context, c *Client, name, id string, obj *T, opts *UpdateOpts[T]) (*T, error) {
	replaced := new(T)

	r := c.rst.R().
		SetContext(ctx).
		SetPathParam("name", name).
		SetPathParam("id", id).
		SetBody(obj).
		SetResult(replaced)

	opts.apply(r)

	resp, err := r.Put("{name}/{id}")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	return replaced, nil
}

func UpdateName[T any](ctx context.Context, c *Client, name, id string, obj *T, opts *UpdateOpts[T]) (*T, error) {
	updated := new(T)

	r := c.rst.R().
		SetContext(ctx).
		SetPathParam("name", name).
		SetPathParam("id", id).
		SetBody(obj).
		SetResult(updated)

	opts.apply(r)

	resp, err := r.Patch("{name}/{id}")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	return updated, nil
}

func StreamGetName[T any](ctx context.Context, c *Client, name, id string, opts *GetOpts[T]) (*GetStream[T], error) {
	r := c.rst.R().
		SetDoNotParseResponse(true).
		SetHeader("Accept", "text/event-stream").
		SetPathParam("name", name).
		SetPathParam("id", id)

	opts.apply(r)

	resp, err := r.Get("{name}/{id}")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	stream := &GetStream[T]{
		ch:   make(chan *T, 100),
		body: resp.RawBody(),
	}

	if opts != nil && opts.Prev != nil {
		stream.prev = opts.Prev
	}

	go stream.process()

	return stream, nil
}

func StreamListName[T any](ctx context.Context, c *Client, name string, opts *ListOpts[T]) (*ListStream[T], error) {
	r := c.rst.R().
		SetDoNotParseResponse(true).
		SetHeader("Accept", "text/event-stream").
		SetPathParam("name", name)

	if opts == nil {
		opts = &ListOpts[T]{}
	}

	err := opts.apply(r)
	if err != nil {
		return nil, err
	}

	resp, err := r.Get("{name}")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	stream := &ListStream[T]{
		ch:   make(chan []*T, 100),
		body: resp.RawBody(),
	}

	if opts != nil && opts.Prev != nil {
		stream.prev = opts.Prev
	}

	switch resp.Header().Get("Stream-Format") {
	case "full":
		go stream.processFull()

	case "diff":
		go stream.processDiff()

	default:
		stream.Close()
		return nil, fmt.Errorf("%s (%w)", resp.Header().Get("Stream-Format"), ErrInvalidStreamFormat)
	}

	return stream, nil
}

type streamEvent[T any] struct {
	eventType string
	params    map[string]string
	data      []byte
}

func newStreamEvent[T any]() *streamEvent[T] {
	return &streamEvent[T]{
		params: map[string]string{},
	}
}

func (ev *streamEvent[T]) decodeObj() (*T, error) {
	obj := new(T)

	err := json.Unmarshal(ev.data, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (ev *streamEvent[T]) decodeList() ([]*T, error) {
	list := []*T{}

	err := json.Unmarshal(ev.data, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

type eventStream[T any] struct {
	scan *bufio.Scanner
}

func newEventStream[T any](scan *bufio.Scanner) *eventStream[T] {
	return &eventStream[T]{
		scan: scan,
	}
}

func (es *eventStream[T]) readEvent() (*streamEvent[T], error) {
	event := newStreamEvent[T]()
	data := [][]byte{}

	for es.scan.Scan() {
		line := es.scan.Text()

		switch {
		case strings.HasPrefix(line, ":"):
			continue

		case strings.HasPrefix(line, "event: "):
			event.eventType = strings.TrimPrefix(line, "event: ")

		case strings.HasPrefix(line, "data: "):
			data = append(data, bytes.TrimPrefix(es.scan.Bytes(), []byte("data: ")))

		case line == "":
			event.data = bytes.Join(data, []byte("\n"))
			return event, nil

		case strings.Contains(line, ": "):
			parts := strings.SplitN(line, ": ", 2)
			event.params[parts[0]] = parts[1]
		}
	}

	return nil, io.EOF
}

type GetStream[T any] struct {
	ch   chan *T
	body io.ReadCloser
	prev *T

	lastEventReceived time.Time
	err               error
	mu                sync.RWMutex
}

func (gs *GetStream[T]) Close() {
	gs.body.Close()
}

func (gs *GetStream[T]) Chan() <-chan *T {
	return gs.ch
}

func (gs *GetStream[T]) Read() *T {
	return <-gs.Chan()
}

func (gs *GetStream[T]) LastEventReceived() time.Time {
	gs.mu.RLock()
	defer gs.mu.RUnlock()

	return gs.lastEventReceived
}

func (gs *GetStream[T]) Error() error {
	gs.mu.RLock()
	defer gs.mu.RUnlock()

	return gs.err
}

func (gs *GetStream[T]) process() {
	scan := bufio.NewScanner(gs.body)
	es := newEventStream[T](scan)

	for {
		event, err := es.readEvent()
		if err != nil {
			gs.writeError(err)
			return
		}

		switch event.eventType {
		case "initial":
			fallthrough
		case "update":
			obj, err := event.decodeObj()
			if err != nil {
				gs.writeError(err)
				return
			}

			gs.writeEvent(obj)

		case "notModified":
			if gs.prev != nil {
				gs.writeEvent(gs.prev)
			} else {
				gs.writeError(fmt.Errorf("notModified without If-None-Match (%w)", ErrInvalidStreamEvent))
				return
			}

		case "heartbeat":
			gs.writeHeartbeat()
		}
	}
}

func (gs *GetStream[T]) writeHeartbeat() {
	gs.mu.Lock()
	gs.lastEventReceived = time.Now()
	gs.mu.Unlock()
}

func (gs *GetStream[T]) writeEvent(obj *T) {
	gs.mu.Lock()
	gs.lastEventReceived = time.Now()
	gs.mu.Unlock()

	gs.ch <- obj
}

func (gs *GetStream[T]) writeError(err error) {
	gs.mu.Lock()
	gs.err = err
	gs.mu.Unlock()

	close(gs.ch)
}

type ListStream[T any] struct {
	ch   chan []*T
	body io.ReadCloser
	prev []*T

	lastEventReceived time.Time

	err error

	mu sync.RWMutex
}

func (ls *ListStream[T]) Close() {
	ls.body.Close()
}

func (ls *ListStream[T]) Chan() <-chan []*T {
	return ls.ch
}

func (ls *ListStream[T]) Read() []*T {
	return <-ls.Chan()
}

func (ls *ListStream[T]) LastEventReceived() time.Time {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	return ls.lastEventReceived
}

func (ls *ListStream[T]) Error() error {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	return ls.err
}

func (ls *ListStream[T]) processFull() {
	scan := bufio.NewScanner(ls.body)
	es := newEventStream[T](scan)

	for {
		event, err := es.readEvent()
		if err != nil {
			ls.writeError(err)
			return
		}

		switch event.eventType {
		case "list":
			list, err := event.decodeList()
			if err != nil {
				ls.writeError(err)
				return
			}

			setListETag(list, fmt.Sprintf(`"%s"`, event.params["id"]))
			ls.writeEvent(list)

		case "notModified":
			ls.writeEvent(ls.prev)

		case "heartbeat":
			ls.writeHeartbeat()
		}
	}
}

func (ls *ListStream[T]) processDiff() {
	scan := bufio.NewScanner(ls.body)
	es := newEventStream[T](scan)
	list := []*T{}

	add := func(event *streamEvent[T]) error {
		obj, err := event.decodeObj()
		if err != nil {
			return err
		}

		pos, err := strconv.Atoi(event.params["new-position"])
		if err != nil {
			return err
		}

		list = slices.Insert(list, pos, obj)

		return nil
	}

	remove := func(event *streamEvent[T]) error {
		pos, err := strconv.Atoi(event.params["old-position"])
		if err != nil {
			return err
		}

		list = slices.Delete(list, pos, pos+1)

		return nil
	}

	for {
		event, err := es.readEvent()
		if err != nil {
			ls.writeError(err)
			return
		}

		switch event.eventType {
		case "add":
			err = add(event)
			if err != nil {
				ls.writeError(err)
				return
			}

		case "update":
			err = remove(event)
			if err != nil {
				ls.writeError(err)
				return
			}

			err = add(event)
			if err != nil {
				ls.writeError(err)
				return
			}

		case "remove":
			err = remove(event)
			if err != nil {
				ls.writeError(err)
				return
			}

		case "sync":
			setListETag(list, fmt.Sprintf(`"%s"`, event.params["id"]))
			ls.writeEvent(list)

		case "notModified":
			list = ls.prev
			ls.writeEvent(list)

		case "heartbeat":
			ls.writeHeartbeat()
		}
	}
}

func (ls *ListStream[T]) writeHeartbeat() {
	ls.mu.Lock()
	ls.lastEventReceived = time.Now()
	ls.mu.Unlock()
}

func (ls *ListStream[T]) writeEvent(list []*T) {
	ls.mu.Lock()
	ls.lastEventReceived = time.Now()
	ls.mu.Unlock()

	ls.ch <- list
}

func (ls *ListStream[T]) writeError(err error) {
	ls.mu.Lock()
	ls.err = err
	ls.mu.Unlock()

	close(ls.ch)
}

//// Internal

func (opts *GetOpts[T]) apply(req *resty.Request) {
	if opts == nil {
		return
	}

	if opts.Prev != nil {
		md := metadata.GetMetadata(opts.Prev)
		req.SetHeader("If-None-Match", fmt.Sprintf(`"%s"`, md.ETag))
	}
}

func (opts *ListOpts[T]) apply(req *resty.Request) error {
	if opts == nil {
		return nil
	}

	etag := getListETag(opts.Prev)
	if etag != "" {
		req.SetHeader("If-None-Match", etag)
	}

	if opts.Stream != "" {
		req.SetQueryParam("_stream", opts.Stream)
	}

	if opts.Limit != 0 {
		req.SetQueryParam("_limit", fmt.Sprintf("%d", opts.Limit))
	}

	if opts.Offset != 0 {
		req.SetQueryParam("_offset", fmt.Sprintf("%d", opts.Offset))
	}

	if opts.After != "" {
		req.SetQueryParam("_after", opts.After)
	}

	for _, filter := range opts.Filters {
		req.SetQueryParam(fmt.Sprintf("%s[%s]", filter.Path, filter.Op), filter.Value)
	}

	sorts := url.Values{}

	for _, sort := range opts.Sorts {
		sorts.Add("_sort", sort)
	}

	req.SetQueryParamsFromValues(sorts)

	return nil
}

func (opts *UpdateOpts[T]) apply(req *resty.Request) {
	if opts == nil {
		return
	}

	if opts.Prev != nil {
		md := metadata.GetMetadata(opts.Prev)
		req.SetHeader("If-Match", fmt.Sprintf(`"%s"`, md.ETag))
	}
}

func (c *Client) fetchMap(ctx context.Context, path string) (map[string]any, error) {
	ret := map[string]any{}

	resp, err := c.rst.R().
		SetContext(ctx).
		SetResult(&ret).
		Get(path)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, jsrest.ReadError(resp.Body())
	}

	return ret, nil
}

func (c *Client) fetchString(ctx context.Context, path string) (string, error) {
	resp, err := c.rst.R().
		SetContext(ctx).
		Get(path)
	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", jsrest.ReadError(resp.Body())
	}

	return resp.String(), nil
}

func getListETag[T any](list []*T) string {
	if len(list) == 0 {
		return ""
	}

	return getListETagField(list).String()
}

func setListETag[T any](list []*T, etag string) {
	if len(list) == 0 {
		return
	}

	getListETagField(list).Set(reflect.ValueOf(etag))
}

func getListETagField[T any](list []*T) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(list[0])).FieldByName("ListETag")
}
