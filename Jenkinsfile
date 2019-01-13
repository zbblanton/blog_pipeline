pipeline {
    agent any
    triggers {
        pollSCM('*/1 * * * *')
    }
    environment {
        OPENFAAS_PASS = credentials('openfaasPass')
    }
    stages {
        stage('Clone Blog Repo') {
            steps {
                dir('blog') {
                    git 'https://github.com/zbblanton/blog'
                }                
            }
        }
        stage('Copy Static Files') {
            steps {
                dir('function') {
                    sh 'cp -R ../blog/public .'
                }
            }
        }
        stage('Build New Function Image') {
            steps {
                sh 'docker build -t localhost:5000/bt-blog .'
            }
        }
        stage('Push New Image') {
            steps {
                sh 'docker push localhost:5000/bt-blog:latest'
            }
        }
        stage('Deploy New Function') {
            steps {
                script {
                    reqJSON = '''
                    {
                        "service": "blog",
                        "image": "localhost:5000/bt-blog",
                        "labels": {
                            "com.openfaas.scale.zero": "true"
                        }
                    }
                    '''
                    resp = sh(returnStdout: true, script: "curl -o /dev/null -s -w '%{http_code}' -X PUT http://192.168.0.10:31112/system/functions -H 'authorization: Basic " + env.OPENFAAS_PASS + "' -d '" + reqJSON + "'").trim()
                    if (resp != "200" && resp != "201" && resp != "202") {
                        error("Updating function failed. Status code: " + resp)
                    }
                    echo resp
                }
            }
        }
        stage('Clean Workspacee') {
            steps {
                cleanWs()
            }
        }
    }
}
