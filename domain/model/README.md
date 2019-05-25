# domain/model

DB に受け渡しをするデータを構造体として定義

```
// User - ユーザーモデル
type User struct {
  Base
	LoginID   string `json:"login_id"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	School    string `json:"school"`
	Role      int    `json:"role"`
	Activated bool   `json:"activated"`
}
```

ディレクトリの分け方

```
model
├── students
│   └── user
└── teachers
    └── user
```

e.g.) model/students/user

```
package model

type User struce {
	...
}
```
