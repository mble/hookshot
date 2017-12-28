package docker

import (
	"net/http"

	docker "github.com/fsouza/go-dockerclient"
)

// DockerSocket defines the Unix Docker socket location
const DockerSocket = "unix:///var/run/docker.sock"

// Image embeds docker.APIImages to make rendering easier
type Image struct {
	docker.APIImages
}

// Render provides render.Renderer compatibility
func (i Image) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// Container represents a deployed container
type Container struct {
	ID     string       `json:"Id"`
	Name   string       `json:"Name"`
	Args   []string     `json:"Args"`
	Image  string       `json:"Image"`
	State  docker.State `json:"State"`
	Status string       `json:"Status"`
}

// Render provides render.Renderer compatibility
func (c Container) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// ListImages returns []APIImages containing image info on the current Docker host
func ListImages() []Image {
	var imgs []Image
	client, err := docker.NewClient(DockerSocket)
	if err != nil {
		panic(err)
	}
	rawImgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range rawImgs {
		imgs = append(imgs, Image{img})
	}
	return imgs
}

// DeployImage deploys the "hello-world" docker image
func DeployImage(imageName string) (container Container, err error) {
	var dockerErr error
	image := docker.PullImageOptions{Repository: imageName, Tag: "latest"}
	auth := docker.AuthConfiguration{}

	client, err := docker.NewClient(DockerSocket)
	if err != nil {
		panic(err)
	}

	dockerErr = client.PullImage(image, auth)
	dockerErr = client.RemoveContainer(docker.RemoveContainerOptions{ID: imageName, Force: true})

	config := docker.Config{Image: imageName}
	hostConfig := docker.HostConfig{PublishAllPorts: true}
	create := docker.CreateContainerOptions{Name: imageName, Config: &config}
	rawContainer, dockerErr := client.CreateContainer(create)
	dockerErr = client.StartContainer(rawContainer.ID, &hostConfig)
	rawContainer, dockerErr = client.InspectContainer(rawContainer.ID)
	cntr := Container{ID: rawContainer.ID, Name: rawContainer.Name, Args: rawContainer.Args, Image: rawContainer.Image, State: rawContainer.State}
	if dockerErr != nil {
		cntr.Status = "error"
	} else {
		cntr.Status = "ok"
	}

	return cntr, dockerErr
}
