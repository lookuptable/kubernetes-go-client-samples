load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_prefix")

go_prefix("github.com/lookuptable/client-go-example")

go_library(
    deps = [
        "@io_k8s_client_go//kubernetes:go_default_library",
        "@io_k8s_client_go//tools/clientcmd:go_default_library",
    ],
    name = "go_default_library",
    srcs = ["utils.go"],
    visibility = ["//visibility:public"],
)
