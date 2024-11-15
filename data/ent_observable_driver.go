package data

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ObservableDriver is a driver that logs and traces all driver operations.
type ObservableDriver struct {
	dialect.Driver                               // underlying driver.
	log            func(context.Context, ...any) // log function. defaults to log.Println.
	loggingEnabled bool
	tracingEnabled bool
}

// NewObservableDriver gets a driver and a logging function, and returns
// a new debugged-driver that prints all outgoing operations with context.
func NewObservableDriver(d dialect.Driver, logger func(context.Context, ...any), loggingEnabled, tracingEnabled bool) dialect.Driver {
	drv := &ObservableDriver{d, logger, loggingEnabled, tracingEnabled}
	return drv
}

func newQuerySpan(ctx context.Context, query string, args ...any) (context.Context, trace.Span) {
	tracer := otel.Tracer("database")
	kind := trace.SpanKindServer
	newCtx, span := tracer.Start(ctx,
		"SQL",
		trace.WithAttributes(
			attribute.String("sql", fmt.Sprintf("query=%v args=%v", query, args)),
		),
		trace.WithSpanKind(kind),
	)
	return newCtx, span
}

func newTxSpan(ctx context.Context, txId string) (context.Context, trace.Span) {
	tracer := otel.Tracer("database")
	kind := trace.SpanKindServer
	newCtx, span := tracer.Start(ctx,
		"Tx",
		trace.WithAttributes(
			attribute.String("txId", txId),
		),
		trace.WithSpanKind(kind),
	)
	return newCtx, span
}

// Exec logs its params and calls the underlying driver Exec method.
func (d *ObservableDriver) Exec(ctx context.Context, query string, args, v any) error {
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("driver.Exec: query=%v args=%v", query, args))
	}
	return d.Driver.Exec(ctx, query, args, v)
}

// ExecContext logs its params and calls the underlying driver ExecContext method if it is supported.
func (d *ObservableDriver) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Driver.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("driver.ExecContext: query=%v args=%v", query, args))
	}
	return drv.ExecContext(ctx, query, args...)
}

// Query logs its params and calls the underlying driver Query method.
func (d *ObservableDriver) Query(ctx context.Context, query string, args, v any) error {
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("driver.Query: query=%v args=%v", query, args))
	}
	return d.Driver.Query(ctx, query, args, v)
}

// QueryContext logs its params and calls the underlying driver QueryContext method if it is supported.
func (d *ObservableDriver) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Driver.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("driver.QueryContext: query=%v args=%v", query, args))
	}
	return drv.QueryContext(ctx, query, args...)
}

// Tx adds an log-id for the transaction and calls the underlying driver Tx command.
func (d *ObservableDriver) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := d.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	if d.tracingEnabled {
		newCtx, span := newTxSpan(ctx, id)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("driver.Tx(%s): started", id))
	}
	return &DebugTx{tx, id, d.log, ctx, d.loggingEnabled, d.tracingEnabled}, nil
}

// BeginTx adds an log-id for the transaction and calls the underlying driver BeginTx command if it is supported.
func (d *ObservableDriver) BeginTx(ctx context.Context, opts *sql.TxOptions) (dialect.Tx, error) {
	drv, ok := d.Driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.BeginTx is not supported")
	}
	tx, err := drv.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	if d.tracingEnabled {
		newCtx, span := newTxSpan(ctx, id)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("driver.BeginTx(%s): started", id))
	}
	return &DebugTx{tx, id, d.log, ctx, d.loggingEnabled, d.tracingEnabled}, nil
}

// DebugTx is a transaction implementation that logs all transaction operations.
type DebugTx struct {
	dialect.Tx                                   // underlying transaction.
	id             string                        // transaction logging id.
	log            func(context.Context, ...any) // log function. defaults to fmt.Println.
	ctx            context.Context               // underlying transaction context.
	loggingEnabled bool
	tracingEnabled bool
}

// Exec logs its params and calls the underlying transaction Exec method.
func (d *DebugTx) Exec(ctx context.Context, query string, args, v any) error {
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("Tx(%s).Exec: query=%v args=%v", d.id, query, args))
	}
	return d.Tx.Exec(ctx, query, args, v)
}

// ExecContext logs its params and calls the underlying transaction ExecContext method if it is supported.
func (d *DebugTx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	drv, ok := d.Tx.(interface {
		ExecContext(context.Context, string, ...any) (sql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.ExecContext is not supported")
	}
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("Tx(%s).ExecContext: query=%v args=%v", d.id, query, args))
	}
	return drv.ExecContext(ctx, query, args...)
}

// Query logs its params and calls the underlying transaction Query method.
func (d *DebugTx) Query(ctx context.Context, query string, args, v any) error {
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("Tx(%s).Query: query=%v args=%v", d.id, query, args))
	}
	return d.Tx.Query(ctx, query, args, v)
}

// QueryContext logs its params and calls the underlying transaction QueryContext method if it is supported.
func (d *DebugTx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	drv, ok := d.Tx.(interface {
		QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.QueryContext is not supported")
	}
	if d.tracingEnabled {
		newCtx, span := newQuerySpan(ctx, query, args)
		defer span.End()
		ctx = newCtx
	}
	if d.loggingEnabled {
		d.log(ctx, fmt.Sprintf("Tx(%s).QueryContext: query=%v args=%v", d.id, query, args))
	}
	return drv.QueryContext(ctx, query, args...)
}

// Commit logs this step and calls the underlying transaction Commit method.
func (d *DebugTx) Commit() error {
	if d.loggingEnabled {
		d.log(d.ctx, fmt.Sprintf("Tx(%s): committed", d.id))
	}
	return d.Tx.Commit()
}

// Rollback logs this step and calls the underlying transaction Rollback method.
func (d *DebugTx) Rollback() error {
	if d.loggingEnabled {
		d.log(d.ctx, fmt.Sprintf("Tx(%s): rollbacked", d.id))
	}
	return d.Tx.Rollback()
}
