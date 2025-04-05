/**
 * @Author: litsky
 * @Date:   2025/4/4
 * @Last Modified by:   litsky
 * @Last Modified time: 2025/4/4
 * @License: MIT
 */

package main

import (
	"fmt"
	"time"
)

func main() {

	ticket := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticket.C:
			fmt.Println(time.Now())
		}

	}

}
