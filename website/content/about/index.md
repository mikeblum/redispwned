---
title: "About"
slug: /about
tags: ["about"]
description: "About Have I Been Redised"
toc: true
---

## What Is RedisPwned aka Have I Been Redised?

RedisPwned seeks to become a [Qualys SSL Labs](https://www.ssllabs.com/ssltest/) for Redis.

> For fun here is this site's score ;) [SSL Labs Score](https://www.ssllabs.com/ssltest/analyze.html?d=redispwned.app&latest)

[Have I Been Redised](https://haveibeenredised.com) ( [redispwned](https://redispwned.app) is much less to type :) ) is a project to help users and operators of Redis launch more secure and hardened databases.

Inspired by sites like [Have I Been Pwned](https://haveibeenpwned.com/) this service seeks to make users of Redis more aware of how to best secure it.

---

## What Is RedisPwned Not?

**A platform for pwning other Redis instances**

There is a reason services such as Shodan, Censys, and their like go to such lengths to gate access to this sort of information. There are a lot of bad actors out there on the internet! Script kiddies, edge lords, and griefers have united to make the public internet a scary place to be running servers especially databases!

**Associated with Redis or RedisLabs**

This site is not associated with, endorsed, or acknowledged by [RediLabs](https://redislabs.com/),  [Redis.io](https://redis.io/), or [@antirez](https://twitter.com/antirez?s=20) (Salvatore Sanfilippo the ex-BDFL of Redis)

---

## Who Made This And Why?

This was built for the [RedisLabs 2021 Hackathon](https://redislabs.com/hackathon-2021/)

I use Redis at work and on side projects and wanted to learn about the state of Redis deployments on the internet today.

- [me on the internet](https://mblum.me)
- [me on Github](https://github.com/mikeblum)
- [me on Twitter](https://twitter.com/roguequery)

---

## How Does This Work?

RedisPwned uses data from a number of sources to scan the internet for publicly exposed Redis servers to then help users scan their Redis infrastructure for issues in turn.

---

## Powered By

RedisPwned is an open-source project that:

#### Runs on ![Redis](/images/redis.svg) [of course ;)]

#### Built with ![Hugo](/images/hugo.svg)

#### Hosted by ![Amazon Web Services (AWS)](/images/aws.svg)

#### Secured by ![Let's Encrypt](/images/letsencrypt.svg)

#### Written in ![Golang](/images/go.svg)

This product includes GeoLite2 data created by MaxMind, available from [http://www.maxmind.com](http://www.maxmind.com)

As well as data from [Censys.io](https://censys.io/), [Rapid7](https://opendata.rapid7.com/), and [Shodan.io](https://shodan.io/) respectively.
