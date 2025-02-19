dev:
	air
up:
	goose up
down:
	goose down
schema:
	jet --source=sqlite -dsn=./beholder.db -schema=dvds -path=./internal/jet
proto:
	protoc --go_out=./ schema/schema.proto
	protoc --es_out=./dashboard/src/ schema/schema.proto