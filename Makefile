air := ${GOPATH}/air
broker:
	cd ./broker-service/cmd/api && ${air}

ui: 
	cd ./front-end/cmd/web && ${air}