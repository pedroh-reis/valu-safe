# Postgres

## Dev

### Database

#### Up database instance
```
docker compose up -d
```

#### Down and clean up database instance
```
docker compose down
docker rm -f $(docker ps -a -q)
docker volume rm $(docker volume ls -q)
```

### Access the database
Access the database container
```
docker exec -it pg-local-container sh
```

From the container, access the database
```
psql -U pg-local-user -d pg-local-db
```
