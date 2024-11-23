.PHONY: build, deploy

build:
	@docker build -t=krr-app-image ./app

deploy:
	@kubectl apply -f ./deploy/krr-app.yaml