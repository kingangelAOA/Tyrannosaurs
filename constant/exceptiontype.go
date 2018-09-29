package constant

import "errors"

var OpenTestFileError = errors.New("open test data file error")
var CVSDataParamError = errors.New("cvs data params is empty")
var InitCVSDataError = errors.New("init cvs data error")
var JsonExtractorJsonFormatError = errors.New("json extractor json format error")
var JsonExtractorPathFormatError = errors.New("json extractor path format error")
var UserParamIsNotExist = errors.New("this param is not exist in user params")
var StoreKeyIsNotExist = errors.New("this key is not exist in store")