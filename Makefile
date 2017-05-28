test:
	go test

release: test docker-image
	docker push ruken/du-standup:latest

docker-image: test
	docker build -t ruken/du-standup .

clean:
	rm du-standup