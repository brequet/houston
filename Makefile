build:
	cd front/ && pnpm run build && cd .. && go build -o houston ./main.go 
