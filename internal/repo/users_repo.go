package repo

import (
	"context"
	"koda-b8-backend1/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db:db,
	}
}

func (r *UserRepo) Create(req *model.CreateUser) error{

	_, err := r.db.Exec(
		context.Background(),`
		INSERT INTO users (name,email,password)
		VALUES ($1,$2,$3) `,
		req.Name,
		req.Email,
		req.Password,
	)
	if err != nil {
		return err
	}
		return nil

	// id := len(*r.data) + 1

	// *r.data = append(*r.data, model.User{
	// 	ID:       int64(id),
	// 	Email:    req.Email,
	// 	Password: req.Password,
	// })
}

func (r *UserRepo) FindAll() ([]model.User, error) {

	rows, err := r.db.Query(
		context.Background(),
		`
		SELECT
			id,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users
		ORDER BY id ASC
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {

		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) FindByEmail(email string) *model.User {

	var user model.User

	err := r.db.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}

		return nil
	}

	return &user
}

func (r *UserRepo) FindByID(id int64) *model.User {

	var user model.User

	err := r.db.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			name,
			email,
			password,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil
		}

		return nil
	}

	return &user
}

func (r *UserRepo) Update(id int64, req *model.UpdateUser) error{
	_, err := r.db.Exec(
		context.Background(),`
		UPDATE users 
		SET
			name=$1,
			email=$2,
			password=$3,
			updated_at= NOW(),
		WHERE id=$4`, 
		req.Name,
		req.Email,
		req.Password,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}