package vpc

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RemoveIPv6TranslatorAclListEntry invokes the vpc.RemoveIPv6TranslatorAclListEntry API synchronously
// api document: https://help.aliyun.com/api/vpc/removeipv6translatoracllistentry.html
func (client *Client) RemoveIPv6TranslatorAclListEntry(request *RemoveIPv6TranslatorAclListEntryRequest) (response *RemoveIPv6TranslatorAclListEntryResponse, err error) {
	response = CreateRemoveIPv6TranslatorAclListEntryResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveIPv6TranslatorAclListEntryWithChan invokes the vpc.RemoveIPv6TranslatorAclListEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/removeipv6translatoracllistentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveIPv6TranslatorAclListEntryWithChan(request *RemoveIPv6TranslatorAclListEntryRequest) (<-chan *RemoveIPv6TranslatorAclListEntryResponse, <-chan error) {
	responseChan := make(chan *RemoveIPv6TranslatorAclListEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveIPv6TranslatorAclListEntry(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// RemoveIPv6TranslatorAclListEntryWithCallback invokes the vpc.RemoveIPv6TranslatorAclListEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/removeipv6translatoracllistentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveIPv6TranslatorAclListEntryWithCallback(request *RemoveIPv6TranslatorAclListEntryRequest, callback func(response *RemoveIPv6TranslatorAclListEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveIPv6TranslatorAclListEntryResponse
		var err error
		defer close(result)
		response, err = client.RemoveIPv6TranslatorAclListEntry(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// RemoveIPv6TranslatorAclListEntryRequest is the request struct for api RemoveIPv6TranslatorAclListEntry
type RemoveIPv6TranslatorAclListEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AclEntryId           string           `position:"Query" name:"AclEntryId"`
}

// RemoveIPv6TranslatorAclListEntryResponse is the response struct for api RemoveIPv6TranslatorAclListEntry
type RemoveIPv6TranslatorAclListEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveIPv6TranslatorAclListEntryRequest creates a request to invoke RemoveIPv6TranslatorAclListEntry API
func CreateRemoveIPv6TranslatorAclListEntryRequest() (request *RemoveIPv6TranslatorAclListEntryRequest) {
	request = &RemoveIPv6TranslatorAclListEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "RemoveIPv6TranslatorAclListEntry", "vpc", "openAPI")
	return
}

// CreateRemoveIPv6TranslatorAclListEntryResponse creates a response to parse from RemoveIPv6TranslatorAclListEntry response
func CreateRemoveIPv6TranslatorAclListEntryResponse() (response *RemoveIPv6TranslatorAclListEntryResponse) {
	response = &RemoveIPv6TranslatorAclListEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
