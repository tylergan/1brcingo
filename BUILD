load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/tylergan/billion_rows
gazelle(name = "gazelle")

go_library(
    name = "billion_rows_lib",
    srcs = ["main.go"],
    importpath = "github.com/tylergan/billion_rows",
    visibility = ["//visibility:private"],
    deps = ["//pkg"],
)

go_binary(
    name = "billion_rows",
    embed = [":billion_rows_lib"],
    visibility = ["//visibility:public"],
)
