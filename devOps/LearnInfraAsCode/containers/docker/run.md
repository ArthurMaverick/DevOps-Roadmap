# CLI COMMAND
docker run [OPTION] IMAGE [COMMAND] [ARG...]

# DESCRIPTION
  Docker run é o conjunto de **/containers/create** com **/containers/(id)/ start** 
  Um container interrompido pode ser reiniciado com todas as alteraçoes anterios intactas usando **docker start**.


# OPTIONS

## Add host
<!-- --add-host [HOST/IP]:[HOST/IP] Adiciona um host ao container. -->


## Attach
--attach -a [STDIN] [STDOUT] [STDERR] - Attach STDIN, STDOUT, and STDERR to a container's STDIO streams.

- Permite que um terminal conecte um contêiner em execução. Ele permite que você conecte para processar 'STDIO em outro terminal.

**Qual é o benefício?**
- Isso permite que você visualize sua saída contínua ou controle-a interativamente, como se os comandos estivessem sendo executados diretamente em seu terminal.

## blkio
<!--  --blkio-weight [WEIGHT] - Set the Block I/O (BLKIO) weight.

--blkio-weight-device [DEVICE]:[WEIGHT] - Set the Block I/O (BLKIO) weight for a device.

--blkio-device-read-bps [DEVICE]:[RATE] - Limit read rate (bytes per second) from a device (e.g., '--blkio-device-read-bps /dev/sda:1').

--blkio-device-write-bps [DEVICE]:[RATE] - Limit write rate (bytes per second) to a device (e.g., '--blkio-device-write-bps /dev/sda:1').

--blkio-device-read-iops [DEVICE]:[RATE] - Limit read rate (IO per second) from a device (e.g., '--blkio-device-read-iops /dev/sda:1').
 -->


## Capabilities
--cap-add [CAPABILITY] - Add Linux capabilities.
--cap-drop [CAPABILITY] - Drop Linux capabilities.

## Cgroup

--cgroup-parent [PATH] - Set the parent cgroup for the container.
--cgroupns [NAMESPACE] - Set the cgroup namespace mode for the container.

# Cpu 
--cpu-period - limit CPU realtime in microseconds
--cpu-rt-runtime -  limit CPU realtime in microseconds
--cpu-shares - CPU shares 
--cpus - number of CPU
--cpuset-cpus - CPU in which to allow execution (0-3,0,1)
--cpuset-mems - MEMs in which t o allow execution (0-3,0,1)



# Detach 
--detach, -d => run container in backgroud and print container ID
--detach-keys => Override the key sequence for detaching a container


