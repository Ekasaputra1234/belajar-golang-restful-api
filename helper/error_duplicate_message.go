package helper

import (
	"gitlab.com/VNEU/logger/exception"
)

var DatabaseErrors = []exception.DatabaseError{
	{
		Contains: "cities.PRIMARY",
		Message:  "city already exists",
	},
	{
		Contains: "outlets.npwp",
		Message:  "npwp already exists",
	},
	{
		Contains: "marketing_structures.idx_marketing_structure_code_period",
		Message:  "code already exists in the same period",
	},
}
