package test_pod

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"gopkg.in/yaml.v2"
	"new_k8s/pkg/container"
	"new_k8s/pkg/test_pod/random"
	"new_k8s/tools/etcd/etcd"
	"time"
)

func (pod *Pod) Start(ctx context.Context, cli *etcd.MyEtcdClient) {

	// 第2步：创建pauseContainer

	pauseID, pauseName, err := createPause(pod, *cli)
	pod.PauseContainer.Name = pauseName
	if pod.PauseContainer == nil {
		fmt.Printf("this is pod start func and pause container is nil\n")
	}
	if err != nil {
		fmt.Printf("create pause container err:%v", err)
	}
	fmt.Printf("pause container create succeed! pause containerID: ", pauseID, "containerName: ", pauseName)
	// Create and Start Common Container
	dockercli, _ := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	defer dockercli.Close()
	containerJSON, err := dockercli.ContainerInspect(context.Background(), pauseID)
	if err != nil {
		panic(err)
	} else {
		// NOTE: 此处有修改！
		//containerIP := containerJSON.N
		containerIP := containerJSON.NetworkSettings.IPAddress
		fmt.Println("PAUSE containerIP: %v\n", containerIP)
		pod.PodIP = containerIP
	}
	pauseMode := "container:" + pauseID
	allContainersRunning := true

	for i := range pod.Configs.Spec.Containers {
		for _, volume := range pod.Configs.Spec.Volumes {
			volumeMount := pod.Configs.Spec.Containers[i].VolumeMount
			if volumeMount.Name == volume.Name {
				pod.Configs.Spec.Containers[i].VolumeMount.MountPath = volume.HostPath.Path
			}
		}

		err := container.EnsureImage(pod.Configs.Spec.Containers[i].Image)
		if err != nil {
			fmt.Printf("Failed to pull container image: %v\n", err)

			allContainersRunning = false
			continue
		}

		volumeBinds := []string{}
		for _, volume := range pod.Configs.Spec.Volumes {
			volumeBinds = append(volumeBinds, fmt.Sprintf("%v:%v", volume.Name, volume.HostPath.Path))

		}
		cont := container.NewContainer()
		cont_randomID := random.GenerateXid()
		cont.Name = pod.Configs.Spec.Containers[i].Name + cont_randomID
		cont.Image = pod.Configs.Spec.Containers[i].Image
		cont.Ports = pod.Configs.Spec.Containers[i].Ports

		ID, contName, err := container.CreateContainer(&pod.Configs.Spec.Containers[i], pauseMode, volumeBinds)

		if err != nil {
			fmt.Printf("Failed to create container: %v\n", err)
			allContainersRunning = false
			break
		}

		err = container.StartContainer(ID)
		if err != nil {
			fmt.Printf("Failed to start container: %s:%v\n", contName, err)
			allContainersRunning = false
			break
		} else {
			fmt.Printf("********* start container succeed! containerName: %s************\n", contName)
		}
	}

	if allContainersRunning {
		pod.PodPhase = Running
	} else {
		pod.PodPhase = Failed
	}
	pod.StartTime = time.Now()

}

func (pod *Pod) Stop(ctx context.Context, cli *client.Client) error {
	allContainerStopped := true

	if pod.Configs.Spec.Containers == nil {
		return fmt.Errorf("containers are nil")
	}

	for i := range pod.Configs.Spec.Containers {
		if err := pod.Configs.Spec.Containers[i].Stop(ctx, cli); err != nil {
			fmt.Printf("Failed to stop container at index %d: %v\n", i, err)
			allContainerStopped = false
		} else {
			pod.Configs.Spec.Containers[i].Status = "stopped"
		}
	}

	if pod.PauseContainer == nil {
		fmt.Println("PauseContainer is nil")
		allContainerStopped = false
	} else {
		if err := pod.PauseContainer.Stop(ctx, cli); err != nil {
			fmt.Printf("Failed to stop pause container: %v\n", err)
			allContainerStopped = false
		} else {
			pod.PauseContainer.Status = "stopped"
		}
	}

	if allContainerStopped {
		pod.PodPhase = Succeeded
	} else {
		pod.PodPhase = Failed
	}

	return nil
}

func (pod *Pod) Remove(ctx context.Context, cli *client.Client) error {
	allContainersRemoved := true
	for i := range pod.Configs.Spec.Containers {
		err := pod.Configs.Spec.Containers[i].Remove(ctx, cli)
		if err != nil {
			fmt.Printf("Failed to remove container: %v\n", err)
			allContainersRemoved = false
		} else {
			pod.Configs.Spec.Containers[i].Status = "removed"
		}
	}
	if err := pod.PauseContainer.Remove(ctx, cli); err != nil {
		fmt.Printf("Failed to remove pause container: %v\n", err)
		allContainersRemoved = false
	} else {
		pod.PauseContainer.Status = "removed"
	}

	if allContainersRemoved {
		pod.PodPhase = Succeeded
	} else {
		pod.PodPhase = Failed
	}
	return nil
}

func (p *Pod) DisplayStatus(pm *PodManager) {
	/**
	podname   podPhase  runTime   podIP   node
	*/
	//for _, c := range p.Configs.Spec.Containers {
	//	fmt.Printf("Container %s is currently %s.\n", c.Name, c.Status)
	//}
	//fmt.Printf("Pod Name: %s\n", p.Configs.Metadata.Name)
	//fmt.Printf("Pod Status: %s\n", p.PodPhase)
	//p.DisplayRunTime()
	//for podName, pod := range pods {
	//	fmt.Println("Pod Name:", podName)
	//	fmt.Printf("Pod Pointer: %p\n", pod)
	//	fmt.Println("Containers:")
	//	for _, container := range pod.Containers {
	//		fmt.Println("\t", container.Name)
	//	}
	//	fmt.Println("=====================")
	//}
	//for podName, pod := range pm.pods {
	//	duration := time.Since(pod.StartTime)
	//	fmt.Println(podName, "         ", "%s", pod.PodPhase, "         ", duration.Truncate(time.Second), "%s\n", pod.podIP)
	//}
	duration := time.Since(p.StartTime)
	fmt.Printf("%s              %s            %s            %s             %s\n", p.Configs.Metadata.Uid, p.Configs.Metadata.Name, p.PodPhase, duration.Truncate(time.Second), p.PodIP)
}

func (p *Pod) DisplayRunTime() {
	if p.PodPhase == Pending {
		fmt.Printf("Pod %s is pending.\n", p.Configs.Metadata.Name)
	} else {
		duration := time.Since(p.StartTime)
		fmt.Printf("Pod %s has been running for %s.\n", p.Configs.Metadata.Name, duration.Truncate(time.Second))
	}
}

func (p *Pod) DescribePod() {
	data, err := yaml.Marshal(p.Configs)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(string(data))
	fmt.Printf("\n")
}
