Custom Resource Definition (CRD) in Kubernetes

This repository contains the complete code and configuration files of the CRD example discussed 
in the blog.

How to get the docker image:
a. Pull it directly from docker hub "docker pull akash125/crdblog"
b. Build it using the make file
  1. Clone this repository in some_folder/src/blog.velotio.com/
  2. Set your GOPATH variable to path of some_folder
  3. Run make from the folder where make file is available, it will generate an image named crd/velotio-blog then tag this image with the name of your docker registry and push, or just change the prefix in Makefile to the name of your docker registry before running make.

Running the image on kubernetes:
You can run the image using the deployment file crd-deployment.yaml available in kubernetes folder of this repo. You need to update the image field in the deployment file if you are building the image youself.
