#!/bin/sh
docker build --network=host -t email:1.1.1 .
docker save -o email.tar email:1.1.1