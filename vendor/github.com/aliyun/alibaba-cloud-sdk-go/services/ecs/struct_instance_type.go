package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// InstanceType is a nested struct in ecs response
type InstanceType struct {
	MemorySize                  float64 `json:"MemorySize" xml:"MemorySize"`
	EniPrivateIpAddressQuantity int     `json:"EniPrivateIpAddressQuantity" xml:"EniPrivateIpAddressQuantity"`
	InstancePpsRx               int64   `json:"InstancePpsRx" xml:"InstancePpsRx"`
	CpuCoreCount                int     `json:"CpuCoreCount" xml:"CpuCoreCount"`
	Cores                       int     `json:"Cores" xml:"Cores"`
	Memory                      int     `json:"Memory" xml:"Memory"`
	InstanceTypeId              string  `json:"InstanceTypeId" xml:"InstanceTypeId"`
	InstanceBandwidthRx         int     `json:"InstanceBandwidthRx" xml:"InstanceBandwidthRx"`
	BaselineCredit              int     `json:"BaselineCredit" xml:"BaselineCredit"`
	InstanceType                string  `json:"InstanceType" xml:"InstanceType"`
	EniQuantity                 int     `json:"EniQuantity" xml:"EniQuantity"`
	GPUAmount                   int     `json:"GPUAmount" xml:"GPUAmount"`
	Generation                  string  `json:"Generation" xml:"Generation"`
	SupportIoOptimized          string  `json:"SupportIoOptimized" xml:"SupportIoOptimized"`
	InstanceTypeFamily          string  `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
	InitialCredit               int     `json:"InitialCredit" xml:"InitialCredit"`
	InstancePpsTx               int64   `json:"InstancePpsTx" xml:"InstancePpsTx"`
	LocalStorageAmount          int     `json:"LocalStorageAmount" xml:"LocalStorageAmount"`
	InstanceFamilyLevel         string  `json:"InstanceFamilyLevel" xml:"InstanceFamilyLevel"`
	LocalStorageCapacity        int64   `json:"LocalStorageCapacity" xml:"LocalStorageCapacity"`
	GPUSpec                     string  `json:"GPUSpec" xml:"GPUSpec"`
	LocalStorageCategory        string  `json:"LocalStorageCategory" xml:"LocalStorageCategory"`
	InstanceBandwidthTx         int     `json:"InstanceBandwidthTx" xml:"InstanceBandwidthTx"`
}
