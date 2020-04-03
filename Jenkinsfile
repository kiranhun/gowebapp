pipeline {
    agnet any
    stages {
        stage('Git Clone') {
            steps{
                checkout scm
            }
	    }
        stage('Build package'){

        }
        stage('Build docker image'){
            steps{
                commit_id = sh(returnStdout: true, script: 'git rev-parse HEAD')
                commit_id = sh(returnStdout: true, script: """echo $commit_id . """).trim()
        
                sh """
                    env && docker build -t ${GO_IMAGE_NAME}:${commit_id} .
                """
            }
        }

        stage('push docker image'){
            steps{
                sh """
                docker login -u ${DOCKER_USER} -p ${DOCKER_PASS}
                docker push ${BACKEND_IMAGE_NAME}:${commit_id}
                """
            }
        }

        stage('update deployment file with latest image name'){
            sh "q write -i deployment.yaml 'spec.template.spec.containers[0].image' ${DOCKER_USER}/webapp-go:${commit_id}"
        }

        stage('Deploy go app'){
            steps{
                script {
                    sh "pwd"
                    withKubeConfig([credentialsId: 'kubernetesCreds', serverUrl: '${url}']){
                        sh"""
                            env
                            kubectl create namespace go > /dev/null 2>&1 || true
                            sudo su
                            kubectl apply -f deployment -n go
                        """
                    }
                }
            }
        }
    }
}