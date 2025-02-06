[![CI](https://github.com/alex-arraga/apple_store/actions/workflows/ci.yml/badge.svg)](https://github.com/alex-arraga/apple_store/actions/workflows/ci.yml)

docker build -t alexarraga/apple_backend:v1.0.0 .
docker push alexarraga/apple_backend:v1.0.0
docker-compose down -v
docker-compose pull
docker-compose up -d