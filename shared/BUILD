load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "shared_deps",
    srcs = glob([
        "*.go"
    ]),
    importpath = "github.com/hariharasudhan-nineleaps/bazel_mono_repo/shared",
    visibility = ["//visibility:public"],
    deps = [
        "@org_uber_go_zap//:go_default_library"
    ]
)