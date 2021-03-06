// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ec2

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

func (c *EC2) WaitUntilBundleTaskComplete(input *DescribeBundleTasksInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeBundleTasks",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "BundleTasks[].State",
				Expected: "complete",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "BundleTasks[].State",
				Expected: "failed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilConversionTaskCancelled(input *DescribeConversionTasksInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeConversionTasks",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "ConversionTasks[].State",
				Expected: "cancelled",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilConversionTaskCompleted(input *DescribeConversionTasksInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeConversionTasks",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "ConversionTasks[].State",
				Expected: "completed",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "ConversionTasks[].State",
				Expected: "cancelled",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "ConversionTasks[].State",
				Expected: "cancelling",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilConversionTaskDeleted(input *DescribeConversionTasksInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeConversionTasks",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "ConversionTasks[].State",
				Expected: "deleted",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilCustomerGatewayAvailable(input *DescribeCustomerGatewaysInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeCustomerGateways",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "CustomerGateways[].State",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "CustomerGateways[].State",
				Expected: "deleted",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "CustomerGateways[].State",
				Expected: "deleting",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilExportTaskCancelled(input *DescribeExportTasksInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeExportTasks",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "ExportTasks[].State",
				Expected: "cancelled",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilExportTaskCompleted(input *DescribeExportTasksInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeExportTasks",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "ExportTasks[].State",
				Expected: "completed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilImageAvailable(input *DescribeImagesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeImages",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Images[].State",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Images[].State",
				Expected: "failed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilImageExists(input *DescribeImagesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeImages",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "length(Images[]) > `0`",
				Expected: true,
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidAMIID.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilInstanceExists(input *DescribeInstancesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstances",
		Delay:       5,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "length(Reservations[]) > `0`",
				Expected: true,
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidInstanceID.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilInstanceRunning(input *DescribeInstancesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstances",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "running",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "shutting-down",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "terminated",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "stopping",
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidInstanceID.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilInstanceStatusOk(input *DescribeInstanceStatusInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstanceStatus",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "InstanceStatuses[].InstanceStatus.Status",
				Expected: "ok",
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidInstanceID.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilInstanceStopped(input *DescribeInstancesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstances",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "stopped",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "pending",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "terminated",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilInstanceTerminated(input *DescribeInstancesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstances",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "terminated",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "pending",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Reservations[].Instances[].State.Name",
				Expected: "stopping",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilKeyPairExists(input *DescribeKeyPairsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeKeyPairs",
		Delay:       5,
		MaxAttempts: 6,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "length(KeyPairs[].KeyName) > `0`",
				Expected: true,
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidKeyPair.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilNatGatewayAvailable(input *DescribeNatGatewaysInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeNatGateways",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "NatGateways[].State",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "NatGateways[].State",
				Expected: "failed",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "NatGateways[].State",
				Expected: "deleting",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "NatGateways[].State",
				Expected: "deleted",
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "NatGatewayNotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilNetworkInterfaceAvailable(input *DescribeNetworkInterfacesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeNetworkInterfaces",
		Delay:       20,
		MaxAttempts: 10,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "NetworkInterfaces[].Status",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidNetworkInterfaceID.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilPasswordDataAvailable(input *GetPasswordDataInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetPasswordData",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "length(PasswordData) > `0`",
				Expected: true,
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilSnapshotCompleted(input *DescribeSnapshotsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeSnapshots",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Snapshots[].State",
				Expected: "completed",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilSpotInstanceRequestFulfilled(input *DescribeSpotInstanceRequestsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeSpotInstanceRequests",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "SpotInstanceRequests[].Status.Code",
				Expected: "fulfilled",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "SpotInstanceRequests[].Status.Code",
				Expected: "schedule-expired",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "SpotInstanceRequests[].Status.Code",
				Expected: "canceled-before-fulfillment",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "SpotInstanceRequests[].Status.Code",
				Expected: "bad-parameters",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "SpotInstanceRequests[].Status.Code",
				Expected: "system-error",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilSubnetAvailable(input *DescribeSubnetsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeSubnets",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Subnets[].State",
				Expected: "available",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilSystemStatusOk(input *DescribeInstanceStatusInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeInstanceStatus",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "InstanceStatuses[].SystemStatus.Status",
				Expected: "ok",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilVolumeAvailable(input *DescribeVolumesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVolumes",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Volumes[].State",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Volumes[].State",
				Expected: "deleted",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilVolumeDeleted(input *DescribeVolumesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVolumes",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Volumes[].State",
				Expected: "deleted",
			},
			{
				State:    "success",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidVolume.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilVolumeInUse(input *DescribeVolumesInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVolumes",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Volumes[].State",
				Expected: "in-use",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "Volumes[].State",
				Expected: "deleted",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilVpcAvailable(input *DescribeVpcsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVpcs",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "Vpcs[].State",
				Expected: "available",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilVpcPeeringConnectionExists(input *DescribeVpcPeeringConnectionsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVpcPeeringConnections",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "status",
				Argument: "",
				Expected: 200,
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "InvalidVpcPeeringConnectionID.NotFound",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilVpnConnectionAvailable(input *DescribeVpnConnectionsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVpnConnections",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "VpnConnections[].State",
				Expected: "available",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "VpnConnections[].State",
				Expected: "deleting",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "VpnConnections[].State",
				Expected: "deleted",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

func (c *EC2) WaitUntilVpnConnectionDeleted(input *DescribeVpnConnectionsInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVpnConnections",
		Delay:       15,
		MaxAttempts: 40,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "pathAll",
				Argument: "VpnConnections[].State",
				Expected: "deleted",
			},
			{
				State:    "failure",
				Matcher:  "pathAny",
				Argument: "VpnConnections[].State",
				Expected: "pending",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
