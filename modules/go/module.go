package module

import (
	"encoding/json"
	"log"
	"strings"

	model "github.com/bifrostcloud/protoc-gen-httpclient/modules/proto"
	pb "github.com/bifrostcloud/protoc-gen-httpclient/proto"
	"github.com/fatih/camelcase"
	proto "github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/palantir/stacktrace"
)

func (p *plugin) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	target := p.Parameters().Str("target")
	if len(target) == 0 || strings.ToLower(target) == "go" {
		for _, file := range targets {
			name := p.Context.OutputPath(file).SetExt(".httpclient.go").String()
			b := model.Base{Package: p.Context.PackageName(file).String()}
			imports := map[string]model.Package{
				"net": model.Package{
					PackagePath: "net/http",
				},
				"stacktrace": model.Package{
					PackageName: "stacktrace",
					PackagePath: "github.com/palantir/stacktrace",
				},
			}
			for _, srv := range file.Services() {
				opt := srv.Descriptor().GetOptions()
				option, err := proto.GetExtension(opt, pb.E_ServiceOptions)
				if err != nil {
					if err == proto.ErrMissingExtension {
						continue
					}
					// log.Fatal(stacktrace.NewError(err.Error()))
				}
				byteData, err := json.Marshal(option)
				if err != nil {
					log.Fatal(stacktrace.NewError(err.Error()))
				}
				srvOpts := pb.ServiceOptions{}
				err = json.Unmarshal(byteData, &srvOpts)
				if err != nil {
					log.Fatal(stacktrace.NewError(err.Error()))
				}
				s := model.Service{}
				s.UpperCamelCaseServiceName = srv.Name().UpperCamelCase().String()
				s.LowerCamelCaseServiceName = srv.Name().LowerCamelCase().String()
				s.Auth = strings.ToLower(srvOpts.Auth)
				if len(s.Auth) == 0 {
					s.Auth = "basic"
				}
				for _, method := range srv.Methods() {

					opt := method.Descriptor().GetOptions()
					option, err := proto.GetExtension(opt, pb.E_RequestOptions)
					if err != nil {
						if err == proto.ErrMissingExtension {
							continue
						}
						// log.Fatal(stacktrace.NewError(err.Error()))
					}
					byteData, err := json.Marshal(option)
					if err != nil {
						log.Fatal(stacktrace.NewError(err.Error()))
					}
					clientOpts := pb.RequestOptions{}
					err = json.Unmarshal(byteData, &clientOpts)
					if err != nil {
						log.Fatal(stacktrace.NewError(err.Error()))
					}
					clientOpts.Target = strings.ToLower(srvOpts.Endpoint + clientOpts.Target)
					if len(clientOpts.ClientType) > 0 {
						clientOpts.ClientType = strings.ToLower(clientOpts.ClientType)
					} else {
						clientOpts.ClientType = "basic"
					}
					clientOpts.Method = strings.ToUpper(clientOpts.Method)
					if clientOpts.Param {
						imports["reflect"] = model.Package{
							PackagePath: "reflect",
						}
						imports["fmt"] = model.Package{
							PackagePath: "fmt",
						}

					}
					if clientOpts.Method == "POST" || clientOpts.Method == "PUT" {
						imports["io"] = model.Package{
							PackagePath: "io",
						}
						imports["ioutil"] = model.Package{
							PackagePath: "io/ioutil",
						}
						imports["url"] = model.Package{
							PackagePath: "net/url",
						}
						imports["strings"] = model.Package{
							PackagePath: "strings",
						}
					}

					imports["utils"] = model.Package{
						PackageName: "utils",
						PackagePath: "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/utils",
					}
					imports["json"] = model.Package{
						PackagePath: "encoding/json",
					}

					if clientOpts.ClientType == "circuit-breaker" {
						imports["circuit-breaker"] = model.Package{
							PackageName: "cb",
							PackagePath: "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/circuit-breaker",
						}
					} else if clientOpts.ClientType == "basic" {
						imports["basic"] = model.Package{
							PackageName: "basic",
							PackagePath: "github.com/bifrostcloud/protoc-gen-httpclient/pkg/go/client/basic",
						}
					}
					ms := p.Context.Name(method).UpperCamelCase().String()
					splitted := camelcase.Split(ms)
					firstElem := strings.ToLower(splitted[0])
					if firstElem == "put" || firstElem == "post" || firstElem == "get" || firstElem == "delete" {
						splitted[0] = ""
					}
					lastElem := strings.ToLower(splitted[len(splitted)-1])
					if lastElem == "put" || lastElem == "post" || lastElem == "get" || lastElem == "delete" {
						splitted[len(splitted)-1] = ""
					}
					ms = strings.Join(splitted, "")
					upperCamelCaseMethodName := pgs.Name(strings.ToLower(clientOpts.Method)).UpperCamelCase().String() + ms

					r := model.Method{}
					r.UpperCamelCaseServiceName = srv.Name().UpperCamelCase().String()
					r.UpperCamelCaseMethodName = upperCamelCaseMethodName

					r.InputType = p.Context.Name(method.Input()).String()

					if !method.Input().BuildTarget() {
						path := p.Context.ImportPath(method.Input()).String()
						imports[path] = model.Package{
							PackagePath: path,
						}

						r.InputType = p.Context.PackageName(method.Input()).String() + "." + p.Context.Name(method.Input()).String()
					}
					inputFields := method.Input().Fields()
					for _, field := range inputFields {
						r.InputFields.Type = r.InputType
						r.InputFields.FieldImport = append(r.InputFields.FieldImport, model.FieldImport{
							Name: field.Name().UpperCamelCase().String(),
							Tag:  field.Name().String(),
						})
						r.InputFields.Base = append(r.InputFields.Base, strings.ToLower(field.Name().String()))
						r.InputFields.Lowercase = append(r.InputFields.Lowercase, strings.ToLower(field.Name().LowerCamelCase().String()))
						r.InputFields.DotNotation = append(r.InputFields.DotNotation, strings.ToLower(field.Name().LowerDotNotation().String()))
						spl := camelcase.Split(field.Name().UpperCamelCase().String())
						r.InputFields.ParamCase = append(r.InputFields.ParamCase, strings.ToLower(strings.Join(spl, "-")))
					}
					r.OutputType = p.Context.Name(method.Output()).String()
					r.Auth = strings.ToLower(srvOpts.Auth)
					r.RequestOptions = clientOpts
					if !method.Output().BuildTarget() {
						path := p.Context.ImportPath(method.Output()).String()
						imports[path] = model.Package{
							PackagePath: path,
						}

						r.OutputType = p.Context.PackageName(method.Output()).String() + "." + p.Context.Name(method.Output()).String()
					}

					s.Methods = append(s.Methods, r)
				}
				b.Services = append(b.Services, s)
			}

			if len(b.Services) == 0 {
				continue
			}

			for _, pkg := range imports {
				b.Imports = append(b.Imports, pkg)
			}

			p.OverwriteGeneratorTemplateFile(
				name,
				template.Lookup("Base_go"),
				&b,
			)

		}
	}

	return p.Artifacts()
}
