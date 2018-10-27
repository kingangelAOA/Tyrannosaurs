package component

import (
	"sort"
	"math/rand"
	. "tyrannosaurs/constant"
)

type UserParams struct {
	Data   map[string][]string
	minLen int
	keys   []string
	Base
}

func (u *UserParams) pre() error {
	var flag []int
	if u.Data == nil {
		return UserParamsDataIsNil
	}
	for k, v := range u.Data {
		u.keys = append(u.keys, k)
		flag = append(flag, len(v))
	}
	if len(flag) == 0 {
		return UserParamsDataIsNil
	}
	sort.Ints(flag)
	u.minLen = flag[0]
	return nil
}

func (u *UserParams) getValues() (result map[string]string, err error) {
	once.Do(func() {
		err = u.pre()
	})
	if err != nil {
		return nil, err
	}
	i := rand.Intn(u.minLen)
	for _, k := range u.keys {
		if v, ok := u.Data[k]; ok {
			result[k] = v[i]
		}
	}
	return result, nil
}

func (u *UserParams) Notify() error {
	for _, o := range u.ObserverList {
		result, err := u.getValues()
		if err != nil {
			return err
		}
		if err := o.Receive(result); err != nil {
			return err
		}
	}
	return nil
}
