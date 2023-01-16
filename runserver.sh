fileServerPath=$(pwd)/fileServer/main.go

imageName=reactblogserv_article
containerName=reactblogserv_articleInstance
port=6123

buildImage() {
    rootDir=$1
    imageName=$2
    containerName=$3

    cd $rootDir
    echo "build image..."
    docker build -t $imageName .
}

deleteImage() {
    imageName=$2
    containerName=$3
    echo "stop running container and delete image.."
    docker stop $containerName &&
        docker rm $containerName &&
        docker rmi $imageName

    echo "delete dangling images"
    docker rmi "$(docker images -f "dangling=true" -q)"
}

deleteImage fileServer fileServerInstance
buildImage $(pwd)/fileServer fileServer fileServerInstance 7123
docker run -p 7123:7123 -d --name fileServerInstance \
-v /soft/react_website/react_website_server/fileServer/assets:/app/assets \
fileServer

deleteImage reactblogserv_article reactblogserv_articleInstance
buildImage $(pwd)/apps/article reactblogserv_article reactblogserv_articleInstance 6211
docker run -p 6211:6211 -d --name reactblogserv_articleInstance reactblogserv_article
