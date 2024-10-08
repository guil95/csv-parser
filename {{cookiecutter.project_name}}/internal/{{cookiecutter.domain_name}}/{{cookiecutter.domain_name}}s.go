package {{cookiecutter.domain_name}}

type {{cookiecutter.domain_struct}} struct {
	Name string `json:"name" db:"name" validate:"required"`
}
