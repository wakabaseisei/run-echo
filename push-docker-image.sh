docker build . --target prod -t wakabaseisei/goapi -f docker/go/Dockerfile
docker push wakabaseisei/goapi:latest