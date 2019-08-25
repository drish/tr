package main

import (
	"net/http"

	"github.com/drish/tr/k8s"
	"github.com/go-chi/render"

	"github.com/go-chi/chi"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/kubernetes"
)

func createJob(w http.ResponseWriter, r *http.Request) {

	render.Respond(w, r, "supppp")
}

type jobCreator struct {
	cs *kubernetes.Clientset
}

func newJobCreator() *jobCreator {
	// config file generated by gcloud get-credentials command
	config, err := clientcmd.BuildConfigFromFlags("", "/home/derich/.kube/config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return &jobCreator{
		cs: clientset,
	}
}

func (jc *jobCreator) createJob(w http.ResponseWriter, r *http.Request) {
	k8s.SpawnJob(jc.cs)
	render.Respond(w, r, "job created")
}

func main() {

	jb := newJobCreator()

	r := chi.NewRouter()
	r.Post("/job", jb.createJob)
	http.ListenAndServe(":3333", r)
}