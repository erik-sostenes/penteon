# User API (Go)

## Endpoints

### Create User

```
curl -X POST "http://localhost:8000/api/v1/users/create" \
  -H "Content-Type: application/json" \
  -d '{"name": "Juan", "age": 30}'
```

```json
{
  "name": "Juan",
  "age": 30
}
```

### Get users 

```
curl 'http://localhost:8000/api/v1/users/get-all'
```