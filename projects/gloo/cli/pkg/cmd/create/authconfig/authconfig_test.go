package authconfig_test

import (
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/create/authconfig"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/solo-io/gloo/pkg/cliutil/testutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
	extauthpb "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("AuthConfig", func() {
	BeforeEach(func() {
		helpers.UseMemoryClients()
	})

	getExtension := func() []*extauthpb.AuthConfig_Config {
		ac, err := helpers.MustAuthConfigClient().Read("gloo-system", "ac1", clients.ReadOpts{})
		Expect(err).NotTo(HaveOccurred())
		return ac.Configs
	}

	getOIDCConfig := func() *extauthpb.OAuth {
		return getExtension()[0].GetOauth()
	}

	getApiKeyConfig := func() *extauthpb.ApiKeyAuth {
		return getExtension()[0].GetApiKeyAuth()
	}
	getOpaConfig := func() *extauthpb.OpaAuth {
		return getExtension()[0].GetOpaAuth()
	}

	get2ndOpaConfig := func() *extauthpb.OpaAuth {
		return getExtension()[1].GetOpaAuth()
	}

	DescribeTable("should create oidc authconfig",
		func(cmd string, expected extauthpb.OAuth) {
			err := testutils.Glooctl(cmd)
			Expect(err).NotTo(HaveOccurred())
			oidc := getOIDCConfig()
			Expect(*oidc).To(Equal(expected))
		},
		Entry("with oid config", "create ac --name ac1 --enable-oidc-auth --oidc-auth-client-id "+
			"1 --oidc-auth-app-url http://app.example.com --oidc-auth-client-secret-name fake "+
			"--oidc-auth-client-secret-namespace fakens --oidc-auth-issuer-url http://issuer.example.com "+
			"--oidc-auth-callback-path /cb",
			extauthpb.OAuth{
				ClientId: "1",
				ClientSecretRef: &core.ResourceRef{
					Name:      "fake",
					Namespace: "fakens",
				},
				CallbackPath: "/cb",
				IssuerUrl:    "http://issuer.example.com",
				AppUrl:       "http://app.example.com",
			}),
		Entry("with default callback", "create ac --name ac1 --enable-oidc-auth --oidc-auth-client-id "+
			"1 --oidc-auth-app-url http://app.example.com --oidc-auth-client-secret-name fake "+
			"--oidc-auth-client-secret-namespace fakens --oidc-auth-issuer-url http://issuer.example.com",
			extauthpb.OAuth{
				ClientId: "1",
				ClientSecretRef: &core.ResourceRef{
					Name:      "fake",
					Namespace: "fakens",
				},
				CallbackPath: "/oidc-gloo-callback",
				IssuerUrl:    "http://issuer.example.com",
				AppUrl:       "http://app.example.com",
			}),
		Entry("with default scopes", "create ac --name ac1 --enable-oidc-auth --oidc-auth-client-id "+
			"1 --oidc-auth-app-url http://app.example.com --oidc-auth-client-secret-name fake "+
			"--oidc-auth-client-secret-namespace fakens --oidc-auth-issuer-url http://issuer.example.com "+
			"--oidc-scope=scope1 --oidc-scope=scope2",
			extauthpb.OAuth{
				ClientId: "1",
				ClientSecretRef: &core.ResourceRef{
					Name:      "fake",
					Namespace: "fakens",
				},
				CallbackPath: "/oidc-gloo-callback",
				IssuerUrl:    "http://issuer.example.com",
				AppUrl:       "http://app.example.com",
				Scopes:       []string{"scope1", "scope2"},
			}),
	)

	DescribeTable("should create apikey authconfig",
		func(cmd string, expected extauthpb.ApiKeyAuth) {
			err := testutils.Glooctl(cmd)
			Expect(err).NotTo(HaveOccurred())
			apiKey := getApiKeyConfig()
			Expect(*apiKey).To(Equal(expected))
		},
		Entry("with apikey config -- label selector", "create ac --name ac1 --enable-apikey-auth "+
			"--apikey-label-selector k1=v1",
			extauthpb.ApiKeyAuth{
				LabelSelector: map[string]string{"k1": "v1"},
			}),

		Entry("with apikey config -- secret refs", "create ac --name ac1 --enable-apikey-auth "+
			"--apikey-secret-namespace ns1 --apikey-secret-name s1 ",
			extauthpb.ApiKeyAuth{
				LabelSelector: nil,
				ApiKeySecretRefs: []*core.ResourceRef{
					{
						Namespace: "ns1",
						Name:      "s1",
					},
				},
			}),
		Entry("with apikey config -- both groups & secret refs", "create ac --name ac1 --enable-apikey-auth "+
			"--apikey-label-selector k1=v1 --apikey-secret-namespace ns1 --apikey-secret-name s1 ",
			extauthpb.ApiKeyAuth{
				LabelSelector: map[string]string{"k1": "v1"},
				ApiKeySecretRefs: []*core.ResourceRef{
					{
						Namespace: "ns1",
						Name:      "s1",
					},
				},
			}),
	)

	Context("ApiKey AuthConfig errors", func() {
		It("throws error if namespace provided and name omitted ", func() {
			_, err := testutils.GlooctlOut("create ac --name ac1 --enable-apikey-auth " +
				"--apikey-secret-namespace ns1")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(authconfig.ProvideNamespaceAndNameError("ns1", "").Error()))
		})
		It("throws error if name provided and namespace omitted ", func() {
			_, err := testutils.GlooctlOut("create ac --name ac1 --enable-apikey-auth " +
				"--apikey-secret-name s1")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal(authconfig.ProvideNamespaceAndNameError("", "s1").Error()))
		})
	})

	DescribeTable("should create opa authconfig",
		func(cmd string, expected extauthpb.OpaAuth) {
			err := testutils.Glooctl(cmd)
			Expect(err).NotTo(HaveOccurred())
			opa := getOpaConfig()
			Expect(*opa).To(Equal(expected))
		},
		Entry("with opa query and no modules", "create ac --name ac1 --enable-opa-auth "+
			"--opa-query test",
			extauthpb.OpaAuth{
				Query: "test",
			}),
		Entry("with opa query and some modules", "create ac --name ac1 --enable-opa-auth "+
			"--opa-query test --opa-module-ref ns1.name1 --opa-module-ref ns2.name2",
			extauthpb.OpaAuth{
				Query:   "test",
				Modules: []*core.ResourceRef{{Namespace: "ns1", Name: "name1"}, {Namespace: "ns2", Name: "name2"}},
			}),
	)

	It("should create opa after oidc", func() {
		err := testutils.Glooctl("create ac --name ac1 " +
			"--enable-oidc-auth --oidc-auth-client-id 1 " +
			"--oidc-auth-app-url http://app.example.com --oidc-auth-client-secret-name fake " +
			"--oidc-auth-client-secret-namespace fakens --oidc-auth-issuer-url http://issuer.example.com " +
			"--oidc-auth-callback-path /cb " +
			"--enable-opa-auth --opa-query test --opa-module-ref ns1.name1 --opa-module-ref ns2.name2")
		Expect(err).NotTo(HaveOccurred())

		expected := extauthpb.OpaAuth{
			Query:   "test",
			Modules: []*core.ResourceRef{{Namespace: "ns1", Name: "name1"}, {Namespace: "ns2", Name: "name2"}},
		}

		oidc := getOIDCConfig()
		Expect(oidc).NotTo(BeNil())

		opa := get2ndOpaConfig()
		Expect(opa).NotTo(BeNil())
		Expect(*opa).To(Equal(expected))

	})

	Context("OPA auth config errors", func() {
		It("throws error if no query provided ", func() {
			_, err := testutils.GlooctlOut("create ac --name ac1 --enable-opa-auth")
			Expect(err).To(MatchError(authconfig.EmptyQueryError))
		})
	})

	Context("Interactive tests", func() {

		It("should create ac with no auth", func() {
			testutil.ExpectInteractive(func(c *testutil.Console) {
				c.ExpectString("do you wish to add oidc auth to the auth config")
				c.SendLine("n")
				c.ExpectString("do you wish to add apikey auth to the auth config")
				c.SendLine("n")
				c.ExpectString("do you wish to add OPA auth to the auth config")
				c.SendLine("n")

				c.ExpectString("Use default namespace (gloo-system)?")
				c.SendLine("")
				c.ExpectString("name of the resource:")
				c.SendLine("default")

				c.ExpectEOF()
			}, func() {
				err := testutils.Glooctl("create ac -i")
				Expect(err).NotTo(HaveOccurred())
				_, err = helpers.MustAuthConfigClient().Read("gloo-system", "default", clients.ReadOpts{})
				Expect(err).NotTo(HaveOccurred())
			})
		})

		It("should create ac with oidc auth", func() {
			testutil.ExpectInteractive(func(c *testutil.Console) {
				c.ExpectString("do you wish to add oidc auth to the auth config")
				c.SendLine("y")
				c.ExpectString("What is your app url?")
				c.SendLine("http://app.example.com")
				c.ExpectString("What is your issuer url?")
				c.SendLine("https://accounts.google.com")
				c.ExpectString("What path (relative to your app url) should we use as a callback from the issuer?")
				c.SendLine("/auth-callback")
				c.ExpectString("What is your client id?")
				c.SendLine("me")
				c.ExpectString("What is your client secret name?")
				c.SendLine("secret-name")
				c.ExpectString("What is your client secret namespace?")
				c.SendLine("gloo-system")
				c.ExpectString("provide additional scopes to request (empty to finish)")
				c.SendLine("scope1")
				c.ExpectString("provide additional scopes to request (empty to finish)")
				c.SendLine("")

				c.ExpectString("do you wish to add apikey auth to the auth config")
				c.SendLine("n")

				c.ExpectString("do you wish to add OPA auth to the auth config")
				c.SendLine("n")

				c.ExpectString("Use default namespace (gloo-system)?")
				c.SendLine("")
				c.ExpectString("name of the resource:")
				c.SendLine("ac1")

				c.ExpectEOF()
			}, func() {
				err := testutils.Glooctl("create ac -i")
				Expect(err).NotTo(HaveOccurred())

				oidc := getOIDCConfig()
				expected := extauthpb.OAuth{
					ClientId: "me",
					ClientSecretRef: &core.ResourceRef{
						Name:      "secret-name",
						Namespace: "gloo-system",
					},
					CallbackPath: "/auth-callback",
					IssuerUrl:    "https://accounts.google.com",
					AppUrl:       "http://app.example.com",
					Scopes:       []string{"scope1"},
				}
				Expect(*oidc).To(Equal(expected))

			})
		})

		It("should create ac with apikey auth", func() {
			testutil.ExpectInteractive(func(c *testutil.Console) {
				c.ExpectString("do you wish to add oidc auth to the auth config")
				c.SendLine("n")
				c.ExpectString("do you wish to add apikey auth to the auth config")
				c.SendLine("y")

				c.ExpectString("provide a label (key=value) to be a part of your label selector (empty to finish)")
				c.SendLine("k1=v1")
				c.ExpectString("provide a label (key=value) to be a part of your label selector (empty to finish)")
				c.SendLine("")
				c.ExpectString("apikey secret name to attach to this auth config? (empty to skip)")
				c.SendLine("s1")
				c.ExpectString("provide a namespace to search for the secret in (empty to finish)")
				c.SendLine("ns1")

				c.ExpectString("do you wish to add OPA auth to the auth config")
				c.SendLine("n")

				c.ExpectString("Use default namespace (gloo-system)?")
				c.SendLine("")
				c.ExpectString("name of the resource:")
				c.SendLine("ac1")

				c.ExpectEOF()
			}, func() {
				err := testutils.Glooctl("create ac -i")
				Expect(err).NotTo(HaveOccurred())

				apiKey := getApiKeyConfig()
				expected := extauthpb.ApiKeyAuth{
					LabelSelector: map[string]string{"k1": "v1"},
					ApiKeySecretRefs: []*core.ResourceRef{
						{
							Namespace: "ns1",
							Name:      "s1",
						},
					},
				}
				Expect(*apiKey).To(Equal(expected))

			})
		})

		It("should create ac with OPA auth", func() {
			testutil.ExpectInteractive(func(c *testutil.Console) {
				c.ExpectString("do you wish to add oidc auth to the auth config")
				c.SendLine("n")
				c.ExpectString("do you wish to add apikey auth to the auth config")
				c.SendLine("n")

				c.ExpectString("do you wish to add OPA auth to the auth config")
				c.SendLine("y")

				c.ExpectString("OPA query to attach to this auth config?")
				c.SendLine("test")
				c.ExpectString("provide references to config maps used as OPA modules in resolving above query (empty to finish)")
				c.SendLine("ns1.name1")
				c.ExpectString("provide references to config maps used as OPA modules in resolving above query (empty to finish)")
				c.SendLine("")

				c.ExpectString("Use default namespace (gloo-system)?")
				c.SendLine("")
				c.ExpectString("name of the resource:")
				c.SendLine("ac1")

				c.ExpectEOF()
			}, func() {
				err := testutils.Glooctl("create ac -i")
				Expect(err).NotTo(HaveOccurred())

				opa := getOpaConfig()
				expected := extauthpb.OpaAuth{
					Query: "test",
					Modules: []*core.ResourceRef{
						{
							Namespace: "ns1",
							Name:      "name1",
						},
					},
				}
				Expect(*opa).To(Equal(expected))

			})
		})

	})

})
