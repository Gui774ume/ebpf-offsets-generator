all: build-ebpf build run

build-ebpf:
	mkdir -p ebpf/bin
	clang -D__KERNEL__ -D__ASM_SYSREG_H \
		-Wno-unused-value \
		-Wno-pointer-sign \
		-Wno-compare-distinct-pointer-types \
		-Wunused \
		-Wall \
		-Werror \
		-I/lib/modules/$$(uname -r)/build/include \
		-I/lib/modules/$$(uname -r)/build/include/uapi \
		-I/lib/modules/$$(uname -r)/build/include/generated/uapi \
		-I/lib/modules/$$(uname -r)/build/arch/x86/include \
		-I/lib/modules/$$(uname -r)/build/arch/x86/include/uapi \
		-I/lib/modules/$$(uname -r)/build/arch/x86/include/generated \
		-O2 -emit-llvm \
		ebpf/main.c \
		-c -o - | llc -march=bpf -filetype=obj -o ebpf/bin/probe.o
	go-bindata -pkg probe -prefix "ebpf/bin" -modtime 1 -o "pkg/probe/probe.go" "ebpf/bin"

build:
	go build -mod vendor -o bin/offsets-generator cmd/main.go

dev:
	go build -mod vendor -o bin/offsets-generator cmd/main.go
	./bin/offsets-generator --offsets="example/offsets_declaration.yaml"

build-generated:
	clang -D__KERNEL__ -D__ASM_SYSREG_H \
		-Wno-unused-value \
		-Wno-pointer-sign \
		-Wno-compare-distinct-pointer-types \
		-Wunused \
		-Wall \
		-I ./tmp/headers2/tmp/usr/src/linux-headers-4.8.0-040800-generic/include \
		-I ./tmp/headers2/tmp/usr/src/linux-headers-4.8.0-040800-generic/includ/uapi \
		-I ./tmp/headers2/tmp/usr/src/linux-headers-4.8.0-040800-generic/includ/generated/uapi \
		-I ./tmp/headers2/tmp/usr/src/linux-headers-4.8.0-040800-generic/arch/x86/include \
		-I ./tmp/headers2/tmp/usr/src/linux-headers-4.8.0-040800-generic/arch/x86/include/uapi \
		-I ./tmp/headers2/tmp/usr/src/linux-headers-4.8.0-040800-generic/arch/x86/include/generated \
		-O2 -emit-llvm \
		tmp/generated_programs/task_struct__pid/CONFIG_MEMCG.c \
		-c -o - | llc -march=bpf -filetype=obj -o ./bin/probe.o
