package db

import (
	"database/sql"
	"strconv"
	"strings"
)

// A PipelineStmt is a simple wrapper for creating a statement consisting of
// a query and a set of arguments to be passed to that query.
type PipelineStmt struct {
	query string
	args  []interface{}
}

func NewPipelineStmt(query string, args ...interface{}) *PipelineStmt {
	return &PipelineStmt{query, args}
}

// Executes the statement within supplied transaction. The literal string `{LAST_INS_ID}`
// will be replaced with the supplied value to make chaining `PipelineStmt` objects together
// simple.
func (ps *PipelineStmt) Exec(tx Transaction, lastInsertId int64) (sql.Result, error) {
	query := strings.Replace(ps.query, "{LAST_INS_ID}", strconv.Itoa(int(lastInsertId)), -1)
	return tx.Exec(query, ps.args...)
}

// Runs the supplied statements within the transaction. If any statement fails, the transaction
// is rolled back, and the original error is returned.
//
// The `LastInsertId` from the previous statement will be passed to `Exec`. The zero-value (0) is
// used initially.
func RunPipeline(tx Transaction, stmts ...*PipelineStmt) (sql.Result, error) {
	var res sql.Result
	var err error
	var lastInsId int64

	for _, ps := range stmts {
		res, err = ps.Exec(tx, lastInsId)
		if err != nil {
			return nil, err
		}

		lastInsId, err = res.LastInsertId()
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
