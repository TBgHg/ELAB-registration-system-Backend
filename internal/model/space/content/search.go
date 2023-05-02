package content

// GET /space/:space_id/<content>

type Query struct {
	Name     string    `form:"name"`
	SpaceId  string    `form:"space_id"`
	Page     int       `form:"page"`
	PageSize int       `form:"page_size"`
	OrderBy  OrderType `form:"order_by"`
}
