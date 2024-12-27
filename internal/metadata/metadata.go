package metadata

var (
	appName    string = "path-filter-action"
	appDesc    string = "Git Path Filter Action CLI"
	authorName string = "Takumi Takahashi"
)

func AppName() string {
	return appName
}

func AppDesc() string {
	return appDesc
}

func AuthorName() string {
	return authorName
}
