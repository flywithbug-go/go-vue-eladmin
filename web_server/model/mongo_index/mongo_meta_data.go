package mongo_index

import (
	"vue-admin/web_server/config"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/model/shareDB"

	"github.com/flywithbug/log4go"
	"gopkg.in/mgo.v2"
)

const (
	CollectionUser           = "user"
	CollectionLogin          = "login"
	CollectionApp            = "application"
	CollectionAppVersion     = "app_version"
	CollectionPermission     = "permission"
	CollectionRole           = "role"
	CollectionRolePermission = "role_permission"
)

type Index struct {
	Name      string
	Index     mgo.Index
	DropIndex []string
}

//唯一约束
var Indexes = []Index{
	{
		Name: CollectionPermission,
		Index: mgo.Index{
			Key:        []string{"alias"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_permission_f_alias_index",
		},
	},
	{
		Name: CollectionPermission,
		Index: mgo.Index{
			Key:        []string{"name"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_permission_f_name_index",
		},
	},
	{
		Name: CollectionUser,
		Index: mgo.Index{
			Key:        []string{"username"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_user_f_username_index",
		},
	},
	{
		Name: CollectionUser,
		Index: mgo.Index{
			Key:        []string{"email"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_user_f_email_index",
		},
	},
	{
		Name: CollectionRole,
		Index: mgo.Index{
			Key:        []string{"name"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_role_f_name_index",
		},
	},
	{
		Name: CollectionRole,
		Index: mgo.Index{
			Key:        []string{"alias"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_role_f_alias_index",
		},
	},
	{
		Name: CollectionApp,
		Index: mgo.Index{
			Key:        []string{"bundle_id"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_app_f_bundle_id_index",
		},
	},
	{
		Name: CollectionApp,
		Index: mgo.Index{
			Key:        []string{"name"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_app_f_name_index",
		},
	},
	{
		Name: CollectionAppVersion,
		Index: mgo.Index{
			Key:        []string{"version", "app_id"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_appVersion_f_version_f_appId_index",
		},
	},
	{
		Name: CollectionRolePermission,
		Index: mgo.Index{
			Key:        []string{"role_id", "permission_id"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_role_f_permission_f_role_index",
		},
	},
}

func CreateMgoIndex() {
	aMCfg := config.Conf().DBConfig

	for _, aMongoIndex := range Indexes {
		c := mongo.Collection(shareDB.DBName(), aMongoIndex.Name)
		if len(aMongoIndex.DropIndex) > 0 {
			for _, idxName := range aMongoIndex.DropIndex {
				if err := c.DropIndexName(idxName); err != nil {
					log4go.Warn(err.Error())
				}
			}
		}
		if aMCfg.ForceSync {
			if err := c.DropIndexName(aMongoIndex.Index.Name); err != nil {
				log4go.Warn(err.Error())
			}
		}
		err := c.EnsureIndex(aMongoIndex.Index)
		if err != nil {
			log4go.Warn(err.Error())
		}
	}

}
