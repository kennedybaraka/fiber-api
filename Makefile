run:
	air

init:
	go install
	go mod tidy
	air init