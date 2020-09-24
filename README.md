sample-container-runtime
========================

This is a simple container-runtime implemented for learning purposes.

**Stay tuned as this evolves!**

## Introduction to container technology

In general, a container is lightweight virtualization that, among others, allows to:
* run a process in separate, isolated sandbox (environment)
* have a separate filesystem mounted to a container
* isolate a container from the security perspective (allow/disallow system calls)
* limit the privileges of the process run in a container
* limit resources allocated for a container

Containers are based on following concepts:

* `namespaces` - Linux (or kernel) namespaces allows to create a new logical instances of an operating system's subsystems. Other definition says that "Linux namespace is a scoped view of your underlying Linux system".
As a result, the new instances are isolated from the OS they are running on and from the instances running in other namespaces.
There are following namespaces in Linux:
  * **UTS namespace** is about isolating hostnames. This namespace allows to set a different hostname for a container.
  * **MNT namespace** allows to mount a separate file system for a container.
  * **PID namespace** gives a container an isolated view on currently running Linux processes. As a result, a container will see only its own processes (processes of host OS will not be visible).
  * **IPC namespace** isolates an inter-process communication. It prevents processes in different namespaces from establishing a shared memory to communicate with each other.
  * **USER namespace** allows to create a separate, (usually) privileged user (technically it's a logical mapping of a user created in host OS, I will explain it later) within the namespace. Users configured in a host OS are not visible from a container.
  * **NET namespace** creates a logical instance of a Linux network stack. A container has its own list of network interfaces, routing table and iptables rules.
* `cgroups` and `setrlimit` - these both mechanisms are used to limit usage of resources (e.g. memory, disk I/O, CPU time) for a container.
* `root capabilities` - capabilities limits the privileges of root user of container.
* `Pivot_root` - the mechanism to change the root file system for a container.

Technically, a container is just a separate process, which is isolated from the host OS by using the concept of Linux namespaces.
Moreover, resources and privileges of this process are limited. All together creates the abstraction of a container.

## Refs

* [Code to accompany the "Namespaces in Go" series of articles](https://github.com/teddyking/ns-process)
* [Linux containers in 500 lines of code](https://blog.lizzie.io/linux-containers-in-500-loc.html)
* [Whitepaper - Understanding and hardening Linux Containers](https://github.com/osinstom/containers-impl-c)
* [Build Your Own Container Using Less than 100 Lines of Go](https://www.infoq.com/articles/build-a-container-golang/)
* [containers-impl-c](https://github.com/osinstom/containers-impl-c)
* [A deep dive into Linux namespaces, part 3](http://ifeanyi.co/posts/linux-namespaces-part-3/#pid-namespaces)
* [Namespaces in operation, part 7: Network namespaces](https://lwn.net/Articles/580893/)
* [Introducing Linux Network Namespaces](https://blog.scottlowe.org/2013/09/04/introducing-linux-network-namespaces/)
* [Run a command in unique namespaces](https://github.com/iffyio/isolate)
* [Mount namespaces and shared subtrees - LWN article](https://lwn.net/Articles/689856/)