## eBPF offsets generator

> A BPF program is a piece of user-provided code which is injected straight into a kernel. Once loaded and verified, BPF programs execute in kernel context. These programs operate inside kernel memory space with access to all the internal kernel state available to it. This is extremely powerful and is one of the reasons why BPF technology is successfully used in so many varied applications. However, this powerful capability also creates the BPF portability pains we have today: BPF programs do not control memory layout of a surrounding kernel environment. They have to work with what they get from independently developed, compiled, and deployed kernels. Additionally, kernel types and data structures are in constant flux. Different kernel versions will have struct fields shuffled around inside a struct, or even moved into a new inner struct. Fields can be renamed or removed, their types changed, either into some trivially-compatible ones or completely different ones. Structs and other types can get renamed, or they can be conditionally compiled out (depending on kernel configuration), or just plain removed between kernel versions.

_Abstract from [BPF Portability and CO-RE](https://facebookmicrosites.github.io/bpf/blog/2020/02/19/bpf-portability-and-co-re.html)_

This project is an attempt to provide an alternative (an clearly less powerful) solution to the CO-RE initiative. Even if CO-RE will probably become the standard for kernels 5.x when it comes to eBPF, we still lack a solution to handle older 4.x kernels that are still widely used and will be widely used for the next few years. Also, CO-RE has some requirements on the kernel side (specifically the `CONFIG_DEBUG_INFO_BTF` kernel flag has to be activated), that major distributions do not ship by default (yet).
In other words, until CO-RE picks up some speed, we are left with solutions that ship a compiler in production environments. Depending on the compliance standards that you might want to certify for, a compiler based solution might become a no-go, and you will have to compile a different version of your eBPF programs for each version of your kernels.

This project hypothesizes that it might be possible to pre-compute offsets for major kernel versions (in a CI for example) and edit the eBPF binaries with the right offsets at runtime.
Obviously, this will not work for custom-built and modified kernels, but it should work for major distributions. For now, only Ubuntu is supported.

This repository contains the eBPF offset generator part of the project. Edition at runtime is handled in [Gui774ume/ebpf](https://github.com/Gui774ume/ebpf).

## Development environment

The project was developed on Linux ubuntu-bionic, kernel version 4.15.0.

## Getting started

1) Vendor go dependencies

```shell script
go mod vendor
```

2) Scrape Ubuntu's package repository to list all available headers.

```shell script
make run-ubuntu-scraper
```

3) Build the project

```shell script
make build
```

## Compute kernel offsets

1) Provide the list of struct members that you want the program to compute, as well as the range of kernels you are interested in. As the name of some members sometimes change overtime, you will soon be able to define a `fallback` member name. Available options are listed below.

```shell script
./bin/offsets-generator --offsets="<path_to_your_offsets_declaration>.yaml" --output="<output_file>.yaml"
```

3) [Gui774ume/ebpf](https://github.com/Gui774ume/ebpf) has an [Editor](https://godoc.org/github.com/Gui774ume/ebpf#Editor) feature that can be used to push the generated offsets in constants at runtime, without having to re-compile your programs.

## Example

We provided an example for the `pid` member of the `task_struct` structure. You can run the example with the following command:

```shell script
./bin/offsets-generator --offsets="example/task_struct__pid/offsets_declaration.yaml" --output="example/task_struct__pid/generated_offsets.yaml"
```

`offsets_declaration.yaml`:

```yaml
version: 1.0
min_kernel: 4.15.0
max_kernel: 4.17.10
offsets:
  - offset_symbol: task_struct__pid
    structure_name: task_struct
    member_name: pid
    include_paths:
      - "<linux/ptrace.h>"
      - "<linux/sched.h>"
```

`generated_offsets.yaml`:

```yaml
version: "1.0"
provider: ubuntu
generated_offsets:
  task_struct__pid:
  - min_kernel: v4.15
    max_kernel: v4.15.5
    offset: 2280
  - min_kernel: v4.15.7
    max_kernel: v4.15.18
    offset: 2280
  - min_kernel: v4.16
    max_kernel: v4.16.3
    offset: 2280
  - min_kernel: v4.16.4
    max_kernel: v4.16.18
    offset: 2216
  - min_kernel: v4.17
    max_kernel: v4.17.10
    offset: 2216
```

## Future work

- Right now the kernel headers pulled for Ubuntu have the kernel config of the `generic` release of the distribution. Unfortunately, Ubuntu has hundreds of sub versions for each of its releases. Even if the different kernel configs are never too far away from each other, it is enough to make those offsets unusable. A feature that shouldn't be too hard to implement is to provide a path to the kernel config of the host, so that the script can apply the right kernel configuration instead of the default generic one.
- Support for Debian and other major linux distributions.
