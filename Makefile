local-build:
	go get .

local-run:
	go run .

docker-build:
	docker build -t bootcamp2 .

docker-run:
	docker run --name foobar -it --rm -p 30000:5000 bootcamp2

docker-stop:
	docker stop foobar

kube-deploy:
	kubectl apply -f foobar.yaml

kube-stop:
	kubectl delete -f foobar.yaml