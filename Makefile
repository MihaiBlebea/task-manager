setup: build up

build:
	docker-compose build

up:
	docker-compose up -d

down: 
	docker-compose stop &&\
	docker-compose rm -f
	
refresh:
	docker-compose stop &&\
	docker-compose rm &&\
	rm -r ./volume &&\
	make local