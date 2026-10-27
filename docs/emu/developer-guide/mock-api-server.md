# Running a Mock eMASS API Server

If you need or want to run a mock eMASS API server to develop `emu` locally (i.e., without access to a real eMASS instance), use the example Dockerfile and Docker Compose file below.

## Dockerfile

```dockerfile
# Set the Dockerfile metadata.
FROM registry.cdso.army.mil/cdso/containers/approved-base/alpine:3.23
LABEL image.authors="Victor Fernandez III, @cyberphor"

# Set the working directory
WORKDIR /emass-api

# Install the Prism-CLI, get the eMASS REST OpenAPI standard, and create a non-root user.
RUN apk upgrade --no-cache &&\
    apk add --no-cache npm &&\
    npm install -g @stoplight/prism-cli &&\
    wget https://raw.githubusercontent.com/mitre/emass_client/main/docs/eMASSRestOpenApi.yaml &&\
    adduser -D emass

# Change to the non-root user.
USER emass

# Start the API.
EXPOSE 4010
CMD [ "prism", "mock", "-d", "-h", "0.0.0.0", "-p", "4010", "eMASSRestOpenApi.yaml" ]
```

## Docker Compose

```yaml
services:
  emass-api:
    build: .
    image: mock/emass-api:latest
    container_name: emass-api
    ports:
      - 4010:4010
```

## Containerizing `emu`

If you need or want to containerize using `emu`, set the environment variables required first. Then, use the provided [Docker Compose file](/docker/cli/compose.yml) and/or [Dockerfile](/docker/cli/Dockerfile) to containerize `emu`.

```bash
docker compose -f docker/compose.yml up --build
```

## References

- [MITRE's eMASS REST API Client Code Repository](https://github.com/mitre/emass_client)
- [MITRE's eMASS REST API Specification](https://github.com/mitre/emass_client/blob/main/docs/eMASSRestOpenApi.yaml)
