root_repos_remote="/home/barriosj/work/src/"


local_repository(
    name = "io_bazel_rules_go",
    path = "/home/barriosj/rules_go",
)

local_repository(
    name = "com_github_bazelbuild_buildtools",
    path = "/home/barriosj/buildtools"
)



load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_sdk",
    "go_rules_dependencies",
    "go_repository",
    "go_host_sdk",
    "go_register_toolchains",
)

go_repository(
name = "org_golang_x_tools",
importpath = "golang.org/x/tools",
commit="3d92dd60033c312e3ae7cac319c792271cf67e37",
remote = root_repos_remote+"golang.org/x/tools",
vcs = "git"
)

go_repository(
    name = "com_github_magiconair_properties",
    commit = "8d7837e64d3c1ee4e54a880c5a920ab4316fc90a",
    importpath = "github.com/magiconair/properties",
    remote = root_repos_remote+"github.com/magiconair/properties",
    vcs = "git",
)




go_repository(
    name = "org_golang_x_crypto",
    commit = "81e90905daefcd6fd217b62423c0908922eadb30",
    importpath = "golang.org/x/crypto",
    remote = root_repos_remote+"golang.org/x/crypto",  # installed by go get
    vcs = "git",
)

go_repository(
    name = "org_golang_x_net",
    commit = "57efc9c3d9f91fb3277f8da1cff370539c4d3dc5",
    importpath = "golang.org/x/net",
    remote = root_repos_remote+"golang.org/x/net",
    vcs = "git",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "0b25a408a50076fbbcae6b7ac0ea5fbb0b085e79",
    importpath = "golang.org/x/sys",
    remote = root_repos_remote+"golang.org/x/sys",
    vcs = "git",
)

go_repository(
    name = "org_golang_x_text",
    commit = "a9a820217f98f7c8a207ec1e45a874e1fe12c478",
    importpath = "golang.org/x/text",
    remote = root_repos_remote+"golang.org/x/text",
    vcs = "git",
)

go_repository(
    name = "org_golang_x_tools",
    commit = "663269851cdddc898f963782f74ea574bcd5c814",
    importpath = "golang.org/x/tools",
    remote = root_repos_remote+"golang.org/x/tools",
    vcs = "git",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    commit = "89742aefa4b206dcf400792f3bd35b542998eb3b",
    importpath = "github.com/sirupsen/logrus",
    remote = root_repos_remote+"github.com/sirupsen/logrus",
    vcs = "git",
)

go_repository(
    name = "com_github_spf13_pflag",
    commit = "a9789e855c7696159b7db0db7f440b449edf2b31",
    importpath = "github.com/spf13/pflag",
    remote = root_repos_remote+"github.com/spf13/pflag",
    vcs = "git",
)

go_repository(
    name = "com_github_spf13_cast",
    commit = "acbeb36b902d72a7a4c18e8f3241075e7ab763e4",
    importpath = "github.com/spf13/cast",
    remote = root_repos_remote+"github.com/spf13/cast",
    vcs = "git",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    commit = "12bd96e66386c1960ab0f74ced1362f66f552f7b",
    importpath = "github.com/spf13/jwalterweatherman",
    remote = root_repos_remote+"github.com/spf13/jwalterweatherman",
    vcs = "git",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "d0303fe809921458f417bcf828397a65db30a7e4",
    importpath = "github.com/mitchellh/mapstructure",
    remote = root_repos_remote+"github.com/mitchellh/mapstructure",
    vcs = "git",
)

go_repository(
    name = "com_github_spf13_afero",
    commit = "e67d870304c4bca21331b02f414f970df13aa694",
    importpath = "github.com/spf13/afero",
    remote = root_repos_remote+"github.com/spf13/afero",
    vcs = "git",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    commit = "2009e44b6f182e34d8ce081ac2767622937ea3d4",
    importpath = "github.com/pelletier/go-toml",
    remote = root_repos_remote+"github.com/pelletier/go-toml",
    vcs = "git",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    commit = "42e33e2d55a0ff1d6263f738896ea8c13571a8d0",
    importpath = "github.com/hashicorp/hcl",
    remote = root_repos_remote+"github.com/hashicorp/hcl",
    vcs = "git",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    commit = "4da3e2cfbabc9f751898f250b49f2439785783a1",
    importpath = "github.com/fsnotify/fsnotify",
    remote = root_repos_remote+"github.com/fsnotify/fsnotify",
    vcs = "git",
)

local_repository(
    name = "com_google_protobuf",
    path = "/home/barriosj/google.com/protobuf",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "eb3733d160e74a9c7e442f435eb3bea458e1d19f",
    importpath = "gopkg.in/yaml.v2",
    remote = root_repos_remote+"gopkg.in/yaml.v2",
    vcs = "git",
)

go_repository(
    name = "com_github_spf13_viper",
    commit = "d9cca5ef33035202efb1586825bdbb15ff9ec3ba",
    importpath = "github.com/spf13/viper",
    remote = root_repos_remote+"github.com/spf13/viper",
    vcs = "git",
)
go_rules_dependencies()

go_register_toolchains(go_version = "host")
