package virtualmachine

import (
	"encoding/xml"
	"github.com/megamsys/opennebula-go/api"
)
const (
		ONE_DISK_SNAPSHOT    = "one.vm.disksaveas"
)
type Query struct {
	VMName string
	VMId   int
	T      *api.Rpc
}

type UserVMs struct {
	UserVM []*UserVM `xml:"VM"`
}

type UserVM struct {
	Id   int    `xml:"ID"`
	Uid  int    `xml:"UID"`
	Name string `xml:"NAME"`
}

// Given a name, this function will return the VM
func (v *Query) GetByName() ([]*UserVM, error) {
	args := []interface{}{v.T.Key, -2, -1, -1, -1}
	VMPool, err := v.T.Call(api.VMPOOL_INFO, args)
	if err != nil {
		return nil, err
	}

	xmlVM := UserVMs{}
	assert, _ := VMPool[1].(string)
	if err = xml.Unmarshal([]byte(assert), &xmlVM); err != nil {
		return nil, err
	}
	var matchedVM = make([]*UserVM, len(xmlVM.UserVM))

	for _, u := range xmlVM.UserVM {
		if u.Name == v.VMName {
			matchedVM[0] = u
		}
	}

	return matchedVM, nil

}
/**
 *
 * VM save as a new Image (DISK_SNAPSHOT)
 *
 **/
func (v *Query) DiskSnap() ([]interface{}, error) {
	args := []interface{}{v.T.Key,v.VMId,0,v.VMName,"",-1}
	res, err := v.T.Call("one.vm.disksaveas", args)
	//close connection
	defer v.T.Client.Close()
	if err != nil {
		return nil, err
	}

	return res, nil
}
