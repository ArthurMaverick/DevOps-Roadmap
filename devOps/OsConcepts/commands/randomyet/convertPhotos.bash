#!bin/bash

PATH=$(pwd)
echo "entrando in $PATH"

cd $(pwd)

for image in *.jpg
do 
  echo "$image" \
  convert $image $image.png
done
