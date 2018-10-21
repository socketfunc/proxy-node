docker-build:
	docker build --no-cache=true --force-rm=true --tag=yuichiko/proxy-node-arm:${TAG} .

docker-push:
	docker push yuichiko/proxy-node-arm:${TAG}

update-proto:
	git subrepo pull proto
