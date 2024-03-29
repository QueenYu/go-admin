package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Dlimingliang/go-admin/core/business"
	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
)

type MenuService struct{}

func (menuService *MenuService) GetMenuTree() ([]model.Menu, error) {
	//获取所有菜单
	var menuList []model.Menu
	err := global.GaDb.Order("sort, id").Find(&menuList).Error
	if err != nil {
		return menuList, err
	}
	//将菜单按照父菜单分组
	menuMap := make(map[int][]model.Menu)
	for _, menu := range menuList {
		menuMap[menu.ParentId] = append(menuMap[menu.ParentId], menu)
	}
	//依次设置子集
	routeMenu := menuMap[0]
	for i := 0; i < len(routeMenu); i++ {
		setChildrenMenu(&routeMenu[i], menuMap)
	}
	return routeMenu, err
}

func setChildrenMenu(menu *model.Menu, menuMap map[int][]model.Menu) {
	menu.Children = menuMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		setChildrenMenu(&menu.Children[i], menuMap)
	}
}

func (menuService MenuService) GetMenuByRole(roleId string) ([]model.Menu, error) {
	var menuList []model.Menu
	err := global.GaDb.Model(&model.Menu{}).Select("menu.*").Joins("left join role_menus on role_menus.menu_id = menu.id").Where("role_menus.role_authority_id = ?", roleId).Order("sort,id").Scan(&menuList).Error
	return menuList, err
}

func (menuService *MenuService) GetMenuById(id int) (model.Menu, error) {
	var menu model.Menu
	err := global.GaDb.First(&menu, id).Error
	return menu, err
}

func (menuService *MenuService) CreateMenu(req model.Menu) (model.Menu, error) {

	var menu model.Menu
	//验证菜单名称重复
	res := global.GaDb.Where("name = ?", req.Meta.Name).First(&menu)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return menu, res.Error
	}
	if res.RowsAffected > 0 {
		return menu, business.New("已存在同名菜单")
	}

	//创建菜单
	err := global.GaDb.Create(&req).Error
	return req, err
}

func (menuService MenuService) AddMenuAuthority(menus []model.Menu, roleId string) error {
	var role model.Role
	role.AuthorityId = roleId
	role.Menus = menus
	err := RoleServiceInstance.SetMenuAuthority(role)
	return err
}

func (menuService *MenuService) UpdateMenu(req model.Menu) error {
	//判断菜单名称是否重复
	var dbMenu model.Menu
	//验证菜单名称重复
	res := global.GaDb.Where("id <> ? AND name = ?", req.ID, req.Meta.Name).First(&dbMenu)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return business.New("已存在同名菜单")
	}

	updateMap := make(map[string]interface{})
	updateMap["parent_id"] = req.ParentId
	updateMap["route_path"] = req.RoutePath
	updateMap["route_name"] = req.RouteName
	updateMap["hidden"] = req.Hidden
	updateMap["component"] = req.Component
	updateMap["sort"] = req.Sort
	updateMap["name"] = req.Name
	updateMap["icon"] = req.Icon
	return global.GaDb.Model(&model.Menu{}).Where("id = ?", req.ID).Updates(updateMap).Error
}

func (menuService *MenuService) DeleteMenu(id int) error {
	res := global.GaDb.Where("parent_id = ?", id).First(&model.Menu{})
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return res.Error
	}
	if res.RowsAffected > 0 {
		return business.New("此菜单存在子菜单不可删除")
	}

	//删除菜单及关联的角色
	var menu model.Menu
	err := global.GaDb.Preload("Roles").Where("id = ?", id).First(&menu).Delete(&menu).Error
	if len(menu.Roles) > 0 {
		err = global.GaDb.Model(&menu).Association("Roles").Delete(&menu.Roles)
	}
	return err
}
