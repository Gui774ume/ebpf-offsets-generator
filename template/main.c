#include <linux/kconfig.h>
#include <linux/version.h>

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Waddress-of-packed-member"
#pragma clang diagnostic ignored "-Wgnu-variable-sized-type-not-at-end"
#include <linux/ptrace.h>
#pragma clang diagnostic pop

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wgnu-variable-sized-type-not-at-end"
{{ range .Offset.IncludePaths }}
#include {{ unescape . }} {{ end }}
#pragma clang diagnostic pop

// Custom eBPF includes
#include "bpf/bpf_map.h"
#include "bpf/bpf.h"
#include "bpf/bpf_helpers.h"
#include "bpf/bpf_endians.h"

{{ range .KernelConfig }}
#ifndef {{ . }}
#define {{ . }} 1
#endif {{ end }}

#ifndef offsetof
#define offsetof(TYPE, MEMBER) ((size_t) &((TYPE *)0)->MEMBER)
#endif

SEC("tracepoint/syscalls/sys_enter_execve")
int bpf_prog(void *ctx)
{
    char a[] = "%d\n";
    bpf_trace_printk(a, sizeof(a), offsetof(struct {{ .Offset.StructureName }}, {{ .Offset.MemberName }}));
    return 0;
}

char _license[] SEC("license") = "GPL";
__u32 _version SEC("version") = 0xFFFFFFFE;
