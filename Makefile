build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.booking . 

run:
	docker service create --name booking-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6025:6025 github.com/turistikrota/service.booking:latest

remove:
	docker service rm booking-api-turistikrota-com

stop:
	docker service scale booking-api-turistikrota-com=0

start:
	docker service scale booking-api-turistikrota-com=1

restart: remove build run
	