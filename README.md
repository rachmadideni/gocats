# GoCats - Category Service

A simple REST API service for managing categories built with Go.

## Running Locally

1. **Prerequisites**: Ensure you have Go installed (version 1.16 or later)

2. **Run the application**:
   ```bash
   go run main.go
   ```

3. **Alternative: Run with Air (live reload)**:
   
   Install Air:
   ```bash
   go install github.com/air-verse/air@latest
   ```
   
   Run with Air:
   ```bash
   air
   ```

4. **Server will start on port 6000**

## API Endpoints

- `GET /health` - Health check
- `GET /api/categories` - Get all categories
- `GET /api/categories/{id}` - Get category by ID

## Example Requests

### Using request.http file (VS Code REST Client)

The repository includes a [request.http](request.http) file that you can use with the [REST Client extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) in VS Code. Simply open the file and click "Send Request" above each request.

### Using curl

```bash
# Health check
curl http://localhost:6000/health

# Get all categories
curl http://localhost:6000/api/categories

# Get category by ID
curl http://localhost:6000/api/categories/1

# Create a new category
curl -X POST http://localhost:6000/api/categories \
  -H "Content-Type: application/json" \
  -d '{"id": 2, "name": "Fashion", "description": "Clothing, shoes, and accessories"}'

# Update a category
curl -X PUT http://localhost:6000/api/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"id": 1, "name": "Electronics Updated", "description": "Updated description for electronics"}'

# Delete a category
curl -X DELETE http://localhost:6000/api/categories/1
```
