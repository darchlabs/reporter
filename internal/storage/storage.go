package storage

import "github.com/go-redis/redis/v9"

type S struct {
	DB *redis.Client
}

func New(URL string) (*S, error) {
	db := redis.NewClient(&redis.Options{
		Addr:	  URL,
		DB:		  0,  // use default DB
	})

	return &S{
		DB: db,
	}, nil
}