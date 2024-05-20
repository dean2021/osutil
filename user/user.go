package user

// UserInfo is a type representing a user's personal information, stored as a key-value pair using a map.
// It allows for storing various types of values, providing flexible and dynamic management of user attributes.
type UserInfo map[string]interface{}

// SetValue sets the value for a specified field in the user information.
// This function enables assigning values dynamically to fields, enhancing code flexibility and extensibility.
// Parameters:
// - fieldName: The field name, a string type, used to identify the field to set.
// - value: An interface{} type, which will be assigned to the specified field.
// Assigns the value to the field identified by the field name, performing a dynamic value setting operation.
func (i *UserInfo) SetValue(fieldName string, value interface{}) {
	(*i)[fieldName] = value
}

// GetValue retrieves the value for a specified field in the user information.
// This method facilitates convenient access to a specific field's value, enabling data retrieval.
// Parameters:
// - fieldName: The field name, a string type, used to identify the field to retrieve.
// Returns:
// - Returns the value of the specified field as an interface{}.
// Since UserInfo stores arbitrary data as key-value pairs, the returned value is also an interface{} and requires type assertion to obtain its actual type.
func (i *UserInfo) GetValue(fieldName string) interface{} {
	return (*i)[fieldName]
}
