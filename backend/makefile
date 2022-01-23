
build:
	nohup ./bin/qpoker > tests/qpoker.out 2>&1 &

run:
	./bin/qpoker

test:
	go build -o bin/qpoker
	nohup ./bin/qpoker > tests/sockettest.out 2>&1 &
	firefox tests/sockettest.html&

kill:
	kill $$(ps -ax | grep qpoker | awk '{print $$1}' | head -n 1)
