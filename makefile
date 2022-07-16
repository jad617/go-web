help:				## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build:			  	## Builds code into a binary
	@go build .

run:				## Runs code localy
	@go run .

dev:				## Dev Mode: will restart app on each code changes
	@reflex -r '.go' -s -- go run main.go

tidy:				## Install && cleans project dependencies in go.mod
	@go mod tidy
