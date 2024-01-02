# WASAPhoto

This repository contains the [Web and Software Architecture](http://gamificationlab.uniroma1.it/en/wasa/) homework project. It uses the [Fantastic Coffee (Decaffeinated)](https://github.com/sapienzaapps/fantastic-coffee-decaffeinated) structure.

## How to build (in development mode)

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## How to run (in development mode)

You can launch the backend only using:

```shell
go run ./cmd/webapi/
```

If you want to launch the WebUI, open a new tab and launch:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

## How to build container images

### Backend

```sh
$ docker build -t wasaphoto-backend:latest -f Dockerfile.backend .
```

### Frontend

```sh
$ docker build -t wasaphoto-frontend:latest -f Dockerfile.frontend .
```

## How to run container images

### Backend

```sh
$ docker run -it --rm -p 3000:3000 wasaphoto-backend:latest
```

### Frontend

```
$ docker run -it --rm -p 8080:80 wasaphoto-frontend:latest
```

## License

See [LICENSE](LICENSE).
