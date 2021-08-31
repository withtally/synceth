package example

//go:generate go run -mod=mod github.com/withtally/ethgen compile --outdir ./artifacts .
//go:generate go run -mod=mod github.com/withtally/ethgen bind --handlers --fakes --outdir ./bindings ./artifacts
