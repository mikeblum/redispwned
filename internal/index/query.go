package index

import (
	"github.com/RediSearch/redisearch-go/redisearch"
)

func (idx *Manager) ServersByCountry() error {
	// `FT.AGGREGATE idx: "*" GROUPBY 1 @country REDUCE COUNT 0 AS num SORTBY 2 @num DESC MAX 25`
	agg := redisearch.NewAggregateQuery().
		GroupBy(*redisearch.NewGroupBy().AddFields("@country").
			Reduce(*redisearch.NewReducerAlias(redisearch.GroupByReducerCount, []string{}, "count"))).
		SortBy([]redisearch.SortingKey{*redisearch.NewSortingKeyDir("@count", false)}).Limit(0, 25)

	results, _, err := idx.Aggregate(agg)
	idx.aggregateToMap(results, []string{"country", "count"})
	return err
}

func (idx *Manager) ServersByVersion() error {
	// `FT.AGGREGATE idx: "*" GROUPBY 1 @redis_version REDUCE COUNT 0 AS num SORTBY 2 @num DESC`
	agg := redisearch.NewAggregateQuery().
		GroupBy(*redisearch.NewGroupBy().AddFields("@version").
			Reduce(*redisearch.NewReducerAlias(redisearch.GroupByReducerCount, []string{}, "count"))).
		SortBy([]redisearch.SortingKey{*redisearch.NewSortingKeyDir("@count", false)}).Limit(0, 5)

	results, _, err := idx.Aggregate(agg)
	idx.aggregateToMap(results, []string{"version", "count"})
	return err
}

func (idx *Manager) aggregateToMap(results [][]string, headers []string) {
	_ = make([]map[string]string, len(results))
	resultMap := make(map[string]string)
	for _, header := range headers {
		resultMap[header] = ""
	}
	for _, row := range results {
		for i, col := range row {
			idx.log.Infof("%d %s", i, col)
		}
	}
}
