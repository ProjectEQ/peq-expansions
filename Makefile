build-dropshift:
	@cd scripts/dropshift && go build && mv dropshift ../../drop/
build-dropstrip:
	@cd scripts/dropstrip && go build && mv dropstrip ../../drop/
build-dropexport:
	@cd scripts/dropexport && go build && mv dropexport ../../drop/
