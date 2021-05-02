# Have I Been Redised

![https://redispwned.app](https://raw.githubusercontent.com/mikeblum/redispwned/redis-marketplace/website/static/images/website.png?raw=true)

![graphs about redis powered by redis](https://raw.githubusercontent.com/mikeblum/redispwned/redis-marketplace/website/static/images/website-graphs.png?raw=true)

![Architecture of redispwned.app](https://raw.githubusercontent.com/mikeblum/redispwned/redis-marketplace/website/static/images/architecture.png?raw=true)

Check it out at [RedisPwned](https://redispwned.app) and [Have I Been Redised](https://haveibeenredised.com)

## How it works

We scan the internet for exposed Redis databases broadcasting to the world on port `6379` - the default Redis port. This data is then corelated to a country of origin, redis version, etc and made searchable.

## How the data is stored:

Each dataset has its own key (`shodan:`, `censys:`, etc) that is then aggregated via a unified search index: `idx:redis-version-by-country-city-geo`. This generic search index allows us to easily aggregate data from different sources.

## How the data is accessed:

The `api.redispwned.app` fufills all aggregation and scan requests from `www.redispwned.app` (`www.haveibeenredised.com` as well). The graph

### Aggregate Exposed Redis Databases

Top 25 By Country

```
FT.AGGREGATE idx: "*" GROUPBY 1 @country REDUCE COUNT 0 AS num SORTBY 2 @num DESC MAX 25
```

Top 5 By Redis Version

```
FT.AGGREGATE idx: "*" GROUPBY 1 @version REDUCE COUNT 0 AS num SORTBY 2 @num DESC
```

To frustrate the script lords and kiddies I've omitting the actual ipv4 data used - all is readily obtainable with a research petition to the below sources.

## Sources

- Censys Data: [Censys Data](https://censys.io/ipv4?q=ports%3A+6379)
- Rapid7 Data: [Rapid7 Open Data](https://opendata.rapid7.com)
- Shodan Data: [Shodan Data](https://www.shodan.io/search?query=product%3ARedis)
