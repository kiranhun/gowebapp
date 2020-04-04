# gowebapp

A light weight webapp built using go lang and deployed on kubernetes

### Jenkins setup

1. Configure kube context in your jenkins instance
2. Create a jenkins pipeline with parameters
3. Add `url, GO_IMAGE_NAME, dockerconfigpath` as parameters to the jenkins job
4. You can access the application using the loadbalancer endpoint of nginx ingress controller
