package definitions

import (
	"github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	"github.com/Kong/kuma/pkg/core/resources/model"
)

var ProxyTemplateWsDefinition = ResourceWsDefinition{
	Name: "ProxyTemplate",
	Path: "proxytemplates",
	ResourceFactory: func() model.Resource {
		return &mesh.ProxyTemplateResource{}
	},
	ResourceListFactory: func() model.ResourceList {
		return &mesh.ProxyTemplateResourceList{}
	},
}
