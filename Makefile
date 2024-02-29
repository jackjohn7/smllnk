tailwind-watch:
	tailwindcss -i input.css -o ./public/styles/utilities.css --watch

tailwind-build:
	tailwindcss -i input.css -o ./public/styles/utilities.css --minify
