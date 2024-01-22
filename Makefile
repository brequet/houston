build:
	cd frontend/ && pnpm run build && cd .. && go build -o ./dist/houston.exe .
