package di

import (
	"koda-b8-backend1/internal/handler"
	"koda-b8-backend1/internal/repo"
	"koda-b8-backend1/internal/svc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	db *pgxpool.Pool

	userRepo    *repo.UserRepo
	userService *svc.UserService
	userHandler *handler.UserHandler
}

func (c *Container) initDeps() {
	c.userRepo = repo.NewUserRepo(c.db)
	c.userService = svc.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func NewContainer(db *pgxpool.Pool) *Container {
	container := &Container{
		db: db,
	}

	container.initDeps()

	return container
}
