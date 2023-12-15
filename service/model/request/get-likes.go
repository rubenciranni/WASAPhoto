package request

type GetLikesRequest struct {
	PathParameters struct {
		PhotoID string `key:"photoID"`
	}
	QueryParameters struct {
		StartID string `key:"startID"`
	}
}

func (request *GetLikesRequest) IsValid() bool {
	return isValidUUID(request.PathParameters.PhotoID) &&
		(request.QueryParameters.StartID == "" || isValidUUID(request.QueryParameters.StartID))
}
