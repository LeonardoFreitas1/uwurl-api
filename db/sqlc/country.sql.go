// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: country.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const getAllCountries = `-- name: GetAllCountries :many
SELECT id, name, official_state_name, tld, iso3166_2_A1, iso3166_2_A3 FROM country
`

type GetAllCountriesRow struct {
	ID                int32          `json:"id"`
	Name              string         `json:"name"`
	OfficialStateName sql.NullString `json:"official_state_name"`
	Tld               string         `json:"tld"`
	Iso31662A1        string         `json:"iso3166_2_a1"`
	Iso31662A3        string         `json:"iso3166_2_a3"`
}

func (q *Queries) GetAllCountries(ctx context.Context) ([]GetAllCountriesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllCountries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllCountriesRow{}
	for rows.Next() {
		var i GetAllCountriesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OfficialStateName,
			&i.Tld,
			&i.Iso31662A1,
			&i.Iso31662A3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCountryById = `-- name: GetCountryById :one
SELECT id, name, official_state_name, tld, iso3166_2_A1, iso3166_2_A3 FROM country where id = $1
`

type GetCountryByIdRow struct {
	ID                int32          `json:"id"`
	Name              string         `json:"name"`
	OfficialStateName sql.NullString `json:"official_state_name"`
	Tld               string         `json:"tld"`
	Iso31662A1        string         `json:"iso3166_2_a1"`
	Iso31662A3        string         `json:"iso3166_2_a3"`
}

func (q *Queries) GetCountryById(ctx context.Context, id int32) (GetCountryByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getCountryById, id)
	var i GetCountryByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.OfficialStateName,
		&i.Tld,
		&i.Iso31662A1,
		&i.Iso31662A3,
	)
	return i, err
}

const getFilteredCountry = `-- name: GetFilteredCountry :many
SELECT DISTINCT ON (ctr.id) id, name, official_state_name, tld, iso3166_2_A1, iso3166_2_A3
FROM country ctr
         LEFT JOIN country_language cl ON ctr.id = cl.country_id
WHERE ($1::int[] IS NULL OR language_id = ANY($1::int[]))
`

type GetFilteredCountryRow struct {
	ID                int32          `json:"id"`
	Name              string         `json:"name"`
	OfficialStateName sql.NullString `json:"official_state_name"`
	Tld               string         `json:"tld"`
	Iso31662A1        string         `json:"iso3166_2_a1"`
	Iso31662A3        string         `json:"iso3166_2_a3"`
}

func (q *Queries) GetFilteredCountry(ctx context.Context, dollar_1 []int32) ([]GetFilteredCountryRow, error) {
	rows, err := q.db.QueryContext(ctx, getFilteredCountry, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFilteredCountryRow{}
	for rows.Next() {
		var i GetFilteredCountryRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.OfficialStateName,
			&i.Tld,
			&i.Iso31662A1,
			&i.Iso31662A3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCountry = `-- name: InsertCountry :one
INSERT INTO country(
    name,
    official_state_name,
    tld,
    iso3166_2_a1,
    iso3166_2_a3,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
`

type InsertCountryParams struct {
	Name              string         `json:"name"`
	OfficialStateName sql.NullString `json:"official_state_name"`
	Tld               string         `json:"tld"`
	Iso31662A1        string         `json:"iso3166_2_a1"`
	Iso31662A3        string         `json:"iso3166_2_a3"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

func (q *Queries) InsertCountry(ctx context.Context, arg InsertCountryParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, insertCountry,
		arg.Name,
		arg.OfficialStateName,
		arg.Tld,
		arg.Iso31662A1,
		arg.Iso31662A3,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
