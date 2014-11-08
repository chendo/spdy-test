build:
	docker build -t spdy-test .
run:
	docker run -p 44300:44300 -p 44301:44301 spdy-test
