include .env

all: stop build start

restart: stop start

start:
	@docker-compose up -d
	(SERVER_PORT=$(VIEWFILE_SERVER_PORT) ./bin/viewfile/viewfile 2>> ./dev/logs/viewfile.log)&
	(SERVER_PORT=$(UPLOAD_SERVER_PORT) MONGODB_ADDR=$(MONGODB_ADDR) ./bin/upload/upload 2>> ./dev/logs/upload.log)&
	(SERVER_PORT=$(ASSET_SERVER_PORT) MONGODB_ADDR=$(MONGODB_ADDR) ./bin/asset/asset 2>> ./dev/logs/asset.log)&

build:
	@echo "rebuilding services"
	@cd bin/upload && go build && cd ../viewfile && go build && cd ../asset && go build

stop:
	@pkill viewfile && echo "viewfile stopped" || echo "viewfile not running"
	@pkill upload && echo "upload stopped" || echo "upload not running"
	@pkill asset && echo "asset stopped" || echo "asset not running"
