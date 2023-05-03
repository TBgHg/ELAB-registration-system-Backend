package user

//
//func NewSpaceNotFoundError() model.ErrorResponse {
//	return model.ErrorResponse{
//		Error:   "space_not_found",
//		Message: "空间不存在",
//	}
//}
//
//func getUserSpaces(ctx *gin.Context) {
//	spaces, err := member.GetUserSpaces(ctx)
//	if err != nil {
//		if err == gorm.ErrRecordNotFound {
//			ctx.JSON(404, NewSpaceNotFoundError())
//			return
//		}
//		slog.Error("user.GetUserSpaces failed %w", err)
//		ctx.JSON(500, model.NewInternalServerError())
//		return
//	}
//	if len(*spaces) == 0 {
//		ctx.JSON(404, NewSpaceNotFoundError())
//		return
//	}
//	ctx.JSON(200, user.SpaceListResponse{
//		Spaces: *spaces,
//	})
//}
