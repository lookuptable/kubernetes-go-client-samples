load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "go_binary")

go_prefix("github.com/lookuptable/client-go-example")

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        "@io_k8s_client_go//kubernetes:go_default_library",
        "@io_k8s_client_go//pkg/api/v1:go_default_library",
        "@io_k8s_client_go//tools/clientcmd:go_default_library",
    ],
)
