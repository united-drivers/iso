build: ../../protobuf/3166-1/3166-1.proto ../../go/3166-1/3166-1.pb.go

../../go/3166-1/3166-1.pb.go:
	@mkdir -p ../../go/3166-1
	docker run -v "$(PWD)/../..:$(PWD)/../.." -w "$(PWD)/../../protobuf/3166-1" znly/protoc -I. --go_out=../../go/3166-1 3166-1.proto

../../protobuf/3166-1/3166-1.proto: .tmp/country-codes.json
	cat $< | go run ./builder.go > $@.tmp
	mv $@.tmp $@

.tmp/country-codes.json:
	mkdir -p .tmp
	cd .tmp && wget -N http://data.okfn.org/data/core/country-codes/r/country-codes.json

clean:
	rm -rf .tmp ../../protobuf/3166-1/3166-1.proto ../../go/3166-1/3166-1.pb.go
