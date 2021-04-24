---
title: "Securing Redis"
slug: /securing-redis
tags: ["how-to", "security", "redis"]
description: "Securing Redis"
---

## Securing Redis on the Internet

It's best to start by reading the manual aka RTFM!

### Articles and Sources

- [Redis Security](https://redis.io/topics/security)
- [DigitalOcean Tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-redis-on-ubuntu-20-04)
- [A few things about Redis security](http://antirez.com/news/96)

All read up on Redis security? Let's recap!

### Rule #1: Don't Expose Redis to the Public Internet!!

On its face this may seem counter-intuitive - RedisLabs runs servers all over the world that are publicly available - why can't you?

As is true of any SaaS provider (AWS with ElastiCache, RedisLabs, etc) - they have dedicated engineering teams to harden and monitor every aspect of running a database like Redis in the cloud. Most side projects, companies, and individual contributors have neither the time, money, nor desire to devote themselves to hardening and continuously monitoring their publicly exposed databases (be it Redis or some other database for that matter).

From the quickstart guide we find:

> Note that a Redis exposed to the internet without any security is very simple to exploit, so make sure you understand the above and apply at least a firewalling layer. After the firewalling is in place, try to connect with redis-cli from an external host in order to prove yourself the instance is actually not reachable.

One might ask though doesn't `Protected Mode` solve this problem for good?

> We expect protected mode to seriously decrease the security issues caused by unprotected Redis instances executed without proper administration, however the system administrator can still ignore the error given by Redis and just disable protected mode or manually bind all the interfaces.

By default the `redis.conf` doesn't ship with a secure-by-default configuration. This makes it easy to get started but opens new challenges when deploying services that use Redis on the internet.

### Rule #2: Use AUTH and ACLs

Redis Access Control Lists (ACLs) are new in Redis 6.x and give much finer0grained control over how users (as well as your application) interact with Redis. I'd give Redis Lab's post on the topic a read:

[Getting Started With Redis 6 ACLs](https://redislabs.com/blog/getting-started-redis-6-access-control-lists-acls/)

as well as the `redis.io docs:

[Redis ACLs](https://redis.io/topics/acl)

### Rule #3: Don't commit your redis.conf to git!

Redis is easy to configure and use. All an operator needs to do is add this to enforce passsword authentication:

```
requirepass "hello world"
```

If you were to have the public IP address of your Redis instance in your conf a malicious attacker now knows where your Redis server is and what the password is. Not bueno!

While testing various git secrets tooling none of these detected `requirepass`:

- [gitleaks](https://github.com/zricethezav/gitleaks)
- [git-secrets](https://github.com/awslabs/git-secrets)

Regardless - its' a bad idea to hard-code secrets into your git repos - especially database passwords!
