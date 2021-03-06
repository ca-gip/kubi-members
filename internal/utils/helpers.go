package utils

import (
	"errors"
	"k8s.io/klog/v2"
	"os"
)

func Check(err error) {
	if err != nil {
		klog.Errorf(err.Error())
	}
}

func Checkf(err error, msg string) {
	if err != nil {
		klog.Errorf("%v : %v", msg, err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetClusterRole(str string) (error, ClusterRole){
	switch str {
	case OpsRole.String():
		return nil, OpsRole
	case AdminRole.String():
		return nil, AdminRole
	case CustomerRole.String():
		return nil, CustomerRole
	case AppRole.String():
		return nil, AppRole
	}
	return errors.New("unknown ClusterRole"), -1
}
