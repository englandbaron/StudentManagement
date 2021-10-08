FROM centos:7.9.2009

RUN cd /tmp && \
  yum -y update && \
  yum -y install git wget && \
  yum -y groupinstall "development tools"

RUN wget https://dl.google.com/go/go1.16.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz && \
    echo "export GOPATH=/root/go" >> /etc/profile && \
    echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile && \
    source /etc/profile && \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct

# 为了加速Git项目的克隆,我们在这里使用了gitee源(每天自动同步1次)
# code-generator推荐使用稳定分支版本

RUN source /etc/profile && \
    mkdir -p ~/go/src && \
    cd ${HOME}/go/src/ && \
    git clone https://gitee.com/facengineer/StudentManagement && \
    mkdir -p ${HOME}/go/src/k8s.io && \
    cd ${HOME}/go/src/k8s.io && \
    git clone https://gitee.com/mirrors/apimachinery && \
    git clone https://gitee.com/facengineer/code-generator -b release-1.20 && \
    cd ${HOME}/go/src/StudentManagement && \
    go mod tidy && \
    ${HOME}/go/src/k8s.io/code-generator/generate-groups.sh all StudentManagement/pkg/client StudentManagement/pkg/apis "esdevops:v1"

CMD ["ls"]
