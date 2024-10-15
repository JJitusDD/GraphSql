package repository

import (
	"context"
	"project-test/internal/domain/model/entities"
)

type IUser interface {
	Create(ctx context.Context, in *entities.User) (ou *entities.User, err error)
	Update(ctx context.Context, id string, in *entities.User) (ou *entities.User, err error)
	Find(ctx context.Context, keysAndValues ...string) (ou []*entities.User, err error)
	FindById(ctx context.Context, id string) (o *entities.User, e error)
	Delete(ctx context.Context, id string) (e error)
}
