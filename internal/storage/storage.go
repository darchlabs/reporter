package storage

import "github.com/go-redis/redis/v9"

type Storage struct {
	DB *redis.Client
}

func New(URL string) (*Storage, error) {
	db := redis.NewClient(&redis.Options{
		Addr:	  URL,
		DB:		  0,  // use default DB
	})

	return &Storage{
		DB: db,
	}, nil
}