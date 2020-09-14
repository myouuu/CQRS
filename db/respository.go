package db
import (
	"context"

	"github.com/myouuu/CQRS/schema"
)

type Repository interface {
	Close()
	InsertMeow(ctx context.Context, CQRS schema.CQRS) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.CQRS, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func InsertMeow(ctx context.Context, meow schema.CQRS) error {
	return impl.InsertMeow(ctx, meow)
}

func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.CQRS, error) {
	return impl.ListMeows(ctx, skip, take)
}