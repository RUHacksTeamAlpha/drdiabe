build:
	docker build -t ruhacks-frontend .

run:
	docker run  --rm -d -p 8080:8080 \
		--name ruhacks-frontend ruhacks-frontend	

kill:
	docker kill ruhacks-frontend

	
cloud:
	docker tag ruhacks-frontend gcr.io/team-dn-htn/ruhacks-frontend
	docker push gcr.io//ruhacks-frontend