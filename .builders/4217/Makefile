build: ../../protobuf/4217/4217.proto ../../go/4217/4217.pb.go

../../go/4217/4217.pb.go:
	@mkdir -p ../../go/4217
	docker run -v "$(PWD)/../..:$(PWD)/../.." -w "$(PWD)/../../protobuf/4217" znly/protoc -I. --go_out=../../go/4217 4217.proto

../../protobuf/4217/4217.proto:
	@mkdir -p .tmp
	cd .tmp && wget -N http://www.currency-iso.org/dam/downloads/lists/list_one.xml
	cat .tmp/list_one.xml | xml2json > .tmp/4217.json
	cat .tmp/4217.json | go run builder.go > $@.tmp
	mv $@.tmp $@

clean:
	rm -rf .tmp ../../protobuf/4217/4217.proto ../../go/4217/4217.pb.go
