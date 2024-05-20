package user

import (
	"bufio"
	"fmt"
	"github.com/dean2021/osutil/misc/array"
	"os/exec"
	"strings"
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
	// Executes the "useradd" command to create a user
	cmd := exec.Command("useradd", username)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Attempts to change the password for the newly created user
	if err := ChangePasswd(username, passwd); err != nil {
		return fmt.Errorf("failed to change password: %w", err)
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
	// Executes the "userdel" command with "-r" to remove the user along with their home directory
	cmd := exec.Command("userdel", "-r", username)
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
	// Executes the "passwd" command to change the user's password
	cmd := exec.Command("passwd", username)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s\n%s\n%s\n", passwd, passwd, passwd))
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to change password: %v", err)
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
	// Executes the "passwd -l" command to lock the user account
	cmd := exec.Command("passwd", "-l", username)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to lock user account: %v", err)
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
	// Executes the "passwd -u" command to unlock the user account
	cmd := exec.Command("passwd", "-u", username)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to unlock user account: %v", err)
	}
	return nil
}

// List retrieves a list of all user information in the system.
// Returns:
//
//	[]UserInfo - A slice containing information for all users.
//	error - Returns an error message if retrieval fails, otherwise nil.
func List() ([]UserInfo, error) {
	// Executes the "getent passwd" command to fetch user information
	cmd := exec.Command("getent", "passwd")
	outputReader, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to get user list: pipe error: %v", err)
	}

	scanner := bufio.NewScanner(outputReader)
	userList := make([]UserInfo, 0)

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to get user list: start error: %v", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		userList = append(userList, UserInfo{
			"Name":   array.Get(parts, 0),
			"Passwd": array.Get(parts, 1),
			"UID":    array.Get(parts, 2),
			"GID":    array.Get(parts, 3),
			"Desc":   array.Get(parts, 4),
			"Home":   array.Get(parts, 5),
			"Shell":  array.Get(parts, 6),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to get user list: scan error: %v", err)
	}
	return userList, nil
}
