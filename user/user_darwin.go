package user

import "errors"

func Create(name string) error {
	return errors.New("implement me")
}

func Delete(name string) error {
	return errors.New("implement me")
}

func ChangePasswd(name string, passwd string) error {
	return errors.New("implement me")
}

func Lock(name string) error {
	return errors.New("implement me")
}

func UnLock(name string) error {
	return errors.New("implement me")
}

func List() ([]UserInfo, error) {
	return nil, errors.New("implement me")
}
