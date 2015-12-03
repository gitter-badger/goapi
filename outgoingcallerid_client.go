// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING;

// Autogenerated by clientgenerator -r OutgoingCallerID -m CRD DO NOT EDIT

package api

import (
	"github.com/exotel/goapi/assets/types"
	"github.com/exotel/goapi/resources"
)

//OutgoingCallerIDService is the struct that defines the outgoingcallerid information and has fubnctions to perform outgoingcallerid level
type OutgoingCallerIDService struct {
	Client
}

//Create sets the action as create
func (__receiver_OService *OutgoingCallerIDService) Create() *OutgoingCallerIDService {
	__receiver_OService.action = types.CREATE
	__receiver_OService.data = resources.OutgoingCallerIDDetails{}
	__receiver_OService.url = resources.OutgoingCallerIDURLS[types.CREATE]
	return __receiver_OService
}

//Delete sets the action as types.DELETE
func (__receiver_OService *OutgoingCallerIDService) Delete() *OutgoingCallerIDService {
	__receiver_OService.action = types.DELETE
	__receiver_OService.data = struct{}{}
	__receiver_OService.url = resources.OutgoingCallerIDURLS[types.DELETE]
	return __receiver_OService
}

//Get sets the action as types.READ
func (__receiver_OService *OutgoingCallerIDService) Get() *OutgoingCallerIDService {
	if len(__receiver_OService.ResourceID) == 0 {
		__receiver_OService.action = types.BULKREAD
		__receiver_OService.data = resources.OutgoingCallerIDFilter{}
		__receiver_OService.url = resources.OutgoingCallerIDURLS[types.BULKREAD]
	} else {
		__receiver_OService.data = struct{}{}
		__receiver_OService.url = resources.OutgoingCallerIDURLS[types.READ]
		__receiver_OService.action = types.READ
	}
	return __receiver_OService
}

//ID sets the id for the resource request
func (__receiver_OService *OutgoingCallerIDService) ID(id string) *OutgoingCallerIDService {
	__receiver_OService.ResourceID = id
	switch __receiver_OService.action {
	case types.BULKREAD:
		__receiver_OService.data = struct{}{}
		__receiver_OService.url = resources.OutgoingCallerIDURLS[types.READ]
		__receiver_OService.action = types.READ

	}
	return __receiver_OService
}

//Filter set the filter for read operation for read
func (__receiver_OService *OutgoingCallerIDService) Filter(filter resources.OutgoingCallerIDFilter) *OutgoingCallerIDService {
	__receiver_OService.data = filter
	return __receiver_OService
}

//Details set the resource details for create resource requests
func (__receiver_OService *OutgoingCallerIDService) Details(details resources.OutgoingCallerIDDetails) *OutgoingCallerIDService {
	__receiver_OService.data = details
	return __receiver_OService
}

//OutgoingCallerID returns an instance of __receiver_OService
func (c *Client) OutgoingCallerID() *OutgoingCallerIDService {
	outgoingcalleridService := OutgoingCallerIDService{Client: *c}
	outgoingcalleridService.validActions = types.CREATE | types.READ | types.BULKREAD | types.DELETE | 0x00
	return &outgoingcalleridService
}
