#!/usr/bin/env bash

docker run -dit \
	   --restart unless-stopped \
	   -e GIN_MODE=release \
	   -v /var/app/redispwned/config.env:/config.env \
	   --network host \
	   redispwned
