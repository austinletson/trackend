package world

// Dot represents a dot file
type Dot struct {
	Id          string `json:"id" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	FilePath    string `json:"filePath" binding:"required"`
}

// DotData represents the functionality required to minipulate an instance of data representing a dot file
type DotOp interface {
	insertDot(dot Dot)
	getUserDots(user User) []Dot
}
