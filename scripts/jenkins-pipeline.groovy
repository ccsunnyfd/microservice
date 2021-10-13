#!groovy
pipeline {
	agent {
	    node {
	        label '193.168.1.211'
	    }
	}
	environment {
		REPOSITORY="ssh://git@172.16.20.198:32203/ccsunnyfd/microservice.git"
		VERSION=sh(returnStdout: true, script: 'git describe --tags --always')
		DEST_PATH="./edgeUser/bin/"
		MODULE_PATH="./app/edge/user/..."
		SCRIPT_PATH="/home/root/scripts"
		MODULE="user-edge"
	}

	stages {
		stage('获取代码') {
			steps {
				echo "start fetch code from git: ${REPOSITORY}"
				deleteDir()
				git branch: 'main', url: "${REPOSITORY}"
			}
		}

		stage('单元测试') {
			steps {
				echo "unit test"
				sh "/usr/local/go/bin/go test -v ./..."
			}
		}

		stage('编译') {
			steps {
				echo "start compile"
				sh "/usr/local/go/bin/go build -ldflags \"-X main.Version=${VERSION}\" -o ${DEST_PATH} ${MODULE_PATH}"
			}
		}

		stage('构建镜像') {
			steps {
				echo "start build image"
				sh "${SCRIPT_PATH}/build-images.sh ${MODULE}"
			}
		}

		stage('发布镜像') {
			steps {
				echo "start deploy"
				sh "${SCRIPT_PATH}/deploy.sh user-service ${MODULE}"
			}
		}
	}
}