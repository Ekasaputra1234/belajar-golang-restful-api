deploy:
	docker-compose down
	docker image rm chick-service-project-template-task_web
	docker-compose up -d