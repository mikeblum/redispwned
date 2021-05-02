# Install Notes

Beware Docker won't bind to `:443` running in `--net=host` so `systemD` is a good alternative.

## Install Script

`git clone git@github.com:mikeblum/redispwned.git`
`./install.sh`

## Let's Encrypt Certificates

Configure Cloudflare DNS Challanges:

https://certbot-dns-cloudflare.readthedocs.io/en/stable/

Create certs group and add both root and redis to it:

```
sudo groupadd certs
sudo usermod -aG certs redis
sudo usermod -aG certs root
sudo chown root:certs -R /etc/certs/
```

Cronjob as `root`:

`cat /usr/local/bin/renew.sh`

```
#!/usr/bin/env bash

certbot renew
systemctl restart redispwned
```

`crontab -e as root`

```
45 2 * * * /usr/local/bin/renew.sh
```
