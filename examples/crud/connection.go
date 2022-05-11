package crud

import (
	"github.com/wlevene/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	_ = mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
}
