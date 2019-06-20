package service

import (
    "context"

    "../db"
    "../schema"
)

func Close(ctx context.Context) {
    db.Close(ctx)
}

func Insert(ctx context.Context, album *schema.Album) (int, error) {
    return db.Insert(ctx, album)
}

func Delete(ctx context.Context, id int) error {
    return db.Delete(ctx, id)
}

func GetAll(ctx context.Context) ([]schema.Album, error) {
    return db.GetAll(ctx)
}