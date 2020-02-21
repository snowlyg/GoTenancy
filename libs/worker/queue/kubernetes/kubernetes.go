package kubernetes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"GoTenancy/libs/worker"
	"github.com/ghodss/yaml"
	"k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var _ worker.Queue = &Kubernetes{}

// Kubernetes implemented a worker Queue based on kubernetes jobs
type Kubernetes struct {
	Clientset *kubernetes.Clientset
	Config    *Config
}

// Config kubernetes config
type Config struct {
	Namespace        string
	JobTemplateMaker func(worker.QorJobInterface) string
	ClusterConfig    *rest.Config
}

// New initialize Kubernetes
func New(config *Config) (*Kubernetes, error) {
	var err error

	if config == nil {
		config = &Config{}
	}

	if config.ClusterConfig == nil {
		config.ClusterConfig, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config.ClusterConfig)
	if err != nil {
		return nil, err
	}

	return &Kubernetes{Clientset: clientset, Config: config}, nil
}

// GetCurrentPod get current pod
func (k8s *Kubernetes) GetCurrentPod() *corev1.Pod {
	var (
		podlist, err = k8s.Clientset.Core().Pods("").List(metav1.ListOptions{})
		localeIP     = GetLocalIP()
	)

	if err == nil {
		for _, item := range podlist.Items {
			if item.Status.PodIP == localeIP {
				return &item
			}
		}
	}
	return nil
}

// GetJobSpec get job spec
func (k8s *Kubernetes) GetJobSpec(qorJob worker.QorJobInterface) (*v1.Job, error) {
	var (
		k8sJob     = &v1.Job{}
		currentPod = k8s.GetCurrentPod()
		namespace  = currentPod.GetNamespace()
	)

	if k8s.Config.Namespace != "" {
		namespace = k8s.Config.Namespace
	}

	if k8s.Config.JobTemplateMaker != nil {
		if err := yaml.Unmarshal([]byte(k8s.Config.JobTemplateMaker(qorJob)), k8sJob); err != nil {
			return nil, err
		}
		if k8sJob.ObjectMeta.Namespace != "" {
			namespace = k8sJob.ObjectMeta.Namespace
		}
	} else {
		if marshaledContainers, err := json.Marshal(currentPod.Spec.Containers); err == nil {
			json.Unmarshal(marshaledContainers, k8sJob.Spec.Template.Spec.Containers)
		}

		if marshaledVolumes, err := json.Marshal(currentPod.Spec.Volumes); err == nil {
			json.Unmarshal(marshaledVolumes, k8sJob.Spec.Template.Spec.Volumes)
		}
	}

	if k8sJob.TypeMeta.Kind == "" {
		k8sJob.TypeMeta.Kind = "Job"
	}

	if k8sJob.TypeMeta.APIVersion == "" {
		k8sJob.TypeMeta.APIVersion = "batch/v1"
	}

	if k8sJob.ObjectMeta.Namespace == "" {
		k8sJob.ObjectMeta.Namespace = namespace
	}

	return k8sJob, nil
}

// Add a job to k8s queue
func (k8s *Kubernetes) Add(qorJob worker.QorJobInterface) error {
	var (
		jobName        = fmt.Sprintf("qor-job-%v", qorJob.GetJobID())
		k8sJob, err    = k8s.GetJobSpec(qorJob)
		currentPath, _ = os.Getwd()
		binaryFile, _  = filepath.Abs(os.Args[0])
	)

	if err == nil {
		k8sJob.ObjectMeta.Name = jobName
		k8sJob.Spec.Template.ObjectMeta.Name = jobName

		if k8sJob.Spec.Template.Spec.RestartPolicy == "" {
			k8sJob.Spec.Template.Spec.RestartPolicy = "Never"
		}

		for idx, container := range k8sJob.Spec.Template.Spec.Containers {
			if len(container.Command) == 0 || k8s.Config.JobTemplateMaker == nil {
				container.Command = []string{binaryFile, "--qor-job", qorJob.GetJobID()}
			}
			if container.WorkingDir == "" || k8s.Config.JobTemplateMaker == nil {
				container.WorkingDir = currentPath
			}

			k8sJob.Spec.Template.Spec.Containers[idx] = container
		}

		_, err = k8s.Clientset.Batch().Jobs(k8sJob.ObjectMeta.GetNamespace()).Create(k8sJob)
	}
	return err
}

// Run a job from k8s queue
func (k8s *Kubernetes) Run(qorJob worker.QorJobInterface) error {
	job := qorJob.GetJob()

	if job.Handler != nil {
		return job.Handler(qorJob.GetSerializableArgument(qorJob), qorJob)
	}

	return errors.New("no handler found for job " + job.Name)
}

// Kill a job from k8s queue
func (k8s *Kubernetes) Kill(qorJob worker.QorJobInterface) error {
	var (
		k8sJob, err = k8s.GetJobSpec(qorJob)
		jobName     = fmt.Sprintf("qor-job-%v", qorJob.GetJobID())
	)

	if err == nil {
		return k8s.Clientset.Batch().Jobs(k8sJob.ObjectMeta.GetNamespace()).Delete(jobName, &metav1.DeleteOptions{})
	}
	return err
}

// Remove a job from k8s queue
func (k8s *Kubernetes) Remove(qorJob worker.QorJobInterface) error {
	var (
		k8sJob, err = k8s.GetJobSpec(qorJob)
		jobName     = fmt.Sprintf("qor-job-%v", qorJob.GetJobID())
	)

	if err == nil {
		// TODO Don't remove if it is already running
		return k8s.Clientset.Batch().Jobs(k8sJob.ObjectMeta.GetNamespace()).Delete(jobName, &metav1.DeleteOptions{})
	}
	return err
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
