package main


import (
    "fmt"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "api-k8s/k8s"
)

func main(){
    //Variables
    var userOption int
    var SelectNamespace string

    kubeconfig := "/home/aurelio/.kube/config"
    
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(err.Error())
    }
    
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }
	
    fmt.Println("Digite o valor númerico para selecionar a opção: ")
    fmt.Println("1 -> Listar todos os namespaces")
    fmt.Println("2 -> Listar todos os ConfigMaps")
    fmt.Println("3 -> Listar todos os secrets")
    fmt.Printf("> ")
    fmt.Scanf("%d", &userOption)
    
    if (userOption == 2 || userOption == 3){
        fmt.Printf("Selecione o Namespace: ")
        fmt.Scanf("%s", &SelectNamespace)
    }
    //List all Namespaces
    if userOption == 1 {
        namespaces, err := configuration.GetAllNamespace(clientset)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(namespaces)
    } else if userOption == 2 {
        configmap, err := configuration.GetAllConfigMaps(clientset, SelectNamespace)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(configmap)
    } else if userOption == 3{
        secret, err := configuration.GetAllSecrets(clientset, SelectNamespace)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(secret)
    }
}