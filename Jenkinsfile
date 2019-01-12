pipeline {
    agent any
<<<<<<< HEAD
=======
    triggers {
        pollSCM('*/1 * * * *')
    }
    environment {
        OPENFAAS_PASS = credentials('openfaasPass')
    }
>>>>>>> c195a96... Added openfaas cred as Jenkins cred
    stages {
        stage('Clone Blog Repo') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Copy Static Files') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Build New Function Image') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Push New Image') {
            steps {
                echo 'Im sure it worked.'
            }
        }
        stage('Deploy New Function') {
            steps {
<<<<<<< HEAD
                echo 'Im sure it worked.'
=======
                sh '''
                    curl -X PUT \
                    http://192.168.0.10:31112/system/functions \
                    -H 'authorization: Basic $OPENFAAS_PASS' \
                    -H 'cache-control: no-cache' \
                    -H 'content-type: application/json' \
                    -H 'postman-token: bbd2fd48-307d-efbb-f4b3-9b73692cead7' \
                    -d '{
                    	    "service": "blog",
                    	    "image": "localhost:5000/bt-blog",
                    	    "labels": {
                    	        "com.openfaas.scale.zero": "true"
                    	    }
                	    }'
                '''
            }
        }
        stage('Clean Workspacee') {
            steps {
                cleanWs()
>>>>>>> c195a96... Added openfaas cred as Jenkins cred
            }
        }
    }
}
