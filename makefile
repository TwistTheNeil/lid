prod:
	rm -rf router/dist
	cd frontend && yarn build
	go generate router/frontend_fs_prod.go
	go build -tags prod
