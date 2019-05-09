package main

import (
	"flag"
	"github.com/kelseyhightower/envconfig"
	"github.com/n3wscott/knap/pkg/config"
	"github.com/n3wscott/knap/pkg/knative"

	//"github.com/n3wscott/knap/pkg/knative"
	"github.com/n3wscott/kpax/poke/pkg/controller"
	"k8s.io/client-go/dynamic"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
)

var (
	cluster    string
	kubeconfig string
)

type envConfig struct {
	// Name of this pod.
	Port string `envconfig:"PORT" default:":8080"`

	// Name of this pod.
	Name string `envconfig:"POD_NAME" required:"true"`

	// Namespace this pod exists in.
	Namespace string `envconfig:"POD_NAMESPACE" required:"true"`
}

func init() {
	flag.StringVar(&cluster, "cluster", "",
		"Provide the cluster to test against. Defaults to the current cluster in kubeconfig.")

	var defaultKubeconfig string
	if usr, err := user.Current(); err == nil {
		defaultKubeconfig = path.Join(usr.HomeDir, ".kube/config")
	}

	flag.StringVar(&kubeconfig, "kubeconfig", defaultKubeconfig,
		"Provide the path to the `kubeconfig` file.")
}

func main() {

	flag.Parse()
	var client dynamic.Interface
	var env envConfig

	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	cfg, err := config.BuildClientConfig(kubeconfig, cluster)
	if err != nil {
		log.Fatalf("Error building kubeconfig", err)
	}

	client = dynamic.NewForConfigOrDie(cfg)

	//kn := knative.New(client)
	var kn *knative.Client

	root := "/Users/nicholss/go/src/github.com/n3wscott/kpax/poke/cmd/poke/kodata"

	c := controller.New(root, kn, "default")

	c.Router().PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(root+"/static"))))

	http.Handle("/", c.Router())
	log.Fatal(http.ListenAndServe(env.Port, nil))
}
