package flavor

import (
	"fmt"
	"github.com/megamsys/opennebula-go/api"
	"github.com/megamsys/opennebula-go/xmlUtil"
)

type FlavorOpts struct {
	TemplateName string
	TemplateId   int
	TemplateData         string
}

/*
 * Given a templateId it would return the XML of that particular template
 *
 */

func (flavor *FlavorOpts) GetTemplate(endpoint string, key string) ([]interface{}, error) {

	client, err := api.RPCClient(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	args := []interface{}{key, -2, flavor.TemplateId, flavor.TemplateId}
	res, err := api.Call(client, "one.templatepool.info", args)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}




func (flavor *FlavorOpts) GetTemplateByName(endpoint string, key string) ([]*xmlUtil.VMTemplate, error) {

	client, err := api.RPCClient(endpoint)
	if err != nil {
		fmt.Println(err)
	}
//get all templates
	args := []interface{}{key, -2, -1, -1}
	templatePool, err := api.Call(client, "one.templatepool.info", args)

//iterate the pool to getbyname
fmt.Println(templatePool)
xmlStrt := xmlUtil.UnmarshallXml(templatePool[1])
//fmt.Println(xmlStrt.VmTemplate[1])
var matchedTemplate = make([]*xmlUtil.VMTemplate, len(xmlStrt.VmTemplate))

		for _, v := range xmlStrt.VmTemplate {
     if v.Name == flavor.TemplateName {
			fmt.Println(v)
         matchedTemplate[0] = v
			}
		}
	return matchedTemplate, nil
}

func (flavor *FlavorOpts) UpdateTemplate(endpoint string, key string) error {
	client, err := api.RPCClient(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	args := []interface{}{key, flavor.TemplateId, flavor.TemplateData, 0}
	templatePool, err := api.Call(client, "one.template.update", args)
  fmt.Println(templatePool)
	fmt.Println(err)
 return nil


}
