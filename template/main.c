#include <linux/kconfig.h>
#include <linux/version.h>

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Waddress-of-packed-member"
#pragma clang diagnostic ignored "-Wgnu-variable-sized-type-not-at-end" {{ range .Offset.IncludePaths }}
#include {{ unescape . }} {{ end }}
#pragma clang diagnostic pop

// Custom eBPF includes
#include "bpf/bpf_map.h"
#include "bpf/bpf.h"
#include "bpf/bpf_helpers.h"
#include "bpf/bpf_endians.h"

#ifndef offsetof
#define offsetof(TYPE, MEMBER) ((size_t) &((TYPE *)0)->MEMBER)
#endif

SEC("tracepoint/syscalls/sys_enter_execve")
int bpf_prog(void *ctx)
{
    int a = offsetof(struct {{ .Offset.StructureName }}, {{ .Offset.MemberName }});
    return a;
}

char _license[] SEC("license") = "GPL";
__u32 _version SEC("version") = 0xFFFFFFFE;
