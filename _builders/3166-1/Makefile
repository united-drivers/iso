build: ../../3166-1/3166-1.proto

../../3166-1/3166-1.proto: .tmp/country-codes.json
	cat $< | go run ./builder.go > $@.tmp
	mv $@.tmp $@

.tmp/country-codes.json:
	mkdir -p .tmp
	cd .tmp && wget -N http://data.okfn.org/data/core/country-codes/r/country-codes.json

clean:
	rm -rf .tmp
