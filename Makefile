build:
	cd frontend/ && pnpm run build && cd .. && go build -o ./dist/houston.exe .

run:
	go build -o ./dist/houston.exe . && ./dist/houston.exe