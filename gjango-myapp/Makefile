BINARY_NAME=gjango-myapp

build:
	echo "Building Gjango App"
	go build -o tmp/${BINARY_NAME} .

run: build
	echo "Starting Gjango App"
	./tmp/${BINARY_NAME} &
	echo "Gjango App Started"

start: run

clean:
	echo "Cleaning"
	go clean
	rm tmp/${BINARY_NAME}
	echo "Cleaned"

test:
	echo "Testing"
	go test ./...
	echo "Done"

stop:
	echo "Stopping Gjango App"
	pkill -SIGTERM -f ./tmp/${BINARY_NAME}
	echo "Stopped Gjango App"

restart: stop start

