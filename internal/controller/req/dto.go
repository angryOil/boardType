package req

type CreateBoardTypeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PatchBoardDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
