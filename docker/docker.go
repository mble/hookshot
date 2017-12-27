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
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, img := range imgs {
		info += fmt.Sprintf("ID: %s\n", img.ID) +
			fmt.Sprintf("RepoTags: %s\n", img.RepoTags) +
			fmt.Sprintf("Created: %s\n", time.Unix(img.Created, 0)) +
			fmt.Sprintf("Labels: %s\n", img.Labels) +
			"\n"
	}
	return info
}
