package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "attrgenmgr.go",
        "common.go",
        "descriptors.go",
        "inventory.go",
        "manager.go",
        "quotasManager.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//mixer/pkg/adapter:go_default_library",
        "//mixer/pkg/aspect/config:go_default_library",
        "//mixer/pkg/attribute:go_default_library",
        "//mixer/pkg/config:go_default_library",
        "//mixer/pkg/config/descriptor:go_default_library",
        "//mixer/pkg/config/proto:go_default_library",
        "//mixer/pkg/expr:go_default_library",
        "//mixer/pkg/status:go_default_library",
        "@com_github_golang_glog//:go_default_library",
        "@com_github_googleapis_googleapis//:google/rpc",
        "@com_github_hashicorp_go_multierror//:go_default_library",
        "@io_istio_api//:mixer/v1/config/descriptor",
    ],
)

go_test(
    name = "go_default_test",
    size = "small",
    srcs = [
        "attrgenmgr_test.go",
        "common_test.go",
        "descriptors_test.go",
        "inventory_test.go",
        "quotasManager_test.go",
    ],
    library = ":go_default_library",
    deps = [
        "//mixer/adapter/noopLegacy:go_default_library",
        "//mixer/pkg/adapter:go_default_library",
        "//mixer/pkg/adapter/test:go_default_library",
        "//mixer/pkg/aspect/config:go_default_library",
        "//mixer/pkg/aspect/test:go_default_library",
        "//mixer/pkg/attribute:go_default_library",
        "//mixer/pkg/config:go_default_library",
        "//mixer/pkg/config/descriptor:go_default_library",
        "//mixer/pkg/config/proto:go_default_library",
        "//mixer/pkg/expr:go_default_library",
        "//mixer/pkg/status:go_default_library",
        "@io_istio_api//:mixer/v1/config/descriptor",
    ],
)
