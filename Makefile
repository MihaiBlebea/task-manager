telegram_token = 1925485610:AAH77lkdTDVMup-IxFB3soCAxSc1NZGQ-AA
ngrok_tunnel = https://3494421cb1cd.ngrok.io

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
	docker-compose rm -f &&\
	rm -r ./volume &&\
	make setup

telegram-webhook:
	curl https://api.telegram.org/bot${telegram_token}/setWebhook?url=${ngrok_tunnel}/api/v1/webhook/telegram