## eBPF offsets generator

> A BPF program is a piece of user-provided code which is injected straight into a kernel. Once loaded and verified, BPF programs execute in kernel context. These programs operate inside kernel memory space with access to all the internal kernel state available to it. This is extremely powerful and is one of the reasons why BPF technology is successfully used in so many varied applications. However, this powerful capability also creates the BPF portability pains we have today: BPF programs do not control memory layout of a surrounding kernel environment. They have to work with what they get from independently developed, compiled, and deployed kernels. Additionally, kernel types and data structures are in constant flux. Different kernel versions will have struct fields shuffled around inside a struct, or even moved into a new inner struct. Fields can be renamed or removed, their types changed, either into some trivially-compatible ones or completely different ones. Structs and other types can get renamed, or they can be conditionally compiled out (depending on kernel configuration), or just plain removed between kernel versions.

_Abstract from [BPF Portability and CO-RE](https://facebookmicrosites.github.io/bpf/blog/2020/02/19/bpf-portability-and-co-re.html)_

This project is an attempt to provide an alternative (an clearly less powerful) solution to the CO-RE initiative. Even if CO-RE will probably become the standard for kernels 4.18+ when it comes to eBPF, we still lack a solution to handle older 4.x kernels that are still widely used and will be widely used for the next few years. Also, CO-RE has some requirements on the kernel side (specifically the `CONFIG_DEBUG_INFO_BTF` kernel flag has to be activated), that major distributions do not ship by default (yet).
In other words, until CO-RE picks up some speed, we are left with solutions that ship a compiler in production environments. Depending on the compliance standards that you might want to certify for, a compiler based solution might become a no-go, and you will have to compile a different version of your eBPF programs for each version of your kernels.

This project hypothesizes that it might be possible to pre-compute offsets for all major kernel versions (in a CI for example) and edit the eBPF binaries with the right offsets at runtime.
Obviously, this will not work for custom built and modified kernels, but it should work for major distributions.

This repository contains the eBPF offset generator part of the project. Edition at runtime is handled in the [Gui774ume/ebpf](https://github.com/Gui774ume/ebpf).

## Requirements

The project was developed on Linux ubuntu-bionic, kernel version 4.15.0. You will also need the following packages:

- git

## Getting started

1) Vendor go dependencies

```shell script
go mod vendor
```

2) Build the project

```shell script
make build-ebpf
make build
```

## Compute kernel offsets

1) Provide the list of struct members that you want the program to compute, as well as the range of kernels you are interested in. As the name of some members sometimes change overtime, you can also define a `fallback` member name. Available options are explained below.

```shell script
blabla
```

2) To compute the offsets, run the following command:

```shell script
blabla
```

3) Package and ship the generated offsets along with your eBPF programs. To load your eBPF program, use the method "" instead of "" like so:

```go
struct yolo
```

## Example

3) Start the project

```shell script
blabla
```
