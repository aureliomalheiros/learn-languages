package configuration

import (
    "context"
    v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//List all Namespaces
func GetAllNamespace(clientset *kubernetes.Clientset) ([]string, error){
    namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})

    if err != nil {
        return nil, err
    }
    var NsList []string
    for _, ns := range namespaces.Items {
        NsList = append(NsList, ns.ObjectMeta.Name)
    }

    return NsList, nil
}

//CONFIGMAPS
//List all ConfigMaps
func GetAllConfigMaps(clientset *kubernetes.Clientset, namespace string) ([]string, error){

	configmaps, err := clientset.CoreV1().ConfigMaps(namespace).List(context.TODO(), v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var ConfigMapsList []string
	for _, cm := range configmaps.Items {
		ConfigMapsList = append(ConfigMapsList, cm.ObjectMeta.Name)
	}

	return ConfigMapsList, nil
}
func GetDescribeConfigMaps(clientset *kubernetes.Clientset, namespace string,) ([]string, error){

    configmap, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	var ConfigMapsList []string
	for _, cm := range configmaps.Items {
		ConfigMapsList = append(ConfigMapsList, cm.ObjectMeta.Name)
	}

	return ConfigMapsList, nil
}


//SECRETS
//List all Secrets
func GetAllSecrets(clientset *kubernetes.Clientset, namespace string) ([]string, error){

	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var SecretsLis []string
	for _, s := range secrets.Items {
		SecretsLis = append(SecretsLis, s.ObjectMeta.Name)
	}

	return SecretsLis, nil
}

func GetDescribeSecrets(clientset *kubernetes.Clientset, namespace string) ([]string, error){

	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	var SecretsLis []string
	for _, s := range secrets.Items {
		SecretsLis = append(SecretsLis, s.ObjectMeta.Name)
	}

	return SecretsLis, nil
}