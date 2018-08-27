package cidl

// Token授权站点类型
type AuthSiteType int

const (
	// 运营端
	AuthSiteTypeAdmin AuthSiteType = 1
	// 城市合伙人端
	AuthSiteTypeOrg AuthSiteType = 2
	// 微信小程序端
	AuthSiteTypeWxXcx AuthSiteType = 3
)

func (m AuthSiteType) String() string {
	switch m {

	case AuthSiteTypeAdmin:
		return "AuthSiteTypeAdmin<enum AuthSiteType>"
	case AuthSiteTypeOrg:
		return "AuthSiteTypeOrg<enum AuthSiteType>"
	case AuthSiteTypeWxXcx:
		return "AuthSiteTypeWxXcx<enum AuthSiteType>"
	default:
		return "UNKNOWN_Name_<AuthSiteType>"
	}
}
