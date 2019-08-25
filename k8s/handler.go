package k8s

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

func SpawnJob(clientset *kubernetes.Clientset) {

	batch := clientset.BatchV1()
	jobs := batch.Jobs("default")
	fmt.Println(jobs.List(metav1.ListOptions{}))

	job := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-job1",
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "test-job-container1",
							Image: "hello-world",
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	j, err := jobs.Create(&job)
	if err != nil {
		klog.Fatal(err)
	}
	fmt.Println(j)
}
