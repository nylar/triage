// package sqrl provides a fluent SQL generator.
//
// See https://github.com/elgris/sqrl for examples.
package sqrl

import (
	"context"
	"database/sql"
	"fmt"
)

// Sqlizer is the interface that wraps the ToSql method.
//
// ToSql returns a SQL representation of the Sqlizer, along with a slice of args
// as passed to e.g. database/sql.Exec. It can also return an error.
type Sqlizer interface {
	ToSql() (string, []interface{}, error)
}

// Execer is the interface that wraps the Exec method.
//
// Exec executes the given query as implemented by database/sql.Exec.
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// ExecerContext is the interface that wraps the Exec method.
//
// ExecContext executes the given query using given context as implemented by database/sql.ExecContext.
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// Queryer is the interface that wraps the Query method.
//
// Query executes the given query as implemented by database/sql.Query.
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// QueryerContext is the interface that wraps the Query method.
//
// QueryerContext executes the given query using given context as implemented by database/sql.QueryContext.
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// QueryRower is the interface that wraps the QueryRow method.
//
// QueryRow executes the given query as implemented by database/sql.QueryRow.
type QueryRower interface {
	QueryRow(query string, args ...interface{}) RowScanner
}

// QueryRowerContext is the interface that wraps the QueryRow method.
//
// QueryRowContext executes the given query using given context as implemented by database/sql.QueryRowContext.
type QueryRowerContext interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) RowScanner
}

// BaseRunner groups the Execer and Queryer interfaces.
type BaseRunner interface {
	Execer
	ExecerContext
	Queryer
	QueryerContext
}

// Runner groups the Execer, Queryer, and QueryRower interfaces.
type Runner interface {
	Execer
	ExecerContext
	Queryer
	QueryerContext
	QueryRower
	QueryRowerContext
}

// ErrRunnerNotSet is returned by methods that need a Runner if it isn't set.
var ErrRunnerNotSet = fmt.Errorf("cannot run; no Runner set (RunWith)")

// ErrRunnerNotQueryRunner is returned by QueryRow if the RunWith value doesn't implement QueryRower.
var ErrRunnerNotQueryRunner = fmt.Errorf("cannot QueryRow; Runner is not a QueryRower")

// ErrRunnerNotQueryRunnerContext is returned by QueryRowContext if the RunWith value doesn't implement QueryRowerContext.
var ErrRunnerNotQueryRunnerContext = fmt.Errorf("cannot QueryRow; Runner is not a QueryRowerContext")

// ExecWith Execs the SQL returned by s with db.
func ExecWith(db Execer, s Sqlizer) (res sql.Result, err error) {
	query, args, err := s.ToSql()
	if err != nil {
		return
	}
	return db.Exec(query, args...)
}

// ExecWithContext Execs the SQL returned by s with db.
func ExecWithContext(ctx context.Context, db ExecerContext, s Sqlizer) (res sql.Result, err error) {
	query, args, err := s.ToSql()
	if err != nil {
		return
	}
	return db.ExecContext(ctx, query, args...)
}

// QueryWith Querys the SQL returned by s with db.
func QueryWith(db Queryer, s Sqlizer) (rows *sql.Rows, err error) {
	query, args, err := s.ToSql()
	if err != nil {
		return
	}
	return db.Query(query, args...)
}

// QueryWithContext Querys the SQL returned by s with db.
func QueryWithContext(ctx context.Context, db QueryerContext, s Sqlizer) (rows *sql.Rows, err error) {
	query, args, err := s.ToSql()
	if err != nil {
		return
	}
	return db.QueryContext(ctx, query, args...)
}

// QueryRowWith QueryRows the SQL returned by s with db.
func QueryRowWith(db QueryRower, s Sqlizer) RowScanner {
	query, args, err := s.ToSql()
	return &Row{RowScanner: db.QueryRow(query, args...), err: err}
}

// QueryRowWithContext QueryRows the SQL returned by s with db.
func QueryRowWithContext(ctx context.Context, db QueryRowerContext, s Sqlizer) RowScanner {
	query, args, err := s.ToSql()
	return &Row{RowScanner: db.QueryRowContext(ctx, query, args...), err: err}
}
