# Have I Been Redised

Check it out at [RedisPwned](https://redispwned.app) and [Have I Been Redised](https://haveibeenredised.com)

![https://redispwned.app](https://raw.githubusercontent.com/mikeblum/redispwned/main/website/static/images/website.png)

![graphs about redis powered by redis](https://raw.githubusercontent.com/mikeblum/redispwned/main/website/static/images/website-graphs.png)

![Architecture of redispwned.app](https://raw.githubusercontent.com/mikeblum/redispwned/main/website/static/images/architecture.png)

## How it works

We scan the internet for exposed Redis databases broadcasting to the world on port `6379` - the default Redis port. This data is then corelated to a country of origin, redis version, etc and made searchable.

## How the data is stored:

Each dataset has its own key (`shodan:`, `censys:`, etc) that is then aggregated via a unified search index: `idx:redis-version-by-country-city-geo`. This generic search index allows us to easily aggregate data from different sources.

## How the data is accessed:

The `api.redispwned.app` fulfills all aggregation and scan requests from `www.redispwned.app` (`www.haveibeenredised.com` as well). The graphs are powered by querying the redispwned Redis database.

### Aggregate Exposed Redis Databases

Top 25 By Country

```
FT.AGGREGATE idx: "*" GROUPBY 1 @country REDUCE COUNT 0 AS num SORTBY 2 @num DESC MAX 25
```

Top 5 By Redis Version

```
FT.AGGREGATE idx: "*" GROUPBY 1 @version REDUCE COUNT 0 AS num SORTBY 2 @num DESC
```

To frustrate the script lords and kiddies I've omitted the actual ipv4 data used - although readily obtainable with a research petition to the below sources.

## Sources

- Censys Data: [Censys Data](https://censys.io/ipv4?q=ports%3A+6379)
- Rapid7 Data: [Rapid7 Open Data](https://opendata.rapid7.com)
- Shodan Data: [Shodan Data](https://www.shodan.io/search?query=product%3ARedis)
