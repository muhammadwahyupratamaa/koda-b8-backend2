package repo

import (
	"context"
	"koda-b8-backend1/internal/model"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(req *model.CreateUser) error {

	_, err := r.db.Exec(
		context.Background(), `
		INSERT INTO users (name,email,password,picture)
		VALUES ($1,$2,$3,$4) `,
		req.Name,
		req.Email,
		req.Password,
		req.Picture,
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

func (r *UserRepo) FindAll(
	page,
	limit int,
	name,
	email,
	sortID,
	sortName,
	sortEmail string,
) ([]model.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password,
			COALESCE(picture, '') AS picture,
			created_at,
			updated_at
		FROM users
	`
	offset := (page - 1) * limit

	var args []any

	if name != "" {
		query += " WHERE name ILIKE $1"
		args = append(args, "%"+name+"%")
	} else if email != "" {
		query += " WHERE email ILIKE $1"
		args = append(args, "%"+email+"%")
	}

	orderBy := "id ASC"

	switch {
	case sortID == "asc":
		orderBy = "id ASC"
	case sortID == "desc":
		orderBy = "id DESC"

	case sortName == "asc":
		orderBy = "name ASC"
	case sortName == "desc":
		orderBy = "name DESC"

	case sortEmail == "asc":
		orderBy = "email ASC"
	case sortEmail == "desc":
		orderBy = "email DESC"
	}

	query += " ORDER BY " + orderBy

	query += " LIMIT $" + strconv.Itoa(len(args)+1)
	args = append(args, limit)

	query += " OFFSET $" + strconv.Itoa(len(args)+1)
	args = append(args, offset)

	rows, err := r.db.Query(
		context.Background(),
		query,
		args...,
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
			&user.Picture,
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
			COALESCE(picture, '') AS picture,
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
		&user.Picture,
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
			COALESCE(picture, '') AS picture,
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
		&user.Picture,
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

func (r *UserRepo) Update(id int64, req *model.UpdateUser) error {
	_, err := r.db.Exec(
		context.Background(), `
		UPDATE users 
		SET
			name=$1,
			email=$2,
			password=$3,
			picture=$4,
			updated_at= NOW()
		WHERE id=$5`,
		req.Name,
		req.Email,
		req.Password,
		req.Picture,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepo) Delete(id int64) error {
	_, err := r.db.Exec(
		context.Background(), `
		DELETE FROM users
		WHERE id=$1`,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
