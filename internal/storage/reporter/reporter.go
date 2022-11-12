package reporterstorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/darchlabs/reporter"
	"github.com/darchlabs/reporter/internal/storage"
	"github.com/go-redis/redis/v9"
)

const prefix = "group-reports"

type Storage struct {
	storage *storage.S
}

func New(s *storage.S) *Storage {
	return &Storage {
		storage: s,
	}
}

func (s *Storage) InsertGroupReport(g * reporter.GroupReport) error {
	ctx := context.Background()

	// format the composed key used in db
	key := fmt.Sprintf("%s:%s:%v", prefix, g.Type, g.CreatedAt.Unix())

	// check if key already exist in database
	current, _ := s.GetGroupReport(g.CreatedAt.Unix(), g.Type)
	if current != nil {
		return fmt.Errorf("key=%s already exist in db", key)
	}

	// parse strut to bytes
	b, err := json.Marshal(g)
	if err != nil {
		return err
	}

	// save in database
	err = s.storage.DB.Set(ctx, key, []byte(b), 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetGroupReport(key int64, t reporter.ServiceType) (*reporter.GroupReport, error) {
	ctx := context.Background()

	// format the composed key used in db
	parsedKey := fmt.Sprintf("%s:%s:%v", prefix, t, key)

	// get bytes by composed key to db
	str, err :=	s.storage.DB.Get(ctx, parsedKey).Result()
	if err == redis.Nil {
		return nil, errors.New("key does not exist")
	} 
	if err != nil {
		return nil, err
	} 
		
	// parse bytes to GroupReport struct
	var groupReport *reporter.GroupReport
	err = json.Unmarshal([]byte(str), &groupReport)
	if err != nil {
		return nil, err
	}

	return groupReport, nil
}

func (s *Storage) ListGroupReports(t reporter.ServiceType) ([]*reporter.GroupReport, error) {
	ctx := context.Background()
	
	// prepare slice of group reports
	groupReports := make([]*reporter.GroupReport, 0)

	// // iterte over db elements and push in slice
	iter := s.storage.DB.Scan(ctx, 0, fmt.Sprintf("%s:%s:*", prefix, t), 0).Iterator()
	for iter.Next(ctx) {
		// split key using prefix
		split := strings.Split(iter.Val(), ":")

		// get int64 key
		key, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return nil, err
		}

		// get value from redis
		groupReport, err := s.GetGroupReport(key, t)
		if err != nil {
			return nil, err
		}

		// append new element to group reports slice
		groupReports = append(groupReports, groupReport)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	// return groupReports, nil
	return groupReports, nil
}