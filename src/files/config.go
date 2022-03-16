package files

import "fmt"

const EditorConfig = `
# EditorConfig is awesome: https://EditorConfig.org

# top-most EditorConfig file
root = true

[*]
indent_style = space
indent_size = 4
end_of_line = lf
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true

`

func Gitignore(projectName *string) string {
	return fmt.Sprintf(`
%s
`, *projectName)
}
