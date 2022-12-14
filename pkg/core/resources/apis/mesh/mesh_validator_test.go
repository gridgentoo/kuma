package mesh

import (
	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	util_proto "github.com/Kong/kuma/pkg/util/proto"
)

var _ = Describe("Mesh", func() {
	type testCase struct {
		mesh     string
		expected string
	}
	Describe("Validate()", func() {
		It("should pass validation", func() {
			// given
			spec := `
            mtls:
              enabledBackend: builtin-1
              backends:
              - name: builtin-1
                type: builtin
            logging:
              backends:
              - name: file-1
                type: file
                config:
                  path: /path/to/file
              - name: file-2
                format: '%START_TIME% %KUMA_SOURCE_SERVICE%'
                type: file
                config:
                  path: /path/to/file2
              - name: tcp-1
                type: tcp
                config:
                  address: kibana:1234
              - name: tcp-2
                format: '%START_TIME% %KUMA_DESTINATION_SERVICE%'
                type: tcp
                config:
                  address: kibana:1234
              defaultBackend: tcp-1
            tracing:
              backends:
              - name: zipkin-us
                sampling: 80.0
                type: zipkin
                config:
                  url: http://zipkin.local:9411/v2/spans
                  traceId128bit: true
                  apiVersion: httpProto
              - name: zipkin-eu
                type: zipkin
                config:
                  url: http://zipkin.local:9411/v2/spans
              defaultBackend: zipkin-us
            metrics:
              enabledBackend: prom-1
              backends:
              - name: prom-1
                type: prometheus
                config:
                  port: 5670
                  path: /metrics
`
			mesh := MeshResource{}

			// when
			err := util_proto.FromYAML([]byte(spec), &mesh.Spec)
			// then
			Expect(err).ToNot(HaveOccurred())

			// when
			err = mesh.Validate()

			// then
			Expect(err).ToNot(HaveOccurred())
		})

		DescribeTable("should validate fields",
			func(given testCase) {
				// given
				mesh := MeshResource{}

				// when
				err := util_proto.FromYAML([]byte(given.mesh), &mesh.Spec)
				// then
				Expect(err).ToNot(HaveOccurred())

				// when
				verr := mesh.Validate()
				// and
				actual, err := yaml.Marshal(verr)

				// then
				Expect(err).ToNot(HaveOccurred())
				// and
				Expect(actual).To(MatchYAML(given.expected))
			},
			Entry("multiple ca backends of the same name", testCase{
				mesh: `
                mtls:
                  enabledBackend: backend-1
                  backends:
                  - name: backend-1
                    type: builtin
                  - name: backend-1
                    type: builtin`,
				expected: `
                violations:
                - field: mtls.backends[1].name
                  message: '"backend-1" name is already used for another backend'`,
			}),
			Entry("enabledBackend of unknown name", testCase{
				mesh: `
                mtls:
                  enabledBackend: backend-2
                  backends:
                  - name: backend-1
                    type: builtin`,
				expected: `
                violations:
                - field: mtls.enabledBackend
                  message: has to be set to one of the backends in the mesh`,
			}),
			Entry("logging backend with empty name", testCase{
				mesh: `
                logging:
                  backends:
                  - name:
                    type: tcp
                    config: 
                      address: kibana:1234`,
				expected: `
                violations:
                - field: logging.backends[0].name
                  message: cannot be empty`,
			}),
			Entry("multiple logging backends of the same name", testCase{
				mesh: `
                logging:
                  backends:
                  - name: backend-1
                    type: tcp
                    config:
                      address: kibana:1234
                  - name: backend-1
                    type: file
                    config:
                      path: /path/to/file
                  defaultBackend: backend-1`,
				expected: `
                violations:
                - field: logging.backends[1].name
                  message: '"backend-1" name is already used for another backend'`,
			}),
			Entry("tcp logging address is empty", testCase{
				mesh: `
                logging:
                  backends:
                  - name: backend-1
                    type: tcp
                    config:
                      address:
                  defaultBackend: backend-1`,
				expected: `
                violations:
                - field: logging.backends[0].config.address
                  message: cannot be empty`,
			}),
			Entry("tcp logging address is invalid", testCase{
				mesh: `
                logging:
                  backends:
                  - name: backend-1
                    type: tcp
                    config:
                      address: wrong-format:234:234
                  defaultBackend: backend-1`,
				expected: `
                violations:
                - field: logging.backends[0].config.address
                  message: has to be in format of HOST:PORT`,
			}),
			Entry("file logging path is empty", testCase{
				mesh: `
                logging:
                  backends:
                  - name: backend-1
                    type: file
                    config:
                      path:
                  defaultBackend: backend-1`,
				expected: `
                violations:
                - field: logging.backends[0].config.path
                  message: cannot be empty`,
			}),
			Entry("invalid access log format", testCase{
				mesh: `
                logging:
                  backends:
                  - name: backend-1
                    format: "%START_TIME% %sent_bytes%"
                    type: file
                    config:
                      path: /var/logs
                  defaultBackend: backend-1`,
				expected: `
                violations:
                - field: logging.backends[0].format
                  message: 'format string is not valid: expected a command operator to start at position 14, instead got: "%sent_bytes%"'`,
			}),
			Entry("default backend has to be set to one of the backends", testCase{
				mesh: `
                logging:
                  backends:
                  - name: backend-1
                    type: file
                    config:
                      path: /path
                  defaultBackend: non-existing-backend`,
				expected: `
                violations:
                - field: logging.defaultBackend
                  message: has to be set to one of the logging backend in mesh`,
			}),
			Entry("tracing backend with empty name", testCase{
				mesh: `
                tracing:
                  backends:
                  - name:
                    type: zipkin
                    config:
                      url: http://zipkin.local:9411/v2/spans`,
				expected: `
                violations:
                - field: tracing.backends[0].name
                  message: cannot be empty`,
			}),
			Entry("multiple tracing backend with the same name", testCase{
				mesh: `
                tracing:
                  backends:
                  - name: zipkin-us
                    type: zipkin
                    config:
                      url: http://zipkin.local:9411/v2/spans
                  - name: zipkin-us
                    type: zipkin
                    config:
                      url: http://zipkin.local:9411/v2/spans`,
				expected: `
                violations:
                - field: tracing.backends[1].name
                  message: '"zipkin-us" name is already used for another backend'`,
			}),
			Entry("tracing with invalid sampling", testCase{
				mesh: `
                tracing:
                  backends:
                  - name: zipkin-us
                    sampling: 100.1
                    type: zipkin
                    config:
                      url: http://zipkin-us.local:9411/v2/spans`,
				expected: `
                violations:
                - field: tracing.backends[0].sampling
                  message: has to be in [0.0 - 100.0] range`,
			}),
			Entry("tracing with zipkin without url", testCase{
				mesh: `
                tracing:
                  backends:
                  - name: zipkin-us
                    type: zipkin
                    config:
                      url: ""`,
				expected: `
                violations:
                - field: tracing.backends[0].config.url
                  message: cannot be empty`,
			}),
			Entry("tracing with zipkin with invalid url", testCase{
				mesh: `
                tracing:
                  backends:
                  - name: zipkin-us
                    type: zipkin
                    config:
                      url: invalid-url`,
				expected: `
                violations:
                - field: tracing.backends[0].config.url
                  message: invalid URL`,
			}),
			Entry("tracing with zipkin with valid url but without port", testCase{
				mesh: `
                tracing:
                  backends:
                  - name: zipkin-us
                    type: zipkin
                    config:
                      url: http://zipkin-us.local/v2/spans`,
				expected: `
                violations:
                - field: tracing.backends[0].config.url
                  message: port has to be explicitly specified`,
			}),
			Entry("tracing with zipkin with invalid apiVersion", testCase{
				mesh: `
                tracing:
                  backends:
                  - name: zipkin-us
                    type: zipkin
                    config:
                      url: http://zipkin-us.local:9411/v2/spans
                      apiVersion: invalid`,
				expected: `
                violations:
                - field: tracing.backends[0].config.apiVersion
                  message: 'has invalid value. Allowed values: httpJsonV1, httpJson, httpProto'`,
			}),
			Entry("default backend has to be set to one of the backends", testCase{
				mesh: `
                tracing:
                  defaultBackend: non-existent
                  backends:
                  - name: zipkin-us
                    type: zipkin
                    config:
                      url: http://zipkin.local:9411/v2/spans`,
				expected: `
                violations:
                - field: tracing.defaultBackend
                  message: has to be set to one of the tracing backend in mesh`,
			}),
			Entry("multiple metrics backends of the same name", testCase{
				mesh: `
                metrics:
                  enabledBackend: backend-1
                  backends:
                  - name: backend-1
                    type: prometheus
                  - name: backend-1
                    type: prometheus`,
				expected: `
                violations:
                - field: metrics.backends[1].name
                  message: '"backend-1" name is already used for another backend'`,
			}),
			Entry("enabledBackend of unknown name", testCase{
				mesh: `
                metrics:
                  enabledBackend: backend-2
                  backends:
                  - name: backend-1
                    type: prometheus`,
				expected: `
                violations:
                - field: metrics.enabledBackend
                  message: has to be set to one of the backends in the mesh`,
			}),
			Entry("unknown backend types", testCase{
				mesh: `
                metrics:
                  backends:
                  - name: backend-1
                    type: xxx
                logging:
                  backends:
                  - name: backend-1
                    type: xxx
                tracing:
                  backends:
                  - name: backend-1
                    type: xxx
`,
				expected: `violations:
                - field: logging.backends[0].type
                  message: 'unknown backend type. Available backends: "tcp", "file"'
                - field: tracing.backends[0].type
                  message: 'unknown backend type. Available backends: "zipkin"'
                - field: metrics.backends[0].type
                  message: 'unknown backend type. Available backends: "prometheus"'`,
			}),
			Entry("multiple errors", testCase{
				mesh: `
                mtls:
                  enabledBackend: invalid-backend
                logging:
                  backends:
                  - name:
                    type: file
                    config:
                      path: /path
                  - name: tcp-1
                    type: file
                    config:
                      address: invalid-address
                  - name: tcp-1
                    type: tcp
                    path:
                      address:
                  defaultBackend: invalid-backend`,
				expected: `
                violations:
                - field: mtls.enabledBackend
                  message: has to be set to one of the backends in the mesh
                - field: logging.backends[0].name
                  message: cannot be empty
                - field: logging.backends[1].config.path
                  message: cannot be empty
                - field: logging.backends[2].config.address
                  message: cannot be empty
                - field: logging.backends[2].name
                  message: '"tcp-1" name is already used for another backend'
                - field: logging.defaultBackend
                  message: has to be set to one of the logging backend in mesh`,
			}),
		)
	})
})
