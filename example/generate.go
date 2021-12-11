package example

//go:generate go run -mod=mod github.com/withtally/synceth compile --outdir ./artifacts .
//go:generate go run -mod=mod github.com/withtally/synceth bind --handlers --fakes --outdir ./bindings ./artifacts
