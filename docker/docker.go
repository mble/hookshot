package docker

import (
	"fmt"
	"time"

	docker "github.com/fsouza/go-dockerclient"
)

// DockerSocket defines the Unix Docker socket location
const DockerSocket = "unix:///var/run/docker.sock"

// ListImages returns a string containing image info on the current Docker host
func ListImages() string {
	var info string

	client, err := docker.NewClient(DockerSocket)
	if err != nil {
		panic(err)
	}
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		info += fmt.Sprintf("ID: %s\n", img.ID) +
			fmt.Sprintf("RepoTags: %s\n", img.RepoTags) +
			fmt.Sprintf("Created: %s\n", time.Unix(img.Created, 0)) +
			fmt.Sprintf("Labels: %s\n", img.Labels) +
			"\n"
	}
	return info
}

// DeployImage deploys the "hello-world" docker image
func DeployImage(imageName string) (containerName string, err error) {
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
	container, dockerErr := client.CreateContainer(create)
	dockerErr = client.StartContainer(container.ID, &hostConfig)

	return container.Name, dockerErr
}
