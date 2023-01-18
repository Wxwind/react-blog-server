echo "delete running images"
docker compose down
echo "build images and deploy"
docker compose build&&
docker compose up -d

