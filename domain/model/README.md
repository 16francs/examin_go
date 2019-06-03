# domain/model

DB に受け渡しをするデータを構造体として定義

## テンプレート

```go
package model

// Sample - Model のサンプル
type Sample struct {
	Message string `json:"message"`
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
