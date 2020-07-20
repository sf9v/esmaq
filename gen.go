package esmaq

import (
	"fmt"
	"io"
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

const pkgPath = "github.com/stevenferrer/esmaq"

// Schema is the state machine schema
type Schema struct {
	// Name is the of the state machine
	Name,
	// Pkg is package of the state machine
	Pkg string
	// States is the states config
	States []StateConfig
}

// Gen generates the state machine
func Gen(cfg Schema, out io.Writer) error {
	name := "StateMachine"
	if len(cfg.Name) > 0 {
		name = cfg.Name
	}

	name = strcase.ToCamel(name)

	pkg := "main"
	if len(cfg.Pkg) > 0 {
		pkg = cfg.Pkg
	}

	f := jen.NewFile(pkg)
	f.PackageComment("Code generated by esmaq, DO NOT EDIT.")

	rcvr := "sm"
	rcvrType := "*" + name

	f.Type().Id("State").Qual(pkgPath, "State")
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, stateCfg := range cfg.States {
			s := string(stateCfg.From)
			sName := strcase.ToCamel(fmt.Sprintf("state_%s", s))
			g.Id(sName).Id("State").Op("=").Lit(s)
		}
	})

	f.Line()
	f.Type().Id("Event").Qual(pkgPath, "Event")
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, stateCfg := range cfg.States {
			for _, trsn := range stateCfg.Transitions {
				e := string(trsn.Event)
				eName := strcase.ToCamel(fmt.Sprintf("event_%s", trsn.Event))
				g.Id(eName).Id("Event").Op("=").Lit(e)
			}
		}
	})

	cbcFields := []jen.Code{}
	methods := []jen.Code{}

	for _, stateCfg := range cfg.States {
		for _, trsn := range stateCfg.Transitions {
			fnName := strcase.ToCamel(string(trsn.Event))
			// cbName := strcase.ToCamel(fmt.Sprintf("%s_cb", fnName))
			// TODO: separate method and callback parameters

			// input args
			ins := []jen.Code{jen.Id("ctx").Qual("context", "Context")}
			// input arg identifiers
			inIDs := []jen.Code{jen.Id("ctx")}
			for id, v := range trsn.Callback.Ins {
				ins = append(ins, getArg(id, reflect.TypeOf(v)))
				inIDs = append(inIDs, jen.Id(id))
			}

			// output args
			outs := []jen.Code{}
			// output arg identifiers
			outIDs := []jen.Code{}
			// return params when error happened
			errRets := []jen.Code{}

			for id, v := range trsn.Callback.Outs {
				t := reflect.TypeOf(v)
				outs = append(outs, getArg(id, t))
				outIDs = append(outIDs, jen.Id(id))
				errRets = append(errRets, getZeroVal(t))
			}

			// return params when no error happened
			okRets := append(outIDs, jen.Nil())
			outs = append(outs, jen.Id("err").Error())
			outIDs = append(outIDs, jen.Id("err"))

			cbName := fnName
			cbcFields = append(cbcFields, jen.Id(cbName).Func().
				Params(ins...).Params(outs...))

			methods = append(methods, jen.Func().
				Params(jen.Id(rcvr).Id(rcvrType)).
				Id(fnName).
				Params(ins...).
				Params(outs...).
				BlockFunc(func(g *jen.Group) {
					jen.Id("next").Qual(pkgPath, "State").
						Op("=").Lit(string(trsn.To))

					// get from in context
					g.List(jen.Id("from"), jen.Id("ok")).Op(":=").
						Id("fromCtx").Call(jen.Id("ctx"))

					g.If(jen.Op("!").Id("ok")).
						BlockFunc(func(g *jen.Group) {
							rets := append(errRets, jen.Qual("errors", "New").
								Call(jen.Lit(`"from" state not set in context`)))
							g.Return(rets...)
						}).Line()

					g.Id("err").Op("=").Id(rcvr).
						Dot("core").
						Dot("Fire").
						Call(
							jen.Qual(pkgPath, "Event").Op("(").Id(strcase.ToCamel("event_"+string(trsn.Event))).Op(")"),
							jen.Qual(pkgPath, "State").Op("(").Id("from").Op(")"),
						)

					g.If(jen.Err().Op("!=").Nil()).
						BlockFunc(func(g *jen.Group) {
							rets := append(errRets, jen.Id("err"))
							g.Return(rets...)
						}).Line()

					g.Id("ctx").Op("=").Id("ctxWtTo").Call(
						jen.Id("ctx"),
						jen.Id(strcase.ToCamel("state_"+string(trsn.To)))).
						Line()

					g.List(outIDs...).Op("=").
						Id(rcvr).
						Dot("cbs").
						Dot(cbName).
						Call(inIDs...)

					g.If(jen.Err().Op("!=").Nil()).
						BlockFunc(func(g *jen.Group) {
							rets := append(errRets, jen.Id("err"))
							g.Return(rets...)
						})

					g.Line().Return(okRets...)
				}))

		}
	}

	f.Line()

	// state machine type def
	f.Type().Id(name).
		Struct(
			jen.Id("core").Op("*").Qual(pkgPath, "Core"),
			jen.Id("cbs").Id("*Callbacks"),
		)

	f.Line()

	// callbacks type def
	f.Type().Id("Callbacks").Struct(cbcFields...)

	for _, c := range methods {
		f.Line()
		f.Add(c)
	}

	f.Line()

	f.Func().Id("New" + strcase.ToCamel(name)).
		Params(jen.Id("cbs").Id("*Callbacks")).
		Params(jen.Id("*" + name)).
		BlockFunc(func(g *jen.Group) {

			g.Id("stateConfigs").Op(":=").Op("[]").
				Qual(pkgPath, "StateConfig").
				BlockFunc(func(g *jen.Group) {

					for _, state := range cfg.States {
						g.BlockFunc(func(g *jen.Group) {
							g.Id("From").Op(":").
								Qual(pkgPath, "State").Op("(").
								Id(strcase.ToCamel("state_" + string(state.From))).
								Op(")").
								Op(",")
							g.Id("Transitions").Op(":").Op("[]").
								Qual(pkgPath, "Transitions").
								BlockFunc(func(g *jen.Group) {
									for _, trsn := range state.Transitions {
										g.BlockFunc(func(g *jen.Group) {
											g.Id("Event").Op(":").
												Qual(pkgPath, "Event").Op("(").
												Id(strcase.ToCamel("event_" + string(trsn.Event))).
												Op(")").
												Op(",")
											g.Id("To").Op(":").
												Qual(pkgPath, "State").Op("(").
												Id(strcase.ToCamel("state_" + string(trsn.To))).
												Op(")").
												Op(",")
										}).Op(",")
									}
								}).Op(",")
						}).Op(",")
					}
				}).Line()

			smVar := strcase.ToLowerCamel(name)
			g.Id(smVar).
				Op(":=").Op("&").Id(name).
				BlockFunc(func(g *jen.Group) {
					g.Id("cbs").Op(":").Id("cbs").Op(",")
					g.Id("core").Op(":").
						Qual(pkgPath, "NewCore").
						Params(jen.Id("stateConfigs")).Op(",")
				}).Line()
			g.Return().Id(smVar)
		})

		// context key types
	f.Type().Id("ctxKey").
		Int()
	f.Const().DefsFunc(func(g *jen.Group) {
		g.Id("fromKey").Id("ctxKey").Op("=").Id("iota")
		g.Id("toKey")
	})

	f.Func().Id("CtxWtFrom").
		Params(jen.Id("ctx").Qual("context", "Context"),
			jen.Id("from").Id("State"),
		).
		Params(jen.Qual("context", "Context")).
		BlockFunc(func(g *jen.Group) {
			g.Return(jen.Qual("context", "WithValue").
				Call(jen.Id("ctx"), jen.Id("fromKey"), jen.Id("from")))
		}).Line()

	f.Func().Id("ctxWtTo").
		Params(jen.Id("ctx").Qual("context", "Context"),
			jen.Id("to").Id("State"),
		).
		Params(jen.Qual("context", "Context")).
		BlockFunc(func(g *jen.Group) {
			g.Return(jen.Qual("context", "WithValue").
				Call(jen.Id("ctx"), jen.Id("toKey"), jen.Id("to")))
		}).Line()

	f.Func().Id("fromCtx").
		Params(jen.Id("ctx").Qual("context", "Context")).
		Params(jen.Id("State"), jen.Bool()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("from"), jen.Id("ok")).
				Op(":=").
				Id("ctx").
				Dot("Value").
				Call(jen.Id("fromKey")).
				Assert(jen.Id("State"))

			g.Return(jen.Id("from"), jen.Id("ok"))
		}).Line()

	f.Func().Id("ToCtx").
		Params(jen.Id("ctx").Qual("context", "Context")).
		Params(jen.Id("State"), jen.Bool()).
		BlockFunc(func(g *jen.Group) {
			g.List(jen.Id("to"), jen.Id("ok")).
				Op(":=").
				Id("ctx").
				Dot("Value").
				Call(jen.Id("toKey")).
				Assert(jen.Id("State"))

			g.Return(jen.Id("to"), jen.Id("ok"))
		}).Line()

	return f.Render(out)
}

func getArg(id string, t reflect.Type) jen.Code {
	c := jen.Id(id)

	// built-in types
	if t.Name() != "" {
		switch t.Name() {
		case "int":
			return c.Int()
		case "int32":
			return c.Int32()
		case "int64":
			return c.Int64()
		case "uint":
			return c.Uint()
		case "uint32":
			return c.Uint32()
		case "uint64":
			return c.Uint64()
		case "float32":
			return c.Float32()
		case "float64":
			return c.Float32()
		case "string":
			return c.String()
		case "error":
			return c.Error()
		}
	}

	return c.Qual(t.PkgPath(), t.Name())
}

func getZeroVal(t reflect.Type) jen.Code {
	// built-in types
	if t.Name() != "" {
		switch t.Name() {
		case "int":
			return jen.Lit(0)
		case "int32":
			return jen.Lit(0)
		case "int64":
			return jen.Lit(0)
		case "uint":
			return jen.Lit(0)
		case "uint32":
			return jen.Lit(0)
		case "uint64":
			return jen.Lit(0)
		case "float32":
			return jen.Lit(0)
		case "float64":
			return jen.Lit(0)
		case "string":
			return jen.Lit("")
		}
	}

	return jen.Nil()
}
