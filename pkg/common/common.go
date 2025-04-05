/**
 * @Author: litsky
 * @Date:   2023/12/26
 * @Last Modified by:   litsky
 * @Last Modified time: 2023/12/26
 * @License: MIT
 */

package common

type ItemType string

const (
	GPU_TYPE    ItemType = "GPU"
	CPU_type    ItemType = "CPU"
	MEMORY_TYPE ItemType = "MEMORY"
	//...
)

type LabelType string

const (
	Product_LabelType     LabelType = "Product"
	BillingItem_LabelType LabelType = "BillingItem"
)
