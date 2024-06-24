package openstack

const (

	// Create params
	imageID          string = "image"
	imageIDDesc      string = "OpenStack image identifier"
	instanceType     string = "flavor"
	instanceTypeDesc string = "OpenStack flavor type for the machine running the cluster. Default is m1.xlarge."
	diskSize         string = "disk-size"
	diskSizeDesc     string = "Disk size in GB for the machine running the cluster. Default is 100."
	networkName      string = "network"
	networkNameDesc  string = "OpenStack network name for the machine running the cluster."
	instanceName     string = "instance-name"
	instanceNameDesc string = "name of instance like 'woodchuck-alln-01'"
	instancePassword string = "instance-password"
	instancePasswordDesc string = "Password for crc instance"


	// default values
	ocpInstanceType               string = "m1.xlarge"
	ocpDefaultRootBlockDeviceSize int    = 100
)
