package shodan

import (
	"github.com/RediSearch/redisearch-go/redisearch"
)

func (s *Client) ServersByCountry() error {
	// `FT.AGGREGATE idx:shodan "*" GROUPBY 1 @country REDUCE COUNT 0 AS num SORTBY 2 @num DESC MAX 25`
	agg := redisearch.NewAggregateQuery().
		SetQuery(redisearch.NewQuery("*")).
		GroupBy(*redisearch.NewGroupBy().AddFields("@country").
			Reduce(*redisearch.NewReducerAlias(redisearch.GroupByReducerCount, []string{}, "count"))).
		SortBy([]redisearch.SortingKey{*redisearch.NewSortingKeyDir("@count", false)}).Limit(0, 25)

	results, _, err := s.idx.Aggregate(agg)
	s.aggregateToMap(results, []string{"country", "count"})
	return err
}

func (s *Client) ServersByVersion() error {
	// `FT.AGGREGATE idx:shodan "*" GROUPBY 1 @redis_version REDUCE COUNT 0 AS num SORTBY 2 @num DESC`
	agg := redisearch.NewAggregateQuery().SetQuery(redisearch.NewQuery("*")).
		GroupBy(*redisearch.NewGroupBy().AddFields("@redis_version").
			Reduce(*redisearch.NewReducerAlias(redisearch.GroupByReducerCount, []string{}, "count"))).
		SortBy([]redisearch.SortingKey{*redisearch.NewSortingKeyDir("@count", false)}).Limit(0, 100)

	results, _, err := s.idx.Aggregate(agg)
	s.aggregateToMap(results, []string{"redis_version", "count"})
	return err
}

func (s *Client) aggregateToMap(results [][]string, headers []string) {
	_ = make([]map[string]string, len(results))
	resultMap := make(map[string]string)
	for _, header := range headers {
		resultMap[header] = ""
	}
	for _, row := range results {
		for _, col := range row {
			s.log.Info(col)
		}
	}
}
