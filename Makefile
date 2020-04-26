build:
	go build -mod vendor -o bin/offsets-generator cmd/offsets-generator/main.go

dev:
	go build -mod vendor -o bin/offsets-generator cmd/offsets-generator/main.go
	./bin/offsets-generator --offsets="example/task_struct__pid/offsets_declaration.yaml" --output="example/task_struct__pid/generated_offsets.yaml"

run-ubuntu-scraper:
	go run -mod vendor cmd/ubuntu-scraper/main.go

clean-tmp:
	rm -rf ./tmp/

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
		tmp/generated_programs/task_struct__pid/task_struct__pid.c \
		-c -o - | llc -march=bpf -filetype=obj -o ./bin/probe.o
