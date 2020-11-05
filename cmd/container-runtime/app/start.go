package app

import (
	log "github.com/sirupsen/logrus"
)

func startContainer(containerName string) {
	containerInfo, err := getContainerInfoByName(containerName)
	if err != nil {
		log.Errorf("Get contaienr info by name %s error %v", containerName, err)
		return
	}
	Run(containerInfo.Detached, containerInfo.Id, containerInfo.Command, containerInfo.ResConf, containerName, containerInfo.Volume, containerInfo.ImageName, containerInfo.Env, containerInfo.Network, containerInfo.PortMapping)
}
