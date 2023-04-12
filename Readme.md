# Fitsphere Backend

This is a monorepo with different backend services for fitsphere.

Fitsphere is a social media for fitness related content.

## Development

docker-compose.dev.yml is used for development. It starts the supporting services required for the backend services to run.

```bash
docker compose -f docker-compose.dev.yml up
```

Please run the following command to start the services:

```bash
scipts/dev.sh

```
currently the env variables will be the defaults for dev enviroinment. This will be changed in the future.