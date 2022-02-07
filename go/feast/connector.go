package feast

import (
	"errors"
	"fmt"
)

func getOnlineStore(config *RepoConfig) (OnlineStore, error) {
	onlineStoreType, ok := getOnlineStoreType(config.OnlineStore)
	if !ok {
		return nil, errors.New(fmt.Sprintf("could not get online store type from online store config: %+v", config.OnlineStore))
	}
	if onlineStoreType == "redis" {
		onlineStore, err := NewRedisOnlineStore(config.Project, config.OnlineStore)
		return onlineStore, err
	} else {
		// TODO(willem): Python connectors here
		return nil, errors.New("not implemented")
	}
}
