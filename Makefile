compose_up:
	docker compose up -d

compose_up_vnc:
	ENABLE_VNC=1 docker compose up -d

compose_down:
	docker compose down

image_build:
	docker build -t yvv4docker/browser-camoufox .

image_push:
	docker push	yvv4docker/browser-camoufox:latest

image_pull:
	docker pull	yvv4docker/browser-camoufox:latest

image_remove:
	docker rmi yvv4docker/browser-camoufox:latest