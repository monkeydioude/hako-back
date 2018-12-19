run:
	./bin/viewfile/viewfile&
	./bin/upload/upload&
	./bin/asset/asset&


build:
	cd bin/upload && go build && cd ../viewfile && go build && cd ../asset && go build

stop:
	pkill viewfile
	pkill upload
	pkill asset
