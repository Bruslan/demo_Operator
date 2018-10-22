package stub

import (
	"fmt"
	"github.com/example-inc/bruslan-operator/pkg/apis/brus/v1alpha1" // "github.com/operator-framework/operator-sdk/pkg/sdk/action"
	sdk "github.com/operator-framework/operator-sdk/pkg/sdk/action"
	"github.com/operator-framework/operator-sdk/pkg/sdk/handler"
	"github.com/operator-framework/operator-sdk/pkg/sdk/types" // "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"                                       // "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	// "github.com/operator-framework/operator-sdk/pkg/k8sclient"
)

func NewHandler() handler.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx types.Context, event types.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Bruslan:

		// err := action.Create(newbusyBoxPod(o))
		// if err != nil && !errors.IsAlreadyExists(err){
		// 	logrus.Error("Failed to create busybox pod: %v",err)
		// 	return err
		// }

		nameSpace := o.Spec.NameSpaceString

		if nameSpace == "all" {

			printPods("")

		} else {
			printPods(nameSpace)
		}

		fmt.Printf("the name of this cr is ", o.Name)
		if o.Spec.BruslanDEMO {
			o.Status.Status = "We have a true configuration"
			fmt.Println("unser Boolean ist auf True")
		} else {
			o.Status.Status = "WE have a False configuration"
			fmt.Println("unser Boolean ist auf false")
		}
		sdk.Update(o)
		// fmt.Printf("%+v\n\n", o)
	}
	return nil
}

func printPods(nameSpace string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods(nameSpace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())

	}
	fmt.Println("all Pods : ", pods)

	fmt.Println("NameSpace: ", nameSpace)
	for _, pod := range pods.Items {
		fmt.Println("Pod NAme: ", pod.Name)
	}

	// fmt.Println("Type: ", reflect.TypeOf(pods))
}

// func getPodNames(pods []corev1.Pod) []string {
// 	var podNames []string
// 	for _, pod := range pods {
// 		podNames = append(podNames, pod.Name)
// 	}
// 	return podNames
// }

// newbusyBoxPod demonstrates how to create a busybox pod
func newbusyBoxPod(cr *v1alpha1.Bruslan) *v1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	return &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box",
			Namespace: "neuesNamespace",
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Bruslan",
				}),
			},
			Labels: labels,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
