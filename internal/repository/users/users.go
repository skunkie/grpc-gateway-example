package users

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/skunkie/grpc-gateway-example/internal/dto"
)

const table = "users"

type Users struct {
	db *sql.DB
}

func New(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) ListUsers(ctx context.Context, ids ...string) ([]*dto.User, error) {
	out := []*dto.User{}

	q := sq.Select("*").From(table).RunWith(u.db)
	if len(ids) > 0 {
		q = q.Where(sq.Eq{"id": ids})
	}

	rows, err := q.QueryContext(ctx)
	if err != nil {
		// it is ok, if nothing is found
		if errors.Is(err, sql.ErrNoRows) {
			return out, nil
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		u := &dto.User{}
		if err := rows.Scan(&u.Id, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		out = append(out, u)
	}

	return out, err
}

func (u *Users) CreateUser(ctx context.Context, in *dto.CreateUserReq) (*dto.User, error) {
	id := uuid.New().String()

	q := sq.Insert(table).Columns("id", "name", "email").Values(id, in.Name, in.Email).RunWith(u.db)
	if _, err := q.ExecContext(ctx); err != nil {
		return nil, err
	}

	return u.GetUser(ctx, id)
}

func (u *Users) GetUser(ctx context.Context, id string) (*dto.User, error) {
	out := &dto.User{}
	q := sq.Select("*").From(table).Where(sq.Eq{"id": id}).RunWith(u.db)
	err := q.ScanContext(ctx, &out.Id, &out.Name, &out.Email)
	if errors.Is(err, sql.ErrNoRows) {
		err = errIdNotFound(id)
	}

	return out, err
}

func (u *Users) UpdateUser(ctx context.Context, in *dto.UpdateUserReq) (*dto.User, error) {
	_, err := u.GetUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	q := sq.Update(table).Where(sq.Eq{"id": in.Id}).RunWith(u.db)
	if in.Name != nil {
		q = q.Set("name", in.Name)
	}
	if in.Email != nil {
		q = q.Set("email", in.Email)
	}

	_, err = q.ExecContext(ctx)
	if err != nil {
		return nil, err
	}

	return u.GetUser(ctx, in.Id)
}

func (u *Users) DeleteUser(ctx context.Context, id string) error {
	q := sq.Delete(table).Where(sq.Eq{"id": id}).RunWith(u.db)
	_, err := q.ExecContext(ctx)

	return err
}
