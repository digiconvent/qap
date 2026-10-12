# qap-middleware

Fiber v3 middleware that guards routes with [QAP](../README.md). A guarded request either runs straight away, or becomes a proposal that other eligible users vote on before it is carried out.

```sh
go get github.com/digiconvent/qap/middleware
```

The package name is `qap`. The module is `github.com/digiconvent/qap`, at the root of the repository, and every release is a tag like `v1.0.0`.

The documentation is the [middleware plan](../docs/middleware/index.md): one page per name in the code, from what a person does down to the atomic parts.
