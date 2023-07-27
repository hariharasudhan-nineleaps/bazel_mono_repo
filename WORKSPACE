workspace(name = "bazel_mono_repo")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

##################
# Pre-requisites #
##################
http_archive(
    name = "bazel_skylib",
    sha256 = "66ffd9315665bfaafc96b52278f57c7e2dd09f5ede279ea6d39b2be471e7e3aa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.4.2/bazel-skylib-1.4.2.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

##################
# Python Support #
##################

# download rules-python
http_archive(
    name = "rules_python",
    sha256 = "0a8003b044294d7840ac7d9d73eef05d6ceb682d7516781a4ec62eeb34702578",
    strip_prefix = "rules_python-0.24.0",
    url = "https://github.com/bazelbuild/rules_python/releases/download/0.24.0/rules_python-0.24.0.tar.gz",
)

# register python toolchain
load("@rules_python//python:repositories.bzl", "python_register_toolchains")

python_register_toolchains(
    name = "python3_11",
    python_version = "3.11",
    register_coverage_tool = True,
)

# install pip packages with toolchain
load("@python3_11//:defs.bzl", "interpreter")
load("@rules_python//python:pip.bzl", "pip_parse")

pip_parse(
    name = "app_deps",
    python_interpreter_target = interpreter,
    requirements_lock = "//:requirements.txt",
)

load("@app_deps//:requirements.bzl", "install_deps")

install_deps()

################
#  GO Support ##
###############
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    version = "1.20.5",
)

# install go packages
http_archive(
    name = "bazel_gazelle",
    sha256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_gin_gonic_gin",
    importpath = "github.com/gin-gonic/gin",
    sum = "h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=",
    version = "v1.9.1",
)

go_repository(
    name = "com_github_bytedance_sonic",
    importpath = "github.com/bytedance/sonic",
    sum = "h1:oDfRZ+4m6AYCOC0GFeOCeYqvBmucy1isvouS2K0cPzo=",
    version = "v1.10.0-rc2",
)

go_repository(
    name = "com_github_chenzhuoyu_base64x",
    importpath = "github.com/chenzhuoyu/base64x",
    sum = "h1:77cEq6EriyTZ0g/qfRdp61a3Uu/AWrgIq2s0ClJV1g0=",
    version = "v0.0.0-20230717121745-296ad89f973d",
)

go_repository(
    name = "com_github_chenzhuoyu_iasm",
    importpath = "github.com/chenzhuoyu/iasm",
    sum = "h1:9fhXjVzq5hUy2gkhhgHl95zG2cEAhw9OSGs8toWWAwo=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_gabriel_vasile_mimetype",
    importpath = "github.com/gabriel-vasile/mimetype",
    sum = "h1:w5qFW6JKBz9Y393Y4q372O9A7cUSequkh1Q7OhCmWKU=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    importpath = "github.com/gin-contrib/sse",
    sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_go_playground_assert_v2",
    importpath = "github.com/go-playground/assert/v2",
    sum = "h1:JvknZsQTYeFEAhQwI4qEt9cyV5ONwRHC+lYKSsYSR8s=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_go_playground_locales",
    importpath = "github.com/go-playground/locales",
    sum = "h1:EWaQ/wswjilfKLTECiXz7Rh+3BjFhfDFKv/oXslEjJA=",
    version = "v0.14.1",
)

go_repository(
    name = "com_github_go_playground_universal_translator",
    importpath = "github.com/go-playground/universal-translator",
    sum = "h1:Bcnm0ZwsGyWbCzImXv+pAJnYK9S473LQFuzCbDbfSFY=",
    version = "v0.18.1",
)

go_repository(
    name = "com_github_go_playground_validator_v10",
    importpath = "github.com/go-playground/validator/v10",
    sum = "h1:9c50NUPC30zyuKprjL3vNZ0m5oG+jU0zvx4AqHGnv4k=",
    version = "v10.14.1",
)

go_repository(
    name = "com_github_goccy_go_json",
    importpath = "github.com/goccy/go-json",
    sum = "h1:CrxCmQqYDkv1z7lO7Wbh2HN93uovUHgrECaO5ZrCXAU=",
    version = "v0.10.2",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:LUVKkCeviFUMKqHa4tXIIij/lbhnMbP7Fn5wKdKkRh4=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:Khx7svrCpmxxtHBq5j2mp/xVjsi8hQMfNLvJFAlrGgU=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_klauspost_cpuid_v2",
    importpath = "github.com/klauspost/cpuid/v2",
    sum = "h1:0E5MSMDEoAulmXNFquVs//DdoomxaoTY1kUhbc/qbZg=",
    version = "v2.2.5",
)

go_repository(
    name = "com_github_knz_go_libedit",
    importpath = "github.com/knz/go-libedit",
    sum = "h1:0pHpWtx9vcvC0xGZqEQlQdfSQs7WRlAjuPvk3fOZDCo=",
    version = "v1.10.1",
)

go_repository(
    name = "com_github_leodido_go_urn",
    importpath = "github.com/leodido/go-urn",
    sum = "h1:XlAE/cm/ms7TE/VMVoduSpNBoyc2dOxHs5MZSwAN63Q=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:JITubQf0MOLdlGRuRq+jtsDlekdYPia9ZFsB8h/APPA=",
    version = "v0.0.19",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_pelletier_go_toml_v2",
    importpath = "github.com/pelletier/go-toml/v2",
    sum = "h1:uH2qQXheeefCCkuBBSLi7jCiSmj3VRh2+Goq2N7Xxu0=",
    version = "v2.0.9",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
    version = "v1.8.4",
)

go_repository(
    name = "com_github_twitchyliquid64_golang_asm",
    importpath = "github.com/twitchyliquid64/golang-asm",
    sum = "h1:SU5vSMR7hnwNxj24w34ZyCi/FmDZTkS4MhqMhdFk5YI=",
    version = "v0.15.1",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:BMaWp1Bb6fHwEtbplGBGJ498wD+LKlNSl25MjdZY4dU=",
    version = "v1.2.11",
)

go_repository(
    name = "com_nullprogram_x_optparse",
    importpath = "nullprogram.com/x/optparse",
    sum = "h1:xGFgVi5ZaWOnYdac2foDT3vg0ZZC9ErXFV57mr4OHrI=",
    version = "v1.0.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "io_rsc_pdf",
    importpath = "rsc.io/pdf",
    sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
    version = "v0.1.1",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:g0LDEJHgrBl9N9r17Ru3sqWhkIx2NB67okBHPwC7hs8=",
    version = "v1.31.0",
)

go_repository(
    name = "org_golang_x_arch",
    importpath = "golang.org/x/arch",
    sum = "h1:A8WCeEWhLwPBKNbFi5Wv5UTCBx5zzubnXDlMOFAzFMc=",
    version = "v0.4.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:6Ewdq3tDic1mg5xRO4milcWCfMVQhI4NkqWWvqejpuA=",
    version = "v0.11.0",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:LUYupSeNrTNCGzR/hVBk2NHZO4hXcVaW1k4Qx7rjPx8=",
    version = "v0.8.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:cfawfvKITfUsFCeJIHJrbSxpeu/E81khclypR0GVT50=",
    version = "v0.12.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:SqMFp9UcQJZa+pmYuAKjd9xq1f0j5rLcDIk0mj4qAsA=",
    version = "v0.10.0",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:3R7pNqamzBraeqj/Tj8qt1aQ2HpmlC+Cx/qL/7hn4/c=",
    version = "v0.10.0",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:LAntKIrcmeSKERyiOh0XMV39LXS8IE9UL2yP7+f5ij4=",
    version = "v0.11.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:BOw41kyTf3PuCW1pVQf8+Cyg8pMlkYB1oo9iJ6D/lKM=",
    version = "v0.6.0",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
    version = "v0.0.0-20191204190536-9bdfabe68543",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:iURUrRGxPUNPdy5/HRSm+Yj6okJ6UtLINN0Q9M4+h3I=",
    version = "v0.8.1",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:ZvwS0R+56ePWxUNi+Atn9dWONBPp/AUETXlHW0DxSjE=",
    version = "v1.11.0",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:wy28qYRKZgnJTxGxvye5/wgWr1EKjmUDGYox5mGlRlI=",
    version = "v1.1.11",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=",
    version = "v1.11.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:FiJd5l1UOLj0wCgbSE0rwwXHzEdAZS6hiiSnxJN/D60=",
    version = "v1.24.0",
)

gazelle_dependencies()
