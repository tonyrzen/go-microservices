AIR := ${GOPATH}/air
broker:
	cd ./broker-service/cmd/api && ${AIR}

ui: 
	cd ./front-end/cmd/web && ${AIR}

auth_service:
	cd ./auth/cmd/api && ${AIR}