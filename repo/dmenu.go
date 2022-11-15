package repo

import (
	"mynozom/model"

	"github.com/jinzhu/gorm"
)

type DmenuRepo struct {
	dbs []*gorm.DB
}

func NewDmenuRepo(dbs []*gorm.DB) DmenuRepo {
	return DmenuRepo{
		dbs: dbs,
	}
}

func (dr *DmenuRepo) MainGroupList(dbIndex uint) (*[]model.MainGroup, error) {
	var resp []model.MainGroup
	rows, err := dr.dbs[dbIndex].Raw("EXEC GroupTypeList").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var group model.MainGroup
		if err := rows.Scan(&group.GroupCode, &group.GroupName, &group.Icon); err != nil {
			return nil, err
		}

		resp = append(resp, group)
	}
	return &resp, nil
}

func (dr *DmenuRepo) SubGroupList(id int, dbIndex uint) (*[]model.SubGroup, error) {
	var resp []model.SubGroup
	rows, err := dr.dbs[dbIndex].Raw("EXEC GroupCodeListByGroupTypeId @GroupTypeID = ?", id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var group model.SubGroup
		if err := rows.Scan(&group.GroupCode, &group.GroupName, &group.ImagePath); err != nil {
			return nil, err
		}
		resp = append(resp, group)
	}
	return &resp, nil
}

func (dr *DmenuRepo) ItemsList(group int, table int, dbIndex uint) (*[]model.DmenuItem, error) {
	var resp []model.DmenuItem
	rows, err := dr.dbs[dbIndex].Raw("EXEC StkMs01ListByMenuAndGroup 	@GroupCode = ? , @TableSerial = ?", group, table).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.DmenuItem
		if err := rows.Scan(&item.ItemSerial, &item.ItemPrice, &item.ImagePath, &item.ItemCode, &item.ItemName, &item.ItemDesc, &item.WithModifier, &item.Qnt); err != nil {
			return nil, err
		}
		resp = append(resp, item)
	}
	return &resp, nil
}

func (dr *DmenuRepo) UpdateImage(req *model.UpdateImageReq, dbIndex uint) (*bool, error) {
	var resp bool
	err := dr.dbs[dbIndex].Raw("EXEC UpdateImage @mainGroup = ? , @subGroup = ?  ,@product = ? , @image = ? ", req.MainGroupCode, req.SubGroupCode, req.ProductCode, req.Image).Row().Scan(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
