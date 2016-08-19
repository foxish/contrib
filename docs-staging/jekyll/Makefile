IMG = docker.io/foxish/jekyll-staging
TAG = $(shell date +v%Y%m%d)

all: push

build:
	docker build -t $(IMG):$(TAG) .
	docker tag -f $(IMG):$(TAG) $(IMG):latest
	@echo Built $(IMG):$(TAG) and tagged with latest

push: build
	docker push $(IMG):$(TAG)
	docker push $(IMG):latest
	@echo Pushed $(IMG) with :latest and :$(TAG) tags