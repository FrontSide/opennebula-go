package compute

import (
	"encoding/xml"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/opennebula-go/api"
	"github.com/megamsys/opennebula-go/template"
	"github.com/megamsys/opennebula-go/virtualmachine"
)

const (
	TEMPLATE_INSTANTIATE = "one.template.instantiate"
	ONE_VM_ACTION        = "one.vm.action"
	DELETE               = "delete"
)

type VirtualMachine struct {
	Name         string
	TemplateName string
	TemplateId   int
	Assembly_id  string
	Cpu          string
	VCpu         string
	Memory       string
	Client       *api.Rpc
}

/**
 *
 * Creates a new VirtualMachine
 *
 **/
func (v *VirtualMachine) Create() ([]interface{}, error) {
  log.Debugf("[one-go] create")
	templateObj := template.TemplateReqs{TemplateName: v.TemplateName, Client: v.Client}

	/*
	 * get a particular template to configure it
	 */
	XMLtemplate, err := templateObj.GetTemplateByName()
	if err != nil {
		return nil, fmt.Errorf("one template list %q failure: %s", v.TemplateName, err)
	}

	/*
	 * Assign Values
	 */

	XMLtemplate[0].Template.Cpu = v.Cpu
	XMLtemplate[0].Template.VCpu = v.VCpu
	XMLtemplate[0].Template.Memory = v.Memory

	XMLtemplate[0].Template.Context.Assembly_id = v.Assembly_id

	finalXML := template.UserTemplates{}
	finalXML.UserTemplate = XMLtemplate

	finalData, _ := xml.Marshal(finalXML.UserTemplate[0].Template)
	data := string(finalData)

	/*
	 * Instantiate a template
	 */
	args := []interface{}{v.Client.Key, finalXML.UserTemplate[0].Id, v.Name, false, data}
	res, err := v.Client.Call(v.Client.RPCClient, TEMPLATE_INSTANTIATE, args)
	if err != nil {
		return nil, fmt.Errorf("one vmcreate failure: %s", err)
	}
	return res, nil
}

/**
 *
 * Deletes a new virtualMachine
 *
 **/
func (v *VirtualMachine) Delete() ([]interface{}, error) {
	log.Debugf("[one-go] delete")

	vmObj := virtualmachine.VirtualMachineReqs{VMName: v.Name, Client: v.Client}

	SingleVM, err := vmObj.GetVirtualMachineByName()
	if err != nil {
		return nil, fmt.Errorf("one vmlist %q failure: %s", v.Name, err)
	}

	args := []interface{}{v.Client.Key, DELETE, SingleVM[0].Id}
	res, err := v.Client.Call(v.Client.RPCClient, ONE_VM_ACTION, args)
	if err != nil {
		return nil, fmt.Errorf("one vmdelete %q failure: %s", v.Name, err)
	}

	return res, nil

}
