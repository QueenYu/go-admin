package model

type Menu struct {
	BaseModel
	Name      string `json:"title" gorm:"type:varchar(20);not null;unique;comment:菜单名称"`             //菜单名称
	Icon      string `json:"icon" gorm:"type:varchar(20);not null;comment:图标"`                       //图表
	RoutePath string `json:"path" gorm:"column:route_path;type:varchar(20);not null;comment:路由path"` //路由path
	RouteName string `json:"name" gorm:"column:route_name;type:varchar(20);not null;comment:路由名称"`   //路由名称
	Hidden    bool   `json:"hidden" gorm:"not null;default:0;comment:是否在列表隐藏"`                       //是否在列表隐藏
	Component string `json:"component" gorm:"type:varchar(40);not null;comment:前端文件路径"`              //前端文件路径
	Sort      int    `json:"sort" gorm:"not null;comment:排序"`                                        //排序
	ParentId  int    `json:"parent_id" gorm:"column:parent_id;not null;default:0;comment:父菜单ID"`     //父菜单id
	Children  []Menu `json:"children" gorm:"foreignKey:ParentId"`
}