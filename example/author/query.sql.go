// Code generated by pggen. DO NOT EDIT.

package author

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

// Querier is a typesafe Go interface backed by SQL queries.
//
// Methods ending with Batch enqueue a query to run later in a pgx.Batch. After
// calling SendBatch on pgx.Conn, pgxpool.Pool, or pgx.Tx, use the Scan methods
// to parse the results.
type Querier interface {
	// FindAuthorById finds one (or zero) authors by ID.
	FindAuthorByID(ctx context.Context, authorID int32) (FindAuthorByIDRow, error)
	// FindAuthorByIDBatch enqueues a FindAuthorByID query into batch to be executed
	// later by the batch.
	FindAuthorByIDBatch(batch *pgx.Batch, authorID int32)
	// FindAuthorByIDScan scans the result of an executed FindAuthorByIDBatch query.
	FindAuthorByIDScan(results pgx.BatchResults) (FindAuthorByIDRow, error)

	// FindAuthors finds authors by first name.
	FindAuthors(ctx context.Context, firstName string) ([]FindAuthorsRow, error)
	// FindAuthorsBatch enqueues a FindAuthors query into batch to be executed
	// later by the batch.
	FindAuthorsBatch(batch *pgx.Batch, firstName string)
	// FindAuthorsScan scans the result of an executed FindAuthorsBatch query.
	FindAuthorsScan(results pgx.BatchResults) ([]FindAuthorsRow, error)

	// FindAuthorNames finds one (or zero) authors by ID.
	FindAuthorNames(ctx context.Context, authorID int32) ([]FindAuthorNamesRow, error)
	// FindAuthorNamesBatch enqueues a FindAuthorNames query into batch to be executed
	// later by the batch.
	FindAuthorNamesBatch(batch *pgx.Batch, authorID int32)
	// FindAuthorNamesScan scans the result of an executed FindAuthorNamesBatch query.
	FindAuthorNamesScan(results pgx.BatchResults) ([]FindAuthorNamesRow, error)

	// DeleteAuthors deletes authors with a first name of "joe".
	DeleteAuthors(ctx context.Context) (pgconn.CommandTag, error)
	// DeleteAuthorsBatch enqueues a DeleteAuthors query into batch to be executed
	// later by the batch.
	DeleteAuthorsBatch(batch *pgx.Batch)
	// DeleteAuthorsScan scans the result of an executed DeleteAuthorsBatch query.
	DeleteAuthorsScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	// DeleteAuthorsByFirstName deletes authors by first name.
	DeleteAuthorsByFirstName(ctx context.Context, firstName string) (pgconn.CommandTag, error)
	// DeleteAuthorsByFirstNameBatch enqueues a DeleteAuthorsByFirstName query into batch to be executed
	// later by the batch.
	DeleteAuthorsByFirstNameBatch(batch *pgx.Batch, firstName string)
	// DeleteAuthorsByFirstNameScan scans the result of an executed DeleteAuthorsByFirstNameBatch query.
	DeleteAuthorsByFirstNameScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	// DeleteAuthorsByFullName deletes authors by the full name.
	DeleteAuthorsByFullName(ctx context.Context, params DeleteAuthorsByFullNameParams) (pgconn.CommandTag, error)
	// DeleteAuthorsByFullNameBatch enqueues a DeleteAuthorsByFullName query into batch to be executed
	// later by the batch.
	DeleteAuthorsByFullNameBatch(batch *pgx.Batch, params DeleteAuthorsByFullNameParams)
	// DeleteAuthorsByFullNameScan scans the result of an executed DeleteAuthorsByFullNameBatch query.
	DeleteAuthorsByFullNameScan(results pgx.BatchResults) (pgconn.CommandTag, error)

	// InsertAuthor inserts an author by name and returns the ID.
	InsertAuthor(ctx context.Context, firstName string, lastName string) (int32, error)
	// InsertAuthorBatch enqueues a InsertAuthor query into batch to be executed
	// later by the batch.
	InsertAuthorBatch(batch *pgx.Batch, firstName string, lastName string)
	// InsertAuthorScan scans the result of an executed InsertAuthorBatch query.
	InsertAuthorScan(results pgx.BatchResults) (int32, error)

	// InsertAuthorSuffix inserts an author by name and suffix and returns the
	// entire row.
	InsertAuthorSuffix(ctx context.Context, params InsertAuthorSuffixParams) (InsertAuthorSuffixRow, error)
	// InsertAuthorSuffixBatch enqueues a InsertAuthorSuffix query into batch to be executed
	// later by the batch.
	InsertAuthorSuffixBatch(batch *pgx.Batch, params InsertAuthorSuffixParams)
	// InsertAuthorSuffixScan scans the result of an executed InsertAuthorSuffixBatch query.
	InsertAuthorSuffixScan(results pgx.BatchResults) (InsertAuthorSuffixRow, error)
}

type DBQuerier struct {
	conn genericConn
}

var _ Querier = &DBQuerier{}

// genericConn is a connection to a Postgres database. This is usually backed by
// *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
type genericConn interface {
	// Query executes sql with args. If there is an error the returned Rows will
	// be returned in an error state. So it is allowed to ignore the error
	// returned from Query and handle it in Rows.
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)

	// QueryRow is a convenience wrapper over Query. Any error that occurs while
	// querying is deferred until calling Scan on the returned Row. That Row will
	// error with pgx.ErrNoRows if no rows are returned.
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row

	// Exec executes sql. sql can be either a prepared statement name or an SQL
	// string. arguments should be referenced positionally from the sql string
	// as $1, $2, etc.
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// NewQuerier creates a DBQuerier that implements Querier. conn is typically
// *pgx.Conn, pgx.Tx, or *pgxpool.Pool.
func NewQuerier(conn genericConn) *DBQuerier {
	return &DBQuerier{
		conn: conn,
	}
}

// WithTx creates a new DBQuerier that uses the transaction to run all queries.
func (q *DBQuerier) WithTx(tx pgx.Tx) (*DBQuerier, error) {
	return &DBQuerier{conn: tx}, nil
}

const findAuthorByIDSQL = `SELECT * FROM author WHERE author_id = $1;`

type FindAuthorByIDRow struct {
	AuthorID  int32       `json:"author_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Suffix    pgtype.Text `json:"suffix"`
}

// FindAuthorByID implements Querier.FindAuthorByID.
func (q *DBQuerier) FindAuthorByID(ctx context.Context, authorID int32) (FindAuthorByIDRow, error) {
	row := q.conn.QueryRow(ctx, findAuthorByIDSQL, authorID)
	var item FindAuthorByIDRow
	if err := row.Scan(&item.AuthorID, &item.FirstName, &item.LastName, &item.Suffix); err != nil {
		return item, fmt.Errorf("query FindAuthorByID: %w", err)
	}
	return item, nil
}

// FindAuthorByIDBatch implements Querier.FindAuthorByIDBatch.
func (q *DBQuerier) FindAuthorByIDBatch(batch *pgx.Batch, authorID int32) {
	batch.Queue(findAuthorByIDSQL, authorID)
}

// FindAuthorByIDScan implements Querier.FindAuthorByIDScan.
func (q *DBQuerier) FindAuthorByIDScan(results pgx.BatchResults) (FindAuthorByIDRow, error) {
	row := results.QueryRow()
	var item FindAuthorByIDRow
	if err := row.Scan(&item.AuthorID, &item.FirstName, &item.LastName, &item.Suffix); err != nil {
		return item, fmt.Errorf("scan FindAuthorByIDBatch row: %w", err)
	}
	return item, nil
}

const findAuthorsSQL = `SELECT * FROM author WHERE first_name = $1;`

type FindAuthorsRow struct {
	AuthorID  int32       `json:"author_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Suffix    pgtype.Text `json:"suffix"`
}

// FindAuthors implements Querier.FindAuthors.
func (q *DBQuerier) FindAuthors(ctx context.Context, firstName string) ([]FindAuthorsRow, error) {
	rows, err := q.conn.Query(ctx, findAuthorsSQL, firstName)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("query FindAuthors: %w", err)
	}
	items := []FindAuthorsRow{}
	for rows.Next() {
		var item FindAuthorsRow
		if err := rows.Scan(&item.AuthorID, &item.FirstName, &item.LastName, &item.Suffix); err != nil {
			return nil, fmt.Errorf("scan FindAuthors row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAuthors rows: %w", err)
	}
	return items, err
}

// FindAuthorsBatch implements Querier.FindAuthorsBatch.
func (q *DBQuerier) FindAuthorsBatch(batch *pgx.Batch, firstName string) {
	batch.Queue(findAuthorsSQL, firstName)
}

// FindAuthorsScan implements Querier.FindAuthorsScan.
func (q *DBQuerier) FindAuthorsScan(results pgx.BatchResults) ([]FindAuthorsRow, error) {
	rows, err := results.Query()
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, err
	}
	items := []FindAuthorsRow{}
	for rows.Next() {
		var item FindAuthorsRow
		if err := rows.Scan(&item.AuthorID, &item.FirstName, &item.LastName, &item.Suffix); err != nil {
			return nil, fmt.Errorf("scan FindAuthorsBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAuthorsBatch rows: %w", err)
	}
	return items, err
}

const findAuthorNamesSQL = `SELECT first_name, last_name FROM author ORDER BY author_id = $1;`

type FindAuthorNamesRow struct {
	FirstName pgtype.Text `json:"first_name"`
	LastName  pgtype.Text `json:"last_name"`
}

// FindAuthorNames implements Querier.FindAuthorNames.
func (q *DBQuerier) FindAuthorNames(ctx context.Context, authorID int32) ([]FindAuthorNamesRow, error) {
	rows, err := q.conn.Query(ctx, findAuthorNamesSQL, authorID)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("query FindAuthorNames: %w", err)
	}
	items := []FindAuthorNamesRow{}
	for rows.Next() {
		var item FindAuthorNamesRow
		if err := rows.Scan(&item.FirstName, &item.LastName); err != nil {
			return nil, fmt.Errorf("scan FindAuthorNames row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAuthorNames rows: %w", err)
	}
	return items, err
}

// FindAuthorNamesBatch implements Querier.FindAuthorNamesBatch.
func (q *DBQuerier) FindAuthorNamesBatch(batch *pgx.Batch, authorID int32) {
	batch.Queue(findAuthorNamesSQL, authorID)
}

// FindAuthorNamesScan implements Querier.FindAuthorNamesScan.
func (q *DBQuerier) FindAuthorNamesScan(results pgx.BatchResults) ([]FindAuthorNamesRow, error) {
	rows, err := results.Query()
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, err
	}
	items := []FindAuthorNamesRow{}
	for rows.Next() {
		var item FindAuthorNamesRow
		if err := rows.Scan(&item.FirstName, &item.LastName); err != nil {
			return nil, fmt.Errorf("scan FindAuthorNamesBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindAuthorNamesBatch rows: %w", err)
	}
	return items, err
}

const deleteAuthorsSQL = `DELETE FROM author WHERE first_name = 'joe';`

// DeleteAuthors implements Querier.DeleteAuthors.
func (q *DBQuerier) DeleteAuthors(ctx context.Context) (pgconn.CommandTag, error) {
	cmdTag, err := q.conn.Exec(ctx, deleteAuthorsSQL)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteAuthors: %w", err)
	}
	return cmdTag, err
}

// DeleteAuthorsBatch implements Querier.DeleteAuthorsBatch.
func (q *DBQuerier) DeleteAuthorsBatch(batch *pgx.Batch) {
	batch.Queue(deleteAuthorsSQL)
}

// DeleteAuthorsScan implements Querier.DeleteAuthorsScan.
func (q *DBQuerier) DeleteAuthorsScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec DeleteAuthorsBatch: %w", err)
	}
	return cmdTag, err
}

const deleteAuthorsByFirstNameSQL = `DELETE FROM author WHERE first_name = $1;`

// DeleteAuthorsByFirstName implements Querier.DeleteAuthorsByFirstName.
func (q *DBQuerier) DeleteAuthorsByFirstName(ctx context.Context, firstName string) (pgconn.CommandTag, error) {
	cmdTag, err := q.conn.Exec(ctx, deleteAuthorsByFirstNameSQL, firstName)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteAuthorsByFirstName: %w", err)
	}
	return cmdTag, err
}

// DeleteAuthorsByFirstNameBatch implements Querier.DeleteAuthorsByFirstNameBatch.
func (q *DBQuerier) DeleteAuthorsByFirstNameBatch(batch *pgx.Batch, firstName string) {
	batch.Queue(deleteAuthorsByFirstNameSQL, firstName)
}

// DeleteAuthorsByFirstNameScan implements Querier.DeleteAuthorsByFirstNameScan.
func (q *DBQuerier) DeleteAuthorsByFirstNameScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec DeleteAuthorsByFirstNameBatch: %w", err)
	}
	return cmdTag, err
}

const deleteAuthorsByFullNameSQL = `DELETE
FROM author
WHERE first_name = $1
  AND last_name = $2
  AND suffix = $3;`

type DeleteAuthorsByFullNameParams struct {
	FirstName string
	LastName  string
	Suffix    string
}

// DeleteAuthorsByFullName implements Querier.DeleteAuthorsByFullName.
func (q *DBQuerier) DeleteAuthorsByFullName(ctx context.Context, params DeleteAuthorsByFullNameParams) (pgconn.CommandTag, error) {
	cmdTag, err := q.conn.Exec(ctx, deleteAuthorsByFullNameSQL, params.FirstName, params.LastName, params.Suffix)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query DeleteAuthorsByFullName: %w", err)
	}
	return cmdTag, err
}

// DeleteAuthorsByFullNameBatch implements Querier.DeleteAuthorsByFullNameBatch.
func (q *DBQuerier) DeleteAuthorsByFullNameBatch(batch *pgx.Batch, params DeleteAuthorsByFullNameParams) {
	batch.Queue(deleteAuthorsByFullNameSQL, params.FirstName, params.LastName, params.Suffix)
}

// DeleteAuthorsByFullNameScan implements Querier.DeleteAuthorsByFullNameScan.
func (q *DBQuerier) DeleteAuthorsByFullNameScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec DeleteAuthorsByFullNameBatch: %w", err)
	}
	return cmdTag, err
}

const insertAuthorSQL = `INSERT INTO author (first_name, last_name)
VALUES ($1, $2)
RETURNING author_id;`

// InsertAuthor implements Querier.InsertAuthor.
func (q *DBQuerier) InsertAuthor(ctx context.Context, firstName string, lastName string) (int32, error) {
	row := q.conn.QueryRow(ctx, insertAuthorSQL, firstName, lastName)
	var item int32
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query InsertAuthor: %w", err)
	}
	return item, nil
}

// InsertAuthorBatch implements Querier.InsertAuthorBatch.
func (q *DBQuerier) InsertAuthorBatch(batch *pgx.Batch, firstName string, lastName string) {
	batch.Queue(insertAuthorSQL, firstName, lastName)
}

// InsertAuthorScan implements Querier.InsertAuthorScan.
func (q *DBQuerier) InsertAuthorScan(results pgx.BatchResults) (int32, error) {
	row := results.QueryRow()
	var item int32
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan InsertAuthorBatch row: %w", err)
	}
	return item, nil
}

const insertAuthorSuffixSQL = `INSERT INTO author (first_name, last_name, suffix)
VALUES ($1, $2, $3)
RETURNING author_id, first_name, last_name, suffix;`

type InsertAuthorSuffixParams struct {
	FirstName string
	LastName  string
	Suffix    string
}

type InsertAuthorSuffixRow struct {
	AuthorID  int32       `json:"author_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Suffix    pgtype.Text `json:"suffix"`
}

// InsertAuthorSuffix implements Querier.InsertAuthorSuffix.
func (q *DBQuerier) InsertAuthorSuffix(ctx context.Context, params InsertAuthorSuffixParams) (InsertAuthorSuffixRow, error) {
	row := q.conn.QueryRow(ctx, insertAuthorSuffixSQL, params.FirstName, params.LastName, params.Suffix)
	var item InsertAuthorSuffixRow
	if err := row.Scan(&item.AuthorID, &item.FirstName, &item.LastName, &item.Suffix); err != nil {
		return item, fmt.Errorf("query InsertAuthorSuffix: %w", err)
	}
	return item, nil
}

// InsertAuthorSuffixBatch implements Querier.InsertAuthorSuffixBatch.
func (q *DBQuerier) InsertAuthorSuffixBatch(batch *pgx.Batch, params InsertAuthorSuffixParams) {
	batch.Queue(insertAuthorSuffixSQL, params.FirstName, params.LastName, params.Suffix)
}

// InsertAuthorSuffixScan implements Querier.InsertAuthorSuffixScan.
func (q *DBQuerier) InsertAuthorSuffixScan(results pgx.BatchResults) (InsertAuthorSuffixRow, error) {
	row := results.QueryRow()
	var item InsertAuthorSuffixRow
	if err := row.Scan(&item.AuthorID, &item.FirstName, &item.LastName, &item.Suffix); err != nil {
		return item, fmt.Errorf("scan InsertAuthorSuffixBatch row: %w", err)
	}
	return item, nil
}
