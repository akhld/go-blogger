This is the written blog post of my previous talk about Introduction to Mesos which I gave at Agoda machine learning meetup.

The post is split into mainly of the following pieces:

- Challenges
- Datacenter partitioning and resource management
- Why Apache Mesos

#### Challenges
Everybody is moving towards microservices, especially If your company is doing good, then you probably are using a bunch of microservices by now. Microservices are built to do one thing and they does it well (just like the unix commands), also these microservices can capture the entire organizational structure very well, they can be built, tested, deployed and rolled back  in isolation without interrupting each other. But, if you are not designing it in the distributed system way, then your service might have reached to a point where it won’t fit on a single machine anymore, because it needs more memory or more computing cores for processing the “big data” or the amount of concurrent users. The users and the data will grow everyday as your  organization start to grow.

#### Datacenter partitioning and resource management

Typically the organization who uses multiple machines hosting multiple services partitions the datacenter as follows:

![cluster-partitioning](https://hacked.work/blog/wp-content/uploads/2017/02/partitioning.png)

Dealing with failures is an important task here. Imagine the top most switch connecting the Cassandra rack fails and that means obviously no one is querying the Cassandra and it’s a blackout.

This is an actual email I got from the work from the infra guys saying the OpenStack hosting a bunch of machine rebooted itself and all the services running on it went down.

![failure-email](https://hacked.work/blog/wp-content/uploads/2017/02/failure-email.png)

Going back to the static partitioning, the failure will look like this:

![failure-look](https://hacked.work/blog/wp-content/uploads/2017/02/failure-looks.jpg)

Another challenge here is the resource utilizations. Mostly, if you look at the resource utilization you could see that the machines holding Hadoop, kafka uses about 30-40% of the resources and the Java applications uses more resources during the day time, but during the night time they are almost idle.

![resource-utilization](https://hacked.work/blog/wp-content/uploads/2017/02/resource-utilization-1.jpg)

As you can see most of the resources are idle here, it will make more sense to have something like the following to obviously reduce the cost and utilize the resources more efficiently.

![resource-utilization](https://hacked.work/blog/wp-content/uploads/2017/02/resource-utilization-2.jpg)

#### Why Apache Mesos?

Using Apache Mesos greatly tackles all the challenges listed above. Apache Mesos is a modern general purpose cluster manager, which is initially started off at UC Berkeley in 2009 and then hardened at Twitter, majority of the stuffs at Twitter are indeed running on top of Mesos over tens of thousands of nodes. Mesos also exposes a rich set of API in various programming languages like Java/c++/Python etc. Also a bunch of top level Apache projects are ported to run on top of Mesos already:

![native-apps](https://hacked.work/blog/wp-content/uploads/2017/02/native-apps.png)

Like most of the distributed systems, Apache Mesos has a master slave architecture:

![mesos-architecture](https://hacked.work/blog/wp-content/uploads/2017/02/mesos-architecture.png)

Where the schedulers/frameworks interact with the master machine to schedule the tasks on the slave machines.

![mesos-scheduler](https://hacked.work/blog/wp-content/uploads/2017/02/mesos-scheduler.png)

##### Scheduling jobs with Mesosphere Marathon
This part is pretty straightforward, you build your application, containerize it, it could be a binary, a tgz or even a docker image. Upload the package to a place where it can be accessed by the slave machines (HDFS/S3/HTTP) ,describe the service and then schedule it. You can use the Marathon Web UI, REST endpoint or the CLI for this.
![marathon-scheduler](https://hacked.work/blog/wp-content/uploads/2017/02/marathon-schedule.png)

Here’s an example JSON:

![mesos-hello-world](https://hacked.work/blog/wp-content/uploads/2017/02/hello-marathon.png)

There’s more to write, you can find the talk slides [from here}(https://docs.google.com/presentation/d/1GGQFFOsidQsRGK-IlQ3X7GpGsqAdLwqxjfz4QARFG1s/edit?usp=sharing).