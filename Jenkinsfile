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
                            kubectl create secret generic myregistrykey --from-file=.dockerconfigjson=${dockerconfigpath}/.docker/config.json --type=kubernetes.io/dockerconfigjson -n go
                            sudo su
                            kubectl apply -f deployment -n go
                            kubectl apply -f load.yaml -n go
                        """
                    }
                }
            }
        }
    }
}