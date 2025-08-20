load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-csharp",
    visibility = ["//visibility:private"],
    deps = [
        "//module:go_default_library",
        "@com_github_lyft_protoc_gen_star_v2//:go_default_library",
        "@org_golang_google_protobuf//types/pluginpb:go_default_library",
    ],
)

go_binary(
    name = "protoc-gen-validate-csharp",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
