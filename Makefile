GIT_HASH := $(shell git log -n 1 --pretty=format:'%h')

test:
	./scripts/giaoduc.sh api test
lint:
	./scripts/giaoduc.sh lint
build:
	./scripts/giaoduc.sh api build
server:
	./scripts/giaoduc.sh api start
infra-up:
	./scripts/giaoduc.sh infra up -d
infra-down:
	./scripts/giaoduc.sh infra down

run_tests:
	bash build.sh

docker_build: docker_login
	sudo docker build -f build/giaoducapi/Dockerfile_prod -t $(REPO) .

docker_upload: docker_build
	sudo docker tag $(REPO):latest $(REPO):$(TRAVIS_BRANCH).$(GIT_HASH)
	sudo docker push $(REPO):$(TRAVIS_BRANCH).$(GIT_HASH)
