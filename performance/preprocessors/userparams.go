package preprocessors

import . "tyrannosaurs/constant"

type UserParams struct {
	Data map[string][]string
	MinLen int
}

func (u *UserParams) AddParam(param string) {
	if _, ok := u.Data[param]; !ok {
		u.Data[param] = []string{}
	}
}

func (u *UserParams) AddUserParam(param string, value []string) {
	u.AddParam(param)
	u.Data[param] = value
	len := len(value)
	if len < u.MinLen {
		u.MinLen = len
	}
}

func (u *UserParams) GetValueByParams(param string) ([]string, error)  {
	if _, ok := u.Data[param]; ok {
		return u.Data[param][0:u.MinLen], nil
	}
	return nil, UserParamIsNotExist
}