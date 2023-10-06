// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-message-center/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-message-center/ent/emaillog"
	"github.com/suyuan32/simple-admin-message-center/ent/emailprovider"
	"github.com/suyuan32/simple-admin-message-center/ent/smslog"
	"github.com/suyuan32/simple-admin-message-center/ent/smsprovider"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// EmailLog is the client for interacting with the EmailLog builders.
	EmailLog *EmailLogClient
	// EmailProvider is the client for interacting with the EmailProvider builders.
	EmailProvider *EmailProviderClient
	// SmsLog is the client for interacting with the SmsLog builders.
	SmsLog *SmsLogClient
	// SmsProvider is the client for interacting with the SmsProvider builders.
	SmsProvider *SmsProviderClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.EmailLog = NewEmailLogClient(c.config)
	c.EmailProvider = NewEmailProviderClient(c.config)
	c.SmsLog = NewSmsLogClient(c.config)
	c.SmsProvider = NewSmsProviderClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		EmailLog:      NewEmailLogClient(cfg),
		EmailProvider: NewEmailProviderClient(cfg),
		SmsLog:        NewSmsLogClient(cfg),
		SmsProvider:   NewSmsProviderClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		EmailLog:      NewEmailLogClient(cfg),
		EmailProvider: NewEmailProviderClient(cfg),
		SmsLog:        NewSmsLogClient(cfg),
		SmsProvider:   NewSmsProviderClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		EmailLog.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.EmailLog.Use(hooks...)
	c.EmailProvider.Use(hooks...)
	c.SmsLog.Use(hooks...)
	c.SmsProvider.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.EmailLog.Intercept(interceptors...)
	c.EmailProvider.Intercept(interceptors...)
	c.SmsLog.Intercept(interceptors...)
	c.SmsProvider.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *EmailLogMutation:
		return c.EmailLog.mutate(ctx, m)
	case *EmailProviderMutation:
		return c.EmailProvider.mutate(ctx, m)
	case *SmsLogMutation:
		return c.SmsLog.mutate(ctx, m)
	case *SmsProviderMutation:
		return c.SmsProvider.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// EmailLogClient is a client for the EmailLog schema.
type EmailLogClient struct {
	config
}

// NewEmailLogClient returns a client for the EmailLog from the given config.
func NewEmailLogClient(c config) *EmailLogClient {
	return &EmailLogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `emaillog.Hooks(f(g(h())))`.
func (c *EmailLogClient) Use(hooks ...Hook) {
	c.hooks.EmailLog = append(c.hooks.EmailLog, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `emaillog.Intercept(f(g(h())))`.
func (c *EmailLogClient) Intercept(interceptors ...Interceptor) {
	c.inters.EmailLog = append(c.inters.EmailLog, interceptors...)
}

// Create returns a builder for creating a EmailLog entity.
func (c *EmailLogClient) Create() *EmailLogCreate {
	mutation := newEmailLogMutation(c.config, OpCreate)
	return &EmailLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EmailLog entities.
func (c *EmailLogClient) CreateBulk(builders ...*EmailLogCreate) *EmailLogCreateBulk {
	return &EmailLogCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EmailLogClient) MapCreateBulk(slice any, setFunc func(*EmailLogCreate, int)) *EmailLogCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EmailLogCreateBulk{err: fmt.Errorf("calling to EmailLogClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EmailLogCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EmailLogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EmailLog.
func (c *EmailLogClient) Update() *EmailLogUpdate {
	mutation := newEmailLogMutation(c.config, OpUpdate)
	return &EmailLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmailLogClient) UpdateOne(el *EmailLog) *EmailLogUpdateOne {
	mutation := newEmailLogMutation(c.config, OpUpdateOne, withEmailLog(el))
	return &EmailLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmailLogClient) UpdateOneID(id uuid.UUID) *EmailLogUpdateOne {
	mutation := newEmailLogMutation(c.config, OpUpdateOne, withEmailLogID(id))
	return &EmailLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EmailLog.
func (c *EmailLogClient) Delete() *EmailLogDelete {
	mutation := newEmailLogMutation(c.config, OpDelete)
	return &EmailLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EmailLogClient) DeleteOne(el *EmailLog) *EmailLogDeleteOne {
	return c.DeleteOneID(el.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EmailLogClient) DeleteOneID(id uuid.UUID) *EmailLogDeleteOne {
	builder := c.Delete().Where(emaillog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmailLogDeleteOne{builder}
}

// Query returns a query builder for EmailLog.
func (c *EmailLogClient) Query() *EmailLogQuery {
	return &EmailLogQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEmailLog},
		inters: c.Interceptors(),
	}
}

// Get returns a EmailLog entity by its id.
func (c *EmailLogClient) Get(ctx context.Context, id uuid.UUID) (*EmailLog, error) {
	return c.Query().Where(emaillog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmailLogClient) GetX(ctx context.Context, id uuid.UUID) *EmailLog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmailLogClient) Hooks() []Hook {
	return c.hooks.EmailLog
}

// Interceptors returns the client interceptors.
func (c *EmailLogClient) Interceptors() []Interceptor {
	return c.inters.EmailLog
}

func (c *EmailLogClient) mutate(ctx context.Context, m *EmailLogMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EmailLogCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EmailLogUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EmailLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EmailLogDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EmailLog mutation op: %q", m.Op())
	}
}

// EmailProviderClient is a client for the EmailProvider schema.
type EmailProviderClient struct {
	config
}

// NewEmailProviderClient returns a client for the EmailProvider from the given config.
func NewEmailProviderClient(c config) *EmailProviderClient {
	return &EmailProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `emailprovider.Hooks(f(g(h())))`.
func (c *EmailProviderClient) Use(hooks ...Hook) {
	c.hooks.EmailProvider = append(c.hooks.EmailProvider, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `emailprovider.Intercept(f(g(h())))`.
func (c *EmailProviderClient) Intercept(interceptors ...Interceptor) {
	c.inters.EmailProvider = append(c.inters.EmailProvider, interceptors...)
}

// Create returns a builder for creating a EmailProvider entity.
func (c *EmailProviderClient) Create() *EmailProviderCreate {
	mutation := newEmailProviderMutation(c.config, OpCreate)
	return &EmailProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EmailProvider entities.
func (c *EmailProviderClient) CreateBulk(builders ...*EmailProviderCreate) *EmailProviderCreateBulk {
	return &EmailProviderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EmailProviderClient) MapCreateBulk(slice any, setFunc func(*EmailProviderCreate, int)) *EmailProviderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EmailProviderCreateBulk{err: fmt.Errorf("calling to EmailProviderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EmailProviderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EmailProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EmailProvider.
func (c *EmailProviderClient) Update() *EmailProviderUpdate {
	mutation := newEmailProviderMutation(c.config, OpUpdate)
	return &EmailProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmailProviderClient) UpdateOne(ep *EmailProvider) *EmailProviderUpdateOne {
	mutation := newEmailProviderMutation(c.config, OpUpdateOne, withEmailProvider(ep))
	return &EmailProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmailProviderClient) UpdateOneID(id uint64) *EmailProviderUpdateOne {
	mutation := newEmailProviderMutation(c.config, OpUpdateOne, withEmailProviderID(id))
	return &EmailProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EmailProvider.
func (c *EmailProviderClient) Delete() *EmailProviderDelete {
	mutation := newEmailProviderMutation(c.config, OpDelete)
	return &EmailProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EmailProviderClient) DeleteOne(ep *EmailProvider) *EmailProviderDeleteOne {
	return c.DeleteOneID(ep.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EmailProviderClient) DeleteOneID(id uint64) *EmailProviderDeleteOne {
	builder := c.Delete().Where(emailprovider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmailProviderDeleteOne{builder}
}

// Query returns a query builder for EmailProvider.
func (c *EmailProviderClient) Query() *EmailProviderQuery {
	return &EmailProviderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEmailProvider},
		inters: c.Interceptors(),
	}
}

// Get returns a EmailProvider entity by its id.
func (c *EmailProviderClient) Get(ctx context.Context, id uint64) (*EmailProvider, error) {
	return c.Query().Where(emailprovider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmailProviderClient) GetX(ctx context.Context, id uint64) *EmailProvider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmailProviderClient) Hooks() []Hook {
	return c.hooks.EmailProvider
}

// Interceptors returns the client interceptors.
func (c *EmailProviderClient) Interceptors() []Interceptor {
	return c.inters.EmailProvider
}

func (c *EmailProviderClient) mutate(ctx context.Context, m *EmailProviderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EmailProviderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EmailProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EmailProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EmailProviderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EmailProvider mutation op: %q", m.Op())
	}
}

// SmsLogClient is a client for the SmsLog schema.
type SmsLogClient struct {
	config
}

// NewSmsLogClient returns a client for the SmsLog from the given config.
func NewSmsLogClient(c config) *SmsLogClient {
	return &SmsLogClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `smslog.Hooks(f(g(h())))`.
func (c *SmsLogClient) Use(hooks ...Hook) {
	c.hooks.SmsLog = append(c.hooks.SmsLog, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `smslog.Intercept(f(g(h())))`.
func (c *SmsLogClient) Intercept(interceptors ...Interceptor) {
	c.inters.SmsLog = append(c.inters.SmsLog, interceptors...)
}

// Create returns a builder for creating a SmsLog entity.
func (c *SmsLogClient) Create() *SmsLogCreate {
	mutation := newSmsLogMutation(c.config, OpCreate)
	return &SmsLogCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SmsLog entities.
func (c *SmsLogClient) CreateBulk(builders ...*SmsLogCreate) *SmsLogCreateBulk {
	return &SmsLogCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SmsLogClient) MapCreateBulk(slice any, setFunc func(*SmsLogCreate, int)) *SmsLogCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SmsLogCreateBulk{err: fmt.Errorf("calling to SmsLogClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SmsLogCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SmsLogCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SmsLog.
func (c *SmsLogClient) Update() *SmsLogUpdate {
	mutation := newSmsLogMutation(c.config, OpUpdate)
	return &SmsLogUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SmsLogClient) UpdateOne(sl *SmsLog) *SmsLogUpdateOne {
	mutation := newSmsLogMutation(c.config, OpUpdateOne, withSmsLog(sl))
	return &SmsLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SmsLogClient) UpdateOneID(id uuid.UUID) *SmsLogUpdateOne {
	mutation := newSmsLogMutation(c.config, OpUpdateOne, withSmsLogID(id))
	return &SmsLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SmsLog.
func (c *SmsLogClient) Delete() *SmsLogDelete {
	mutation := newSmsLogMutation(c.config, OpDelete)
	return &SmsLogDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SmsLogClient) DeleteOne(sl *SmsLog) *SmsLogDeleteOne {
	return c.DeleteOneID(sl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SmsLogClient) DeleteOneID(id uuid.UUID) *SmsLogDeleteOne {
	builder := c.Delete().Where(smslog.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SmsLogDeleteOne{builder}
}

// Query returns a query builder for SmsLog.
func (c *SmsLogClient) Query() *SmsLogQuery {
	return &SmsLogQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSmsLog},
		inters: c.Interceptors(),
	}
}

// Get returns a SmsLog entity by its id.
func (c *SmsLogClient) Get(ctx context.Context, id uuid.UUID) (*SmsLog, error) {
	return c.Query().Where(smslog.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SmsLogClient) GetX(ctx context.Context, id uuid.UUID) *SmsLog {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SmsLogClient) Hooks() []Hook {
	return c.hooks.SmsLog
}

// Interceptors returns the client interceptors.
func (c *SmsLogClient) Interceptors() []Interceptor {
	return c.inters.SmsLog
}

func (c *SmsLogClient) mutate(ctx context.Context, m *SmsLogMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SmsLogCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SmsLogUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SmsLogUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SmsLogDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SmsLog mutation op: %q", m.Op())
	}
}

// SmsProviderClient is a client for the SmsProvider schema.
type SmsProviderClient struct {
	config
}

// NewSmsProviderClient returns a client for the SmsProvider from the given config.
func NewSmsProviderClient(c config) *SmsProviderClient {
	return &SmsProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `smsprovider.Hooks(f(g(h())))`.
func (c *SmsProviderClient) Use(hooks ...Hook) {
	c.hooks.SmsProvider = append(c.hooks.SmsProvider, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `smsprovider.Intercept(f(g(h())))`.
func (c *SmsProviderClient) Intercept(interceptors ...Interceptor) {
	c.inters.SmsProvider = append(c.inters.SmsProvider, interceptors...)
}

// Create returns a builder for creating a SmsProvider entity.
func (c *SmsProviderClient) Create() *SmsProviderCreate {
	mutation := newSmsProviderMutation(c.config, OpCreate)
	return &SmsProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SmsProvider entities.
func (c *SmsProviderClient) CreateBulk(builders ...*SmsProviderCreate) *SmsProviderCreateBulk {
	return &SmsProviderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SmsProviderClient) MapCreateBulk(slice any, setFunc func(*SmsProviderCreate, int)) *SmsProviderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SmsProviderCreateBulk{err: fmt.Errorf("calling to SmsProviderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SmsProviderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SmsProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SmsProvider.
func (c *SmsProviderClient) Update() *SmsProviderUpdate {
	mutation := newSmsProviderMutation(c.config, OpUpdate)
	return &SmsProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SmsProviderClient) UpdateOne(sp *SmsProvider) *SmsProviderUpdateOne {
	mutation := newSmsProviderMutation(c.config, OpUpdateOne, withSmsProvider(sp))
	return &SmsProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SmsProviderClient) UpdateOneID(id uint64) *SmsProviderUpdateOne {
	mutation := newSmsProviderMutation(c.config, OpUpdateOne, withSmsProviderID(id))
	return &SmsProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SmsProvider.
func (c *SmsProviderClient) Delete() *SmsProviderDelete {
	mutation := newSmsProviderMutation(c.config, OpDelete)
	return &SmsProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SmsProviderClient) DeleteOne(sp *SmsProvider) *SmsProviderDeleteOne {
	return c.DeleteOneID(sp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SmsProviderClient) DeleteOneID(id uint64) *SmsProviderDeleteOne {
	builder := c.Delete().Where(smsprovider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SmsProviderDeleteOne{builder}
}

// Query returns a query builder for SmsProvider.
func (c *SmsProviderClient) Query() *SmsProviderQuery {
	return &SmsProviderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSmsProvider},
		inters: c.Interceptors(),
	}
}

// Get returns a SmsProvider entity by its id.
func (c *SmsProviderClient) Get(ctx context.Context, id uint64) (*SmsProvider, error) {
	return c.Query().Where(smsprovider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SmsProviderClient) GetX(ctx context.Context, id uint64) *SmsProvider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SmsProviderClient) Hooks() []Hook {
	return c.hooks.SmsProvider
}

// Interceptors returns the client interceptors.
func (c *SmsProviderClient) Interceptors() []Interceptor {
	return c.inters.SmsProvider
}

func (c *SmsProviderClient) mutate(ctx context.Context, m *SmsProviderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SmsProviderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SmsProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SmsProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SmsProviderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SmsProvider mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		EmailLog, EmailProvider, SmsLog, SmsProvider []ent.Hook
	}
	inters struct {
		EmailLog, EmailProvider, SmsLog, SmsProvider []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
