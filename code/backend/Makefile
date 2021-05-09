build:
	DOCKER_BUILDKIT=1 docker build -t ruhacks-backend .

pull:
	docker pull alphakilo07/ruhacks-backend

push:
	docker tag ruhacks-backend alphakilo07/ruhacks-backend
	docker push alphakilo07/ruhacks-backend

cloud:
	
	docker tag ruhacks-backend gcr.io/diabetech-515ed/ruhacks-backend
	docker push gcr.io/diabetech-515ed/ruhacks-backend

run:
	docker run  --rm -d -p 8081:8081 -e PORT='8081' \
		--name ruhacks-backend ruhacks-backend

kill:
	docker kill ruhacks-backend
	