package models

import (
	"context"

	"github.com/pkg/errors"
)

// Create creates a user in the database
func (u *User) Create(ctx context.Context) error {
	q := `
	INSERT INTO users
	(
		username,
		created_at,
		updated_at
	)
	VALUES
	(
		:username,
		CURRENT_TIMESTAMP,
		CURRENT_TIMESTAMP
	)
	RETURNING id;`

	r, err := db.NamedQueryContext(ctx, q, u)
	if err != nil {
		return errors.Wrap(err, "cannot create user")
	}

	r.Next()
	if err := r.Scan(&u.ID); err != nil {
		return errors.Wrap(err, "cannot scan row into id")
	}

	return nil
}

// func (u *User) update(ctx context.Context) error {
// 	q := `
// 	UPDATE users
// 	SET
// 	;`

// 	r, err := db.NamedExecContext(ctx, q, u)

// 	if err != nil {
// 		return errors.Wrap(err, "cannot create user")
// 	}

// 	id, err := r.LastInsertId()
// 	if err != nil {
// 		return errors.Wrap(err, "cannot get last inserted id")
// 	}

// 	rows, err := r.RowsAffected()
// 	if err != nil {
// 		return errors.Wrap(err, "error getting rows affected")
// 	}

// 	if rows != 1 {
// 		return errors.Errorf("rows affected is not 1 (acutal: %v)", rows)
// 	}

// 	u.ID = int(id)

// 	return nil
// }

// getUserByID returns a user and an optional error by ID
func getUserByID(ctx context.Context, id int) (*User, error) {
	q := `
	SELECT
		id,
		username
	FROM
		users
	WHERE
		id = $1;`

	var u User
	if err := db.GetContext(ctx, &u, q, id); err != nil {
		return nil, errors.Wrap(err, "error executing query with ctx")
	}

	return &u, nil
}
