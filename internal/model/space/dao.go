package space

import "gorm.io/gen"

type Querier interface {
	// Exists 判断空间是否存在
	//
	// EXISTS(SELECT * FROM @@table WHERE id = @id)
	Exists(id string) (bool, error)
	// SearchByName 根据名称搜索空间
	// 需要注意，这个方法应该返回两种空间：一种为公开空间，一种为用户所在的空间。
	//
	// SELECT * FROM @@table WHERE name LIKE @name
	// UNION

}

type ContentQuerier interface {
	// GetByQuery 根据查询条件获取内容（包括分页和分页限制），此时应该返回多个结果。
	//
	// SELECT * FROM @@table
	// {{where}}
	//   space_id=@query.SpaceId
	//	 type=@query.Type
	//   {{if @query.OrderBy == "time_desc"}}
	//     ORDER BY last_update_at DESC
	//   {{else if @query.OrderBy == "like_desc"}}
	//     ORDER BY JSON_EXTRACT(meta, "$.likes") DESC
	//   {{end}}
	//   {{if @query.Name != ""}}
	//     name LIKE @query.Name
	//   {{end}}
	// {{end}}
	// LIMIT @query.Limit OFFSET @query.Offset
	GetByQuery(query *ContentQuery) ([]*gen.T, error)
	// GetById 根据Id精确查询
	//
	// SELECT * FROM @@table
	// {{where}}
	//   id=@id
	//   space_id=@spaceId
	//   type=@type
	// {{end}}
	// LIMIT 1
	GetById(id string, spaceId string, contentType ContentType) (*gen.T, error)
}

type ContentHistoryQuerier interface {
	// GetById 根据内容的唯一标识符获取内容历史
	//
	// SELECT * FROM @@table
	// {{where}}
	//   space_id=@spaceId
	//   content_id=@contentId
	//   type=@type
	// {{end}}
	GetById(spaceId string, contentId string, contentType ContentType) ([]*gen.T, error)
	// GetLatestById 根据内容的唯一标识符获取最新的一份内容历史
	//
	// SELECT * FROM @@table
	// {{where}}
	//   space_id=@spaceId
	//   content_id=@contentId
	//   type=@type
	// {{end}}
	// ORDER BY time DESC
	// LIMIT 1
	GetLatestById(spaceId string, contentId string, contentType ContentType) (*gen.T, error)
}
