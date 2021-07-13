build:
	@echo " > Building [marvelapi]..."
	@cd ./cmd/marvelapi/ && go build -o ../../bin && cd ../..	
	@echo " > Finished building [marvelapi]"

# Test
test:
	@echo " > Testing starts..."
	@go test -race -cover ./...
	@echo " > Finished testing"

# RUN
run: build
	@echo " > Running [marvelapi]..."
	@ ./bin/marvelapi
	@echo " > Finished running [marvelapi]"
