package constant

import "errors"

var (
	OpenTestFileError               = errors.New("open test data file error")
	CVSDataParamError               = errors.New("cvs data params is empty")
	InitCVSDataError                = errors.New("init cvs data error")
	JsonExtractorJsonFormatError    = errors.New("json extractor json format error")
	JsonExtractorPathFormatError    = errors.New("json extractor path format error")
	UserParamIsNotExist             = errors.New("this param is not exist in user params")
	StoreKeyIsNotExist              = errors.New("this key is not exist in store")
	SliceExistsError                = errors.New("SliceExists given a non-slice type")
	JsonAssertFormatError           = errors.New("parameter reception is not a json string")
	JsonAssertReceiveIsNotString    = errors.New("parameter reception is not a string")
	HttpMergeCacheError             = errors.New("http merge cache error: repeating key")
	JsonExtractorReceiveFormatError = errors.New("json extractor receive data is not string")
	UserParamsDataIsNil             = errors.New("user params init data error")
	OpenEnvfileError                = errors.New("open env file failed or env file is not exists")
	EnvToStructError                = errors.New("env to struct error")
	EnvError                        = errors.New("env is not alpha beta prod")
)
