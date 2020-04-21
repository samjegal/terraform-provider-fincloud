package locks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
)

var fincloudMutexKV = mutexkv.NewMutexKV()

func ByID(id string) {
	fincloudMutexKV.Lock(id)
}

func ByName(name string, resourceType string) {
	updatedName := resourceType + "." + name
	fincloudMutexKV.Lock(updatedName)
}

func ByHash(hash int) {
	fincloudMutexKV.Lock(string(hash))
}

func UnlockByID(id string) {
	fincloudMutexKV.Unlock(id)
}

func UnlockByName(name string, resourceType string) {
	updatedName := resourceType + "." + name
	fincloudMutexKV.Unlock(updatedName)
}

func UnlockByHash(hash int) {
	fincloudMutexKV.Unlock(string(hash))
}
