PACKAGE_NAME := github.com/TypicalAM/mines

.PHONY: release-dry-run
release-dry-run:
	@docker build -t mines .
	@docker run \
		--rm \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		mines \
	  --clean --snapshot

.PHONY: release
release:
	@if [ ! -f ".env-release" ]; then\
		echo ".env-release is required for release";\
		exit 1;\
	fi
	@docker build -t mines .
	@docker run \
		--rm \
		-e CGO_ENABLED=1 \
		--env-file .env-release \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		mines \
	  release --clean
