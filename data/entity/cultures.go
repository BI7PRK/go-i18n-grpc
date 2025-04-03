// entity/cultures.go
package entity

// cultures
type CulturesResources struct {
	ID        int32  `xorm:"pk autoincr 'id'" json:"id"`     //主键ID
	Name      string `xorm:"varchar(50) 'name'" json:"name"` //名称
	Code      string `xorm:"varchar(10) 'code'" json:"code"` //代码
	IsDefault bool   `xorm:"'is_default'" json:"is_default"` //是否默认
}

type CulturesResourceTypes struct {
	ID     int32  `xorm:"pk autoincr 'id'" json:"id"`          //主键ID
	Name   string `xorm:"varchar(50) 'name'" json:"name"`      //名称
	Remark string `xorm:"varchar(255) 'remark'" json:"remark"` //备注
}

type CulturesResourceKeys struct {
	ID     int32  `xorm:"pk autoincr 'id'" json:"id"`      //主键ID
	Name   string `xorm:"varchar(100) 'name'" json:"name"` //名称
	TypeID int32  `xorm:"'type_id'" json:"type_id"`        //类型ID
}

type CulturesResourceLangs struct {
	ID        int64  `xorm:"pk autoincr 'id'" json:"id"`      //主键ID
	KeyID     int32  `xorm:"'key_id'" json:"key_id"`          //keyID
	CultureID int32  `xorm:"'culture_id'" json:"culture_id"`  //cultureID
	Text      string `xorm:"varchar(500) 'text'" json:"text"` //文本
}
