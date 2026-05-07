package registry

import (
	todo_module "gin-backend/internal/modules/todo"
	user_module "gin-backend/internal/modules/user"
	"gin-backend/internal/repository/mem_repo"
	todo_usecase "gin-backend/internal/domain/todo/usecase"
	user_usecase "gin-backend/internal/domain/user/usecase"
	"github.com/gin-gonic/gin"
)

type Container struct {
	TodoRouter *todo_module.Router
	UserRouter *user_module.Router
}

func NewContainer() *Container {
	// 1. Repositories
	todoRepo := mem_repo.NewTodoRepository()
	userRepo := mem_repo.NewUserRepository()

	// 2. Use Cases
	createTodoUC := todo_usecase.NewCreateTodoUseCase(todoRepo, userRepo)
	createUserUC := user_usecase.NewCreateUserUseCase(userRepo)

	// 3. Modules (Mutations/Queries)
	todoMutation := todo_module.NewMutation(createTodoUC)
	userMutation := user_module.NewMutation(createUserUC)

	// 4. Routers
	todoRouter := todo_module.NewRouter(todoMutation)
	userRouter := user_module.NewRouter(userMutation)

	return &Container{
		TodoRouter: todoRouter,
		UserRouter: userRouter,
	}
}

func (c *Container) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	c.TodoRouter.Register(api)
	c.UserRouter.Register(api)
}
