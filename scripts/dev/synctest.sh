#!/bin/bash

function invoke {
    sleep 2
    let limit=$1
    let index=0
    while [  $index -lt $limit ]; do
        ./invoke.sh i
        let index=index+1
    done
}

./runyafabric.sh -f 1 -c p -r clearall

sleep 1
./invoke.sh d
invoke 1
invoke 1
invoke 1

sleep 10
./runyafabric.sh -k 3

invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5

sleep 5
./runyafabric.sh -f 1 -c p -i 3

sleep 5
#invoke 1



exit

invoke 5
invoke 5
invoke 5
invoke 5
invoke 5
invoke 5