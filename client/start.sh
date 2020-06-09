#!/bin/sh
echo 'Checking environment...'
echo $ENV
if [ $ENV = production ]; then
    echo 'Spinning up production nginx server'
    nginx -g 'daemon off;'
else
    echo 'Spinning up development node server'
    ng serve --host 0.0.0.0
fi;