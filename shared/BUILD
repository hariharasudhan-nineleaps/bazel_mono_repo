load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "shared_deps",
    srcs = glob(["*.go"], exclude=["*_test.go"]),
    importpath = "github.com/hariharasudhan-nineleaps/bazel_mono_repo/shared",
    visibility = ["//visibility:public"],
    deps = [
        "@org_uber_go_zap//:go_default_library"
    ]
)

go_test(
    name = "shared_test",
    srcs = ["jwt_test.go", "logger_test.go"],
    embed = [
        ":shared_deps"
    ]
)