package shodan

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/go-redis/redis/v8"
)

const shodanDataJSONPath = "/etc/data/shodan/shodan-export.json"

func (s *Client) ImportShodanData(redisClient *redis.Client) error {
	ctx := context.Background()
	var err error
	dump, err := s.LoadFile(shodanDataJSONPath)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(dump)
	decoder := json.NewDecoder(reader)
	numRecords := 0
	pipe := redisClient.TxPipeline()
	for {
		var meta Redis
		if err = decoder.Decode(&meta); err == io.EOF {
			break
		} else if err != nil {
			s.log.Error("Error reading Shodan data: ", err)
			break
		}
		meta.ToHSet(ctx, pipe)
		numRecords++
	}
	results, err := pipe.Exec(ctx)
	if err != nil {
		s.log.Error("Failed to commit Shodan data: ", err)
	} else {
		s.log.Infof("Loaded %d / %d Shodan records", numRecords, len(results))
	}
	dump.Close()
	pipe.Close()
	return nil
}

func (s *Client) LoadFile(path string) (*os.File, error) {
	dump, err := os.Open(path)
	if err != nil {
		s.log.Error("Failed to load Shodan export data")
	}
	return dump, err
}
