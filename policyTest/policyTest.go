package policyTest

import (
	"fmt"
	"jasondrogba/alluxio-cacheTest/alluxioTest"
	"jasondrogba/alluxio-cacheTest/metricsTest"
	"jasondrogba/alluxio-cacheTest/sshTest"
	"jasondrogba/alluxio-cacheTest/startTest"
)

func PolicyTest(instanceMap map[string]string, policy string, count int) (float64, float64) {

	fmt.Println("start Alluxio:", policy)
	startTest.StartTest(instanceMap["Ec2Cluster-default-masters-0"], policy)
	fmt.Println("LOAD Alluxio:", "worker0:1~13,worker1:1~10:worker2:1~5")
	sshTest.SshTest(instanceMap)
	fmt.Println("READ Alluxio:", count)
	alluxioTest.ReadAlluxio(instanceMap["Ec2Cluster-default-masters-0"], count, false)
	fmt.Println("METRIC Alluxio:")
	resultRemote, resultUFS := metricsTest.BackProcess(instanceMap)
	return resultRemote, resultUFS
}

func DynamicTest(instanceMap map[string]string, count int) (float64, float64) {
	fmt.Println("start Alluxio:", "REPLICA")
	startTest.StartTest(instanceMap["Ec2Cluster-default-masters-0"], "REPLICA")
	fmt.Println("LOAD Alluxio:", "worker0:1~13,worker1:1~10:worker2:1~5")
	sshTest.SshTest(instanceMap)
	fmt.Println("READ Alluxio:", count)
	alluxioTest.ReadAlluxio(instanceMap["Ec2Cluster-default-masters-0"], count, true)
	fmt.Println("METRIC Alluxio:")
	resultRemote, resultUFS := metricsTest.BackProcess(instanceMap)
	return resultRemote, resultUFS

}
