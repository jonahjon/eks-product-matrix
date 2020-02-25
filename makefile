SHELL := /bin/bash
.PHONY: compose

compose:
	docker-compose rm -v -f
	docker-compose up
