compose-up:
	docker compose up -d

compose-down:
	docker compose down

image-build:
	docker build -t yvv4docker/browser-camoufox .

image-push:
	docker push	yvv4docker/browser-camoufox:latest

image-pull:
	docker pull	yvv4docker/browser-camoufox:latest

image-remove:
	docker rmi yvv4docker/browser-camoufox:latest