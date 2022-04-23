BINARY_NAME=gojango-myapp

build:
	echo "Building Gojango App"
	go build -o tmp/${BINARY_NAME} .

run: build
	echo "Starting Gojango App"
	./tmp/${BINARY_NAME} &
	echo "Gojango App Started"

start: run

clean:
	echo "Cleaning"
	go clean
	rm tmp/${BINARY_NAME}
	echo "Cleaned"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

test:
	echo "Testing"
	go test ./...
	echo "Done"

stop:
	echo "Stopping Gojango App"
	pkill -SIGTERM -f ./tmp/${BINARY_NAME}
	echo "Stopped Gojango App"

restart: stop start


