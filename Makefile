tailwind-watch:
	tailwindcss -i input.css -o ./public/styles/utilities.css --watch

tailwind-build:
	tailwindcss -i input.css -o ./public/styles/utilities.css --minify

podman-containers:
	./scripts/start-database-podman.sh && ./scripts/start-redis-podman.sh
