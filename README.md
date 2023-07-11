# wiki

Implementation of a simple wiki http server by golang

## Usage

```bash
# Build image
docker build -t wiki-app:latest .

# Use image without docker-compose.
docker run -itd -p 8080:8080 wiki-app:latest

# Use image by docker-compose.
docker-compose up -d

# For down docker-compose.
docker-compose down -v
```