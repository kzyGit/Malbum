package service

import (
    "context"

    "../db"
    "../schema"
)

func Close(ctx context.Context) {
    db.Close(ctx)
}

func Insert(ctx context.Context, album *schema.Album) (string, schema.Album, error) {
    return db.Insert(ctx, album)
}

func Delete(ctx context.Context, id int) error {
    return db.Delete(ctx, id)
}

func UpdateSong(ctx context.Context, id int, album *schema.Album)(string, schema.Album, error){
    return db.UpdateSong(ctx, id, album)
}
func GetAll(ctx context.Context) ([]schema.Album, error) {
    return db.GetAll(ctx)
}

func GetOne(ctx context.Context, id int) (schema.Album, error) {
    return db.GetOne(ctx, id)
}