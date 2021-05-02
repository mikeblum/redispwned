#!/usr/bin/env bash

go mod download
go mod verify
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/redispwned
# allow binding to :443
sudo setcap 'cap_net_bind_service=+ep' /usr/local/bin/redispwned
sudo chmod 755 /usr/local/bin/redispwned
# install SystemD service
sudo cp redispwned.service /lib/systemd/system/
sudo chmod 755 /lib/systemd/system/redispwned.service
sudo systemctl daemon-reload
sudo systemctl enable redispwned.service
sudo systemctl start redispwned
