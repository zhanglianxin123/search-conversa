package gpt

type Conversa struct {
	ID      string `gorm:"column:id" db:"id" json:"id" form:"id"`                     //  主键id
	Role    string `gorm:"column:role" db:"role" json:"role" form:"role"`             //  对话人
	Content string `gorm:"column:content" db:"content" json:"content" form:"content"` //  对话人
}

func (Conversa) TableName() string {
	return "conversa"
}
