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
	CollectionUserRole       = "user_role"
	CollectionVerify         = "verify"
	CollectionMenu           = "menu"
	CollectionMenuRole       = "menu_role"
	CollectionFile           = "file"
	CollectionPicture        = "picture"
)

// monitor
const (
	MonitorDBName   = "monitor"
	CollectionLog   = "log"
	CollectionVisit = "visit"
)

type Index struct {
	Collection string
	DBName     string
	Index      mgo.Index
	DropIndex  []string
}

//唯一约束
var Indexes = []Index{
	{
		DBName:     shareDB.DBName(),
		Collection: CollectionPermission,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionPermission,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionUser,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionUser,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionRole,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionRole,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionApp,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionApp,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionAppVersion,
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
		DBName:     shareDB.DBName(),
		Collection: CollectionRolePermission,
		Index: mgo.Index{
			Key:        []string{"role_id", "permission_id"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_role_f_permission_f_role_index",
		},
	},
	{
		DBName:     shareDB.DBName(),
		Collection: CollectionUserRole,
		Index: mgo.Index{
			Key:        []string{"role_id", "user_id"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_user_f_role_index",
		},
	},
	{
		DBName:     shareDB.DBName(),
		Collection: CollectionMenuRole,
		Index: mgo.Index{
			Key:        []string{"role_id", "menu_id"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_menu_f_role_index",
		},
	},
	{
		DBName:     shareDB.DBName(),
		Collection: CollectionFile,
		Index: mgo.Index{
			Key:        []string{"md5"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_file_f_md5_index",
		},
	},
	{
		DBName:     MonitorDBName,
		Collection: CollectionVisit,
		Index: mgo.Index{
			Key:        []string{"client_ip", "uuid"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_uuid_f_client_ip_index",
		},
	},

	{
		DBName:     MonitorDBName,
		Collection: CollectionLog,
		Index: mgo.Index{
			Key:        []string{"request_id"},
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
			Name:       "c_request_id_f_index",
		},
	},
}

func CreateMgoIndex() {
	aMCfg := config.Conf().DBConfig
	for _, aMongoIndex := range Indexes {
		_, c := mongo.Collection(aMongoIndex.DBName, aMongoIndex.Collection)
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
