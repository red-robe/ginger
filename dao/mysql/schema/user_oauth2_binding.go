package schema

/*
   This code is generated for gendry
   ginger-cli schema -d ginger_db -t user_oauth2_binding
*/

const UserOauth2BindingTableName = "user_oauth2_binding"

// UserOauth2Binding is a mapping object for user_oauth2_binding table in mysql
type UserOauth2Binding struct {
	UserID      int `ddb:"user_id" json:"user_id"`
	OauthUserID int `ddb:"oauth_user_id" json:"oauth_user_id"`
}

func (*UserOauth2Binding) TableName() string {
	return UserOauth2BindingTableName
}

