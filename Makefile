build:
		go mod vendor
		git add --all
		git diff-index --quiet HEAD || git commit -a -m 'fix'
		git push origin master

registry:
		docker build -t eu.gcr.io/hprofits/api-getway:latest .
		docker push eu.gcr.io/hprofits/api-getway:latest

deploy:
	go mod vendor
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	git add --all
	git diff-index --quiet HEAD || git commit -a -m 'fix'
	git push origin master
	docker build -t eu.gcr.io/hprofits/api-getway:latest .
	docker push eu.gcr.io/hprofits/api-getway:latest
	kubectl replace -f ./deployments/deployment.yml