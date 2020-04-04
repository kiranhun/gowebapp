pipeline {
    agent any
    stages {
        stage('Git Clone') {
            steps{
                checkout scm
            }
	    }
        stage('Build docker image'){
            steps{        
                sh """
                    env && docker build -t ${GO_IMAGE_NAME}:${GIT_COMMIT} .
                """
            }
        }

        stage('push docker image'){
            steps{
                sh """
                docker login -u ${DOCKER_USER} -p ${DOCKER_PASS}
                docker push ${GO_IMAGE_NAME}:${GIT_COMMIT}
                """
            }
        }

        stage('update deployment file with latest image name'){
            steps{
                script{
                    sh "yq write -i deployment.yaml 'spec.template.spec.containers[0].image' ${DOCKER_USER}/webapp-go:${GIT_COMMIT}"
                } 
            }
        }

        stage('Deploy go app'){
            steps{
                script {
                    sh "pwd"
                    withKubeConfig([credentialsId: 'kubernetesCreds', serverUrl: '${url}']){
                        sh"""
                            env
                            kubectl create namespace go > /dev/null 2>&1 || true
                            pwd
                            kubectl create secret generic myregistrykey --from-file=.dockerconfigjson=${dockerconfigpath}/.docker/config.json --type=kubernetes.io/dockerconfigjson -n go > /dev/null 2>&1 || true
                            kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
                            kubectl apply -f nginx-ingress.yaml
                            sudo su
                            kubectl apply -f deployment.yaml -n go
                            kubectl apply -f service.yaml
                            kubectl apply -f ingress.yaml -n go
                        """
                    }
                }
            }
        }
    }
}