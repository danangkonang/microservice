# Deployment

```bash
# start service as container
docker-compose up --build -d

# migrate user
./migrate -database postgres://postgres:postgres@localhost:5000/user?sslmode=disable -path ./user/db/migrations up
# migrate product
./migrate -database postgres://postgres:postgres@localhost:5001/product?sslmode=disable -path ./product/db/migrations up
```
