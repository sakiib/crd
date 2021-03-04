package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/klog/v2"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/davecgh/go-spew/spew"
	"github.com/sakiib/crd/pkg/apis/book.com/v1alpha1"

	sakibv1alpha1 "github.com/sakiib/crd/pkg/client/clientset/versioned"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	indexer   cache.Indexer
	queue     workqueue.RateLimitingInterface
	informer  cache.Controller
	crdClient sakibv1alpha1.Interface
	kClient   kubernetes.Interface
}

func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller, crdClient sakibv1alpha1.Interface, kClient kubernetes.Interface) *Controller {
	return &Controller{
		indexer:   indexer,
		queue:     queue,
		informer:  informer,
		crdClient: crdClient,
		kClient:   kClient,
	}
}

func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()

	defer c.queue.ShutDown()
	fmt.Println("Starting BookAPI Controller")

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waiting for caches to sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	fmt.Println("Stopping BookAPI Controller")

}

func (c *Controller) runWorker() {
	for c.processNextItem() {

	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	defer c.queue.Done(key)

	err := c.reconcileFunc(key.(string))
	c.handleErr(err, key)

	return true
}

func (c *Controller) reconcileFunc(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		fmt.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		fmt.Printf("BookAPI %s does not exist anymore\n", key)
	} else {
		fmt.Printf("Sync/Add/Update for BookAPI %s\n", obj.(*v1alpha1.BookAPI).GetName())
		bookapiObj := obj.(*v1alpha1.BookAPI).DeepCopy()
		c.process(bookapiObj)
	}

	return nil
}

func (c *Controller) process(bookapiObj *v1alpha1.BookAPI) {
	deploymentClient := c.kClient.AppsV1().Deployments(apiv1.NamespaceDefault)
	serviceClient := c.kClient.CoreV1().Services(apiv1.NamespaceDefault)

	deploymentName := bookapiObj.ObjectMeta.Name

	tpmnt, err := deploymentClient.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	errorMessage := "deployments.apps" + " " + "\"" + deploymentName + "\"" + " not found"

	if err != nil {
		if err.Error() == errorMessage {
			spew.Dump(tpmnt)
			deployment := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name: deploymentName,
				},
				Spec: appsv1.DeploymentSpec{
					Replicas: int32Ptr(int32(*bookapiObj.Spec.Replica)),
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{
							"app": "demo",
						},
					},
					Template: apiv1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{
								"app": "demo",
							},
						},
						Spec: apiv1.PodSpec{
							Containers: []apiv1.Container{
								{
									Name:  deploymentName,
									Image: bookapiObj.Spec.Image,
									Ports: []apiv1.ContainerPort{
										{
											Name:          deploymentName,
											Protocol:      apiv1.ProtocolTCP,
											ContainerPort: int32(bookapiObj.Spec.Port),
										},
									},
								},
							},
						},
					},
				},
			}

			fmt.Println("Creating BookAPI deployment...")
			result, err := deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})

			if err != nil {
				panic(err)
			}
			fmt.Printf("Created BookAPI deployment %q.\n", result.GetObjectMeta().GetName())

			service := &apiv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: bookapiObj.ObjectMeta.Name,
					OwnerReferences: []metav1.OwnerReference{
						*metav1.NewControllerRef(bookapiObj, v1alpha1.SchemeGroupVersion.WithKind("BookAPI")),
					},
				},
				Spec: apiv1.ServiceSpec{
					Ports: []apiv1.ServicePort{
						{
							Protocol: apiv1.ProtocolTCP,
							Port:     int32(bookapiObj.Spec.Port),
							TargetPort: intstr.IntOrString{
								IntVal: int32(bookapiObj.Spec.Port),
							},
							NodePort: bookapiObj.Spec.NodePort,
						},
					},
					Type: apiv1.ServiceType(bookapiObj.Spec.ServiceType),
				},
			}
			res, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
			if err != nil {
				panic(err)
			}
			fmt.Printf("Created BookAPI service %q.\n", res.GetObjectMeta().GetName())
		} else {
			fmt.Printf("%v", err.Error())
		}
		return
	}
}

func (c *Controller) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	if c.queue.NumRequeues(key) < 5 {
		klog.Infof("Error syncing BookAPI %v: %v", key, err)
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	runtime.HandleError(err)
	klog.Infof("Dropping BookAPI %q out of the queue: %v", key, err)
}

func main() {
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := sakibv1alpha1.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	bookapiListWatcher := cache.NewListWatchFromClient(clientset.BookV1alpha1().RESTClient(), "bookapis", "default", fields.Everything())

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	indexer, informer := cache.NewIndexerInformer(bookapiListWatcher, &v1alpha1.BookAPI{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	kClient := kubernetes.NewForConfigOrDie(config)
	controller := NewController(queue, indexer, informer, clientset, kClient)

	indexer.Add(&v1alpha1.BookAPI{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bookapis",
			Namespace: "default",
		},
	})

	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(1, stop)

	select {}
}

func int32Ptr(i int32) *int32 {
	return &i
}
