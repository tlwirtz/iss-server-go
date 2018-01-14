GOOS=linux go build
docker build -t tlwirtz/iss-server-go:latest .
docker run -p8080:8080 tlwirtz/iss-server-go:latest