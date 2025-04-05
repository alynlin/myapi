/**
 * @Author: litsky
 * @Date:   2023/12/23
 * @Last Modified by:   litsky
 * @Last Modified time: 2023/12/23
 * @License: MIT
 */

package errors

type ErrorCode int

const (
	DB_ERROR ErrorCode = iota + 1000
	ErrRecordNotFound
)

const (
	OTHER ErrorCode = iota + 2000
	TYPE_CAST_ERROR
	LAGO_ERROR
	PRICE_ERROR
)
