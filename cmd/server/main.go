package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// CONNECT K8S
	// Use cli
	kubeconfig := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "jphuc96", "dev8s", "secret", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	// Use rest api
	// config := &rest.Config{}
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// CONNECT MONGO
	clientOptions := options.Client().ApplyURI("mongodb://root:dev8s@localhost:27017")
	MongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = MongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := MongoClient.Database("dev8s_test").Collection("numbers")
	res, err := collection.InsertOne(context.Background(), bson.M{"name": "pi", "value": 3.14})
	pp.Println("Inserted", res.InsertedID)

	pods, err := k8sClient.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range pods.Items {
		pp.Println(v.Name)
	}
	// podsJSON, err := json.MarshalIndent(pods, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ioutil.WriteFile("pods.json", podsJSON, 0644)
}
