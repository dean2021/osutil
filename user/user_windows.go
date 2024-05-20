package user

import (
	"fmt"
	"os/exec"
)

// Create is a function that creates a new user and sets their password.
// Parameters:
//
//	username - The username of the new user.
//	passwd - The password to set for the new user.
//
// Returns:
//
//	error - An error if the user creation or password change fails, otherwise nil.
func Create(username string, passwd string) error {
	cmd := exec.Command("net", "user", username, passwd, "/ADD")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// Delete removes a user.
// Parameters:
//
//	username - The name of the user to be deleted.
//
// Returns:
//
//	error - Returns an error message if deletion fails, otherwise nil.
func Delete(username string) error {
	cmd := exec.Command("net", "user", username, "/delete")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// ChangePasswd changes a user's password.
// Parameters:
//
//	username - The name of the user whose password is to be changed.
//	passwd - The new password.
//
// Returns:
//
//	error - Returns an error message if the password change fails, otherwise nil.
func ChangePasswd(username string, passwd string) error {
	cmd := exec.Command("net", "user", username, passwd)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to change user password: %w", err)
	}
	return nil
}

// Lock locks a user account.
// Parameters:
//
//	username - The name of the user whose account is to be locked.
//
// Returns:
//
//	error - Returns an error message if locking fails, otherwise nil.
func Lock(username string) error {
	cmd := exec.Command("net", "user", username, "/active:no")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to lock user: %w", err)
	}
	return nil
}

// Unlock unlocks a user account.
// Parameters:
//
//	username - The name of the user whose account is to be unlocked.
//
// Returns:
//
//	error - Returns an error message if unlocking fails, otherwise nil.
func Unlock(username string) error {
	cmd := exec.Command("net", "user", username, "/active:yes")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to activate user: %w", err)
	}
	return nil
}

type Win32UserAccount struct {
	AccountType        int64  `json:"accountType"`
	Caption            string `json:"caption"`
	Description        string `json:"description"`
	Disabled           bool   `json:"disabled"`
	Domain             string `json:"domain"`
	FullName           string `json:"fullName"`
	InstallDate        string `json:"installDate"`
	LocalAccount       bool   `json:"localAccount"`
	Lockout            bool   `json:"lockout"`
	PasswordChangeable bool   `json:"passwordChangeable"`
	PasswordExpires    bool   `json:"passwordExpires"`
	PasswordRequired   bool   `json:"passwordRequired"`
	Name               string `json:"name"`
	SID                string `json:"sid"`
	SIDType            int64  `json:"sidType"`
	Status             string `json:"status"`
}

func getWin32UserAccount() ([]Win32UserAccount, error) {
	var s []Win32UserAccount
	err := wmi.Query("SELECT * FROM Win32_UserAccount", &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// List retrieves a list of all user information in the system.
// Returns:
//
//	[]UserInfo - A slice containing information for all users.
//	error - Returns an error message if retrieval fails, otherwise nil.
func List() ([]UserInfo, error) {
	accounts, err := getWin32UserAccount()
	if err != nil {
		return nil, err
	}
	userList := make([]UserInfo, 0)
	for _, account := range accounts {
		userList = append(userList, UserInfo{
			"AccountType":  account.AccountType,
			"Caption":      account.Caption,
			"Description":  account.Description,
			"Disabled":     account.Disabled,
			"Domain":       account.Domain,
			"FullName":     account.FullName,
			"InstallDate":  account.InstallDate,
			"LocalAccount": account.LocalAccount,
			"Lockout":      account.Lockout,
			"Name":         account.Name,
			"SID":          account.SID,
			"SIDType":      account.SIDType,
			"Status":       account.Status,
		})
	}
	return userList, nil
}
