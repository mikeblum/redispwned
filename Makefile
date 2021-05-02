.PHONY: test
test:
	go test -count=1 -race ./... -test.v

.PHONY: run
run:
	go run -race main.go

.PHONY: docker-build
docker-build: build
	docker build -t redispwned .

docker-run: run
	docker run -dit \
	   --restart unless-stopped \
	   -e GIN_MODE=debug \
	   -v /var/app/redispwned/config.env:/config.env \
	   --network host \
	   redispwned

.PHONY: install
install:
	go mod download
	go mod verify
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o redispwned
	sudo mv redispwned /usr/local/bin/
	sudo chown redis:redis /usr/local/bin/redispwned
	# allow binding to :443
	sudo setcap 'cap_net_bind_service=+ep' /usr/local/bin/redispwned
	sudo chmod 755 /usr/local/bin/redispwned
	# install SystemD service
	sudo cp redispwned.service /lib/systemd/system/
	sudo chmod 755 /lib/systemd/system/redispwned.service
	sudo systemctl daemon-reload
	sudo systemctl enable redispwned.service
	sudo systemctl start redispwned
