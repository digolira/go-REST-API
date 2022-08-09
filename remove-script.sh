
function deleting_containers() {
    sudo docker rm -f first-mysql
}

function deleting_volumes(){
    sudo docker volume rm -f go-rest_myvolume
}

deleting_containers
deleting_volumes