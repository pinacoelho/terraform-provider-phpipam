package phpipam

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourcePHPIPAMReserveFirstFreeAddress() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePHPIPAMReserveFirstFreeAddressRead,
		Schema: map[string]*schema.Schema{
			"subnet_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePHPIPAMReserveFirstFreeAddressRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*ProviderPHPIPAMClient).subnetsController
	out, err := c.ReserveFirstFreeAddress(d.Get("subnet_id").(int))
	if err != nil {
		return err
	}
	if out == "" {
		return errors.New("Subnet has no free IP addresses")
	}

	d.SetId(out)
	d.Set("ip_address", out)

	return nil
}
