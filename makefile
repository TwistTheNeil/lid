prod:
	go generate router/frontend_fs_prod.go
	go build -tags prod
