### 基于code-generator自动生成CRD客户端代码

#### 需求分析
```
由于疫情原因，学校需要对每个学生(Student)进行封闭管理，只允许学生从事：学习、娱乐、吃饭、睡觉四项活动.
请使用golang + K8S CRD来设计一套管理系统实现如下功能：

1 获取学生花名册（列举所有学生姓名、状态）
2 新生的默认状态为吃饭
3 修改某一学生当前的状态（吃完饭就睡觉）
```
#### 一些预先准备工作

第一步：克隆项目（留意目录结构）
```bash
tang@zhang ~/ % cd ${HOME}/go/src/
tang@zhang src % git clone https://github.com/englandbaron/StudentManagement

tang@zhang ~/ % cd ~/
tang@zhang ~/ % mkdir -p ${HOME}/go/src/k8s.io
tang@zhang ~/ % cd ${HOME}/go/src/k8s.io
tang@zhang k8s.io % git clone https://github.com/kubernetes/apimachinery
tang@zhang k8s.io % git clone https://github.com/kubernetes/code-generator
```

第二步：自动生成代码
```
tang@zhang StudentManagement % cd ${HOME}/go/src/StudentManagement
tang@zhang StudentManagement % ${HOME}/go/src/k8s.io/code-generator/generate-groups.sh \
    all \
    StudentManagement/pkg/client \
    StudentManagement/pkg/apis \
    "esdevops:v1"

tang@zhang StudentManagement % #这里需要再一次同步包
tang@zhang StudentManagement % go mod tidy
```

第三步：引用CRD资源文件
```
tang@tangdeMacBook-Pro ～/ % cd ${HOME}/go/src/StudentManagement
tang@zhang StudentManagement % kubectl create -f crd.yaml
```

#### 快速安装

参考Dockerfile文件

#### 案例

请参考example.go文件
环境变量配置请参考.vscode目录
