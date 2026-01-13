# Onyx

**Onyx** is a SaaS which automates workflow of taking pictures for stock items or any other set of images which require labeling and exporting.

# Development

Create `secrets/` in this project directory with 2 files named `minio_passwd` and `pg_passwd`. Set random passwords to them.

```sh
docker compose -f docker-compose.dev.yml watch
```
