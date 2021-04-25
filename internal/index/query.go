package index

import (
	"strconv"
	"strings"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type CountResult struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

func (idx *Manager) ServersByCountry() ([]CountResult, error) {
	// `FT.AGGREGATE idx: "*" GROUPBY 1 @country REDUCE COUNT 0 AS num SORTBY 2 @num DESC MAX 25`
	agg := redisearch.NewAggregateQuery().
		GroupBy(*redisearch.NewGroupBy().AddFields("@country").
			Reduce(*redisearch.NewReducerAlias(redisearch.GroupByReducerCount, []string{}, "count"))).
		SortBy([]redisearch.SortingKey{*redisearch.NewSortingKeyDir("@count", false)}).Limit(0, 25)

	results, _, err := idx.Aggregate(agg)
	mapping := idx.aggregateCountToMap(results, []string{"country", "count"})
	// Titleize country names
	for i, data := range mapping {
		data.Value = strings.Title(data.Value)
		mapping[i] = data
	}
	return mapping, err
}

func (idx *Manager) ServersByVersion() ([]CountResult, error) {
	// `FT.AGGREGATE idx: "*" GROUPBY 1 @version REDUCE COUNT 0 AS num SORTBY 2 @num DESC`
	agg := redisearch.NewAggregateQuery().
		GroupBy(*redisearch.NewGroupBy().AddFields("@version").
			Reduce(*redisearch.NewReducerAlias(redisearch.GroupByReducerCount, []string{}, "count"))).
		SortBy([]redisearch.SortingKey{*redisearch.NewSortingKeyDir("@count", false)}).Limit(0, 5)

	results, _, err := idx.Aggregate(agg)
	mapping := idx.aggregateCountToMap(results, []string{"version", "count"})
	return mapping, err
}

func (idx *Manager) aggregateCountToMap(results [][]string, headers []string) []CountResult {
	mapping := make([]CountResult, len(results))
	headerMap := make(map[string]string)
	for _, header := range headers {
		headerMap[header] = ""
	}
	for r, row := range results {
		result := CountResult{}
		for i, col := range row {
			if _, ok := headerMap[col]; ok {
				if strings.EqualFold(col, "count") {
					if count, err := strconv.Atoi(row[i+1]); err == nil {
						result.Count = count
					}
				} else {
					result.Value = row[i+1]
				}
			}
		}
		mapping[r] = result
	}
	return mapping
}
