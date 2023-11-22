package helpers

func Messagess(MessageCode int) string {
	var messages string
	switch MessageCode {
	case 1000:
		messages = "Data found successfully"
	case 1001:
		messages = "Data not found "
	case 1002:
		messages = "Data Inserted successfully"
	case 1003:
		messages = "Data Updated Successfully"
	case 1004:
		messages = "Data Delere Successfully"
	case 2000:
		messages = "ID not found"
	case 2001:
		messages = "Error In update"
	case 2002:
		messages = "Error in create"
	case 5001:
		messages = "Unauthorized User"
	case 5002:
		messages = "Tokan Not Found"
	default:
		messages = "okk"
	}
	return messages
}
