package main

import (
	v1 "StudentManagement/pkg/apis/esdevops/v1"
	clientset "StudentManagement/pkg/client/clientset/versioned"
	"context"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type patchUInt32Value struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value bool   `json:"value"`
}

func main() {
	// Step-1. 初始化Config配置
	cfg, err := rest.InClusterConfig()
	if err != nil {
		klog.Error("Get Config Error", err)
		return
	}
	// Step-2. 基于Config配置初始化client客户端
	// 这里的客户端代码即为我们使用code-generator自动
	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Error("Get Client Error", err)
		return
	}
	// Step-3. 获取学生列表 - default的namespace下
	_student_list, err := client.EsdevopsV1().Students("default").List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		klog.Error("List Student Error: ", err)
		return
	}
	for _, student := range _student_list.Items {
		klog.Infof(
			"Student Name: %s Student Status",
			student.Spec.Name,
			student.StudentStatus,
		)
	}
	// Step-4. 创建新生
	student := &v1.Student{}
	cr_name := "6076e124-5866-4c1d-af47-21bf022e95a6"
	student.Name = cr_name
	student.Spec.Name = "小明"
	student.StudentStatus.Eating = true
	create_resp, err := client.EsdevopsV1().Students("default").Create(
		context.TODO(),
		student, metav1.CreateOptions{},
	)
	if err != nil {
		klog.Error("Create Student Error: ", err)
		// return
	}
	klog.Infof("Student Create Result: %s - <Type.Student>", create_resp)
	// Step-5. 修改学生状态 - 学生吃完饭就睡觉啦
	payload, err := json.Marshal([]patchUInt32Value{
		{
			Op:    "replace",
			Path:  "/status/eating",
			Value: false,
		},
		{
			Op:    "replace",
			Path:  "/status/sleeping",
			Value: true,
		},
	})
	if err != nil {
		klog.Errorf("Json Render Error: %s", err)
		return
	}
	patch_resp, err := client.EsdevopsV1().Students("default").Patch(
		context.TODO(),
		cr_name,
		types.JSONPatchType,
		payload,
		metav1.PatchOptions{},
	)
	if err != nil {
		klog.Errorf("Patch Error: %s", err)
		return
	}
	klog.Infof("Patch Student: %s Succeed", cr_name, patch_resp)
}
