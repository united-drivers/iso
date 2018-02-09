build: ../../4217/4217.proto

../../4217/4217.proto:
	@mkdir -p .tmp
	cd .tmp && wget -N http://www.currency-iso.org/dam/downloads/lists/list_one.xml
	cat .tmp/list_one.xml | xml2json > .tmp/4217.json
	cat .tmp/4217.json | go run builder.go > $@.tmp
	mv $@.tmp $@

clean:
	rm -rf .tmp ../../4217/4217.proto
