package db

import (
    "context"

    "../schema"
)

const keyRepository = "Repository"

type Repository interface {
    Close()
    Insert(album *schema.Album) (int, error)
    Delete(id int) error
    GetAll() ([]schema.Album, error)
}

func SetRepository(ctx context.Context, repository Repository) context.Context {
    return context.WithValue(ctx, keyRepository, repository)
}

func Close(ctx context.Context) {
    getRepository(ctx).Close()
}

func Insert(ctx context.Context, album *schema.Album) (int, error) {
    return getRepository(ctx).Insert(album)
}

func Delete(ctx context.Context, id int) error {
    return getRepository(ctx).Delete(id)
}

func GetAll(ctx context.Context) ([]schema.Album, error) {
    return getRepository(ctx).GetAll()
}

func getRepository(ctx context.Context) Repository {
    return ctx.Value(keyRepository).(Repository)
}