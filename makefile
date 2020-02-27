.PHONY: table compose

compose:
	docker-compose rm -v -f
	docker-compose up --build

table:
	rm -rf TABLE.md && node table/index.js

	