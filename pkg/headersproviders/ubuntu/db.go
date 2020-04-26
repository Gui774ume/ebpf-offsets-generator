
package ubuntu

import "github.com/Gui774ume/ebpf-offsets-generator/pkg/model"

var db = model.UbuntuDatabase{
	Packages: map[model.KernelVersion]*model.UbuntuPackages{ 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 82,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.82/linux-headers-3.2.82-030282_3.2.82-030282.201608231034_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.82/linux-headers-3.2.82-030282-generic_3.2.82-030282.201608231034_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 83,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.83/linux-headers-3.2.83-030283_3.2.83-030283.201610210432_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 84,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.84/linux-headers-3.2.84-030284_3.2.84-030284.201611200532_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 85,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.85/linux-headers-3.2.85-030285_3.2.85-030285.201702230332_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 86,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.86/linux-headers-3.2.86-030286_3.2.86-030286.201702270232_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 87,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.87/linux-headers-3.2.87-030287_3.2.87-030287.201703200659_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.87/linux-headers-3.2.87-030287-generic_3.2.87-030287.201703200659_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 88,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.88/linux-headers-3.2.88-030288_3.2.88-030288.201704050532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.88/linux-headers-3.2.88-030288-generic_3.2.88-030288.201704050532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 89,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.89/linux-headers-3.2.89-030289_3.2.89-030289.201706060537_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.89/linux-headers-3.2.89-030289-generic_3.2.89-030289.201706060537_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 90,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.90/linux-headers-3.2.90-030290_3.2.90-030290.201707030332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.90/linux-headers-3.2.90-030290-generic_3.2.90-030290.201707030332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 91,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.91/linux-headers-3.2.91-030291_3.2.91-030291.201707181532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.91/linux-headers-3.2.91-030291-generic_3.2.91-030291.201707181532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 92,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.92/linux-headers-3.2.92-030292_3.2.92-030292.201708260632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.92/linux-headers-3.2.92-030292-generic_3.2.92-030292.201708260632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 93,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.93/linux-headers-3.2.93-030293_3.2.93-030293.201709151532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.93/linux-headers-3.2.93-030293-generic_3.2.93-030293.201709151532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 94,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.94/linux-headers-3.2.94-030294_3.2.94-030294.201710121635_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.94/linux-headers-3.2.94-030294-generic_3.2.94-030294.201710121635_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 95,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.95/linux-headers-3.2.95-030295_3.2.95-030295.201711130935_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.95/linux-headers-3.2.95-030295-generic_3.2.95-030295.201711130935_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 96,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.96/linux-headers-3.2.96-030296_3.2.96-030296.201711261635_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.96/linux-headers-3.2.96-030296-generic_3.2.96-030296.201711261635_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 97,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.97/linux-headers-3.2.97-030297_3.2.97-030297.201801021031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.97/linux-headers-3.2.97-030297-generic_3.2.97-030297.201801021031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 98,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.98/linux-headers-3.2.98-030298_3.2.98-030298.201801072130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.98/linux-headers-3.2.98-030298-generic_3.2.98-030298.201801072130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 99,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.99/linux-headers-3.2.99-030299_3.2.99-030299.201802141031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.99/linux-headers-3.2.99-030299-generic_3.2.99-030299.201802141031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 100,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.100/linux-headers-3.2.100-0302100_3.2.100-0302100.201803041028_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.100/linux-headers-3.2.100-0302100-generic_3.2.100-0302100.201803041028_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 2,
			Patch: 101,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.101/linux-headers-3.2.101-0302101_3.2.101-0302101.201803192131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.2.101/linux-headers-3.2.101-0302101-generic_3.2.101-0302101.201803192131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 4,
			Patch: 113,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.4.113/linux-headers-3.4.113-0304113_3.4.113-0304113.201610261546_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.4.113/linux-headers-3.4.113-0304113-generic_3.4.113-0304113.201610261546_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 8,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.8.3/linux-headers-3.8.3-030803_3.8.3-030803.201805132101_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.8.3/linux-headers-3.8.3-030803-generic_3.8.3-030803.201805132101_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 8,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.8.7/linux-headers-3.8.7-030807_3.8.7-030807.201805131709_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.8.7/linux-headers-3.8.7-030807-generic_3.8.7-030807.201805131709_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 62,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.62/linux-headers-3.12.62-031262_3.12.62-031262.201607211632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.62/linux-headers-3.12.62-031262-generic_3.12.62-031262.201607211632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 63,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.63/linux-headers-3.12.63-031263_3.12.63-031263.201609061133_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.63/linux-headers-3.12.63-031263-generic_3.12.63-031263.201609061133_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 64,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.64/linux-headers-3.12.64-031264_3.12.64-031264.201610030943_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.64/linux-headers-3.12.64-031264-generic_3.12.64-031264.201610030943_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 65,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.65/linux-headers-3.12.65-031265_3.12.65-031265.201610190831_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 66,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.66/linux-headers-3.12.66-031266_3.12.66-031266.201610210522_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.66/linux-headers-3.12.66-031266-generic_3.12.66-031266.201610210522_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 67,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.67/linux-headers-3.12.67-031267_3.12.67-031267.201611100231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.67/linux-headers-3.12.67-031267-generic_3.12.67-031267.201611100231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 68,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.68/linux-headers-3.12.68-031268_3.12.68-031268.201611291231_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 69,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.69/linux-headers-3.12.69-031269_3.12.69-031269.201612180431_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 70,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.70/linux-headers-3.12.70-031270_3.12.70-031270.201702010831_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 71,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.71/linux-headers-3.12.71-031271_3.12.71-031271.201703091634_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 72,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.72/linux-headers-3.12.72-031272_3.12.72-031272.201703162318_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 73,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.73/linux-headers-3.12.73-031273_3.12.73-031273.201704131533_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.73/linux-headers-3.12.73-031273-generic_3.12.73-031273.201704131533_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 12,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.74/linux-headers-3.12.74-031274_3.12.74-031274.201705101001_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.12.74/linux-headers-3.12.74-031274-generic_3.12.74-031274.201705101001_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 14,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.74/linux-headers-3.14.74-031474_3.14.74-031474.201607271333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.74/linux-headers-3.14.74-031474-generic_3.14.74-031474.201607271333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 14,
			Patch: 75,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.75/linux-headers-3.14.75-031475_3.14.75-031475.201608100532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.75/linux-headers-3.14.75-031475-generic_3.14.75-031475.201608100532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 14,
			Patch: 77,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.77/linux-headers-3.14.77-031477_3.14.77-031477.201608200632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.77/linux-headers-3.14.77-031477-generic_3.14.77-031477.201608200632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 14,
			Patch: 78,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.78/linux-headers-3.14.78-031478_3.14.78-031478.201609070333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.78/linux-headers-3.14.78-031478-generic_3.14.78-031478.201609070333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 14,
			Patch: 79,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.79/linux-headers-3.14.79-031479_3.14.79-031479.201609110530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.14.79/linux-headers-3.14.79-031479-generic_3.14.79-031479.201609110530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 37,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.37/linux-headers-3.16.37-031637_3.16.37-031637.201608231033_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.37/linux-headers-3.16.37-031637-generic_3.16.37-031637.201608231033_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 38,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.38/linux-headers-3.16.38-031638_3.16.38-031638.201610210434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.38/linux-headers-3.16.38-031638-generic_3.16.38-031638.201610210434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 39,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.39/linux-headers-3.16.39-031639_3.16.39-031639.201611200534_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.39/linux-headers-3.16.39-031639-generic_3.16.39-031639.201611200534_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 40,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.40/linux-headers-3.16.40-031640_3.16.40-031640.201702230332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.40/linux-headers-3.16.40-031640-generic_3.16.40-031640.201702230332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 41,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.41/linux-headers-3.16.41-031641_3.16.41-031641.201702270232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.41/linux-headers-3.16.41-031641-generic_3.16.41-031641.201702270232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 42,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.42/linux-headers-3.16.42-031642_3.16.42-031642.201703160332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.42/linux-headers-3.16.42-031642-generic_3.16.42-031642.201703160332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 43,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.43/linux-headers-3.16.43-031643_3.16.43-031643.201704050532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.43/linux-headers-3.16.43-031643-generic_3.16.43-031643.201704050532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 44,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.44/linux-headers-3.16.44-031644_3.16.44-031644.201706060532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.44/linux-headers-3.16.44-031644-generic_3.16.44-031644.201706060532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 45,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.45/linux-headers-3.16.45-031645_3.16.45-031645.201707030336_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.45/linux-headers-3.16.45-031645-generic_3.16.45-031645.201707030336_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 46,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.46/linux-headers-3.16.46-031646_3.16.46-031646.201707181533_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.46/linux-headers-3.16.46-031646-generic_3.16.46-031646.201707181533_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 47,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.47/linux-headers-3.16.47-031647_3.16.47-031647.201708260658_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.47/linux-headers-3.16.47-031647-generic_3.16.47-031647.201708260658_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 48,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.48/linux-headers-3.16.48-031648_3.16.48-031648.201709151559_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.48/linux-headers-3.16.48-031648-generic_3.16.48-031648.201709151559_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 49,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.49/linux-headers-3.16.49-031649_3.16.49-031649.201710121231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.49/linux-headers-3.16.49-031649-generic_3.16.49-031649.201710121231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 50,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.50/linux-headers-3.16.50-031650_3.16.50-031650.201711130432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.50/linux-headers-3.16.50-031650-generic_3.16.50-031650.201711130432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 51,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.51/linux-headers-3.16.51-031651_3.16.51-031651.201711261147_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.51/linux-headers-3.16.51-031651-generic_3.16.51-031651.201711261147_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 52,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.52/linux-headers-3.16.52-031652_3.16.52-031652.201801021102_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.52/linux-headers-3.16.52-031652-generic_3.16.52-031652.201801021102_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 53,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.53/linux-headers-3.16.53-031653_3.16.53-031653.201801090931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.53/linux-headers-3.16.53-031653-generic_3.16.53-031653.201801090931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 54,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.54/linux-headers-3.16.54-031654_3.16.54-031654.201802141044_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.54/linux-headers-3.16.54-031654-generic_3.16.54-031654.201802141044_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 55,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.55/linux-headers-3.16.55-031655_3.16.55-031655.201803041545_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.55/linux-headers-3.16.55-031655-generic_3.16.55-031655.201803041545_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 56,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.56/linux-headers-3.16.56-031656_3.16.56-031656.201803191751_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.56/linux-headers-3.16.56-031656-generic_3.16.56-031656.201803191751_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 57,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.57/linux-headers-3.16.57-031657_3.16.57-031657.201806170831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.57/linux-headers-3.16.57-031657-generic_3.16.57-031657.201806170831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 58,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.58/linux-headers-3.16.58-031658_3.16.58-031658.201809270632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.58/linux-headers-3.16.58-031658-generic_3.16.58-031658.201809270632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 59,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.59/linux-headers-3.16.59-031659_3.16.59-031659.201810031232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.59/linux-headers-3.16.59-031659-generic_3.16.59-031659.201810031232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 60,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.60/linux-headers-3.16.60-031660_3.16.60-031660.201810220732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.60/linux-headers-3.16.60-031660-generic_3.16.60-031660.201810220732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 61,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.61/linux-headers-3.16.61-031661_3.16.61-031661.201811210619_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.61/linux-headers-3.16.61-031661-generic_3.16.61-031661.201811210619_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 62,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.62/linux-headers-3.16.62-031662_3.16.62-031662.201812170533_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.62/linux-headers-3.16.62-031662-generic_3.16.62-031662.201812170533_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 63,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.63/linux-headers-3.16.63-031663_3.16.63-031663.201902112033_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.63/linux-headers-3.16.63-031663-generic_3.16.63-031663.201902112033_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 64,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.64/linux-headers-3.16.64-031664_3.16.64-031664.201903252034_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.64/linux-headers-3.16.64-031664-generic_3.16.64-031664.201903252034_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 65,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.65/linux-headers-3.16.65-031665_3.16.65-031665.201904041732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.65/linux-headers-3.16.65-031665-generic_3.16.65-031665.201904041732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 66,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.66/linux-headers-3.16.66-031666_3.16.66-031666.201905030733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.66/linux-headers-3.16.66-031666-generic_3.16.66-031666.201905030733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 67,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.67/linux-headers-3.16.67-031667_3.16.67-031667.201905120732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.67/linux-headers-3.16.67-031667-generic_3.16.67-031667.201905120732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 68,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.68/linux-headers-3.16.68-031668_3.16.68-031668.201906051332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.68/linux-headers-3.16.68-031668-generic_3.16.68-031668.201906051332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 69,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.69/linux-headers-3.16.69-031669_3.16.69-031669.201906201833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.69/linux-headers-3.16.69-031669-generic_3.16.69-031669.201906201833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 70,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.70/linux-headers-3.16.70-031670_3.16.70-031670.201907100927_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.70/linux-headers-3.16.70-031670-generic_3.16.70-031670.201907100927_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 71,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.71/linux-headers-3.16.71-031671_3.16.71-031671.201907240334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.71/linux-headers-3.16.71-031671-generic_3.16.71-031671.201907240334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 72,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.72/linux-headers-3.16.72-031672_3.16.72-031672.201908131119_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.72/linux-headers-3.16.72-031672-generic_3.16.72-031672.201908131119_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 73,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.73/linux-headers-3.16.73-031673_3.16.73-031673.201908201731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.73/linux-headers-3.16.73-031673-generic_3.16.73-031673.201908201731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.74/linux-headers-3.16.74-031674_3.16.74-031674.201909240534_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.74/linux-headers-3.16.74-031674-generic_3.16.74-031674.201909240534_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 75,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.75/linux-headers-3.16.75-031675_3.16.75-031675.201910051235_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.75/linux-headers-3.16.75-031675-generic_3.16.75-031675.201910051235_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 76,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.76/linux-headers-3.16.76-031676_3.16.76-031676.201911010534_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.76/linux-headers-3.16.76-031676-generic_3.16.76-031676.201911010534_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 77,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.77/linux-headers-3.16.77-031677_3.16.77-031677.201911150335_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.77/linux-headers-3.16.77-031677-generic_3.16.77-031677.201911150335_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 78,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.78/linux-headers-3.16.78-031678_3.16.78-031678.201911231034_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.78/linux-headers-3.16.78-031678-generic_3.16.78-031678.201911231034_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 79,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.79/linux-headers-3.16.79-031679_3.16.79-031679.201912101634_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.79/linux-headers-3.16.79-031679-generic_3.16.79-031679.201912101634_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 80,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.80/linux-headers-3.16.80-031680_3.16.80-031680.201912191731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.80/linux-headers-3.16.80-031680-generic_3.16.80-031680.201912191731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 81,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.81/linux-headers-3.16.81-031681_3.16.81-031681.202001110932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.81/linux-headers-3.16.81-031681-generic_3.16.81-031681.202001110932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 16,
			Patch: 82,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.82/linux-headers-3.16.82-031682_3.16.82-031682.202002112136_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.16.82/linux-headers-3.16.82-031682-generic_3.16.82-031682.202002112136_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 37,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.37/linux-headers-3.18.37-031837_3.18.37-031837.201607131932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.37/linux-headers-3.18.37-031837-generic_3.18.37-031837.201607131932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 38,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.38/linux-headers-3.18.38-031838_3.18.38-031838.201607301232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.38/linux-headers-3.18.38-031838-generic_3.18.38-031838.201607301232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 39,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.39/linux-headers-3.18.39-031839_3.18.39-031839.201608091532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.39/linux-headers-3.18.39-031839-generic_3.18.39-031839.201608091532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 40,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.40/linux-headers-3.18.40-031840_3.18.40-031840.201608221832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.40/linux-headers-3.18.40-031840-generic_3.18.40-031840.201608221832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 41,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.41/linux-headers-3.18.41-031841_3.18.41-031841.201609050346_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.41/linux-headers-3.18.41-031841-generic_3.18.41-031841.201609050346_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 42,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.42/linux-headers-3.18.42-031842_3.18.42-031842.201609180430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.42/linux-headers-3.18.42-031842-generic_3.18.42-031842.201609180430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 43,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.43/linux-headers-3.18.43-031843_3.18.43-031843.201610120425_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.43/linux-headers-3.18.43-031843-generic_3.18.43-031843.201610120425_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 44,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.44/linux-headers-3.18.44-031844_3.18.44-031844.201610241522_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.44/linux-headers-3.18.44-031844-generic_3.18.44-031844.201610241522_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 45,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.45/linux-headers-3.18.45-031845_3.18.45-031845.201611300620_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.45/linux-headers-3.18.45-031845-generic_3.18.45-031845.201611300620_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 46,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.46/linux-headers-3.18.46-031846_3.18.46-031846.201612271322_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.46/linux-headers-3.18.46-031846-generic_3.18.46-031846.201612271322_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 47,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.47/linux-headers-3.18.47-031847_3.18.47-031847.201701181631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.47/linux-headers-3.18.47-031847-generic_3.18.47-031847.201701181631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 48,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.48/linux-headers-3.18.48-031848_3.18.48-031848.201702080431_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 49,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.49/linux-headers-3.18.49-031849_3.18.49-031849.201704180352_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.49/linux-headers-3.18.49-031849-generic_3.18.49-031849.201704180352_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 50,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.50/linux-headers-3.18.50-031850_3.18.50-031850.201704220231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.50/linux-headers-3.18.50-031850-generic_3.18.50-031850.201704220231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 51,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.51/linux-headers-3.18.51-031851_3.18.51-031851.201704300047_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.51/linux-headers-3.18.51-031851-generic_3.18.51-031851.201704300047_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 52,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.52/linux-headers-3.18.52-031852_3.18.52-031852.201705080449_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.52/linux-headers-3.18.52-031852-generic_3.18.52-031852.201705080449_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 53,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.53/linux-headers-3.18.53-031853_3.18.53-031853.201705150431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.53/linux-headers-3.18.53-031853-generic_3.18.53-031853.201705150431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 54,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.54/linux-headers-3.18.54-031854_3.18.54-031854.201705201151_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.54/linux-headers-3.18.54-031854-generic_3.18.54-031854.201705201151_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 55,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.55/linux-headers-3.18.55-031855_3.18.55-031855.201705251256_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 56,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.56/linux-headers-3.18.56-031856_3.18.56-031856.201706070753_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.56/linux-headers-3.18.56-031856-generic_3.18.56-031856.201706070753_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 57,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.57/linux-headers-3.18.57-031857_3.18.57-031857.201706141152_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.57/linux-headers-3.18.57-031857-generic_3.18.57-031857.201706141152_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 58,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.58/linux-headers-3.18.58-031858_3.18.58-031858.201706260232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.58/linux-headers-3.18.58-031858-generic_3.18.58-031858.201706260232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 59,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.59/linux-headers-3.18.59-031859_3.18.59-031859.201706290917_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.59/linux-headers-3.18.59-031859-generic_3.18.59-031859.201706290917_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 60,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.60/linux-headers-3.18.60-031860_3.18.60-031860.201707051015_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.60/linux-headers-3.18.60-031860-generic_3.18.60-031860.201707051015_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 61,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.61/linux-headers-3.18.61-031861_3.18.61-031861.201707150531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.61/linux-headers-3.18.61-031861-generic_3.18.61-031861.201707150531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 62,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.62/linux-headers-3.18.62-031862_3.18.62-031862.201707210456_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.62/linux-headers-3.18.62-031862-generic_3.18.62-031862.201707210456_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 63,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.63/linux-headers-3.18.63-031863_3.18.63-031863.201707272009_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.63/linux-headers-3.18.63-031863-generic_3.18.63-031863.201707272009_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 64,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.64/linux-headers-3.18.64-031864_3.18.64-031864.201708111343_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.64/linux-headers-3.18.64-031864-generic_3.18.64-031864.201708111343_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 66,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.66/linux-headers-3.18.66-031866_3.18.66-031866.201708161841_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.66/linux-headers-3.18.66-031866-generic_3.18.66-031866.201708161841_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 67,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.67/linux-headers-3.18.67-031867_3.18.67-031867.201708250059_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.67/linux-headers-3.18.67-031867-generic_3.18.67-031867.201708250059_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 68,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.68/linux-headers-3.18.68-031868_3.18.68-031868.201708300712_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.68/linux-headers-3.18.68-031868-generic_3.18.68-031868.201708300712_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 69,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.69/linux-headers-3.18.69-031869_3.18.69-031869.201709020313_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.69/linux-headers-3.18.69-031869-generic_3.18.69-031869.201709020313_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 70,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.70/linux-headers-3.18.70-031870_3.18.70-031870.201709070441_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.70/linux-headers-3.18.70-031870-generic_3.18.70-031870.201709070441_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 71,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.71/linux-headers-3.18.71-031871_3.18.71-031871.201709132337_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.71/linux-headers-3.18.71-031871-generic_3.18.71-031871.201709132337_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 72,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.72/linux-headers-3.18.72-031872_3.18.72-031872.201709270532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.72/linux-headers-3.18.72-031872-generic_3.18.72-031872.201709270532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 73,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.73/linux-headers-3.18.73-031873_3.18.73-031873.201710050658_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.73/linux-headers-3.18.73-031873-generic_3.18.73-031873.201710050658_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.74/linux-headers-3.18.74-031874_3.18.74-031874.201710092356_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.74/linux-headers-3.18.74-031874-generic_3.18.74-031874.201710092356_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 75,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.75/linux-headers-3.18.75-031875_3.18.75-031875.201710120836_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.75/linux-headers-3.18.75-031875-generic_3.18.75-031875.201710120836_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 76,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.76/linux-headers-3.18.76-031876_3.18.76-031876.201710180902_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.76/linux-headers-3.18.76-031876-generic_3.18.76-031876.201710180902_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 77,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.77/linux-headers-3.18.77-031877_3.18.77-031877.201710211703_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.77/linux-headers-3.18.77-031877-generic_3.18.77-031877.201710211703_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 78,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.78/linux-headers-3.18.78-031878_3.18.78-031878.201710271003_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.78/linux-headers-3.18.78-031878-generic_3.18.78-031878.201710271003_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 80,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.80/linux-headers-3.18.80-031880_3.18.80-031880.201711081112_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.80/linux-headers-3.18.80-031880-generic_3.18.80-031880.201711081112_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 81,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.81/linux-headers-3.18.81-031881_3.18.81-031881.201711151031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.81/linux-headers-3.18.81-031881-generic_3.18.81-031881.201711151031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 82,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.82/linux-headers-3.18.82-031882_3.18.82-031882.201805131757_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.82/linux-headers-3.18.82-031882-generic_3.18.82-031882.201805131757_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 83,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.83/linux-headers-3.18.83-031883_3.18.83-031883.201805132109_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.83/linux-headers-3.18.83-031883-generic_3.18.83-031883.201805132109_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 84,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.84/linux-headers-3.18.84-031884_3.18.84-031884.201711240425_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.84/linux-headers-3.18.84-031884-generic_3.18.84-031884.201711240425_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 85,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.85/linux-headers-3.18.85-031885_3.18.85-031885.201711301031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.85/linux-headers-3.18.85-031885-generic_3.18.85-031885.201711301031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 86,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.86/linux-headers-3.18.86-031886_3.18.86-031886.201712051202_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.86/linux-headers-3.18.86-031886-generic_3.18.86-031886.201712051202_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 87,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.87/linux-headers-3.18.87-031887_3.18.87-031887.201712091831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.87/linux-headers-3.18.87-031887-generic_3.18.87-031887.201712091831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 88,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.88/linux-headers-3.18.88-031888_3.18.88-031888.201712161031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.88/linux-headers-3.18.88-031888-generic_3.18.88-031888.201712161031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 89,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.89/linux-headers-3.18.89-031889_3.18.89-031889.201712201102_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.89/linux-headers-3.18.89-031889-generic_3.18.89-031889.201712201102_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 90,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.90/linux-headers-3.18.90-031890_3.18.90-031890.201712251629_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.90/linux-headers-3.18.90-031890-generic_3.18.90-031890.201712251629_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 91,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.91/linux-headers-3.18.91-031891_3.18.91-031891.201801022233_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.91/linux-headers-3.18.91-031891-generic_3.18.91-031891.201801022233_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 93,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.93/linux-headers-3.18.93-031893_3.18.93-031893.201801311433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.93/linux-headers-3.18.93-031893-generic_3.18.93-031893.201801311433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 94,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.94/linux-headers-3.18.94-031894_3.18.94-031894.201802072309_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.94/linux-headers-3.18.94-031894-generic_3.18.94-031894.201802072309_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 95,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.95/linux-headers-3.18.95-031895_3.18.95-031895.201802162336_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.95/linux-headers-3.18.95-031895-generic_3.18.95-031895.201802162336_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 97,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.97/linux-headers-3.18.97-031897_3.18.97-031897.201802280618_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.97/linux-headers-3.18.97-031897-generic_3.18.97-031897.201802280618_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 98,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.98/linux-headers-3.18.98-031898_3.18.98-031898.201803041521_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.98/linux-headers-3.18.98-031898-generic_3.18.98-031898.201803041521_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 99,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.99/linux-headers-3.18.99-031899_3.18.99-031899.201803111720_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.99/linux-headers-3.18.99-031899-generic_3.18.99-031899.201803111720_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 100,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.100/linux-headers-3.18.100-0318100_3.18.100-0318100.201803180731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.100/linux-headers-3.18.100-0318100-generic_3.18.100-0318100.201803180731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 101,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.101/linux-headers-3.18.101-0318101_3.18.101-0318101.201803221011_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.101/linux-headers-3.18.101-0318101-generic_3.18.101-0318101.201803221011_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 102,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.102/linux-headers-3.18.102-0318102_3.18.102-0318102.201803250949_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.102/linux-headers-3.18.102-0318102-generic_3.18.102-0318102.201803250949_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 103,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.103/linux-headers-3.18.103-0318103_3.18.103-0318103.201804080643_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.103/linux-headers-3.18.103-0318103-generic_3.18.103-0318103.201804080643_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 104,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.104/linux-headers-3.18.104-0318104_3.18.104-0318104.201804100430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.104/linux-headers-3.18.104-0318104-generic_3.18.104-0318104.201804100430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 105,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.105/linux-headers-3.18.105-0318105_3.18.105-0318105.201804131736_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.105/linux-headers-3.18.105-0318105-generic_3.18.105-0318105.201804131736_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 106,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.106/linux-headers-3.18.106-0318106_3.18.106-0318106.201804240526_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.106/linux-headers-3.18.106-0318106-generic_3.18.106-0318106.201804240526_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 107,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.107/linux-headers-3.18.107-0318107_3.18.107-0318107.201804300838_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.107/linux-headers-3.18.107-0318107-generic_3.18.107-0318107.201804300838_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 108,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.108/linux-headers-3.18.108-0318108_3.18.108-0318108.201805021230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.108/linux-headers-3.18.108-0318108-generic_3.18.108-0318108.201805021230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 109,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.109/linux-headers-3.18.109-0318109_3.18.109-0318109.201805160726_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.109/linux-headers-3.18.109-0318109-generic_3.18.109-0318109.201805160726_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 110,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.110/linux-headers-3.18.110-0318110_3.18.110-0318110.201805250938_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.110/linux-headers-3.18.110-0318110-generic_3.18.110-0318110.201805250938_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 111,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.111/linux-headers-3.18.111-0318111_3.18.111-0318111.201805300350_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.111/linux-headers-3.18.111-0318111-generic_3.18.111-0318111.201805300350_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 112,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.112/linux-headers-3.18.112-0318112_3.18.112-0318112.201805302221_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.112/linux-headers-3.18.112-0318112-generic_3.18.112-0318112.201805302221_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 113,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.113/linux-headers-3.18.113-0318113_3.18.113-0318113.201806131532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.113/linux-headers-3.18.113-0318113-generic_3.18.113-0318113.201806131532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 114,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.114/linux-headers-3.18.114-0318114_3.18.114-0318114.201807031319_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.114/linux-headers-3.18.114-0318114-generic_3.18.114-0318114.201807031319_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 115,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.115/linux-headers-3.18.115-0318115_3.18.115-0318115.201807111039_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.115/linux-headers-3.18.115-0318115-generic_3.18.115-0318115.201807111039_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 116,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.116/linux-headers-3.18.116-0318116_3.18.116-0318116.201807221231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.116/linux-headers-3.18.116-0318116-generic_3.18.116-0318116.201807221231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 117,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.117/linux-headers-3.18.117-0318117_3.18.117-0318117.201807280642_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.117/linux-headers-3.18.117-0318117-generic_3.18.117-0318117.201807280642_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 118,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.118/linux-headers-3.18.118-0318118_3.18.118-0318118.201808091240_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.118/linux-headers-3.18.118-0318118-generic_3.18.118-0318118.201808091240_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 119,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.119/linux-headers-3.18.119-0318119_3.18.119-0318119.201808171540_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.119/linux-headers-3.18.119-0318119-generic_3.18.119-0318119.201808171540_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 120,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.120/linux-headers-3.18.120-0318120_3.18.120-0318120.201808280231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.120/linux-headers-3.18.120-0318120-generic_3.18.120-0318120.201808280231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 121,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.121/linux-headers-3.18.121-0318121_3.18.121-0318121.201809050938_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.121/linux-headers-3.18.121-0318121-generic_3.18.121-0318121.201809050938_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 122,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.122/linux-headers-3.18.122-0318122_3.18.122-0318122.201809091618_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.122/linux-headers-3.18.122-0318122-generic_3.18.122-0318122.201809091618_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 123,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.123/linux-headers-3.18.123-0318123_3.18.123-0318123.201809260843_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.123/linux-headers-3.18.123-0318123-generic_3.18.123-0318123.201809260843_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 124,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.124/linux-headers-3.18.124-0318124_3.18.124-0318124.201810130934_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.124/linux-headers-3.18.124-0318124-generic_3.18.124-0318124.201810130934_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 125,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.125/linux-headers-3.18.125-0318125_3.18.125-0318125.201811101736_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.125/linux-headers-3.18.125-0318125-generic_3.18.125-0318125.201811101736_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 126,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.126/linux-headers-3.18.126-0318126_3.18.126-0318126.201811220732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.126/linux-headers-3.18.126-0318126-generic_3.18.126-0318126.201811220732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 127,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.127/linux-headers-3.18.127-0318127_3.18.127-0318127.201811271243_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.127/linux-headers-3.18.127-0318127-generic_3.18.127-0318127.201811271243_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 128,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.128/linux-headers-3.18.128-0318128_3.18.128-0318128.201812010548_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.128/linux-headers-3.18.128-0318128-generic_3.18.128-0318128.201812010548_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 129,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.129/linux-headers-3.18.129-0318129_3.18.129-0318129.201812131051_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.129/linux-headers-3.18.129-0318129-generic_3.18.129-0318129.201812131051_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 130,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.130/linux-headers-3.18.130-0318130_3.18.130-0318130.201812171049_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.130/linux-headers-3.18.130-0318130-generic_3.18.130-0318130.201812171049_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 131,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.131/linux-headers-3.18.131-0318131_3.18.131-0318131.201812210939_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.131/linux-headers-3.18.131-0318131-generic_3.18.131-0318131.201812210939_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 132,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.132/linux-headers-3.18.132-0318132_3.18.132-0318132.201901130632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.132/linux-headers-3.18.132-0318132-generic_3.18.132-0318132.201901130632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 133,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.133/linux-headers-3.18.133-0318133_3.18.133-0318133.201901260555_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.133/linux-headers-3.18.133-0318133-generic_3.18.133-0318133.201901260555_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 134,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.134/linux-headers-3.18.134-0318134_3.18.134-0318134.201902061410_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.134/linux-headers-3.18.134-0318134-generic_3.18.134-0318134.201902061410_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 135,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.135/linux-headers-3.18.135-0318135_3.18.135-0318135.201902200737_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.135/linux-headers-3.18.135-0318135-generic_3.18.135-0318135.201902200737_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 136,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.136/linux-headers-3.18.136-0318136_3.18.136-0318136.201902230608_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.136/linux-headers-3.18.136-0318136-generic_3.18.136-0318136.201902230608_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 137,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.137/linux-headers-3.18.137-0318137_3.18.137-0318137.201903230434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.137/linux-headers-3.18.137-0318137-generic_3.18.137-0318137.201903230434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 138,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.138/linux-headers-3.18.138-0318138_3.18.138-0318138.201904030310_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.138/linux-headers-3.18.138-0318138-generic_3.18.138-0318138.201904030310_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 139,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.139/linux-headers-3.18.139-0318139_3.18.139-0318139.201904270612_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.139/linux-headers-3.18.139-0318139-generic_3.18.139-0318139.201904270612_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 3,
			Minor: 18,
			Patch: 140,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.140/linux-headers-3.18.140-0318140_3.18.140-0318140.201905160849_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v3.18.140/linux-headers-3.18.140-0318140-generic_3.18.140-0318140.201905160849_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 28,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.28/linux-headers-4.1.28-040128_4.1.28-040128.201607131931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.28/linux-headers-4.1.28-040128-generic_4.1.28-040128.201607131931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 29,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.29/linux-headers-4.1.29-040129_4.1.29-040129.201607301232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.29/linux-headers-4.1.29-040129-generic_4.1.29-040129.201607301232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 30,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.30/linux-headers-4.1.30-040130_4.1.30-040130.201608091532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.30/linux-headers-4.1.30-040130-generic_4.1.30-040130.201608091532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 31,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.31/linux-headers-4.1.31-040131_4.1.31-040131.201608221832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.31/linux-headers-4.1.31-040131-generic_4.1.31-040131.201608221832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 32,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.32/linux-headers-4.1.32-040132_4.1.32-040132.201609050346_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.32/linux-headers-4.1.32-040132-generic_4.1.32-040132.201609050346_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 33,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.33/linux-headers-4.1.33-040133_4.1.33-040133.201609180430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.33/linux-headers-4.1.33-040133-generic_4.1.33-040133.201609180430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 34,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.34/linux-headers-4.1.34-040134_4.1.34-040134.201610120331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.34/linux-headers-4.1.34-040134-generic_4.1.34-040134.201610120331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 35,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.35/linux-headers-4.1.35-040135_4.1.35-040135.201610241431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.35/linux-headers-4.1.35-040135-generic_4.1.35-040135.201610241431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 36,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.36/linux-headers-4.1.36-040136_4.1.36-040136.201611300531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.36/linux-headers-4.1.36-040136-generic_4.1.36-040136.201611300531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 37,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.37/linux-headers-4.1.37-040137_4.1.37-040137.201612271231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.37/linux-headers-4.1.37-040137-generic_4.1.37-040137.201612271231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 38,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.38/linux-headers-4.1.38-040138_4.1.38-040138.201701181632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.38/linux-headers-4.1.38-040138-generic_4.1.38-040138.201701181632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 39,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.39/linux-headers-4.1.39-040139_4.1.39-040139.201703131937_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.39/linux-headers-4.1.39-040139-generic_4.1.39-040139.201703131937_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 40,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.40/linux-headers-4.1.40-040140_4.1.40-040140.201705290739_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.40/linux-headers-4.1.40-040140-generic_4.1.40-040140.201705290739_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 41,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.41/linux-headers-4.1.41-040141_4.1.41-040141.201706151642_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.41/linux-headers-4.1.41-040141-generic_4.1.41-040141.201706151642_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 42,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.42/linux-headers-4.1.42-040142_4.1.42-040142.201706291340_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.42/linux-headers-4.1.42-040142-generic_4.1.42-040142.201706291340_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 43,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.43/linux-headers-4.1.43-040143_4.1.43-040143.201708060041_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.43/linux-headers-4.1.43-040143-generic_4.1.43-040143.201708060041_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 44,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.44/linux-headers-4.1.44-040144_4.1.44-040144.201709190942_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.44/linux-headers-4.1.44-040144-generic_4.1.44-040144.201709190942_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 45,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.45/linux-headers-4.1.45-040145_4.1.45-040145.201906051658_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.45/linux-headers-4.1.45-040145-generic_4.1.45-040145.201906051658_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 46,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.46/linux-headers-4.1.46-040146_4.1.46-040146.201906051259_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.46/linux-headers-4.1.46-040146-generic_4.1.46-040146.201906051259_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 47,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.47/linux-headers-4.1.47-040147_4.1.47-040147.201712071344_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.47/linux-headers-4.1.47-040147-generic_4.1.47-040147.201712071344_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 48,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.48/linux-headers-4.1.48-040148_4.1.48-040148.201712151631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.48/linux-headers-4.1.48-040148-generic_4.1.48-040148.201712151631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 49,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.49/linux-headers-4.1.49-040149_4.1.49-040149.201906051326_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.49/linux-headers-4.1.49-040149-generic_4.1.49-040149.201906051326_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 50,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.50/linux-headers-4.1.50-040150_4.1.50-040150.201906052210_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.50/linux-headers-4.1.50-040150-generic_4.1.50-040150.201906052210_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 51,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.51/linux-headers-4.1.51-040151_4.1.51-040151.201906091239_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.51/linux-headers-4.1.51-040151-generic_4.1.51-040151.201906091239_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 1,
			Patch: 52,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.52/linux-headers-4.1.52-040152_4.1.52-040152.201805290846_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.1.52/linux-headers-4.1.52-040152-generic_4.1.52-040152.201805290846_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.15/linux-headers-4.4.15-040415_4.4.15-040415.201607111333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.15/linux-headers-4.4.15-040415-generic_4.4.15-040415.201607111333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.16/linux-headers-4.4.16-040416_4.4.16-040416.201607271333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.16/linux-headers-4.4.16-040416-generic_4.4.16-040416.201607271333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.17/linux-headers-4.4.17-040417_4.4.17-040417.201608100634_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.17/linux-headers-4.4.17-040417-generic_4.4.17-040417.201608100634_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.19/linux-headers-4.4.19-040419_4.4.19-040419.201608201335_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.19/linux-headers-4.4.19-040419-generic_4.4.19-040419.201608201335_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.20/linux-headers-4.4.20-040420_4.4.20-040420.201609070334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.20/linux-headers-4.4.20-040420-generic_4.4.20-040420.201609070334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.21/linux-headers-4.4.21-040421_4.4.21-040421.201609150330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.21/linux-headers-4.4.21-040421-generic_4.4.21-040421.201609150330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 22,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.22/linux-headers-4.4.22-040422_4.4.22-040422.201609240532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.22/linux-headers-4.4.22-040422-generic_4.4.22-040422.201609240532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 23,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.23/linux-headers-4.4.23-040423_4.4.23-040423.201609300709_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.23/linux-headers-4.4.23-040423-generic_4.4.23-040423.201609300709_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 24,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.24/linux-headers-4.4.24-040424_4.4.24-040424.201610071138_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.24/linux-headers-4.4.24-040424-generic_4.4.24-040424.201610071138_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 25,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.25/linux-headers-4.4.25-040425_4.4.25-040425.201610161231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.25/linux-headers-4.4.25-040425-generic_4.4.25-040425.201610161231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 26,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.26/linux-headers-4.4.26-040426_4.4.26-040426.201610200745_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.26/linux-headers-4.4.26-040426-generic_4.4.26-040426.201610200745_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 27,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.27/linux-headers-4.4.27-040427_4.4.27-040427.201610220949_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.27/linux-headers-4.4.27-040427-generic_4.4.27-040427.201610220949_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 28,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.28/linux-headers-4.4.28-040428_4.4.28-040428.201610280549_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.28/linux-headers-4.4.28-040428-generic_4.4.28-040428.201610280549_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 29,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.29/linux-headers-4.4.29-040429_4.4.29-040429.201610310948_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.29/linux-headers-4.4.29-040429-generic_4.4.29-040429.201610310948_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 30,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.30/linux-headers-4.4.30-040430_4.4.30-040430.201611010007_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.30/linux-headers-4.4.30-040430-generic_4.4.30-040430.201611010007_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 31,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.31/linux-headers-4.4.31-040431_4.4.31-040431.201611101249_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.31/linux-headers-4.4.31-040431-generic_4.4.31-040431.201611101249_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 32,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.32/linux-headers-4.4.32-040432_4.4.32-040432.201611150347_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.32/linux-headers-4.4.32-040432-generic_4.4.32-040432.201611150347_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 33,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.33/linux-headers-4.4.33-040433_4.4.33-040433.201611180742_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.33/linux-headers-4.4.33-040433-generic_4.4.33-040433.201611180742_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 34,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.34/linux-headers-4.4.34-040434_4.4.34-040434.201611210641_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.34/linux-headers-4.4.34-040434-generic_4.4.34-040434.201611210641_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 35,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.35/linux-headers-4.4.35-040435_4.4.35-040435.201611260431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.35/linux-headers-4.4.35-040435-generic_4.4.35-040435.201611260431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 36,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.36/linux-headers-4.4.36-040436_4.4.36-040436.201612020543_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.36/linux-headers-4.4.36-040436-generic_4.4.36-040436.201612020543_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 37,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.37/linux-headers-4.4.37-040437_4.4.37-040437.201612081650_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.37/linux-headers-4.4.37-040437-generic_4.4.37-040437.201612081650_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 38,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.38/linux-headers-4.4.38-040438_4.4.38-040438.201612101547_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.38/linux-headers-4.4.38-040438-generic_4.4.38-040438.201612101547_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 39,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.39/linux-headers-4.4.39-040439_4.4.39-040439.201612151346_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.39/linux-headers-4.4.39-040439-generic_4.4.39-040439.201612151346_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 40,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.40/linux-headers-4.4.40-040440_4.4.40-040440.201701060808_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.40/linux-headers-4.4.40-040440-generic_4.4.40-040440.201701060808_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 41,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.41/linux-headers-4.4.41-040441_4.4.41-040441.201701090549_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.41/linux-headers-4.4.41-040441-generic_4.4.41-040441.201701090549_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 42,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.42/linux-headers-4.4.42-040442_4.4.42-040442.201701120631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.42/linux-headers-4.4.42-040442-generic_4.4.42-040442.201701120631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 43,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.43/linux-headers-4.4.43-040443_4.4.43-040443.201701150831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.43/linux-headers-4.4.43-040443-generic_4.4.43-040443.201701150831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 44,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.44/linux-headers-4.4.44-040444_4.4.44-040444.201701200532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.44/linux-headers-4.4.44-040444-generic_4.4.44-040444.201701200532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 45,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.45/linux-headers-4.4.45-040445_4.4.45-040445.201701260331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.45/linux-headers-4.4.45-040445-generic_4.4.45-040445.201701260331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 46,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.46/linux-headers-4.4.46-040446_4.4.46-040446.201702010331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.46/linux-headers-4.4.46-040446-generic_4.4.46-040446.201702010331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 47,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.47/linux-headers-4.4.47-040447_4.4.47-040447.201702040433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.47/linux-headers-4.4.47-040447-generic_4.4.47-040447.201702040433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 48,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.48/linux-headers-4.4.48-040448_4.4.48-040448.201702090334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.48/linux-headers-4.4.48-040448-generic_4.4.48-040448.201702090334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 49,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.49/linux-headers-4.4.49-040449_4.4.49-040449.201702141931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.49/linux-headers-4.4.49-040449-generic_4.4.49-040449.201702141931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 50,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.50/linux-headers-4.4.50-040450_4.4.50-040450.201702181144_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.50/linux-headers-4.4.50-040450-generic_4.4.50-040450.201702181144_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 51,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.51/linux-headers-4.4.51-040451_4.4.51-040451.201702231231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.51/linux-headers-4.4.51-040451-generic_4.4.51-040451.201702231231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 52,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.52/linux-headers-4.4.52-040452_4.4.52-040452.201702260631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.52/linux-headers-4.4.52-040452-generic_4.4.52-040452.201702260631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 53,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.53/linux-headers-4.4.53-040453_4.4.53-040453.201703120131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.53/linux-headers-4.4.53-040453-generic_4.4.53-040453.201703120131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 54,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.54/linux-headers-4.4.54-040454_4.4.54-040454.201703142331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.54/linux-headers-4.4.54-040454-generic_4.4.54-040454.201703142331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 55,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.55/linux-headers-4.4.55-040455_4.4.55-040455.201703180831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.55/linux-headers-4.4.55-040455-generic_4.4.55-040455.201703180831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 56,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.56/linux-headers-4.4.56-040456_4.4.56-040456.201703231535_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.56/linux-headers-4.4.56-040456-generic_4.4.56-040456.201703231535_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 57,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.57/linux-headers-4.4.57-040457_4.4.57-040457.201703260734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.57/linux-headers-4.4.57-040457-generic_4.4.57-040457.201703260734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 58,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.58/linux-headers-4.4.58-040458_4.4.58-040458.201703300433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.58/linux-headers-4.4.58-040458-generic_4.4.58-040458.201703300433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 59,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.59/linux-headers-4.4.59-040459_4.4.59-040459.201703310531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.59/linux-headers-4.4.59-040459-generic_4.4.59-040459.201703310531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 60,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.60/linux-headers-4.4.60-040460_4.4.60-040460.201704080434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.60/linux-headers-4.4.60-040460-generic_4.4.60-040460.201704080434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 61,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.61/linux-headers-4.4.61-040461_4.4.61-040461.201704120732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.61/linux-headers-4.4.61-040461-generic_4.4.61-040461.201704120732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 62,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.62/linux-headers-4.4.62-040462_4.4.62-040462.201704180232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.62/linux-headers-4.4.62-040462-generic_4.4.62-040462.201704180232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 63,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.63/linux-headers-4.4.63-040463_4.4.63-040463.201704210431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.63/linux-headers-4.4.63-040463-generic_4.4.63-040463.201704210431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 64,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.64/linux-headers-4.4.64-040464_4.4.64-040464.201704270431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.64/linux-headers-4.4.64-040464-generic_4.4.64-040464.201704270431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 65,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.65/linux-headers-4.4.65-040465_4.4.65-040465.201704300046_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.65/linux-headers-4.4.65-040465-generic_4.4.65-040465.201704300046_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 66,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.66/linux-headers-4.4.66-040466_4.4.66-040466.201705031344_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.66/linux-headers-4.4.66-040466-generic_4.4.66-040466.201705031344_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 67,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.67/linux-headers-4.4.67-040467_4.4.67-040467.201705080411_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.67/linux-headers-4.4.67-040467-generic_4.4.67-040467.201705080411_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 68,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.68/linux-headers-4.4.68-040468_4.4.68-040468.201705140831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.68/linux-headers-4.4.68-040468-generic_4.4.68-040468.201705140831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 69,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.69/linux-headers-4.4.69-040469_4.4.69-040469.201705200931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.69/linux-headers-4.4.69-040469-generic_4.4.69-040469.201705200931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 70,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.70/linux-headers-4.4.70-040470_4.4.70-040470.201705251131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.70/linux-headers-4.4.70-040470-generic_4.4.70-040470.201705251131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 71,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.71/linux-headers-4.4.71-040471_4.4.71-040471.201706070735_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.71/linux-headers-4.4.71-040471-generic_4.4.71-040471.201706070735_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 72,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.72/linux-headers-4.4.72-040472_4.4.72-040472.201706141033_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.72/linux-headers-4.4.72-040472-generic_4.4.72-040472.201706141033_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 73,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.73/linux-headers-4.4.73-040473_4.4.73-040473.201706170414_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.73/linux-headers-4.4.73-040473-generic_4.4.73-040473.201706170414_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.74/linux-headers-4.4.74-040474_4.4.74-040474.201706260232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.74/linux-headers-4.4.74-040474-generic_4.4.74-040474.201706260232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 75,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.75/linux-headers-4.4.75-040475_4.4.75-040475.201706290836_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.75/linux-headers-4.4.75-040475-generic_4.4.75-040475.201706290836_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 76,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.76/linux-headers-4.4.76-040476_4.4.76-040476.201707050936_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.76/linux-headers-4.4.76-040476-generic_4.4.76-040476.201707050936_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 77,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.77/linux-headers-4.4.77-040477_4.4.77-040477.201707150734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.77/linux-headers-4.4.77-040477-generic_4.4.77-040477.201707150734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 78,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.78/linux-headers-4.4.78-040478_4.4.78-040478.201707210237_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.78/linux-headers-4.4.78-040478-generic_4.4.78-040478.201707210237_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 79,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.79/linux-headers-4.4.79-040479_4.4.79-040479.201707271932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.79/linux-headers-4.4.79-040479-generic_4.4.79-040479.201707271932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 80,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.80/linux-headers-4.4.80-040480_4.4.80-040480.201708062341_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.80/linux-headers-4.4.80-040480-generic_4.4.80-040480.201708062341_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 81,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.81/linux-headers-4.4.81-040481_4.4.81-040481.201708111333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.81/linux-headers-4.4.81-040481-generic_4.4.81-040481.201708111333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 83,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.83/linux-headers-4.4.83-040483_4.4.83-040483.201708161731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.83/linux-headers-4.4.83-040483-generic_4.4.83-040483.201708161731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 84,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.84/linux-headers-4.4.84-040484_4.4.84-040484.201708242241_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.84/linux-headers-4.4.84-040484-generic_4.4.84-040484.201708242241_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 85,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.85/linux-headers-4.4.85-040485_4.4.85-040485.201708300532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.85/linux-headers-4.4.85-040485-generic_4.4.85-040485.201708300532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 86,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.86/linux-headers-4.4.86-040486_4.4.86-040486.201709020231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.86/linux-headers-4.4.86-040486-generic_4.4.86-040486.201709020231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 87,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.87/linux-headers-4.4.87-040487_4.4.87-040487.201709070334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.87/linux-headers-4.4.87-040487-generic_4.4.87-040487.201709070334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 88,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.88/linux-headers-4.4.88-040488_4.4.88-040488.201709131950_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.88/linux-headers-4.4.88-040488-generic_4.4.88-040488.201709131950_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 89,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.89/linux-headers-4.4.89-040489_4.4.89-040489.201709270634_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.89/linux-headers-4.4.89-040489-generic_4.4.89-040489.201709270634_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 90,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.90/linux-headers-4.4.90-040490_4.4.90-040490.201710050517_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.90/linux-headers-4.4.90-040490-generic_4.4.90-040490.201710050517_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 91,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.91/linux-headers-4.4.91-040491_4.4.91-040491.201710092313_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.91/linux-headers-4.4.91-040491-generic_4.4.91-040491.201710092313_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 92,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.92/linux-headers-4.4.92-040492_4.4.92-040492.201710120634_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.92/linux-headers-4.4.92-040492-generic_4.4.92-040492.201710120634_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 93,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.93/linux-headers-4.4.93-040493_4.4.93-040493.201710180435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.93/linux-headers-4.4.93-040493-generic_4.4.93-040493.201710180435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 94,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.94/linux-headers-4.4.94-040494_4.4.94-040494.201710211231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.94/linux-headers-4.4.94-040494-generic_4.4.94-040494.201710211231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 95,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.95/linux-headers-4.4.95-040495_4.4.95-040495.201710270532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.95/linux-headers-4.4.95-040495-generic_4.4.95-040495.201710270532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 96,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.96/linux-headers-4.4.96-040496_4.4.96-040496.201711020932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.96/linux-headers-4.4.96-040496-generic_4.4.96-040496.201711020932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 97,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.97/linux-headers-4.4.97-040497_4.4.97-040497.201711081042_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.97/linux-headers-4.4.97-040497-generic_4.4.97-040497.201711081042_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 98,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.98/linux-headers-4.4.98-040498_4.4.98-040498.201711151231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.98/linux-headers-4.4.98-040498-generic_4.4.98-040498.201711151231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 99,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.99/linux-headers-4.4.99-040499_4.4.99-040499.201805131725_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.99/linux-headers-4.4.99-040499-generic_4.4.99-040499.201805131725_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 100,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.100/linux-headers-4.4.100-0404100_4.4.100-0404100.201805132143_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.100/linux-headers-4.4.100-0404100-generic_4.4.100-0404100.201805132143_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 101,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.101/linux-headers-4.4.101-0404101_4.4.101-0404101.201711240830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.101/linux-headers-4.4.101-0404101-generic_4.4.101-0404101.201711240830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 102,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.102/linux-headers-4.4.102-0404102_4.4.102-0404102.201711240631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.102/linux-headers-4.4.102-0404102-generic_4.4.102-0404102.201711240631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 103,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.103/linux-headers-4.4.103-0404103_4.4.103-0404103.201711300432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.103/linux-headers-4.4.103-0404103-generic_4.4.103-0404103.201711300432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 104,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.104/linux-headers-4.4.104-0404104_4.4.104-0404104.201712051130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.104/linux-headers-4.4.104-0404104-generic_4.4.104-0404104.201712051130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 105,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.105/linux-headers-4.4.105-0404105_4.4.105-0404105.201712091332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.105/linux-headers-4.4.105-0404105-generic_4.4.105-0404105.201712091332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 106,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.106/linux-headers-4.4.106-0404106_4.4.106-0404106.201712160531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.106/linux-headers-4.4.106-0404106-generic_4.4.106-0404106.201712160531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 107,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.107/linux-headers-4.4.107-0404107_4.4.107-0404107.201712201032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.107/linux-headers-4.4.107-0404107-generic_4.4.107-0404107.201712201032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 108,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.108/linux-headers-4.4.108-0404108_4.4.108-0404108.201712251510_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.108/linux-headers-4.4.108-0404108-generic_4.4.108-0404108.201712251510_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 109,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.109/linux-headers-4.4.109-0404109_4.4.109-0404109.201801022112_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.109/linux-headers-4.4.109-0404109-generic_4.4.109-0404109.201801022112_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 110,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.110/linux-headers-4.4.110-0404110_4.4.110-0404110.201801051613_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.110/linux-headers-4.4.110-0404110-generic_4.4.110-0404110.201801051613_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 111,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.111/linux-headers-4.4.111-0404111_4.4.111-0404111.201801100931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.111/linux-headers-4.4.111-0404111-generic_4.4.111-0404111.201801100931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 113,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.113/linux-headers-4.4.113-0404113_4.4.113-0404113.201801231435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.113/linux-headers-4.4.113-0404113-generic_4.4.113-0404113.201801231435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 114,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.114/linux-headers-4.4.114-0404114_4.4.114-0404114.201801310736_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.114/linux-headers-4.4.114-0404114-generic_4.4.114-0404114.201801310736_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 115,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.115/linux-headers-4.4.115-0404115_4.4.115-0404115.201802031230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.115/linux-headers-4.4.115-0404115-generic_4.4.115-0404115.201802031230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 116,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.116/linux-headers-4.4.116-0404116_4.4.116-0404116.201802162136_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.116/linux-headers-4.4.116-0404116-generic_4.4.116-0404116.201802162136_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 117,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.117/linux-headers-4.4.117-0404117_4.4.117-0404117.201802261801_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.117/linux-headers-4.4.117-0404117-generic_4.4.117-0404117.201802261801_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 119,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.119/linux-headers-4.4.119-0404119_4.4.119-0404119.201802281030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.119/linux-headers-4.4.119-0404119-generic_4.4.119-0404119.201802281030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 120,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.120/linux-headers-4.4.120-0404120_4.4.120-0404120.201803040932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.120/linux-headers-4.4.120-0404120-generic_4.4.120-0404120.201803040932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 121,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.121/linux-headers-4.4.121-0404121_4.4.121-0404121.201803111232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.121/linux-headers-4.4.121-0404121-generic_4.4.121-0404121.201803111232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 122,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.122/linux-headers-4.4.122-0404122_4.4.122-0404122.201803180732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.122/linux-headers-4.4.122-0404122-generic_4.4.122-0404122.201803180732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 123,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.123/linux-headers-4.4.123-0404123_4.4.123-0404123.201803220531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.123/linux-headers-4.4.123-0404123-generic_4.4.123-0404123.201803220531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 124,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.124/linux-headers-4.4.124-0404124_4.4.124-0404124.201803250430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.124/linux-headers-4.4.124-0404124-generic_4.4.124-0404124.201803250430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 125,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.125/linux-headers-4.4.125-0404125_4.4.125-0404125.201803281340_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.125/linux-headers-4.4.125-0404125-generic_4.4.125-0404125.201803281340_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 126,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.126/linux-headers-4.4.126-0404126_4.4.126-0404126.201803311331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.126/linux-headers-4.4.126-0404126-generic_4.4.126-0404126.201803311331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 127,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.127/linux-headers-4.4.127-0404127_4.4.127-0404127.201804081036_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.127/linux-headers-4.4.127-0404127-generic_4.4.127-0404127.201804081036_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 128,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.128/linux-headers-4.4.128-0404128_4.4.128-0404128.201804131733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.128/linux-headers-4.4.128-0404128-generic_4.4.128-0404128.201804131733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 129,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.129/linux-headers-4.4.129-0404129_4.4.129-0404129.201804240833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.129/linux-headers-4.4.129-0404129-generic_4.4.129-0404129.201804240833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 130,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.130/linux-headers-4.4.130-0404130_4.4.130-0404130.201805131657_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.130/linux-headers-4.4.130-0404130-generic_4.4.130-0404130.201805131657_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 131,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.131/linux-headers-4.4.131-0404131_4.4.131-0404131.201805021136_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.131/linux-headers-4.4.131-0404131-generic_4.4.131-0404131.201805021136_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 132,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.132/linux-headers-4.4.132-0404132_4.4.132-0404132.201805211041_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.132/linux-headers-4.4.132-0404132-generic_4.4.132-0404132.201805211041_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 133,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.133/linux-headers-4.4.133-0404133_4.4.133-0404133.201805260734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.133/linux-headers-4.4.133-0404133-generic_4.4.133-0404133.201805260734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 134,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.134/linux-headers-4.4.134-0404134_4.4.134-0404134.201805300240_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.134/linux-headers-4.4.134-0404134-generic_4.4.134-0404134.201805300240_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 135,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.135/linux-headers-4.4.135-0404135_4.4.135-0404135.201805301732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.135/linux-headers-4.4.135-0404135-generic_4.4.135-0404135.201805301732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 136,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.136/linux-headers-4.4.136-0404136_4.4.136-0404136.201806061537_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.136/linux-headers-4.4.136-0404136-generic_4.4.136-0404136.201806061537_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 137,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.137/linux-headers-4.4.137-0404137_4.4.137-0404137.201806131133_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.137/linux-headers-4.4.137-0404137-generic_4.4.137-0404137.201806131133_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 138,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.138/linux-headers-4.4.138-0404138_4.4.138-0404138.201806160436_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.138/linux-headers-4.4.138-0404138-generic_4.4.138-0404138.201806160436_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 139,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.139/linux-headers-4.4.139-0404139_4.4.139-0404139.201807030736_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.139/linux-headers-4.4.139-0404139-generic_4.4.139-0404139.201807030736_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 140,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.140/linux-headers-4.4.140-0404140_4.4.140-0404140.201807111137_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.140/linux-headers-4.4.140-0404140-generic_4.4.140-0404140.201807111137_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 141,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.141/linux-headers-4.4.141-0404141_4.4.141-0404141.201807170633_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.141/linux-headers-4.4.141-0404141-generic_4.4.141-0404141.201807170633_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 142,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.142/linux-headers-4.4.142-0404142_4.4.142-0404142.201807191435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.142/linux-headers-4.4.142-0404142-generic_4.4.142-0404142.201807191435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 143,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.143/linux-headers-4.4.143-0404143_4.4.143-0404143.201807221334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.143/linux-headers-4.4.143-0404143-generic_4.4.143-0404143.201807221334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 144,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.144/linux-headers-4.4.144-0404144_4.4.144-0404144.201807261050_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.144/linux-headers-4.4.144-0404144-generic_4.4.144-0404144.201807261050_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 145,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.145/linux-headers-4.4.145-0404145_4.4.145-0404145.201807280852_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.145/linux-headers-4.4.145-0404145-generic_4.4.145-0404145.201807280852_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 146,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.146/linux-headers-4.4.146-0404146_4.4.146-0404146.201808061143_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.146/linux-headers-4.4.146-0404146-generic_4.4.146-0404146.201808061143_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 147,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.147/linux-headers-4.4.147-0404147_4.4.147-0404147.201808090732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.147/linux-headers-4.4.147-0404147-generic_4.4.147-0404147.201808090732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 148,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.148/linux-headers-4.4.148-0404148_4.4.148-0404148.201808151240_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.148/linux-headers-4.4.148-0404148-generic_4.4.148-0404148.201808151240_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 149,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.149/linux-headers-4.4.149-0404149_4.4.149-0404149.201808171532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.149/linux-headers-4.4.149-0404149-generic_4.4.149-0404149.201808171532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 150,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.150/linux-headers-4.4.150-0404150_4.4.150-0404150.201808180530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.150/linux-headers-4.4.150-0404150-generic_4.4.150-0404150.201808180530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 151,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.151/linux-headers-4.4.151-0404151_4.4.151-0404151.201808220242_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.151/linux-headers-4.4.151-0404151-generic_4.4.151-0404151.201808220242_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 152,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.152/linux-headers-4.4.152-0404152_4.4.152-0404152.201808241240_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.152/linux-headers-4.4.152-0404152-generic_4.4.152-0404152.201808241240_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 153,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.153/linux-headers-4.4.153-0404153_4.4.153-0404153.201808280632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.153/linux-headers-4.4.153-0404153-generic_4.4.153-0404153.201808280632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 154,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.154/linux-headers-4.4.154-0404154_4.4.154-0404154.201809050440_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.154/linux-headers-4.4.154-0404154-generic_4.4.154-0404154.201809050440_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 155,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.155/linux-headers-4.4.155-0404155_4.4.155-0404155.201809091944_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.155/linux-headers-4.4.155-0404155-generic_4.4.155-0404155.201809091944_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 156,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.156/linux-headers-4.4.156-0404156_4.4.156-0404156.201809150434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.156/linux-headers-4.4.156-0404156-generic_4.4.156-0404156.201809150434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 157,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.157/linux-headers-4.4.157-0404157_4.4.157-0404157.201809191751_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.157/linux-headers-4.4.157-0404157-generic_4.4.157-0404157.201809191751_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 158,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.158/linux-headers-4.4.158-0404158_4.4.158-0404158.201809260355_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.158/linux-headers-4.4.158-0404158-generic_4.4.158-0404158.201809260355_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 159,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.159/linux-headers-4.4.159-0404159_4.4.159-0404159.201809290743_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.159/linux-headers-4.4.159-0404159-generic_4.4.159-0404159.201809290743_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 160,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.160/linux-headers-4.4.160-0404160_4.4.160-0404160.201810100335_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.160/linux-headers-4.4.160-0404160-generic_4.4.160-0404160.201810100335_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 161,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.161/linux-headers-4.4.161-0404161_4.4.161-0404161.201810130434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.161/linux-headers-4.4.161-0404161-generic_4.4.161-0404161.201810130434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 162,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.162/linux-headers-4.4.162-0404162_4.4.162-0404162.201810200432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.162/linux-headers-4.4.162-0404162-generic_4.4.162-0404162.201810200432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 163,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.163/linux-headers-4.4.163-0404163_4.4.163-0404163.201811101146_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.163/linux-headers-4.4.163-0404163-generic_4.4.163-0404163.201811101146_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 164,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.164/linux-headers-4.4.164-0404164_4.4.164-0404164.201811210507_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.164/linux-headers-4.4.164-0404164-generic_4.4.164-0404164.201811210507_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 165,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.165/linux-headers-4.4.165-0404165_4.4.165-0404165.201811271134_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.165/linux-headers-4.4.165-0404165-generic_4.4.165-0404165.201811271134_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 166,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.166/linux-headers-4.4.166-0404166_4.4.166-0404166.201812010431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.166/linux-headers-4.4.166-0404166-generic_4.4.166-0404166.201812010431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 167,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.167/linux-headers-4.4.167-0404167_4.4.167-0404167.201812130444_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.167/linux-headers-4.4.167-0404167-generic_4.4.167-0404167.201812130444_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 168,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.168/linux-headers-4.4.168-0404168_4.4.168-0404168.201812172143_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.168/linux-headers-4.4.168-0404168-generic_4.4.168-0404168.201812172143_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 169,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.169/linux-headers-4.4.169-0404169_4.4.169-0404169.201812210931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.169/linux-headers-4.4.169-0404169-generic_4.4.169-0404169.201812210931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 170,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.170/linux-headers-4.4.170-0404170_4.4.170-0404170.201901131034_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.170/linux-headers-4.4.170-0404170-generic_4.4.170-0404170.201901131034_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 171,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.171/linux-headers-4.4.171-0404171_4.4.171-0404171.201901162233_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.171/linux-headers-4.4.171-0404171-generic_4.4.171-0404171.201901162233_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 172,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.172/linux-headers-4.4.172-0404172_4.4.172-0404172.201901260447_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.172/linux-headers-4.4.172-0404172-generic_4.4.172-0404172.201901260447_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 173,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.173/linux-headers-4.4.173-0404173_4.4.173-0404173.201902061458_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.173/linux-headers-4.4.173-0404173-generic_4.4.173-0404173.201902061458_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 174,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.174/linux-headers-4.4.174-0404174_4.4.174-0404174.201902080647_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.174/linux-headers-4.4.174-0404174-generic_4.4.174-0404174.201902080647_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 175,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.175/linux-headers-4.4.175-0404175_4.4.175-0404175.201902200548_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.175/linux-headers-4.4.175-0404175-generic_4.4.175-0404175.201902200548_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 176,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.176/linux-headers-4.4.176-0404176_4.4.176-0404176.201902230445_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.176/linux-headers-4.4.176-0404176-generic_4.4.176-0404176.201902230445_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 177,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.177/linux-headers-4.4.177-0404177_4.4.177-0404177.201903230846_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.177/linux-headers-4.4.177-0404177-generic_4.4.177-0404177.201903230846_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 178,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.178/linux-headers-4.4.178-0404178_4.4.178-0404178.201904030153_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.178/linux-headers-4.4.178-0404178-generic_4.4.178-0404178.201904030153_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 179,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.179/linux-headers-4.4.179-0404179_4.4.179-0404179.201904270438_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.179/linux-headers-4.4.179-0404179-generic_4.4.179-0404179.201904270438_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 180,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.180/linux-headers-4.4.180-0404180_4.4.180-0404180.201905161546_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.180/linux-headers-4.4.180-0404180-generic_4.4.180-0404180.201905161546_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 181,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.181/linux-headers-4.4.181-0404181_4.4.181-0404181.201906110738_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.181/linux-headers-4.4.181-0404181-generic_4.4.181-0404181.201906110738_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 182,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.182/linux-headers-4.4.182-0404182_4.4.182-0404182.201906171910_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.182/linux-headers-4.4.182-0404182-generic_4.4.182-0404182.201906171910_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 183,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.183/linux-headers-4.4.183-0404183_4.4.183-0404183.201906220342_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.183/linux-headers-4.4.183-0404183-generic_4.4.183-0404183.201906220342_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 184,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.184/linux-headers-4.4.184-0404184_4.4.184-0404184.201906262148_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.184/linux-headers-4.4.184-0404184-generic_4.4.184-0404184.201906262148_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 185,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.185/linux-headers-4.4.185-0404185_4.4.185-0404185.201907100448_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.185/linux-headers-4.4.185-0404185-generic_4.4.185-0404185.201907100448_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 186,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.186/linux-headers-4.4.186-0404186_4.4.186-0404186.201907210435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.186/linux-headers-4.4.186-0404186-generic_4.4.186-0404186.201907210435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 187,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.187/linux-headers-4.4.187-0404187_4.4.187-0404187.201908040457_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.187/linux-headers-4.4.187-0404187-generic_4.4.187-0404187.201908040457_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 188,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.188/linux-headers-4.4.188-0404188_4.4.188-0404188.201908061332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.188/linux-headers-4.4.188-0404188-generic_4.4.188-0404188.201908061332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 189,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.189/linux-headers-4.4.189-0404189_4.4.189-0404189.201908111150_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.189/linux-headers-4.4.189-0404189-generic_4.4.189-0404189.201908111150_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 190,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.190/linux-headers-4.4.190-0404190_4.4.190-0404190.201908250946_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.190/linux-headers-4.4.190-0404190-generic_4.4.190-0404190.201908250946_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 191,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.191/linux-headers-4.4.191-0404191_4.4.191-0404191.201909060546_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.191/linux-headers-4.4.191-0404191-generic_4.4.191-0404191.201909060546_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 192,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.192/linux-headers-4.4.192-0404192_4.4.192-0404192.201909100635_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.192/linux-headers-4.4.192-0404192-generic_4.4.192-0404192.201909100635_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 193,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.193/linux-headers-4.4.193-0404193_4.4.193-0404193.201909160345_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.193/linux-headers-4.4.193-0404193-generic_4.4.193-0404193.201909160345_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 194,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.194/linux-headers-4.4.194-0404194_4.4.194-0404194.201909210240_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.194/linux-headers-4.4.194-0404194-generic_4.4.194-0404194.201909210240_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 195,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.195/linux-headers-4.4.195-0404195_4.4.195-0404195.201910051141_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.195/linux-headers-4.4.195-0404195-generic_4.4.195-0404195.201910051141_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 196,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.196/linux-headers-4.4.196-0404196_4.4.196-0404196.201910071632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.196/linux-headers-4.4.196-0404196-generic_4.4.196-0404196.201910071632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 197,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.197/linux-headers-4.4.197-0404197_4.4.197-0404197.201910172110_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.197/linux-headers-4.4.197-0404197-generic_4.4.197-0404197.201910172110_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 198,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.198/linux-headers-4.4.198-0404198_4.4.198-0404198.201910290540_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.198/linux-headers-4.4.198-0404198-generic_4.4.198-0404198.201910290540_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 199,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.199/linux-headers-4.4.199-0404199_4.4.199-0404199.201911060804_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.199/linux-headers-4.4.199-0404199-generic_4.4.199-0404199.201911060804_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 200,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.200/linux-headers-4.4.200-0404200_4.4.200-0404200.201911100636_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.200/linux-headers-4.4.200-0404200-generic_4.4.200-0404200.201911100636_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 201,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.201/linux-headers-4.4.201-0404201_4.4.201-0404201.201911122313_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.201/linux-headers-4.4.201-0404201-generic_4.4.201-0404201.201911122313_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 202,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.202/linux-headers-4.4.202-0404202_4.4.202-0404202.201911160550_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.202/linux-headers-4.4.202-0404202-generic_4.4.202-0404202.201911160550_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 203,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.203/linux-headers-4.4.203-0404203_4.4.203-0404203.201911251040_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.203/linux-headers-4.4.203-0404203-generic_4.4.203-0404203.201911251040_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 204,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.204/linux-headers-4.4.204-0404204_4.4.204-0404204.201911282131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.204/linux-headers-4.4.204-0404204-generic_4.4.204-0404204.201911282131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 205,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.205/linux-headers-4.4.205-0404205_4.4.205-0404205.201911290436_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.205/linux-headers-4.4.205-0404205-generic_4.4.205-0404205.201911290436_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 206,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.206/linux-headers-4.4.206-0404206_4.4.206-0404206.201912051035_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.206/linux-headers-4.4.206-0404206-generic_4.4.206-0404206.201912051035_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 207,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.207/linux-headers-4.4.207-0404207_4.4.207-0404207.201912210540_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.207/linux-headers-4.4.207-0404207-generic_4.4.207-0404207.201912210540_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 208,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.208/linux-headers-4.4.208-0404208_4.4.208-0404208.202001041343_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.208/linux-headers-4.4.208-0404208-generic_4.4.208-0404208.202001041343_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 209,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.209/linux-headers-4.4.209-0404209_4.4.209-0404209.202001121140_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.209/linux-headers-4.4.209-0404209-generic_4.4.209-0404209.202001121140_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 210,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.210/linux-headers-4.4.210-0404210_4.4.210-0404210.202001142032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.210/linux-headers-4.4.210-0404210-generic_4.4.210-0404210.202001142032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 211,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.211/linux-headers-4.4.211-0404211_4.4.211-0404211.202001230339_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.211/linux-headers-4.4.211-0404211-generic_4.4.211-0404211.202001230339_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 212,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.212/linux-headers-4.4.212-0404212_4.4.212-0404212.202001300050_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.212/linux-headers-4.4.212-0404212-generic_4.4.212-0404212.202001300050_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 213,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.213/linux-headers-4.4.213-0404213_4.4.213-0404213.202002050938_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.213/linux-headers-4.4.213-0404213-generic_4.4.213-0404213.202002050938_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 214,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.214/linux-headers-4.4.214-0404214_4.4.214-0404214.202002150640_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.214/linux-headers-4.4.214-0404214-generic_4.4.214-0404214.202002150640_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 215,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.215/linux-headers-4.4.215-0404215_4.4.215-0404215.202002281539_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.215/linux-headers-4.4.215-0404215-generic_4.4.215-0404215.202002281539_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 216,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.216/linux-headers-4.4.216-0404216_4.4.216-0404216.202003111524_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.216/linux-headers-4.4.216-0404216-generic_4.4.216-0404216.202003111524_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 217,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.217/linux-headers-4.4.217-0404217_4.4.217-0404217.202003200541_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.217/linux-headers-4.4.217-0404217-generic_4.4.217-0404217.202003200541_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 218,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.218/linux-headers-4.4.218-0404218_4.4.218-0404218.202004021841_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.218/linux-headers-4.4.218-0404218-generic_4.4.218-0404218.202004021841_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 4,
			Patch: 219,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.219/linux-headers-4.4.219-0404219_4.4.219-0404219.202004130538_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.219/linux-headers-4.4.219-0404219-generic_4.4.219-0404219.202004130538_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 6,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.4/linux-headers-4.6.4-040604_4.6.4-040604.201607111332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.4/linux-headers-4.6.4-040604-generic_4.6.4-040604.201607111332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 6,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.5/linux-headers-4.6.5-040605_4.6.5-040605.201607271333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.5/linux-headers-4.6.5-040605-generic_4.6.5-040605.201607271333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 6,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.6/linux-headers-4.6.6-040606_4.6.6-040606.201608100733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.6/linux-headers-4.6.6-040606-generic_4.6.6-040606.201608100733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 6,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.7/linux-headers-4.6.7-040607_4.6.7-040607.201608160432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.6.7/linux-headers-4.6.7-040607-generic_4.6.7-040607.201608160432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7/linux-headers-4.7.0-040700_4.7.0-040700.201608021801_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7/linux-headers-4.7.0-040700-generic_4.7.0-040700.201608021801_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.1/linux-headers-4.7.1-040701_4.7.1-040701.201608160432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.1/linux-headers-4.7.1-040701-generic_4.7.1-040701.201608160432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.2/linux-headers-4.7.2-040702_4.7.2-040702.201608201334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.2/linux-headers-4.7.2-040702-generic_4.7.2-040702.201608201334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.3/linux-headers-4.7.3-040703_4.7.3-040703.201609070334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.3/linux-headers-4.7.3-040703-generic_4.7.3-040703.201609070334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.4/linux-headers-4.7.4-040704_4.7.4-040704.201609150330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.4/linux-headers-4.7.4-040704-generic_4.7.4-040704.201609150330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.5/linux-headers-4.7.5-040705_4.7.5-040705.201609240533_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.5/linux-headers-4.7.5-040705-generic_4.7.5-040705.201609240533_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.6/linux-headers-4.7.6-040706_4.7.6-040706.201609300531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.6/linux-headers-4.7.6-040706-generic_4.7.6-040706.201609300531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.7/linux-headers-4.7.7-040707_4.7.7-040707.201610071031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.7/linux-headers-4.7.7-040707-generic_4.7.7-040707.201610071031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.8/linux-headers-4.7.8-040708_4.7.8-040708.201610161547_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.8/linux-headers-4.7.8-040708-generic_4.7.8-040708.201610161547_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.9/linux-headers-4.7.9-040709_4.7.9-040709.201610200645_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.9/linux-headers-4.7.9-040709-generic_4.7.9-040709.201610200645_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 7,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.10/linux-headers-4.7.10-040710_4.7.10-040710.201610220847_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.7.10/linux-headers-4.7.10-040710-generic_4.7.10-040710.201610220847_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8/linux-headers-4.8.0-040800_4.8.0-040800.201610022031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8/linux-headers-4.8.0-040800-generic_4.8.0-040800.201610022031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.1/linux-headers-4.8.1-040801_4.8.1-040801.201610170913_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.1/linux-headers-4.8.1-040801-generic_4.8.1-040801.201610170913_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.2/linux-headers-4.8.2-040802_4.8.2-040802.201610161339_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.2/linux-headers-4.8.2-040802-generic_4.8.2-040802.201610161339_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.3/linux-headers-4.8.3-040803_4.8.3-040803.201610200531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.3/linux-headers-4.8.3-040803-generic_4.8.3-040803.201610200531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.4/linux-headers-4.8.4-040804_4.8.4-040804.201610220733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.4/linux-headers-4.8.4-040804-generic_4.8.4-040804.201610220733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.5/linux-headers-4.8.5-040805_4.8.5-040805.201610280434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.5/linux-headers-4.8.5-040805-generic_4.8.5-040805.201610280434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.6/linux-headers-4.8.6-040806_4.8.6-040806.201610310831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.6/linux-headers-4.8.6-040806-generic_4.8.6-040806.201610310831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.7/linux-headers-4.8.7-040807_4.8.7-040807.201611101131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.7/linux-headers-4.8.7-040807-generic_4.8.7-040807.201611101131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.8/linux-headers-4.8.8-040808_4.8.8-040808.201611150231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.8/linux-headers-4.8.8-040808-generic_4.8.8-040808.201611150231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.9/linux-headers-4.8.9-040809_4.8.9-040809.201611180631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.9/linux-headers-4.8.9-040809-generic_4.8.9-040809.201611180631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.10/linux-headers-4.8.10-040810_4.8.10-040810.201611210531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.10/linux-headers-4.8.10-040810-generic_4.8.10-040810.201611210531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.11/linux-headers-4.8.11-040811_4.8.11-040811.201611260431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.11/linux-headers-4.8.11-040811-generic_4.8.11-040811.201611260431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.12/linux-headers-4.8.12-040812_4.8.12-040812.201612020431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.12/linux-headers-4.8.12-040812-generic_4.8.12-040812.201612020431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.13/linux-headers-4.8.13-040813_4.8.13-040813.201612081531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.13/linux-headers-4.8.13-040813-generic_4.8.13-040813.201612081531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.14/linux-headers-4.8.14-040814_4.8.14-040814.201612101431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.14/linux-headers-4.8.14-040814-generic_4.8.14-040814.201612101431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.15/linux-headers-4.8.15-040815_4.8.15-040815.201612151231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.15/linux-headers-4.8.15-040815-generic_4.8.15-040815.201612151231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.16/linux-headers-4.8.16-040816_4.8.16-040816.201701060650_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.16/linux-headers-4.8.16-040816-generic_4.8.16-040816.201701060650_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 8,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.17/linux-headers-4.8.17-040817_4.8.17-040817.201701090438_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.8.17/linux-headers-4.8.17-040817-generic_4.8.17-040817.201701090438_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9/linux-headers-4.9.0-040900_4.9.0-040900.201612111631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9/linux-headers-4.9.0-040900-generic_4.9.0-040900.201612111631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.1/linux-headers-4.9.1-040901_4.9.1-040901.201701060531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.1/linux-headers-4.9.1-040901-generic_4.9.1-040901.201701060531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.2/linux-headers-4.9.2-040902_4.9.2-040902.201701090331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.2/linux-headers-4.9.2-040902-generic_4.9.2-040902.201701090331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.3/linux-headers-4.9.3-040903_4.9.3-040903.201701120631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.3/linux-headers-4.9.3-040903-generic_4.9.3-040903.201701120631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.4/linux-headers-4.9.4-040904_4.9.4-040904.201701150831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.4/linux-headers-4.9.4-040904-generic_4.9.4-040904.201701150831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.5/linux-headers-4.9.5-040905_4.9.5-040905.201701200532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.5/linux-headers-4.9.5-040905-generic_4.9.5-040905.201701200532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.6/linux-headers-4.9.6-040906_4.9.6-040906.201701260330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.6/linux-headers-4.9.6-040906-generic_4.9.6-040906.201701260330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.7/linux-headers-4.9.7-040907_4.9.7-040907.201702010331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.7/linux-headers-4.9.7-040907-generic_4.9.7-040907.201702010331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.8/linux-headers-4.9.8-040908_4.9.8-040908.201702040431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.8/linux-headers-4.9.8-040908-generic_4.9.8-040908.201702040431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.9/linux-headers-4.9.9-040909_4.9.9-040909.201702090333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.9/linux-headers-4.9.9-040909-generic_4.9.9-040909.201702090333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.10/linux-headers-4.9.10-040910_4.9.10-040910.201702141931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.10/linux-headers-4.9.10-040910-generic_4.9.10-040910.201702141931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.11/linux-headers-4.9.11-040911_4.9.11-040911.201702181031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.11/linux-headers-4.9.11-040911-generic_4.9.11-040911.201702181031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.12/linux-headers-4.9.12-040912_4.9.12-040912.201702231232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.12/linux-headers-4.9.12-040912-generic_4.9.12-040912.201702231232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.13/linux-headers-4.9.13-040913_4.9.13-040913.201702260631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.13/linux-headers-4.9.13-040913-generic_4.9.13-040913.201702260631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.14/linux-headers-4.9.14-040914_4.9.14-040914.201703120131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.14/linux-headers-4.9.14-040914-generic_4.9.14-040914.201703120131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.15/linux-headers-4.9.15-040915_4.9.15-040915.201703142331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.15/linux-headers-4.9.15-040915-generic_4.9.15-040915.201703142331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.16/linux-headers-4.9.16-040916_4.9.16-040916.201703180831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.16/linux-headers-4.9.16-040916-generic_4.9.16-040916.201703180831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.17/linux-headers-4.9.17-040917_4.9.17-040917.201703220831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.17/linux-headers-4.9.17-040917-generic_4.9.17-040917.201703220831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.18/linux-headers-4.9.18-040918_4.9.18-040918.201703260832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.18/linux-headers-4.9.18-040918-generic_4.9.18-040918.201703260832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.19/linux-headers-4.9.19-040919_4.9.19-040919.201703300431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.19/linux-headers-4.9.19-040919-generic_4.9.19-040919.201703300431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.20/linux-headers-4.9.20-040920_4.9.20-040920.201703310531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.20/linux-headers-4.9.20-040920-generic_4.9.20-040920.201703310531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.21/linux-headers-4.9.21-040921_4.9.21-040921.201704080434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.21/linux-headers-4.9.21-040921-generic_4.9.21-040921.201704080434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 22,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.22/linux-headers-4.9.22-040922_4.9.22-040922.201704120731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.22/linux-headers-4.9.22-040922-generic_4.9.22-040922.201704120731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 23,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.23/linux-headers-4.9.23-040923_4.9.23-040923.201704180232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.23/linux-headers-4.9.23-040923-generic_4.9.23-040923.201704180232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 24,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.24/linux-headers-4.9.24-040924_4.9.24-040924.201704210431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.24/linux-headers-4.9.24-040924-generic_4.9.24-040924.201704210431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 25,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.25/linux-headers-4.9.25-040925_4.9.25-040925.201705041424_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.25/linux-headers-4.9.25-040925-generic_4.9.25-040925.201705041424_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 26,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.26/linux-headers-4.9.26-040926_4.9.26-040926.201705031231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.26/linux-headers-4.9.26-040926-generic_4.9.26-040926.201705031231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 27,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.27/linux-headers-4.9.27-040927_4.9.27-040927.201711051759_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.27/linux-headers-4.9.27-040927-generic_4.9.27-040927.201711051759_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 28,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.28/linux-headers-4.9.28-040928_4.9.28-040928.201705140931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.28/linux-headers-4.9.28-040928-generic_4.9.28-040928.201705140931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 29,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.29/linux-headers-4.9.29-040929_4.9.29-040929.201705200932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.29/linux-headers-4.9.29-040929-generic_4.9.29-040929.201705200932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 30,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.30/linux-headers-4.9.30-040930_4.9.30-040930.201705251131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.30/linux-headers-4.9.30-040930-generic_4.9.30-040930.201705251131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 31,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.31/linux-headers-4.9.31-040931_4.9.31-040931.201706070929_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.31/linux-headers-4.9.31-040931-generic_4.9.31-040931.201706070929_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 32,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.32/linux-headers-4.9.32-040932_4.9.32-040932.201706141032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.32/linux-headers-4.9.32-040932-generic_4.9.32-040932.201706141032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 33,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.33/linux-headers-4.9.33-040933_4.9.33-040933.201706170303_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.33/linux-headers-4.9.33-040933-generic_4.9.33-040933.201706170303_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 34,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.34/linux-headers-4.9.34-040934_4.9.34-040934.201711052051_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.34/linux-headers-4.9.34-040934-generic_4.9.34-040934.201711052051_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 35,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.35/linux-headers-4.9.35-040935_4.9.35-040935.201706290837_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.35/linux-headers-4.9.35-040935-generic_4.9.35-040935.201706290837_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 36,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.36/linux-headers-4.9.36-040936_4.9.36-040936.201710131400_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.36/linux-headers-4.9.36-040936-generic_4.9.36-040936.201710131400_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 37,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.37/linux-headers-4.9.37-040937_4.9.37-040937.201707121132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.37/linux-headers-4.9.37-040937-generic_4.9.37-040937.201707121132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 38,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.38/linux-headers-4.9.38-040938_4.9.38-040938.201707150734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.38/linux-headers-4.9.38-040938-generic_4.9.38-040938.201707150734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 39,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.39/linux-headers-4.9.39-040939_4.9.39-040939.201707210235_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.39/linux-headers-4.9.39-040939-generic_4.9.39-040939.201707210235_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 40,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.40/linux-headers-4.9.40-040940_4.9.40-040940.201707271932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.40/linux-headers-4.9.40-040940-generic_4.9.40-040940.201707271932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 41,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.41/linux-headers-4.9.41-040941_4.9.41-040941.201710131114_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.41/linux-headers-4.9.41-040941-generic_4.9.41-040941.201710131114_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 42,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.42/linux-headers-4.9.42-040942_4.9.42-040942.201711060132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.42/linux-headers-4.9.42-040942-generic_4.9.42-040942.201711060132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 43,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.43/linux-headers-4.9.43-040943_4.9.43-040943.201708122332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.43/linux-headers-4.9.43-040943-generic_4.9.43-040943.201708122332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 44,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.44/linux-headers-4.9.44-040944_4.9.44-040944.201708161731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.44/linux-headers-4.9.44-040944-generic_4.9.44-040944.201708161731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 45,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.45/linux-headers-4.9.45-040945_4.9.45-040945.201708242131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.45/linux-headers-4.9.45-040945-generic_4.9.45-040945.201708242131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 46,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.46/linux-headers-4.9.46-040946_4.9.46-040946.201708300532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.46/linux-headers-4.9.46-040946-generic_4.9.46-040946.201708300532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 47,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.47/linux-headers-4.9.47-040947_4.9.47-040947.201709020231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.47/linux-headers-4.9.47-040947-generic_4.9.47-040947.201709020231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 48,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.48/linux-headers-4.9.48-040948_4.9.48-040948.201710131005_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.48/linux-headers-4.9.48-040948-generic_4.9.48-040948.201710131005_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 49,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.49/linux-headers-4.9.49-040949_4.9.49-040949.201711060420_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.49/linux-headers-4.9.49-040949-generic_4.9.49-040949.201711060420_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 50,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.50/linux-headers-4.9.50-040950_4.9.50-040950.201709131832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.50/linux-headers-4.9.50-040950-generic_4.9.50-040950.201709131832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 51,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.51/linux-headers-4.9.51-040951_4.9.51-040951.201709200331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.51/linux-headers-4.9.51-040951-generic_4.9.51-040951.201709200331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 52,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.52/linux-headers-4.9.52-040952_4.9.52-040952.201710131251_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.52/linux-headers-4.9.52-040952-generic_4.9.52-040952.201710131251_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 53,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.53/linux-headers-4.9.53-040953_4.9.53-040953.201711051759_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.53/linux-headers-4.9.53-040953-generic_4.9.53-040953.201711051759_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 54,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.54/linux-headers-4.9.54-040954_4.9.54-040954.201711052237_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.54/linux-headers-4.9.54-040954-generic_4.9.54-040954.201711052237_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 55,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.55/linux-headers-4.9.55-040955_4.9.55-040955.201711052316_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.55/linux-headers-4.9.55-040955-generic_4.9.55-040955.201711052316_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 56,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.56/linux-headers-4.9.56-040956_4.9.56-040956.201711052144_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.56/linux-headers-4.9.56-040956-generic_4.9.56-040956.201711052144_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 57,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.57/linux-headers-4.9.57-040957_4.9.57-040957.201711052158_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.57/linux-headers-4.9.57-040957-generic_4.9.57-040957.201711052158_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 58,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.58/linux-headers-4.9.58-040958_4.9.58-040958.201711060210_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.58/linux-headers-4.9.58-040958-generic_4.9.58-040958.201711060210_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 59,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.59/linux-headers-4.9.59-040959_4.9.59-040959.201711052038_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.59/linux-headers-4.9.59-040959-generic_4.9.59-040959.201711052038_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 61,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.61/linux-headers-4.9.61-040961_4.9.61-040961.201711080535_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.61/linux-headers-4.9.61-040961-generic_4.9.61-040961.201711080535_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 62,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.62/linux-headers-4.9.62-040962_4.9.62-040962.201711151031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.62/linux-headers-4.9.62-040962-generic_4.9.62-040962.201711151031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 63,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.63/linux-headers-4.9.63-040963_4.9.63-040963.201711210644_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.63/linux-headers-4.9.63-040963-generic_4.9.63-040963.201711210644_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 64,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.64/linux-headers-4.9.64-040964_4.9.64-040964.201711211144_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.64/linux-headers-4.9.64-040964-generic_4.9.64-040964.201711211144_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 65,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.65/linux-headers-4.9.65-040965_4.9.65-040965.201711240331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.65/linux-headers-4.9.65-040965-generic_4.9.65-040965.201711240331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 66,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.66/linux-headers-4.9.66-040966_4.9.66-040966.201711300932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.66/linux-headers-4.9.66-040966-generic_4.9.66-040966.201711300932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 67,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.67/linux-headers-4.9.67-040967_4.9.67-040967.201712050630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.67/linux-headers-4.9.67-040967-generic_4.9.67-040967.201712050630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 68,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.68/linux-headers-4.9.68-040968_4.9.68-040968.201712091734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.68/linux-headers-4.9.68-040968-generic_4.9.68-040968.201712091734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 69,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.69/linux-headers-4.9.69-040969_4.9.69-040969.201712140430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.69/linux-headers-4.9.69-040969-generic_4.9.69-040969.201712140430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 70,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.70/linux-headers-4.9.70-040970_4.9.70-040970.201712161132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.70/linux-headers-4.9.70-040970-generic_4.9.70-040970.201712161132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 71,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.71/linux-headers-4.9.71-040971_4.9.71-040971.201804300911_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.71/linux-headers-4.9.71-040971-generic_4.9.71-040971.201804300911_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 72,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.72/linux-headers-4.9.72-040972_4.9.72-040972.201712251431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.72/linux-headers-4.9.72-040972-generic_4.9.72-040972.201712251431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 73,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.73/linux-headers-4.9.73-040973_4.9.73-040973.201712291730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.73/linux-headers-4.9.73-040973-generic_4.9.73-040973.201712291730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.74/linux-headers-4.9.74-040974_4.9.74-040974.201801022030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.74/linux-headers-4.9.74-040974-generic_4.9.74-040974.201801022030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 75,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.75/linux-headers-4.9.75-040975_4.9.75-040975.201801051530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.75/linux-headers-4.9.75-040975-generic_4.9.75-040975.201801051530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 76,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.76/linux-headers-4.9.76-040976_4.9.76-040976.201801100432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.76/linux-headers-4.9.76-040976-generic_4.9.76-040976.201801100432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 77,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.77/linux-headers-4.9.77-040977_4.9.77-040977.201801170430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.77/linux-headers-4.9.77-040977-generic_4.9.77-040977.201801170430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 78,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.78/linux-headers-4.9.78-040978_4.9.78-040978.201801231931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.78/linux-headers-4.9.78-040978-generic_4.9.78-040978.201801231931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 79,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.79/linux-headers-4.9.79-040979_4.9.79-040979.201801311230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.79/linux-headers-4.9.79-040979-generic_4.9.79-040979.201801311230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 80,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.80/linux-headers-4.9.80-040980_4.9.80-040980.201802031730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.80/linux-headers-4.9.80-040980-generic_4.9.80-040980.201802031730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 81,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.81/linux-headers-4.9.81-040981_4.9.81-040981.201802131230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.81/linux-headers-4.9.81-040981-generic_4.9.81-040981.201802131230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 82,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.82/linux-headers-4.9.82-040982_4.9.82-040982.201802171330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.82/linux-headers-4.9.82-040982-generic_4.9.82-040982.201802171330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 83,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.83/linux-headers-4.9.83-040983_4.9.83-040983.201802261609_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.83/linux-headers-4.9.83-040983-generic_4.9.83-040983.201802261609_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 85,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.85/linux-headers-4.9.85-040985_4.9.85-040985.201802280533_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.85/linux-headers-4.9.85-040985-generic_4.9.85-040985.201802280533_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 86,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.86/linux-headers-4.9.86-040986_4.9.86-040986.201803041431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.86/linux-headers-4.9.86-040986-generic_4.9.86-040986.201803041431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 87,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.87/linux-headers-4.9.87-040987_4.9.87-040987.201803111631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.87/linux-headers-4.9.87-040987-generic_4.9.87-040987.201803111631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 88,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.88/linux-headers-4.9.88-040988_4.9.88-040988.201803181131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.88/linux-headers-4.9.88-040988-generic_4.9.88-040988.201803181131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 89,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.89/linux-headers-4.9.89-040989_4.9.89-040989.201803220931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.89/linux-headers-4.9.89-040989-generic_4.9.89-040989.201803220931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 90,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.90/linux-headers-4.9.90-040990_4.9.90-040990.201803250830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.90/linux-headers-4.9.90-040990-generic_4.9.90-040990.201803250830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 91,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.91/linux-headers-4.9.91-040991_4.9.91-040991.201803281733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.91/linux-headers-4.9.91-040991-generic_4.9.91-040991.201803281733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 92,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.92/linux-headers-4.9.92-040992_4.9.92-040992.201803311730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.92/linux-headers-4.9.92-040992-generic_4.9.92-040992.201803311730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 93,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.93/linux-headers-4.9.93-040993_4.9.93-040993.201804080736_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.93/linux-headers-4.9.93-040993-generic_4.9.93-040993.201804080736_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 94,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.94/linux-headers-4.9.94-040994_4.9.94-040994.201804132130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.94/linux-headers-4.9.94-040994-generic_4.9.94-040994.201804132130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 95,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.95/linux-headers-4.9.95-040995_4.9.95-040995.201804200730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.95/linux-headers-4.9.95-040995-generic_4.9.95-040995.201804200730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 96,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.96/linux-headers-4.9.96-040996_4.9.96-040996.201804240437_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.96/linux-headers-4.9.96-040996-generic_4.9.96-040996.201804240437_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 97,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.97/linux-headers-4.9.97-040997_4.9.97-040997.201804300507_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.97/linux-headers-4.9.97-040997-generic_4.9.97-040997.201804300507_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 98,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.98/linux-headers-4.9.98-040998_4.9.98-040998.201805021530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.98/linux-headers-4.9.98-040998-generic_4.9.98-040998.201805021530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 99,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.99/linux-headers-4.9.99-040999_4.9.99-040999.201805090831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.99/linux-headers-4.9.99-040999-generic_4.9.99-040999.201805090831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 100,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.100/linux-headers-4.9.100-0409100_4.9.100-0409100.201805160931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.100/linux-headers-4.9.100-0409100-generic_4.9.100-0409100.201805160931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 101,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.101/linux-headers-4.9.101-0409101_4.9.101-0409101.201805190930_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.101/linux-headers-4.9.101-0409101-generic_4.9.101-0409101.201805190930_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 102,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.102/linux-headers-4.9.102-0409102_4.9.102-0409102.201805221335_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.102/linux-headers-4.9.102-0409102-generic_4.9.102-0409102.201805221335_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 103,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.103/linux-headers-4.9.103-0409103_4.9.103-0409103.201805251141_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.103/linux-headers-4.9.103-0409103-generic_4.9.103-0409103.201805251141_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 104,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.104/linux-headers-4.9.104-0409104_4.9.104-0409104.201805300631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.104/linux-headers-4.9.104-0409104-generic_4.9.104-0409104.201805300631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 105,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.105/linux-headers-4.9.105-0409105_4.9.105-0409105.201805302130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.105/linux-headers-4.9.105-0409105-generic_4.9.105-0409105.201805302130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 106,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.106/linux-headers-4.9.106-0409106_4.9.106-0409106.201806051226_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.106/linux-headers-4.9.106-0409106-generic_4.9.106-0409106.201806051226_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 107,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.107/linux-headers-4.9.107-0409107_4.9.107-0409107.201806061130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.107/linux-headers-4.9.107-0409107-generic_4.9.107-0409107.201806061130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 108,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.108/linux-headers-4.9.108-0409108_4.9.108-0409108.201806131131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.108/linux-headers-4.9.108-0409108-generic_4.9.108-0409108.201806131131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 109,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.109/linux-headers-4.9.109-0409109_4.9.109-0409109.201806160833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.109/linux-headers-4.9.109-0409109-generic_4.9.109-0409109.201806160833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 110,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.110/linux-headers-4.9.110-0409110_4.9.110-0409110.201806260131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.110/linux-headers-4.9.110-0409110-generic_4.9.110-0409110.201806260131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 111,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.111/linux-headers-4.9.111-0409111_4.9.111-0409111.201807031131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.111/linux-headers-4.9.111-0409111-generic_4.9.111-0409111.201807031131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 112,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.112/linux-headers-4.9.112-0409112_4.9.112-0409112.201807111536_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.112/linux-headers-4.9.112-0409112-generic_4.9.112-0409112.201807111536_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 113,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.113/linux-headers-4.9.113-0409113_4.9.113-0409113.201807171031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.113/linux-headers-4.9.113-0409113-generic_4.9.113-0409113.201807171031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 114,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.114/linux-headers-4.9.114-0409114_4.9.114-0409114.201807220942_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.114/linux-headers-4.9.114-0409114-generic_4.9.114-0409114.201807220942_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 115,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.115/linux-headers-4.9.115-0409115_4.9.115-0409115.201807261127_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.115/linux-headers-4.9.115-0409115-generic_4.9.115-0409115.201807261127_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 116,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.116/linux-headers-4.9.116-0409116_4.9.116-0409116.201807280342_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.116/linux-headers-4.9.116-0409116-generic_4.9.116-0409116.201807280342_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 117,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.117/linux-headers-4.9.117-0409117_4.9.117-0409117.201808030631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.117/linux-headers-4.9.117-0409117-generic_4.9.117-0409117.201808030631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 118,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.118/linux-headers-4.9.118-0409118_4.9.118-0409118.201808061531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.118/linux-headers-4.9.118-0409118-generic_4.9.118-0409118.201808061531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 119,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.119/linux-headers-4.9.119-0409119_4.9.119-0409119.201808091131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.119/linux-headers-4.9.119-0409119-generic_4.9.119-0409119.201808091131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 120,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.120/linux-headers-4.9.120-0409120_4.9.120-0409120.201808151731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.120/linux-headers-4.9.120-0409120-generic_4.9.120-0409120.201808151731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 121,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.121/linux-headers-4.9.121-0409121_4.9.121-0409121.201808172032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.121/linux-headers-4.9.121-0409121-generic_4.9.121-0409121.201808172032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 122,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.122/linux-headers-4.9.122-0409122_4.9.122-0409122.201808180931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.122/linux-headers-4.9.122-0409122-generic_4.9.122-0409122.201808180931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 123,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.123/linux-headers-4.9.123-0409123_4.9.123-0409123.201808220630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.123/linux-headers-4.9.123-0409123-generic_4.9.123-0409123.201808220630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 124,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.124/linux-headers-4.9.124-0409124_4.9.124-0409124.201808240839_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.124/linux-headers-4.9.124-0409124-generic_4.9.124-0409124.201808240839_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 125,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.125/linux-headers-4.9.125-0409125_4.9.125-0409125.201809050432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.125/linux-headers-4.9.125-0409125-generic_4.9.125-0409125.201809050432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 126,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.126/linux-headers-4.9.126-0409126_4.9.126-0409126.201809091531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.126/linux-headers-4.9.126-0409126-generic_4.9.126-0409126.201809091531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 127,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.127/linux-headers-4.9.127-0409127_4.9.127-0409127.201809150838_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.127/linux-headers-4.9.127-0409127-generic_4.9.127-0409127.201809150838_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 128,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.128/linux-headers-4.9.128-0409128_4.9.128-0409128.201809192132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.128/linux-headers-4.9.128-0409128-generic_4.9.128-0409128.201809192132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 129,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.129/linux-headers-4.9.129-0409129_4.9.129-0409129.201809260731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.129/linux-headers-4.9.129-0409129-generic_4.9.129-0409129.201809260731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 130,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.130/linux-headers-4.9.130-0409130_4.9.130-0409130.201809291131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.130/linux-headers-4.9.130-0409130-generic_4.9.130-0409130.201809291131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 131,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.131/linux-headers-4.9.131-0409131_4.9.131-0409131.201810032146_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.131/linux-headers-4.9.131-0409131-generic_4.9.131-0409131.201810032146_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 132,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.132/linux-headers-4.9.132-0409132_4.9.132-0409132.201810100731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.132/linux-headers-4.9.132-0409132-generic_4.9.132-0409132.201810100731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 133,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.133/linux-headers-4.9.133-0409133_4.9.133-0409133.201810130831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.133/linux-headers-4.9.133-0409133-generic_4.9.133-0409133.201810130831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 134,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.134/linux-headers-4.9.134-0409134_4.9.134-0409134.201810180830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.134/linux-headers-4.9.134-0409134-generic_4.9.134-0409134.201810180830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 135,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.135/linux-headers-4.9.135-0409135_4.9.135-0409135.201810200830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.135/linux-headers-4.9.135-0409135-generic_4.9.135-0409135.201810200830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 136,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.136/linux-headers-4.9.136-0409136_4.9.136-0409136.201811101632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.136/linux-headers-4.9.136-0409136-generic_4.9.136-0409136.201811101632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 137,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.137/linux-headers-4.9.137-0409137_4.9.137-0409137.201811131541_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.137/linux-headers-4.9.137-0409137-generic_4.9.137-0409137.201811131541_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 138,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.138/linux-headers-4.9.138-0409138_4.9.138-0409138.201811210934_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.138/linux-headers-4.9.138-0409138-generic_4.9.138-0409138.201811210934_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 139,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.139/linux-headers-4.9.139-0409139_4.9.139-0409139.201811230830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.139/linux-headers-4.9.139-0409139-generic_4.9.139-0409139.201811230830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 140,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.140/linux-headers-4.9.140-0409140_4.9.140-0409140.201811231231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.140/linux-headers-4.9.140-0409140-generic_4.9.140-0409140.201811231231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 141,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.141/linux-headers-4.9.141-0409141_4.9.141-0409141.201811291439_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.141/linux-headers-4.9.141-0409141-generic_4.9.141-0409141.201811291439_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 142,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.142/linux-headers-4.9.142-0409142_4.9.142-0409142.201812010931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.142/linux-headers-4.9.142-0409142-generic_4.9.142-0409142.201812010931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 143,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.143/linux-headers-4.9.143-0409143_4.9.143-0409143.201812051441_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.143/linux-headers-4.9.143-0409143-generic_4.9.143-0409143.201812051441_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 144,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.144/linux-headers-4.9.144-0409144_4.9.144-0409144.201812081330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.144/linux-headers-4.9.144-0409144-generic_4.9.144-0409144.201812081330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 145,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.145/linux-headers-4.9.145-0409145_4.9.145-0409145.201812130932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.145/linux-headers-4.9.145-0409145-generic_4.9.145-0409145.201812130932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 146,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.146/linux-headers-4.9.146-0409146_4.9.146-0409146.201812170932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.146/linux-headers-4.9.146-0409146-generic_4.9.146-0409146.201812170932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 147,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.147/linux-headers-4.9.147-0409147_4.9.147-0409147.201812211431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.147/linux-headers-4.9.147-0409147-generic_4.9.147-0409147.201812211431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 149,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.149/linux-headers-4.9.149-0409149_4.9.149-0409149.201901100216_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.149/linux-headers-4.9.149-0409149-generic_4.9.149-0409149.201901100216_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 150,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.150/linux-headers-4.9.150-0409150_4.9.150-0409150.201901130532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.150/linux-headers-4.9.150-0409150-generic_4.9.150-0409150.201901130532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 151,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.151/linux-headers-4.9.151-0409151_4.9.151-0409151.201901161743_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.151/linux-headers-4.9.151-0409151-generic_4.9.151-0409151.201901161743_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 152,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.152/linux-headers-4.9.152-0409152_4.9.152-0409152.201901230832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.152/linux-headers-4.9.152-0409152-generic_4.9.152-0409152.201901230832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 153,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.153/linux-headers-4.9.153-0409153_4.9.153-0409153.201901260933_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.153/linux-headers-4.9.153-0409153-generic_4.9.153-0409153.201901260933_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 154,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.154/linux-headers-4.9.154-0409154_4.9.154-0409154.201901310831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.154/linux-headers-4.9.154-0409154-generic_4.9.154-0409154.201901310831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 155,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.155/linux-headers-4.9.155-0409155_4.9.155-0409155.201902061234_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.155/linux-headers-4.9.155-0409155-generic_4.9.155-0409155.201902061234_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 156,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.156/linux-headers-4.9.156-0409156_4.9.156-0409156.201902121937_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.156/linux-headers-4.9.156-0409156-generic_4.9.156-0409156.201902121937_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 157,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.157/linux-headers-4.9.157-0409157_4.9.157-0409157.201902150831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.157/linux-headers-4.9.157-0409157-generic_4.9.157-0409157.201902150831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 158,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.158/linux-headers-4.9.158-0409158_4.9.158-0409158.201902150951_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.158/linux-headers-4.9.158-0409158-generic_4.9.158-0409158.201902150951_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 159,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.159/linux-headers-4.9.159-0409159_4.9.159-0409159.201902201034_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.159/linux-headers-4.9.159-0409159-generic_4.9.159-0409159.201902201034_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 160,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.160/linux-headers-4.9.160-0409160_4.9.160-0409160.201902230931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.160/linux-headers-4.9.160-0409160-generic_4.9.160-0409160.201902230931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 161,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.161/linux-headers-4.9.161-0409161_4.9.161-0409161.201902271031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.161/linux-headers-4.9.161-0409161-generic_4.9.161-0409161.201902271031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 162,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.162/linux-headers-4.9.162-0409162_4.9.162-0409162.201903051732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.162/linux-headers-4.9.162-0409162-generic_4.9.162-0409162.201903051732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 163,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.163/linux-headers-4.9.163-0409163_4.9.163-0409163.201903131846_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.163/linux-headers-4.9.163-0409163-generic_4.9.163-0409163.201903131846_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 164,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.164/linux-headers-4.9.164-0409164_4.9.164-0409164.201903190933_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.164/linux-headers-4.9.164-0409164-generic_4.9.164-0409164.201903190933_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 165,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.165/linux-headers-4.9.165-0409165_4.9.165-0409165.201903230935_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.165/linux-headers-4.9.165-0409165-generic_4.9.165-0409165.201903230935_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 166,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.166/linux-headers-4.9.166-0409166_4.9.166-0409166.201903271616_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.166/linux-headers-4.9.166-0409166-generic_4.9.166-0409166.201903271616_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 167,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.167/linux-headers-4.9.167-0409167_4.9.167-0409167.201904030146_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.167/linux-headers-4.9.167-0409167-generic_4.9.167-0409167.201904030146_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 168,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.168/linux-headers-4.9.168-0409168_4.9.168-0409168.201904051739_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.168/linux-headers-4.9.168-0409168-generic_4.9.168-0409168.201904051739_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 169,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.169/linux-headers-4.9.169-0409169_4.9.169-0409169.201904170334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.169/linux-headers-4.9.169-0409169-generic_4.9.169-0409169.201904170334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 170,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.170/linux-headers-4.9.170-0409170_4.9.170-0409170.201904200430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.170/linux-headers-4.9.170-0409170-generic_4.9.170-0409170.201904200430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 171,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.171/linux-headers-4.9.171-0409171_4.9.171-0409171.201904270449_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.171/linux-headers-4.9.171-0409171-generic_4.9.171-0409171.201904270449_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 172,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.172/linux-headers-4.9.172-0409172_4.9.172-0409172.201905020555_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.172/linux-headers-4.9.172-0409172-generic_4.9.172-0409172.201905020555_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 173,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.173/linux-headers-4.9.173-0409173_4.9.173-0409173.201906091313_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.173/linux-headers-4.9.173-0409173-generic_4.9.173-0409173.201906091313_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 174,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.174/linux-headers-4.9.174-0409174_4.9.174-0409174.201905080231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.174/linux-headers-4.9.174-0409174-generic_4.9.174-0409174.201905080231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 175,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.175/linux-headers-4.9.175-0409175_4.9.175-0409175.201905101647_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.175/linux-headers-4.9.175-0409175-generic_4.9.175-0409175.201905101647_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 176,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.176/linux-headers-4.9.176-0409176_4.9.176-0409176.201905141431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.176/linux-headers-4.9.176-0409176-generic_4.9.176-0409176.201905141431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 177,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.177/linux-headers-4.9.177-0409177_4.9.177-0409177.201905161442_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.177/linux-headers-4.9.177-0409177-generic_4.9.177-0409177.201905161442_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 178,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.178/linux-headers-4.9.178-0409178_4.9.178-0409178.201905211331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.178/linux-headers-4.9.178-0409178-generic_4.9.178-0409178.201905211331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 179,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.179/linux-headers-4.9.179-0409179_4.9.179-0409179.201906051459_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.179/linux-headers-4.9.179-0409179-generic_4.9.179-0409179.201906051459_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 180,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.180/linux-headers-4.9.180-0409180_4.9.180-0409180.201906051839_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.180/linux-headers-4.9.180-0409180-generic_4.9.180-0409180.201906051839_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 181,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.181/linux-headers-4.9.181-0409181_4.9.181-0409181.201906111132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.181/linux-headers-4.9.181-0409181-generic_4.9.181-0409181.201906111132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 182,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.182/linux-headers-4.9.182-0409182_4.9.182-0409182.201906171853_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.182/linux-headers-4.9.182-0409182-generic_4.9.182-0409182.201906171853_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 183,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.183/linux-headers-4.9.183-0409183_4.9.183-0409183.201906220732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.183/linux-headers-4.9.183-0409183-generic_4.9.183-0409183.201906220732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 184,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.184/linux-headers-4.9.184-0409184_4.9.184-0409184.201906262150_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.184/linux-headers-4.9.184-0409184-generic_4.9.184-0409184.201906262150_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 185,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.185/linux-headers-4.9.185-0409185_4.9.185-0409185.201907100835_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.185/linux-headers-4.9.185-0409185-generic_4.9.185-0409185.201907100835_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 186,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.186/linux-headers-4.9.186-0409186_4.9.186-0409186.201907210453_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.186/linux-headers-4.9.186-0409186-generic_4.9.186-0409186.201907210453_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 187,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.187/linux-headers-4.9.187-0409187_4.9.187-0409187.201908040854_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.187/linux-headers-4.9.187-0409187-generic_4.9.187-0409187.201908040854_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 188,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.188/linux-headers-4.9.188-0409188_4.9.188-0409188.201908061734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.188/linux-headers-4.9.188-0409188-generic_4.9.188-0409188.201908061734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 189,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.189/linux-headers-4.9.189-0409189_4.9.189-0409189.201908111148_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.189/linux-headers-4.9.189-0409189-generic_4.9.189-0409189.201908111148_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 190,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.190/linux-headers-4.9.190-0409190_4.9.190-0409190.201908250934_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.190/linux-headers-4.9.190-0409190-generic_4.9.190-0409190.201908250934_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 191,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.191/linux-headers-4.9.191-0409191_4.9.191-0409191.201909060556_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.191/linux-headers-4.9.191-0409191-generic_4.9.191-0409191.201909060556_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 192,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.192/linux-headers-4.9.192-0409192_4.9.192-0409192.201909101051_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.192/linux-headers-4.9.192-0409192-generic_4.9.192-0409192.201909101051_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 193,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.193/linux-headers-4.9.193-0409193_4.9.193-0409193.201909160733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.193/linux-headers-4.9.193-0409193-generic_4.9.193-0409193.201909160733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 194,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.194/linux-headers-4.9.194-0409194_4.9.194-0409194.201909210254_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.194/linux-headers-4.9.194-0409194-generic_4.9.194-0409194.201909210254_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 195,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.195/linux-headers-4.9.195-0409195_4.9.195-0409195.201910051133_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.195/linux-headers-4.9.195-0409195-generic_4.9.195-0409195.201910051133_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 196,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.196/linux-headers-4.9.196-0409196_4.9.196-0409196.201910071732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.196/linux-headers-4.9.196-0409196-generic_4.9.196-0409196.201910071732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 197,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.197/linux-headers-4.9.197-0409197_4.9.197-0409197.201910180045_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.197/linux-headers-4.9.197-0409197-generic_4.9.197-0409197.201910180045_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 199,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.199/linux-headers-4.9.199-0409199_4.9.199-0409199.201911061233_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.199/linux-headers-4.9.199-0409199-generic_4.9.199-0409199.201911061233_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 200,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.200/linux-headers-4.9.200-0409200_4.9.200-0409200.201911100657_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.200/linux-headers-4.9.200-0409200-generic_4.9.200-0409200.201911100657_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 201,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.201/linux-headers-4.9.201-0409201_4.9.201-0409201.201911121750_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.201/linux-headers-4.9.201-0409201-generic_4.9.201-0409201.201911121750_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 202,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.202/linux-headers-4.9.202-0409202_4.9.202-0409202.201911161036_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.202/linux-headers-4.9.202-0409202-generic_4.9.202-0409202.201911161036_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 203,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.203/linux-headers-4.9.203-0409203_4.9.203-0409203.201911250935_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.203/linux-headers-4.9.203-0409203-generic_4.9.203-0409203.201911250935_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 204,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.204/linux-headers-4.9.204-0409204_4.9.204-0409204.201911282120_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.204/linux-headers-4.9.204-0409204-generic_4.9.204-0409204.201911282120_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 205,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.205/linux-headers-4.9.205-0409205_4.9.205-0409205.201911290931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.205/linux-headers-4.9.205-0409205-generic_4.9.205-0409205.201911290931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 206,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.206/linux-headers-4.9.206-0409206_4.9.206-0409206.201912051108_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.206/linux-headers-4.9.206-0409206-generic_4.9.206-0409206.201912051108_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 207,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.207/linux-headers-4.9.207-0409207_4.9.207-0409207.201912211037_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.207/linux-headers-4.9.207-0409207-generic_4.9.207-0409207.201912211037_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 208,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.208/linux-headers-4.9.208-0409208_4.9.208-0409208.202001041332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.208/linux-headers-4.9.208-0409208-generic_4.9.208-0409208.202001041332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 209,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.209/linux-headers-4.9.209-0409209_4.9.209-0409209.202001121135_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.209/linux-headers-4.9.209-0409209-generic_4.9.209-0409209.202001121135_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 210,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.210/linux-headers-4.9.210-0409210_4.9.210-0409210.202001142031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.210/linux-headers-4.9.210-0409210-generic_4.9.210-0409210.202001142031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 211,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.211/linux-headers-4.9.211-0409211_4.9.211-0409211.202001230833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.211/linux-headers-4.9.211-0409211-generic_4.9.211-0409211.202001230833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 212,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.212/linux-headers-4.9.212-0409212_4.9.212-0409212.202001300045_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.212/linux-headers-4.9.212-0409212-generic_4.9.212-0409212.202001300045_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 213,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.213/linux-headers-4.9.213-0409213_4.9.213-0409213.202002051432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.213/linux-headers-4.9.213-0409213-generic_4.9.213-0409213.202002051432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 214,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.214/linux-headers-4.9.214-0409214_4.9.214-0409214.202002151113_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.214/linux-headers-4.9.214-0409214-generic_4.9.214-0409214.202002151113_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 215,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.215/linux-headers-4.9.215-0409215_4.9.215-0409215.202002281204_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.215/linux-headers-4.9.215-0409215-generic_4.9.215-0409215.202002281204_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 216,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.216/linux-headers-4.9.216-0409216_4.9.216-0409216.202003111208_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.216/linux-headers-4.9.216-0409216-generic_4.9.216-0409216.202003111208_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 217,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.217/linux-headers-4.9.217-0409217_4.9.217-0409217.202003201525_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.217/linux-headers-4.9.217-0409217-generic_4.9.217-0409217.202003201525_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 218,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.218/linux-headers-4.9.218-0409218_4.9.218-0409218.202004021416_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.218/linux-headers-4.9.218-0409218-generic_4.9.218-0409218.202004021416_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 9,
			Patch: 219,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.219/linux-headers-4.9.219-0409219_4.9.219-0409219.202004131105_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.9.219/linux-headers-4.9.219-0409219-generic_4.9.219-0409219.202004131105_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10/linux-headers-4.10.0-041000_4.10.0-041000.201702191831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10/linux-headers-4.10.0-041000-generic_4.10.0-041000.201702191831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.1/linux-headers-4.10.1-041001_4.10.1-041001.201702260735_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.1/linux-headers-4.10.1-041001-generic_4.10.1-041001.201702260735_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.2/linux-headers-4.10.2-041002_4.10.2-041002.201703120131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.2/linux-headers-4.10.2-041002-generic_4.10.2-041002.201703120131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.3/linux-headers-4.10.3-041003_4.10.3-041003.201703142331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.3/linux-headers-4.10.3-041003-generic_4.10.3-041003.201703142331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.4/linux-headers-4.10.4-041004_4.10.4-041004.201703180831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.4/linux-headers-4.10.4-041004-generic_4.10.4-041004.201703180831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.5/linux-headers-4.10.5-041005_4.10.5-041005.201703220931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.5/linux-headers-4.10.5-041005-generic_4.10.5-041005.201703220931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.6/linux-headers-4.10.6-041006_4.10.6-041006.201703260832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.6/linux-headers-4.10.6-041006-generic_4.10.6-041006.201703260832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.7/linux-headers-4.10.7-041007_4.10.7-041007.201703300510_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.7/linux-headers-4.10.7-041007-generic_4.10.7-041007.201703300510_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.8/linux-headers-4.10.8-041008_4.10.8-041008.201703310531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.8/linux-headers-4.10.8-041008-generic_4.10.8-041008.201703310531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.9/linux-headers-4.10.9-041009_4.10.9-041009.201704080516_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.9/linux-headers-4.10.9-041009-generic_4.10.9-041009.201704080516_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.10/linux-headers-4.10.10-041010_4.10.10-041010.201704120813_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.10/linux-headers-4.10.10-041010-generic_4.10.10-041010.201704120813_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.11/linux-headers-4.10.11-041011_4.10.11-041011.201704180310_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.11/linux-headers-4.10.11-041011-generic_4.10.11-041011.201704180310_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.12/linux-headers-4.10.12-041012_4.10.12-041012.201704210512_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.12/linux-headers-4.10.12-041012-generic_4.10.12-041012.201704210512_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.13/linux-headers-4.10.13-041013_4.10.13-041013.201705041649_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.13/linux-headers-4.10.13-041013-generic_4.10.13-041013.201705041649_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.14/linux-headers-4.10.14-041014_4.10.14-041014.201705031501_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.14/linux-headers-4.10.14-041014-generic_4.10.14-041014.201705031501_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.15/linux-headers-4.10.15-041015_4.10.15-041015.201705080411_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.15/linux-headers-4.10.15-041015-generic_4.10.15-041015.201705080411_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.16/linux-headers-4.10.16-041016_4.10.16-041016.201710130856_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.16/linux-headers-4.10.16-041016-generic_4.10.16-041016.201710130856_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 10,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.17/linux-headers-4.10.17-041017_4.10.17-041017.201705201051_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.10.17/linux-headers-4.10.17-041017-generic_4.10.17-041017.201705201051_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11/linux-headers-4.11.0-041100_4.11.0-041100.201705041534_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11/linux-headers-4.11.0-041100-generic_4.11.0-041100.201705041534_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.1/linux-headers-4.11.1-041101_4.11.1-041101.201705140931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.1/linux-headers-4.11.1-041101-generic_4.11.1-041101.201705140931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.2/linux-headers-4.11.2-041102_4.11.2-041102.201705201036_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.2/linux-headers-4.11.2-041102-generic_4.11.2-041102.201705201036_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.3/linux-headers-4.11.3-041103_4.11.3-041103.201705251233_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.3/linux-headers-4.11.3-041103-generic_4.11.3-041103.201705251233_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.4/linux-headers-4.11.4-041104_4.11.4-041104.201706071003_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.4/linux-headers-4.11.4-041104-generic_4.11.4-041104.201706071003_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.5/linux-headers-4.11.5-041105_4.11.5-041105.201706141137_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.5/linux-headers-4.11.5-041105-generic_4.11.5-041105.201706141137_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.6/linux-headers-4.11.6-041106_4.11.6-041106.201706170517_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.6/linux-headers-4.11.6-041106-generic_4.11.6-041106.201706170517_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.7/linux-headers-4.11.7-041107_4.11.7-041107.201706240231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.7/linux-headers-4.11.7-041107-generic_4.11.7-041107.201706240231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.8/linux-headers-4.11.8-041108_4.11.8-041108.201706290836_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.8/linux-headers-4.11.8-041108-generic_4.11.8-041108.201706290836_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.9/linux-headers-4.11.9-041109_4.11.9-041109.201707050933_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.9/linux-headers-4.11.9-041109-generic_4.11.9-041109.201707050933_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.10/linux-headers-4.11.10-041110_4.11.10-041110.201707121132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.10/linux-headers-4.11.10-041110-generic_4.11.10-041110.201707121132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.11/linux-headers-4.11.11-041111_4.11.11-041111.201707150838_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.11/linux-headers-4.11.11-041111-generic_4.11.11-041111.201707150838_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 11,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.12/linux-headers-4.11.12-041112_4.11.12-041112.201707210350_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.11.12/linux-headers-4.11.12-041112-generic_4.11.12-041112.201707210350_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12/linux-headers-4.12.0-041200_4.12.0-041200.201707022031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12/linux-headers-4.12.0-041200-generic_4.12.0-041200.201707022031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.1/linux-headers-4.12.1-041201_4.12.1-041201.201707121132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.1/linux-headers-4.12.1-041201-generic_4.12.1-041201.201707121132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.2/linux-headers-4.12.2-041202_4.12.2-041202.201707150832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.2/linux-headers-4.12.2-041202-generic_4.12.2-041202.201707150832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.3/linux-headers-4.12.3-041203_4.12.3-041203.201707210343_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.3/linux-headers-4.12.3-041203-generic_4.12.3-041203.201707210343_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.4/linux-headers-4.12.4-041204_4.12.4-041204.201707271932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.4/linux-headers-4.12.4-041204-generic_4.12.4-041204.201707271932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.5/linux-headers-4.12.5-041205_4.12.5-041205.201708061334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.5/linux-headers-4.12.5-041205-generic_4.12.5-041205.201708061334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.6/linux-headers-4.12.6-041206_4.12.6-041206.201708111232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.6/linux-headers-4.12.6-041206-generic_4.12.6-041206.201708111232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.7/linux-headers-4.12.7-041207_4.12.7-041207.201708160856_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.7/linux-headers-4.12.7-041207-generic_4.12.7-041207.201708160856_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.8/linux-headers-4.12.8-041208_4.12.8-041208.201708161815_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.8/linux-headers-4.12.8-041208-generic_4.12.8-041208.201708161815_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.9/linux-headers-4.12.9-041209_4.12.9-041209.201708242344_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.9/linux-headers-4.12.9-041209-generic_4.12.9-041209.201708242344_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.10/linux-headers-4.12.10-041210_4.12.10-041210.201708300614_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.10/linux-headers-4.12.10-041210-generic_4.12.10-041210.201708300614_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.11/linux-headers-4.12.11-041211_4.12.11-041211.201709070418_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.11/linux-headers-4.12.11-041211-generic_4.12.11-041211.201709070418_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.12/linux-headers-4.12.12-041212_4.12.12-041212.201709100316_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.12/linux-headers-4.12.12-041212-generic_4.12.12-041212.201709100316_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.13/linux-headers-4.12.13-041213_4.12.13-041213.201709132217_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.13/linux-headers-4.12.13-041213-generic_4.12.13-041213.201709132217_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 12,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.14/linux-headers-4.12.14-041214_4.12.14-041214.201709200843_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.12.14/linux-headers-4.12.14-041214-generic_4.12.14-041214.201709200843_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13/linux-headers-4.13.0-041300_4.13.0-041300.201709031731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13/linux-headers-4.13.0-041300-generic_4.13.0-041300.201709031731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.1/linux-headers-4.13.1-041301_4.13.1-041301.201709100232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.1/linux-headers-4.13.1-041301-generic_4.13.1-041301.201709100232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.2/linux-headers-4.13.2-041302_4.13.2-041302.201709132057_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.2/linux-headers-4.13.2-041302-generic_4.13.2-041302.201709132057_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.3/linux-headers-4.13.3-041303_4.13.3-041303.201709200606_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.3/linux-headers-4.13.3-041303-generic_4.13.3-041303.201709200606_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.4/linux-headers-4.13.4-041304_4.13.4-041304.201709270931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.4/linux-headers-4.13.4-041304-generic_4.13.4-041304.201709270931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.5/linux-headers-4.13.5-041305_4.13.5-041305.201710050600_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.5/linux-headers-4.13.5-041305-generic_4.13.5-041305.201710050600_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.6/linux-headers-4.13.6-041306_4.13.6-041306.201711060046_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.6/linux-headers-4.13.6-041306-generic_4.13.6-041306.201711060046_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.7/linux-headers-4.13.7-041307_4.13.7-041307.201711060248_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.7/linux-headers-4.13.7-041307-generic_4.13.7-041307.201711060248_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.8/linux-headers-4.13.8-041308_4.13.8-041308.201710180430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.8/linux-headers-4.13.8-041308-generic_4.13.8-041308.201710180430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.9/linux-headers-4.13.9-041309_4.13.9-041309.201710211231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.9/linux-headers-4.13.9-041309-generic_4.13.9-041309.201710211231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.10/linux-headers-4.13.10-041310_4.13.10-041310.201710270531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.10/linux-headers-4.13.10-041310-generic_4.13.10-041310.201710270531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.11/linux-headers-4.13.11-041311_4.13.11-041311.201711020532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.11/linux-headers-4.13.11-041311-generic_4.13.11-041311.201711020532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.12/linux-headers-4.13.12-041312_4.13.12-041312.201711080535_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.12/linux-headers-4.13.12-041312-generic_4.13.12-041312.201711080535_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.13/linux-headers-4.13.13-041313_4.13.13-041313.201711150531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.13/linux-headers-4.13.13-041313-generic_4.13.13-041313.201711150531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.14/linux-headers-4.13.14-041314_4.13.14-041314.201711180632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.14/linux-headers-4.13.14-041314-generic_4.13.14-041314.201711180632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.15/linux-headers-4.13.15-041315_4.13.15-041315.201711211030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.15/linux-headers-4.13.15-041315-generic_4.13.15-041315.201711211030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 13,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.16/linux-headers-4.13.16-041316_4.13.16-041316.201711240901_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.13.16/linux-headers-4.13.16-041316-generic_4.13.16-041316.201711240901_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14/linux-headers-4.14.0-041400_4.14.0-041400.201711122031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14/linux-headers-4.14.0-041400-generic_4.14.0-041400.201711122031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.1/linux-headers-4.14.1-041401_4.14.1-041401.201711210430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.1/linux-headers-4.14.1-041401-generic_4.14.1-041401.201711210430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.2/linux-headers-4.14.2-041402_4.14.2-041402.201711240330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.2/linux-headers-4.14.2-041402-generic_4.14.2-041402.201711240330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.3/linux-headers-4.14.3-041403_4.14.3-041403.201711300431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.3/linux-headers-4.14.3-041403-generic_4.14.3-041403.201711300431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.4/linux-headers-4.14.4-041404_4.14.4-041404.201712050630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.4/linux-headers-4.14.4-041404-generic_4.14.4-041404.201712050630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.5/linux-headers-4.14.5-041405_4.14.5-041405.201712101332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.5/linux-headers-4.14.5-041405-generic_4.14.5-041405.201712101332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.6/linux-headers-4.14.6-041406_4.14.6-041406.201712140930_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.6/linux-headers-4.14.6-041406-generic_4.14.6-041406.201712140930_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.7/linux-headers-4.14.7-041407_4.14.7-041407.201712171031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.7/linux-headers-4.14.7-041407-generic_4.14.7-041407.201712171031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.8/linux-headers-4.14.8-041408_4.14.8-041408.201712200555_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.8/linux-headers-4.14.8-041408-generic_4.14.8-041408.201712200555_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.9/linux-headers-4.14.9-041409_4.14.9-041409.201712251541_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.9/linux-headers-4.14.9-041409-generic_4.14.9-041409.201712251541_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.10/linux-headers-4.14.10-041410_4.14.10-041410.201712291810_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.10/linux-headers-4.14.10-041410-generic_4.14.10-041410.201712291810_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.11/linux-headers-4.14.11-041411_4.14.11-041411.201801022143_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.11/linux-headers-4.14.11-041411-generic_4.14.11-041411.201801022143_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.12/linux-headers-4.14.12-041412_4.14.12-041412.201801051649_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.12/linux-headers-4.14.12-041412-generic_4.14.12-041412.201801051649_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.13/linux-headers-4.14.13-041413_4.14.13-041413.201801101001_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.13/linux-headers-4.14.13-041413-generic_4.14.13-041413.201801101001_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.14/linux-headers-4.14.14-041414_4.14.14-041414.201801201219_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.14/linux-headers-4.14.14-041414-generic_4.14.14-041414.201801201219_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.15/linux-headers-4.14.15-041415_4.14.15-041415.201801231530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.15/linux-headers-4.14.15-041415-generic_4.14.15-041415.201801231530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.16/linux-headers-4.14.16-041416_4.14.16-041416.201801310931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.16/linux-headers-4.14.16-041416-generic_4.14.16-041416.201801310931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.17/linux-headers-4.14.17-041417_4.14.17-041417.201802031230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.17/linux-headers-4.14.17-041417-generic_4.14.17-041417.201802031230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.18/linux-headers-4.14.18-041418_4.14.18-041418.201802071730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.18/linux-headers-4.14.18-041418-generic_4.14.18-041418.201802071730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.19/linux-headers-4.14.19-041419_4.14.19-041419.201802131030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.19/linux-headers-4.14.19-041419-generic_4.14.19-041419.201802131030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.20/linux-headers-4.14.20-041420_4.14.20-041420.201802162247_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.20/linux-headers-4.14.20-041420-generic_4.14.20-041420.201802162247_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.21/linux-headers-4.14.21-041421_4.14.21-041421.201802221610_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.21/linux-headers-4.14.21-041421-generic_4.14.21-041421.201802221610_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 23,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.23/linux-headers-4.14.23-041423_4.14.23-041423.201802281111_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.23/linux-headers-4.14.23-041423-generic_4.14.23-041423.201802281111_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 24,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.24/linux-headers-4.14.24-041424_4.14.24-041424.201803040931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.24/linux-headers-4.14.24-041424-generic_4.14.24-041424.201803040931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 25,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.25/linux-headers-4.14.25-041425_4.14.25-041425.201803091130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.25/linux-headers-4.14.25-041425-generic_4.14.25-041425.201803091130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 26,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.26/linux-headers-4.14.26-041426_4.14.26-041426.201803111320_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.26/linux-headers-4.14.26-041426-generic_4.14.26-041426.201803111320_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 27,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.27/linux-headers-4.14.27-041427_4.14.27-041427.201803151737_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.27/linux-headers-4.14.27-041427-generic_4.14.27-041427.201803151737_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 28,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.28/linux-headers-4.14.28-041428_4.14.28-041428.201803190830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.28/linux-headers-4.14.28-041428-generic_4.14.28-041428.201803190830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 29,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.29/linux-headers-4.14.29-041429_4.14.29-041429.201803211330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.29/linux-headers-4.14.29-041429-generic_4.14.29-041429.201803211330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 30,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.30/linux-headers-4.14.30-041430_4.14.30-041430.201803250538_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.30/linux-headers-4.14.30-041430-generic_4.14.30-041430.201803250538_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 31,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.31/linux-headers-4.14.31-041431_4.14.31-041431.201803281824_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.31/linux-headers-4.14.31-041431-generic_4.14.31-041431.201803281824_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 32,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.32/linux-headers-4.14.32-041432_4.14.32-041432.201803311811_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.32/linux-headers-4.14.32-041432-generic_4.14.32-041432.201803311811_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 33,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.33/linux-headers-4.14.33-041433_4.14.33-041433.201804080932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.33/linux-headers-4.14.33-041433-generic_4.14.33-041433.201804080932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 34,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.34/linux-headers-4.14.34-041434_4.14.34-041434.201804120731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.34/linux-headers-4.14.34-041434-generic_4.14.34-041434.201804120731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 35,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.35/linux-headers-4.14.35-041435_4.14.35-041435.201804190330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.35/linux-headers-4.14.35-041435-generic_4.14.35-041435.201804190330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 36,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.36/linux-headers-4.14.36-041436_4.14.36-041436.201804240906_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.36/linux-headers-4.14.36-041436-generic_4.14.36-041436.201804240906_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 37,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.37/linux-headers-4.14.37-041437_4.14.37-041437.201804261030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.37/linux-headers-4.14.37-041437-generic_4.14.37-041437.201804261030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 38,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.38/linux-headers-4.14.38-041438_4.14.38-041438.201806042303_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.38/linux-headers-4.14.38-041438-generic_4.14.38-041438.201806042303_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 39,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.39/linux-headers-4.14.39-041439_4.14.39-041439.201806051201_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.39/linux-headers-4.14.39-041439-generic_4.14.39-041439.201806051201_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 40,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.40/linux-headers-4.14.40-041440_4.14.40-041440.201806042114_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.40/linux-headers-4.14.40-041440-generic_4.14.40-041440.201806042114_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 41,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.41/linux-headers-4.14.41-041441_4.14.41-041441.201806042208_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.41/linux-headers-4.14.41-041441-generic_4.14.41-041441.201806042208_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 42,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.42/linux-headers-4.14.42-041442_4.14.42-041442.201806042157_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.42/linux-headers-4.14.42-041442-generic_4.14.42-041442.201806042157_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 43,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.43/linux-headers-4.14.43-041443_4.14.43-041443.201806041806_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.43/linux-headers-4.14.43-041443-generic_4.14.43-041443.201806041806_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 44,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.44/linux-headers-4.14.44-041444_4.14.44-041444.201806042228_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.44/linux-headers-4.14.44-041444-generic_4.14.44-041444.201806042228_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 45,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.45/linux-headers-4.14.45-041445_4.14.45-041445.201806050317_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.45/linux-headers-4.14.45-041445-generic_4.14.45-041445.201806050317_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 46,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.46/linux-headers-4.14.46-041446_4.14.46-041446.201806050100_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.46/linux-headers-4.14.46-041446-generic_4.14.46-041446.201806050100_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 47,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.47/linux-headers-4.14.47-041447_4.14.47-041447.201806050359_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.47/linux-headers-4.14.47-041447-generic_4.14.47-041447.201806050359_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 48,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.48/linux-headers-4.14.48-041448_4.14.48-041448.201806051318_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.48/linux-headers-4.14.48-041448-generic_4.14.48-041448.201806051318_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 49,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.49/linux-headers-4.14.49-041449_4.14.49-041449.201806112130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.49/linux-headers-4.14.49-041449-generic_4.14.49-041449.201806112130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 50,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.50/linux-headers-4.14.50-041450_4.14.50-041450.201806160924_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.50/linux-headers-4.14.50-041450-generic_4.14.50-041450.201806160924_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 51,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.51/linux-headers-4.14.51-041451_4.14.51-041451.201806201631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.51/linux-headers-4.14.51-041451-generic_4.14.51-041451.201806201631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 52,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.52/linux-headers-4.14.52-041452_4.14.52-041452.201806252131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.52/linux-headers-4.14.52-041452-generic_4.14.52-041452.201806252131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 53,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.53/linux-headers-4.14.53-041453_4.14.53-041453.201807030857_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.53/linux-headers-4.14.53-041453-generic_4.14.53-041453.201807030857_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 54,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.54/linux-headers-4.14.54-041454_4.14.54-041454.201807081031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.54/linux-headers-4.14.54-041454-generic_4.14.54-041454.201807081031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 55,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.55/linux-headers-4.14.55-041455_4.14.55-041455.201807111244_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.55/linux-headers-4.14.55-041455-generic_4.14.55-041455.201807111244_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 56,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.56/linux-headers-4.14.56-041456_4.14.56-041456.201807170758_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.56/linux-headers-4.14.56-041456-generic_4.14.56-041456.201807170758_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 57,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.57/linux-headers-4.14.57-041457_4.14.57-041457.201807221412_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.57/linux-headers-4.14.57-041457-generic_4.14.57-041457.201807221412_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 58,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.58/linux-headers-4.14.58-041458_4.14.58-041458.201807261230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.58/linux-headers-4.14.58-041458-generic_4.14.58-041458.201807261230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 59,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.59/linux-headers-4.14.59-041459_4.14.59-041459.201807280929_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.59/linux-headers-4.14.59-041459-generic_4.14.59-041459.201807280929_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 60,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.60/linux-headers-4.14.60-041460_4.14.60-041460.201808030232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.60/linux-headers-4.14.60-041460-generic_4.14.60-041460.201808030232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 61,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.61/linux-headers-4.14.61-041461_4.14.61-041461.201808061634_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.61/linux-headers-4.14.61-041461-generic_4.14.61-041461.201808061634_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 62,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.62/linux-headers-4.14.62-041462_4.14.62-041462.201808090814_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.62/linux-headers-4.14.62-041462-generic_4.14.62-041462.201808090814_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 63,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.63/linux-headers-4.14.63-041463_4.14.63-041463.201808151830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.63/linux-headers-4.14.63-041463-generic_4.14.63-041463.201808151830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 64,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.64/linux-headers-4.14.64-041464_4.14.64-041464.201808171730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.64/linux-headers-4.14.64-041464-generic_4.14.64-041464.201808171730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 65,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.65/linux-headers-4.14.65-041465_4.14.65-041465.201808181033_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.65/linux-headers-4.14.65-041465-generic_4.14.65-041465.201808181033_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 66,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.66/linux-headers-4.14.66-041466_4.14.66-041466.201808220733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.66/linux-headers-4.14.66-041466-generic_4.14.66-041466.201808220733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 67,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.67/linux-headers-4.14.67-041467_4.14.67-041467.201808240939_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.67/linux-headers-4.14.67-041467-generic_4.14.67-041467.201808240939_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 68,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.68/linux-headers-4.14.68-041468_4.14.68-041468.201809050531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.68/linux-headers-4.14.68-041468-generic_4.14.68-041468.201809050531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 69,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.69/linux-headers-4.14.69-041469_4.14.69-041469.201809091431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.69/linux-headers-4.14.69-041469-generic_4.14.69-041469.201809091431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 70,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.70/linux-headers-4.14.70-041470_4.14.70-041470.201809150531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.70/linux-headers-4.14.70-041470-generic_4.14.70-041470.201809150531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 72,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.72/linux-headers-4.14.72-041472_4.14.72-041472.201809260435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.72/linux-headers-4.14.72-041472-generic_4.14.72-041472.201809260435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 73,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.73/linux-headers-4.14.73-041473_4.14.73-041473.201809291240_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.73/linux-headers-4.14.73-041473-generic_4.14.73-041473.201809291240_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.74/linux-headers-4.14.74-041474_4.14.74-041474.201810040137_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.74/linux-headers-4.14.74-041474-generic_4.14.74-041474.201810040137_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 75,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.75/linux-headers-4.14.75-041475_4.14.75-041475.201810100435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.75/linux-headers-4.14.75-041475-generic_4.14.75-041475.201810100435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 76,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.76/linux-headers-4.14.76-041476_4.14.76-041476.201810130531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.76/linux-headers-4.14.76-041476-generic_4.14.76-041476.201810130531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 77,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.77/linux-headers-4.14.77-041477_4.14.77-041477.201810180434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.77/linux-headers-4.14.77-041477-generic_4.14.77-041477.201810180434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 78,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.78/linux-headers-4.14.78-041478_4.14.78-041478.201810200529_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.78/linux-headers-4.14.78-041478-generic_4.14.78-041478.201810200529_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 79,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.79/linux-headers-4.14.79-041479_4.14.79-041479.201811040934_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.79/linux-headers-4.14.79-041479-generic_4.14.79-041479.201811040934_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 80,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.80/linux-headers-4.14.80-041480_4.14.80-041480.201811101233_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.80/linux-headers-4.14.80-041480-generic_4.14.80-041480.201811101233_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 81,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.81/linux-headers-4.14.81-041481_4.14.81-041481.201811132127_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.81/linux-headers-4.14.81-041481-generic_4.14.81-041481.201811132127_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 82,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.82/linux-headers-4.14.82-041482_4.14.82-041482.201811210543_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.82/linux-headers-4.14.82-041482-generic_4.14.82-041482.201811210543_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 83,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.83/linux-headers-4.14.83-041483_4.14.83-041483.201811230331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.83/linux-headers-4.14.83-041483-generic_4.14.83-041483.201811230331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 84,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.84/linux-headers-4.14.84-041484_4.14.84-041484.201811271734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.84/linux-headers-4.14.84-041484-generic_4.14.84-041484.201811271734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 85,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.85/linux-headers-4.14.85-041485_4.14.85-041485.201812010530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.85/linux-headers-4.14.85-041485-generic_4.14.85-041485.201812010530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 86,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.86/linux-headers-4.14.86-041486_4.14.86-041486.201812052124_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.86/linux-headers-4.14.86-041486-generic_4.14.86-041486.201812052124_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 87,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.87/linux-headers-4.14.87-041487_4.14.87-041487.201812080833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.87/linux-headers-4.14.87-041487-generic_4.14.87-041487.201812080833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 88,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.88/linux-headers-4.14.88-041488_4.14.88-041488.201812130534_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.88/linux-headers-4.14.88-041488-generic_4.14.88-041488.201812130534_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 89,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.89/linux-headers-4.14.89-041489_4.14.89-041489.201812170432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.89/linux-headers-4.14.89-041489-generic_4.14.89-041489.201812170432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 91,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.91/linux-headers-4.14.91-041491_4.14.91-041491.201812290831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.91/linux-headers-4.14.91-041491-generic_4.14.91-041491.201812290831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 92,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.92/linux-headers-4.14.92-041492_4.14.92-041492.201901092326_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.92/linux-headers-4.14.92-041492-generic_4.14.92-041492.201901092326_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 93,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.93/linux-headers-4.14.93-041493_4.14.93-041493.201901131115_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.93/linux-headers-4.14.93-041493-generic_4.14.93-041493.201901131115_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 94,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.94/linux-headers-4.14.94-041494_4.14.94-041494.201901161855_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.94/linux-headers-4.14.94-041494-generic_4.14.94-041494.201901161855_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 95,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.95/linux-headers-4.14.95-041495_4.14.95-041495.201901230332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.95/linux-headers-4.14.95-041495-generic_4.14.95-041495.201901230332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 96,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.96/linux-headers-4.14.96-041496_4.14.96-041496.201901261054_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.96/linux-headers-4.14.96-041496-generic_4.14.96-041496.201901261054_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 97,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.97/linux-headers-4.14.97-041497_4.14.97-041497.201901310429_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.97/linux-headers-4.14.97-041497-generic_4.14.97-041497.201901310429_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 98,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.98/linux-headers-4.14.98-041498_4.14.98-041498.201902061900_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.98/linux-headers-4.14.98-041498-generic_4.14.98-041498.201902061900_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 99,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.99/linux-headers-4.14.99-041499_4.14.99-041499.201902121444_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.99/linux-headers-4.14.99-041499-generic_4.14.99-041499.201902121444_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 100,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.100/linux-headers-4.14.100-0414100_4.14.100-0414100.201902150437_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.100/linux-headers-4.14.100-0414100-generic_4.14.100-0414100.201902150437_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 101,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.101/linux-headers-4.14.101-0414101_4.14.101-0414101.201902151111_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.101/linux-headers-4.14.101-0414101-generic_4.14.101-0414101.201902151111_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 102,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.102/linux-headers-4.14.102-0414102_4.14.102-0414102.201902201154_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.102/linux-headers-4.14.102-0414102-generic_4.14.102-0414102.201902201154_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 103,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.103/linux-headers-4.14.103-0414103_4.14.103-0414103.201902231052_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.103/linux-headers-4.14.103-0414103-generic_4.14.103-0414103.201902231052_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 104,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.104/linux-headers-4.14.104-0414104_4.14.104-0414104.201902270640_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.104/linux-headers-4.14.104-0414104-generic_4.14.104-0414104.201902270640_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 105,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.105/linux-headers-4.14.105-0414105_4.14.105-0414105.201903051236_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.105/linux-headers-4.14.105-0414105-generic_4.14.105-0414105.201903051236_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 106,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.106/linux-headers-4.14.106-0414106_4.14.106-0414106.201903131950_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.106/linux-headers-4.14.106-0414106-generic_4.14.106-0414106.201903131950_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 107,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.107/linux-headers-4.14.107-0414107_4.14.107-0414107.201903191434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.107/linux-headers-4.14.107-0414107-generic_4.14.107-0414107.201903191434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 108,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.108/linux-headers-4.14.108-0414108_4.14.108-0414108.201903231437_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.108/linux-headers-4.14.108-0414108-generic_4.14.108-0414108.201903231437_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 109,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.109/linux-headers-4.14.109-0414109_4.14.109-0414109.201903271718_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.109/linux-headers-4.14.109-0414109-generic_4.14.109-0414109.201903271718_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 110,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.110/linux-headers-4.14.110-0414110_4.14.110-0414110.201904030252_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.110/linux-headers-4.14.110-0414110-generic_4.14.110-0414110.201904030252_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 111,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.111/linux-headers-4.14.111-0414111_4.14.111-0414111.201904052241_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.111/linux-headers-4.14.111-0414111-generic_4.14.111-0414111.201904052241_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 112,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.112/linux-headers-4.14.112-0414112_4.14.112-0414112.201904170835_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.112/linux-headers-4.14.112-0414112-generic_4.14.112-0414112.201904170835_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 113,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.113/linux-headers-4.14.113-0414113_4.14.113-0414113.201904200930_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.113/linux-headers-4.14.113-0414113-generic_4.14.113-0414113.201904200930_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 114,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.114/linux-headers-4.14.114-0414114_4.14.114-0414114.201904270558_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.114/linux-headers-4.14.114-0414114-generic_4.14.114-0414114.201904270558_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 115,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.115/linux-headers-4.14.115-0414115_4.14.115-0414115.201906051732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.115/linux-headers-4.14.115-0414115-generic_4.14.115-0414115.201906051732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 116,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.116/linux-headers-4.14.116-0414116_4.14.116-0414116.201905040435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.116/linux-headers-4.14.116-0414116-generic_4.14.116-0414116.201905040435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 117,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.117/linux-headers-4.14.117-0414117_4.14.117-0414117.201905080729_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.117/linux-headers-4.14.117-0414117-generic_4.14.117-0414117.201905080729_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 118,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.118/linux-headers-4.14.118-0414118_4.14.118-0414118.201905101231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.118/linux-headers-4.14.118-0414118-generic_4.14.118-0414118.201905101231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 119,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.119/linux-headers-4.14.119-0414119_4.14.119-0414119.201905141530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.119/linux-headers-4.14.119-0414119-generic_4.14.119-0414119.201905141530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 120,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.120/linux-headers-4.14.120-0414120_4.14.120-0414120.201905161610_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.120/linux-headers-4.14.120-0414120-generic_4.14.120-0414120.201905161610_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 121,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.121/linux-headers-4.14.121-0414121_4.14.121-0414121.201905211331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.121/linux-headers-4.14.121-0414121-generic_4.14.121-0414121.201905211331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 122,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.122/linux-headers-4.14.122-0414122_4.14.122-0414122.201906051405_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.122/linux-headers-4.14.122-0414122-generic_4.14.122-0414122.201906051405_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 123,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.123/linux-headers-4.14.123-0414123_4.14.123-0414123.201906051942_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.123/linux-headers-4.14.123-0414123-generic_4.14.123-0414123.201906051942_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 124,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.124/linux-headers-4.14.124-0414124_4.14.124-0414124.201906090432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.124/linux-headers-4.14.124-0414124-generic_4.14.124-0414124.201906090432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 125,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.125/linux-headers-4.14.125-0414125_4.14.125-0414125.201906111227_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.125/linux-headers-4.14.125-0414125-generic_4.14.125-0414125.201906111227_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 126,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.126/linux-headers-4.14.126-0414126_4.14.126-0414126.201906150631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.126/linux-headers-4.14.126-0414126-generic_4.14.126-0414126.201906150631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 127,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.127/linux-headers-4.14.127-0414127_4.14.127-0414127.201906171942_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.127/linux-headers-4.14.127-0414127-generic_4.14.127-0414127.201906171942_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 128,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.128/linux-headers-4.14.128-0414128_4.14.128-0414128.201906190332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.128/linux-headers-4.14.128-0414128-generic_4.14.128-0414128.201906190332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 129,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.129/linux-headers-4.14.129-0414129_4.14.129-0414129.201906220821_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.129/linux-headers-4.14.129-0414129-generic_4.14.129-0414129.201906220821_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 130,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.130/linux-headers-4.14.130-0414130_4.14.130-0414130.201906250038_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.130/linux-headers-4.14.130-0414130-generic_4.14.130-0414130.201906250038_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 131,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.131/linux-headers-4.14.131-0414131_4.14.131-0414131.201906270136_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.131/linux-headers-4.14.131-0414131-generic_4.14.131-0414131.201906270136_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 132,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.132/linux-headers-4.14.132-0414132_4.14.132-0414132.201907030842_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.132/linux-headers-4.14.132-0414132-generic_4.14.132-0414132.201907030842_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 133,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.133/linux-headers-4.14.133-0414133_4.14.133-0414133.201907100917_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.133/linux-headers-4.14.133-0414133-generic_4.14.133-0414133.201907100917_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 134,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.134/linux-headers-4.14.134-0414134_4.14.134-0414134.201907210912_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.134/linux-headers-4.14.134-0414134-generic_4.14.134-0414134.201907210912_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 135,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.135/linux-headers-4.14.135-0414135_4.14.135-0414135.201907310237_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 136,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.136/linux-headers-4.14.136-0414136_4.14.136-0414136.201908040442_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.136/linux-headers-4.14.136-0414136-generic_4.14.136-0414136.201908040442_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 137,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.137/linux-headers-4.14.137-0414137_4.14.137-0414137.201908061439_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.137/linux-headers-4.14.137-0414137-generic_4.14.137-0414137.201908061439_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 138,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.138/linux-headers-4.14.138-0414138_4.14.138-0414138.201908091240_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.138/linux-headers-4.14.138-0414138-generic_4.14.138-0414138.201908091240_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 139,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.139/linux-headers-4.14.139-0414139_4.14.139-0414139.201908160545_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.139/linux-headers-4.14.139-0414139-generic_4.14.139-0414139.201908160545_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 140,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.140/linux-headers-4.14.140-0414140_4.14.140-0414140.201908250540_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.140/linux-headers-4.14.140-0414140-generic_4.14.140-0414140.201908250540_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 141,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.141/linux-headers-4.14.141-0414141_4.14.141-0414141.201908290342_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.141/linux-headers-4.14.141-0414141-generic_4.14.141-0414141.201908290342_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 142,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.142/linux-headers-4.14.142-0414142_4.14.142-0414142.201909061035_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.142/linux-headers-4.14.142-0414142-generic_4.14.142-0414142.201909061035_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 143,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.143/linux-headers-4.14.143-0414143_4.14.143-0414143.201909101129_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.143/linux-headers-4.14.143-0414143-generic_4.14.143-0414143.201909101129_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 144,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.144/linux-headers-4.14.144-0414144_4.14.144-0414144.201909160828_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.144/linux-headers-4.14.144-0414144-generic_4.14.144-0414144.201909160828_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 145,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.145/linux-headers-4.14.145-0414145_4.14.145-0414145.201909190433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.145/linux-headers-4.14.145-0414145-generic_4.14.145-0414145.201909190433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 146,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.146/linux-headers-4.14.146-0414146_4.14.146-0414146.201909210738_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.146/linux-headers-4.14.146-0414146-generic_4.14.146-0414146.201909210738_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 147,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.147/linux-headers-4.14.147-0414147_4.14.147-0414147.201910050734_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.147/linux-headers-4.14.147-0414147-generic_4.14.147-0414147.201910050734_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 148,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.148/linux-headers-4.14.148-0414148_4.14.148-0414148.201910071338_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.148/linux-headers-4.14.148-0414148-generic_4.14.148-0414148.201910071338_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 149,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.149/linux-headers-4.14.149-0414149_4.14.149-0414149.201910111332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.149/linux-headers-4.14.149-0414149-generic_4.14.149-0414149.201910111332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 150,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.150/linux-headers-4.14.150-0414150_4.14.150-0414150.201910172201_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.150/linux-headers-4.14.150-0414150-generic_4.14.150-0414150.201910172201_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 151,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.151/linux-headers-4.14.151-0414151_4.14.151-0414151.201910290701_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.151/linux-headers-4.14.151-0414151-generic_4.14.151-0414151.201910290701_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 152,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.152/linux-headers-4.14.152-0414152_4.14.152-0414152.201911060751_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.152/linux-headers-4.14.152-0414152-generic_4.14.152-0414152.201911060751_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 153,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.153/linux-headers-4.14.153-0414153_4.14.153-0414153.201911101449_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.153/linux-headers-4.14.153-0414153-generic_4.14.153-0414153.201911101449_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 154,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.154/linux-headers-4.14.154-0414154_4.14.154-0414154.201911121843_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.154/linux-headers-4.14.154-0414154-generic_4.14.154-0414154.201911121843_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 155,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.155/linux-headers-4.14.155-0414155_4.14.155-0414155.201911201336_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.155/linux-headers-4.14.155-0414155-generic_4.14.155-0414155.201911201336_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 156,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.156/linux-headers-4.14.156-0414156_4.14.156-0414156.201911240436_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.156/linux-headers-4.14.156-0414156-generic_4.14.156-0414156.201911240436_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 157,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.157/linux-headers-4.14.157-0414157_4.14.157-0414157.201912010440_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.157/linux-headers-4.14.157-0414157-generic_4.14.157-0414157.201912010440_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 158,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.158/linux-headers-4.14.158-0414158_4.14.158-0414158.201912051542_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.158/linux-headers-4.14.158-0414158-generic_4.14.158-0414158.201912051542_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 159,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.159/linux-headers-4.14.159-0414159_4.14.159-0414159.201912171622_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.159/linux-headers-4.14.159-0414159-generic_4.14.159-0414159.201912171622_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 160,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.160/linux-headers-4.14.160-0414160_4.14.160-0414160.201912210542_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.160/linux-headers-4.14.160-0414160-generic_4.14.160-0414160.201912210542_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 161,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.161/linux-headers-4.14.161-0414161_4.14.161-0414161.201912311234_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.161/linux-headers-4.14.161-0414161-generic_4.14.161-0414161.201912311234_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 162,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.162/linux-headers-4.14.162-0414162_4.14.162-0414162.202001040937_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.162/linux-headers-4.14.162-0414162-generic_4.14.162-0414162.202001040937_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 163,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.163/linux-headers-4.14.163-0414163_4.14.163-0414163.202001090543_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.163/linux-headers-4.14.163-0414163-generic_4.14.163-0414163.202001090543_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 164,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.164/linux-headers-4.14.164-0414164_4.14.164-0414164.202001121244_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.164/linux-headers-4.14.164-0414164-generic_4.14.164-0414164.202001121244_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 165,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.165/linux-headers-4.14.165-0414165_4.14.165-0414165.202001142138_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.165/linux-headers-4.14.165-0414165-generic_4.14.165-0414165.202001142138_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 166,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.166/linux-headers-4.14.166-0414166_4.14.166-0414166.202001172029_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.166/linux-headers-4.14.166-0414166-generic_4.14.166-0414166.202001172029_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 167,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.167/linux-headers-4.14.167-0414167_4.14.167-0414167.202001230928_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.167/linux-headers-4.14.167-0414167-generic_4.14.167-0414167.202001230928_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 168,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.168/linux-headers-4.14.168-0414168_4.14.168-0414168.202001271437_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.168/linux-headers-4.14.168-0414168-generic_4.14.168-0414168.202001271437_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 169,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.169/linux-headers-4.14.169-0414169_4.14.169-0414169.202001300343_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.169/linux-headers-4.14.169-0414169-generic_4.14.169-0414169.202001300343_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 170,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.170/linux-headers-4.14.170-0414170_4.14.170-0414170.202002051535_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.170/linux-headers-4.14.170-0414170-generic_4.14.170-0414170.202002051535_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 171,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.171/linux-headers-4.14.171-0414171_4.14.171-0414171.202002151207_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.171/linux-headers-4.14.171-0414171-generic_4.14.171-0414171.202002151207_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 172,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.172/linux-headers-4.14.172-0414172_4.14.172-0414172.202002281642_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.172/linux-headers-4.14.172-0414172-generic_4.14.172-0414172.202002281642_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 173,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.173/linux-headers-4.14.173-0414173_4.14.173-0414173.202003111439_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.173/linux-headers-4.14.173-0414173-generic_4.14.173-0414173.202003111439_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 174,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.174/linux-headers-4.14.174-0414174_4.14.174-0414174.202003200742_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.174/linux-headers-4.14.174-0414174-generic_4.14.174-0414174.202003200742_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 175,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.175/linux-headers-4.14.175-0414175_4.14.175-0414175.202004021211_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.175/linux-headers-4.14.175-0414175-generic_4.14.175-0414175.202004021211_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 14,
			Patch: 176,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.176/linux-headers-4.14.176-0414176_4.14.176-0414176.202004130536_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.14.176/linux-headers-4.14.176-0414176-generic_4.14.176-0414176.202004130536_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15/linux-headers-4.15.0-041500_4.15.0-041500.201802011154_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15/linux-headers-4.15.0-041500-generic_4.15.0-041500.201802011154_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.1/linux-headers-4.15.1-041501_4.15.1-041501.201802031831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.1/linux-headers-4.15.1-041501-generic_4.15.1-041501.201802031831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.2/linux-headers-4.15.2-041502_4.15.2-041502.201802072230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.2/linux-headers-4.15.2-041502-generic_4.15.2-041502.201802072230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.3/linux-headers-4.15.3-041503_4.15.3-041503.201802120730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.3/linux-headers-4.15.3-041503-generic_4.15.3-041503.201802120730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.4/linux-headers-4.15.4-041504_4.15.4-041504.201802162207_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.4/linux-headers-4.15.4-041504-generic_4.15.4-041504.201802162207_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.5/linux-headers-4.15.5-041505_4.15.5-041505.201802261304_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.5/linux-headers-4.15.5-041505-generic_4.15.5-041505.201802261304_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.7/linux-headers-4.15.7-041507_4.15.7-041507.201802280530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.7/linux-headers-4.15.7-041507-generic_4.15.7-041507.201802280530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.8/linux-headers-4.15.8-041508_4.15.8-041508.201803091630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.8/linux-headers-4.15.8-041508-generic_4.15.8-041508.201803091630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.9/linux-headers-4.15.9-041509_4.15.9-041509.201803111231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.9/linux-headers-4.15.9-041509-generic_4.15.9-041509.201803111231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.10/linux-headers-4.15.10-041510_4.15.10-041510.201803152130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.10/linux-headers-4.15.10-041510-generic_4.15.10-041510.201803152130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.11/linux-headers-4.15.11-041511_4.15.11-041511.201803190530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.11/linux-headers-4.15.11-041511-generic_4.15.11-041511.201803190530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.12/linux-headers-4.15.12-041512_4.15.12-041512.201803211230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.12/linux-headers-4.15.12-041512-generic_4.15.12-041512.201803211230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.13/linux-headers-4.15.13-041513_4.15.13-041513.201803250910_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.13/linux-headers-4.15.13-041513-generic_4.15.13-041513.201803250910_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.14/linux-headers-4.15.14-041514_4.15.14-041514.201803281351_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.14/linux-headers-4.15.14-041514-generic_4.15.14-041514.201803281351_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.15/linux-headers-4.15.15-041515_4.15.15-041515.201803311331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.15/linux-headers-4.15.15-041515-generic_4.15.15-041515.201803311331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.16/linux-headers-4.15.16-041516_4.15.16-041516.201804080932_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.16/linux-headers-4.15.16-041516-generic_4.15.16-041516.201804080932_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.17/linux-headers-4.15.17-041517_4.15.17-041517.201804120730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.17/linux-headers-4.15.17-041517-generic_4.15.17-041517.201804120730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 15,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.18/linux-headers-4.15.18-041518_4.15.18-041518.201804190330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.15.18/linux-headers-4.15.18-041518-generic_4.15.18-041518.201804190330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16/linux-headers-4.16.0-041600_4.16.0-041600.201804012230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16/linux-headers-4.16.0-041600-generic_4.16.0-041600.201804012230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.1/linux-headers-4.16.1-041601_4.16.1-041601.201804081334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.1/linux-headers-4.16.1-041601-generic_4.16.1-041601.201804081334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.2/linux-headers-4.16.2-041602_4.16.2-041602.201804121130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.2/linux-headers-4.16.2-041602-generic_4.16.2-041602.201804121130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.3/linux-headers-4.16.3-041603_4.16.3-041603.201804190730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.3/linux-headers-4.16.3-041603-generic_4.16.3-041603.201804190730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.4/linux-headers-4.16.4-041604_4.16.4-041604.201804240433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.4/linux-headers-4.16.4-041604-generic_4.16.4-041604.201804240433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.5/linux-headers-4.16.5-041605_4.16.5-041605.201804260630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.5/linux-headers-4.16.5-041605-generic_4.16.5-041605.201804260630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.6/linux-headers-4.16.6-041606_4.16.6-041606.201806042022_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.6/linux-headers-4.16.6-041606-generic_4.16.6-041606.201806042022_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.7/linux-headers-4.16.7-041607_4.16.7-041607.201806042046_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.7/linux-headers-4.16.7-041607-generic_4.16.7-041607.201806042046_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.8/linux-headers-4.16.8-041608_4.16.8-041608.201806041956_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.8/linux-headers-4.16.8-041608-generic_4.16.8-041608.201806041956_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.9/linux-headers-4.16.9-041609_4.16.9-041609.201806041905_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.9/linux-headers-4.16.9-041609-generic_4.16.9-041609.201806041905_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.10/linux-headers-4.16.10-041610_4.16.10-041610.201806050235_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.10/linux-headers-4.16.10-041610-generic_4.16.10-041610.201806050235_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.11/linux-headers-4.16.11-041611_4.16.11-041611.201806050359_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.11/linux-headers-4.16.11-041611-generic_4.16.11-041611.201806050359_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.12/linux-headers-4.16.12-041612_4.16.12-041612.201806042137_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.12/linux-headers-4.16.12-041612-generic_4.16.12-041612.201806042137_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.13/linux-headers-4.16.13-041613_4.16.13-041613.201806050742_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.13/linux-headers-4.16.13-041613-generic_4.16.13-041613.201806050742_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.14/linux-headers-4.16.14-041614_4.16.14-041614.201806051643_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.14/linux-headers-4.16.14-041614-generic_4.16.14-041614.201806051643_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.15/linux-headers-4.16.15-041615_4.16.15-041615.201806111730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.15/linux-headers-4.16.15-041615-generic_4.16.15-041615.201806111730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.16/linux-headers-4.16.16-041616_4.16.16-041616.201806160515_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.16/linux-headers-4.16.16-041616-generic_4.16.16-041616.201806160515_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.17/linux-headers-4.16.17-041617_4.16.17-041617.201806201630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.17/linux-headers-4.16.17-041617-generic_4.16.17-041617.201806201630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 16,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.18/linux-headers-4.16.18-041618_4.16.18-041618.201806252030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.16.18/linux-headers-4.16.18-041618-generic_4.16.18-041618.201806252030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17/linux-headers-4.17.0-041700_4.17.0-041700.201806041953_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17/linux-headers-4.17.0-041700-generic_4.17.0-041700.201806041953_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.1/linux-headers-4.17.1-041701_4.17.1-041701.201806111730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.1/linux-headers-4.17.1-041701-generic_4.17.1-041701.201806111730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.2/linux-headers-4.17.2-041702_4.17.2-041702.201806160433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.2/linux-headers-4.17.2-041702-generic_4.17.2-041702.201806160433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.3/linux-headers-4.17.3-041703_4.17.3-041703.201806252030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.3/linux-headers-4.17.3-041703-generic_4.17.3-041703.201806252030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.4/linux-headers-4.17.4-041704_4.17.4-041704.201807031235_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.4/linux-headers-4.17.4-041704-generic_4.17.4-041704.201807031235_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.5/linux-headers-4.17.5-041705_4.17.5-041705.201807081431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.5/linux-headers-4.17.5-041705-generic_4.17.5-041705.201807081431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.6/linux-headers-4.17.6-041706_4.17.6-041706.201807111639_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.6/linux-headers-4.17.6-041706-generic_4.17.6-041706.201807111639_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.7/linux-headers-4.17.7-041707_4.17.7-041707.201807171133_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.7/linux-headers-4.17.7-041707-generic_4.17.7-041707.201807171133_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.8/linux-headers-4.17.8-041708_4.17.8-041708.201807180730_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.8/linux-headers-4.17.8-041708-generic_4.17.8-041708.201807180730_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.9/linux-headers-4.17.9-041709_4.17.9-041709.201807221110_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.9/linux-headers-4.17.9-041709-generic_4.17.9-041709.201807221110_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.10/linux-headers-4.17.10-041710_4.17.10-041710.201807260825_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.10/linux-headers-4.17.10-041710-generic_4.17.10-041710.201807260825_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.11/linux-headers-4.17.11-041711_4.17.11-041711.201807280505_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.11/linux-headers-4.17.11-041711-generic_4.17.11-041711.201807280505_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.12/linux-headers-4.17.12-041712_4.17.12-041712.201808030231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.12/linux-headers-4.17.12-041712-generic_4.17.12-041712.201808030231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.13/linux-headers-4.17.13-041713_4.17.13-041713.201808061132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.13/linux-headers-4.17.13-041713-generic_4.17.13-041713.201808061132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.14/linux-headers-4.17.14-041714_4.17.14-041714.201808090733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.14/linux-headers-4.17.14-041714-generic_4.17.14-041714.201808090733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.15/linux-headers-4.17.15-041715_4.17.15-041715.201808151332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.15/linux-headers-4.17.15-041715-generic_4.17.15-041715.201808151332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.16/linux-headers-4.17.16-041716_4.17.16-041716.201808171644_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.16/linux-headers-4.17.16-041716-generic_4.17.16-041716.201808171644_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.17/linux-headers-4.17.17-041717_4.17.17-041717.201808180630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.17/linux-headers-4.17.17-041717-generic_4.17.17-041717.201808180630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.18/linux-headers-4.17.18-041718_4.17.18-041718.201808220323_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.18/linux-headers-4.17.18-041718-generic_4.17.18-041718.201808220323_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 17,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.19/linux-headers-4.17.19-041719_4.17.19-041719.201808240919_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.17.19/linux-headers-4.17.19-041719-generic_4.17.19-041719.201808240919_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18/linux-headers-4.18.0-041800_4.18.0-041800.201808122131_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18/linux-headers-4.18.0-041800-generic_4.18.0-041800.201808122131_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.1/linux-headers-4.18.1-041801_4.18.1-041801.201808151233_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.1/linux-headers-4.18.1-041801-generic_4.18.1-041801.201808151233_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.2/linux-headers-4.18.2-041802_4.18.2-041802.201808171631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.2/linux-headers-4.18.2-041802-generic_4.18.2-041802.201808171631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.3/linux-headers-4.18.3-041803_4.18.3-041803.201808180530_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.3/linux-headers-4.18.3-041803-generic_4.18.3-041803.201808180530_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.4/linux-headers-4.18.4-041804_4.18.4-041804.201808220230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.4/linux-headers-4.18.4-041804-generic_4.18.4-041804.201808220230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.5/linux-headers-4.18.5-041805_4.18.5-041805.201808241320_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.5/linux-headers-4.18.5-041805-generic_4.18.5-041805.201808241320_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.6/linux-headers-4.18.6-041806_4.18.6-041806.201809050847_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.6/linux-headers-4.18.6-041806-generic_4.18.6-041806.201809050847_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.7/linux-headers-4.18.7-041807_4.18.7-041807.201809090930_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.7/linux-headers-4.18.7-041807-generic_4.18.7-041807.201809090930_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.8/linux-headers-4.18.8-041808_4.18.8-041808.201809150431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.8/linux-headers-4.18.8-041808-generic_4.18.8-041808.201809150431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.10/linux-headers-4.18.10-041810_4.18.10-041810.201809260332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.10/linux-headers-4.18.10-041810-generic_4.18.10-041810.201809260332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.11/linux-headers-4.18.11-041811_4.18.11-041811.201809290731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.11/linux-headers-4.18.11-041811-generic_4.18.11-041811.201809290731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.12/linux-headers-4.18.12-041812_4.18.12-041812.201810032137_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.12/linux-headers-4.18.12-041812-generic_4.18.12-041812.201810032137_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.13/linux-headers-4.18.13-041813_4.18.13-041813.201810100332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.13/linux-headers-4.18.13-041813-generic_4.18.13-041813.201810100332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.14/linux-headers-4.18.14-041814_4.18.14-041814.201810130431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.14/linux-headers-4.18.14-041814-generic_4.18.14-041814.201810130431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.15/linux-headers-4.18.15-041815_4.18.15-041815.201810180430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.15/linux-headers-4.18.15-041815-generic_4.18.15-041815.201810180430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.16/linux-headers-4.18.16-041816_4.18.16-041816.201810200431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.16/linux-headers-4.18.16-041816-generic_4.18.16-041816.201810200431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.17/linux-headers-4.18.17-041817_4.18.17-041817.201812031201_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.17/linux-headers-4.18.17-041817-generic_4.18.17-041817.201812031201_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.18/linux-headers-4.18.18-041818_4.18.18-041818.201812031529_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.18/linux-headers-4.18.18-041818-generic_4.18.18-041818.201812031529_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.19/linux-headers-4.18.19-041819_4.18.19-041819.201812030823_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.19/linux-headers-4.18.19-041819-generic_4.18.19-041819.201812030823_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 18,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.20/linux-headers-4.18.20-041820_4.18.20-041820.201812030624_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.18.20/linux-headers-4.18.20-041820-generic_4.18.20-041820.201812030624_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19/linux-headers-4.19.0-041900_4.19.0-041900.201810221809_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19/linux-headers-4.19.0-041900-generic_4.19.0-041900.201810221809_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.1/linux-headers-4.19.1-041901_4.19.1-041901.201812031343_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.1/linux-headers-4.19.1-041901-generic_4.19.1-041901.201812031343_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.2/linux-headers-4.19.2-041902_4.19.2-041902.201812031253_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.2/linux-headers-4.19.2-041902-generic_4.19.2-041902.201812031253_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.3/linux-headers-4.19.3-041903_4.19.3-041903.201812030715_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.3/linux-headers-4.19.3-041903-generic_4.19.3-041903.201812030715_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.4/linux-headers-4.19.4-041904_4.19.4-041904.201812030922_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.4/linux-headers-4.19.4-041904-generic_4.19.4-041904.201812030922_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.5/linux-headers-4.19.5-041905_4.19.5-041905.201812031110_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.5/linux-headers-4.19.5-041905-generic_4.19.5-041905.201812031110_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.6/linux-headers-4.19.6-041906_4.19.6-041906.201812030857_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.6/linux-headers-4.19.6-041906-generic_4.19.6-041906.201812030857_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.7/linux-headers-4.19.7-041907_4.19.7-041907.201812052010_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.7/linux-headers-4.19.7-041907-generic_4.19.7-041907.201812052010_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.8/linux-headers-4.19.8-041908_4.19.8-041908.201812080831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.8/linux-headers-4.19.8-041908-generic_4.19.8-041908.201812080831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.9/linux-headers-4.19.9-041909_4.19.9-041909.201812130432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.9/linux-headers-4.19.9-041909-generic_4.19.9-041909.201812130432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.10/linux-headers-4.19.10-041910_4.19.10-041910.201812170433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.10/linux-headers-4.19.10-041910-generic_4.19.10-041910.201812170433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.11/linux-headers-4.19.11-041911_4.19.11-041911.201812191931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.11/linux-headers-4.19.11-041911-generic_4.19.11-041911.201812191931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.13/linux-headers-4.19.13-041913_4.19.13-041913.201812291331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.13/linux-headers-4.19.13-041913-generic_4.19.13-041913.201812291331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.14/linux-headers-4.19.14-041914_4.19.14-041914.201901100338_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.14/linux-headers-4.19.14-041914-generic_4.19.14-041914.201901100338_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.15/linux-headers-4.19.15-041915_4.19.15-041915.201901130432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.15/linux-headers-4.19.15-041915-generic_4.19.15-041915.201901130432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.16/linux-headers-4.19.16-041916_4.19.16-041916.201901162330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.16/linux-headers-4.19.16-041916-generic_4.19.16-041916.201901162330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.17/linux-headers-4.19.17-041917_4.19.17-041917.201901221712_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.17/linux-headers-4.19.17-041917-generic_4.19.17-041917.201901221712_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.18/linux-headers-4.19.18-041918_4.19.18-041918.201901260539_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.18/linux-headers-4.19.18-041918-generic_4.19.18-041918.201901260539_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.19/linux-headers-4.19.19-041919_4.19.19-041919.201901310331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.19/linux-headers-4.19.19-041919-generic_4.19.19-041919.201901310331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.20/linux-headers-4.19.20-041920_4.19.20-041920.201902061341_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.20/linux-headers-4.19.20-041920-generic_4.19.20-041920.201902061341_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.21/linux-headers-4.19.21-041921_4.19.21-041921.201902121440_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.21/linux-headers-4.19.21-041921-generic_4.19.21-041921.201902121440_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 22,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.22/linux-headers-4.19.22-041922_4.19.22-041922.201902150332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.22/linux-headers-4.19.22-041922-generic_4.19.22-041922.201902150332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 23,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.23/linux-headers-4.19.23-041923_4.19.23-041923.201902150553_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.23/linux-headers-4.19.23-041923-generic_4.19.23-041923.201902150553_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 24,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.24/linux-headers-4.19.24-041924_4.19.24-041924.201902200632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.24/linux-headers-4.19.24-041924-generic_4.19.24-041924.201902200632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 25,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.25/linux-headers-4.19.25-041925_4.19.25-041925.201902230535_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.25/linux-headers-4.19.25-041925-generic_4.19.25-041925.201902230535_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 26,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.26/linux-headers-4.19.26-041926_4.19.26-041926.201902270533_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.26/linux-headers-4.19.26-041926-generic_4.19.26-041926.201902270533_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 27,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.27/linux-headers-4.19.27-041927_4.19.27-041927.201903051835_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.27/linux-headers-4.19.27-041927-generic_4.19.27-041927.201903051835_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 28,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.28/linux-headers-4.19.28-041928_4.19.28-041928.201903100334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.28/linux-headers-4.19.28-041928-generic_4.19.28-041928.201903100334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 29,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.29/linux-headers-4.19.29-041929_4.19.29-041929.201903132331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.29/linux-headers-4.19.29-041929-generic_4.19.29-041929.201903132331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 30,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.30/linux-headers-4.19.30-041930_4.19.30-041930.201903191032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.30/linux-headers-4.19.30-041930-generic_4.19.30-041930.201903191032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 31,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.31/linux-headers-4.19.31-041931_4.19.31-041931.201903231635_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.31/linux-headers-4.19.31-041931-generic_4.19.31-041931.201903231635_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 32,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.32/linux-headers-4.19.32-041932_4.19.32-041932.201903271306_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.32/linux-headers-4.19.32-041932-generic_4.19.32-041932.201903271306_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 33,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.33/linux-headers-4.19.33-041933_4.19.33-041933.201904030634_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.33/linux-headers-4.19.33-041933-generic_4.19.33-041933.201904030634_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 34,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.34/linux-headers-4.19.34-041934_4.19.34-041934.201904051741_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.34/linux-headers-4.19.34-041934-generic_4.19.34-041934.201904051741_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 35,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.35/linux-headers-4.19.35-041935_4.19.35-041935.201904170334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.35/linux-headers-4.19.35-041935-generic_4.19.35-041935.201904170334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 36,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.36/linux-headers-4.19.36-041936_4.19.36-041936.201904200430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.36/linux-headers-4.19.36-041936-generic_4.19.36-041936.201904200430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 37,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.37/linux-headers-4.19.37-041937_4.19.37-041937.201904270929_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.37/linux-headers-4.19.37-041937-generic_4.19.37-041937.201904270929_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 38,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.38/linux-headers-4.19.38-041938_4.19.38-041938.201905020956_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.38/linux-headers-4.19.38-041938-generic_4.19.38-041938.201905020956_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 39,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.39/linux-headers-4.19.39-041939_4.19.39-041939.201905040435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.39/linux-headers-4.19.39-041939-generic_4.19.39-041939.201905040435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 40,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.40/linux-headers-4.19.40-041940_4.19.40-041940.201906051454_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.40/linux-headers-4.19.40-041940-generic_4.19.40-041940.201906051454_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 41,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.41/linux-headers-4.19.41-041941_4.19.41-041941.201905080231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.41/linux-headers-4.19.41-041941-generic_4.19.41-041941.201905080231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 42,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.42/linux-headers-4.19.42-041942_4.19.42-041942.201905101231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.42/linux-headers-4.19.42-041942-generic_4.19.42-041942.201905101231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 43,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.43/linux-headers-4.19.43-041943_4.19.43-041943.201905141927_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.43/linux-headers-4.19.43-041943-generic_4.19.43-041943.201905141927_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 44,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.44/linux-headers-4.19.44-041944_4.19.44-041944.201905161958_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.44/linux-headers-4.19.44-041944-generic_4.19.44-041944.201905161958_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 45,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.45/linux-headers-4.19.45-041945_4.19.45-041945.201905220231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.45/linux-headers-4.19.45-041945-generic_4.19.45-041945.201905220231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 46,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.46/linux-headers-4.19.46-041946_4.19.46-041946.201905251332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.46/linux-headers-4.19.46-041946-generic_4.19.46-041946.201905251332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 47,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.47/linux-headers-4.19.47-041947_4.19.47-041947.201905311431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.47/linux-headers-4.19.47-041947-generic_4.19.47-041947.201905311431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 48,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.48/linux-headers-4.19.48-041948_4.19.48-041948.201906040331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.48/linux-headers-4.19.48-041948-generic_4.19.48-041948.201906040331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 49,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.49/linux-headers-4.19.49-041949_4.19.49-041949.201906090431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.49/linux-headers-4.19.49-041949-generic_4.19.49-041949.201906090431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 50,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.50/linux-headers-4.19.50-041950_4.19.50-041950.201906110735_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.50/linux-headers-4.19.50-041950-generic_4.19.50-041950.201906110735_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 51,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.51/linux-headers-4.19.51-041951_4.19.51-041951.201906151031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.51/linux-headers-4.19.51-041951-generic_4.19.51-041951.201906151031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 52,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.52/linux-headers-4.19.52-041952_4.19.52-041952.201906171539_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.52/linux-headers-4.19.52-041952-generic_4.19.52-041952.201906171539_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 53,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.53/linux-headers-4.19.53-041953_4.19.53-041953.201906190731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.53/linux-headers-4.19.53-041953-generic_4.19.53-041953.201906190731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 54,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.54/linux-headers-4.19.54-041954_4.19.54-041954.201906220333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.54/linux-headers-4.19.54-041954-generic_4.19.54-041954.201906220333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 55,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.55/linux-headers-4.19.55-041955_4.19.55-041955.201906221031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.55/linux-headers-4.19.55-041955-generic_4.19.55-041955.201906221031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 56,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.56/linux-headers-4.19.56-041956_4.19.56-041956.201906250431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.56/linux-headers-4.19.56-041956-generic_4.19.56-041956.201906250431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 57,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.57/linux-headers-4.19.57-041957_4.19.57-041957.201907031234_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.57/linux-headers-4.19.57-041957-generic_4.19.57-041957.201907031234_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 58,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.58/linux-headers-4.19.58-041958_4.19.58-041958.201907100436_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.58/linux-headers-4.19.58-041958-generic_4.19.58-041958.201907100436_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 59,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.59/linux-headers-4.19.59-041959_4.19.59-041959.201907140332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.59/linux-headers-4.19.59-041959-generic_4.19.59-041959.201907140332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 60,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.60/linux-headers-4.19.60-041960_4.19.60-041960.201907210912_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.60/linux-headers-4.19.60-041960-generic_4.19.60-041960.201907210912_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 61,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.61/linux-headers-4.19.61-041961_4.19.61-041961.201907260439_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.61/linux-headers-4.19.61-041961-generic_4.19.61-041961.201907260439_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 62,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.62/linux-headers-4.19.62-041962_4.19.62-041962.201907280331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.62/linux-headers-4.19.62-041962-generic_4.19.62-041962.201907280331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 63,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.63/linux-headers-4.19.63-041963_4.19.63-041963.201907310632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.63/linux-headers-4.19.63-041963-generic_4.19.63-041963.201907310632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 65,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.65/linux-headers-4.19.65-041965_4.19.65-041965.201908061833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.65/linux-headers-4.19.65-041965-generic_4.19.65-041965.201908061833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 66,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.66/linux-headers-4.19.66-041966_4.19.66-041966.201908091231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.66/linux-headers-4.19.66-041966-generic_4.19.66-041966.201908091231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 67,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.67/linux-headers-4.19.67-041967_4.19.67-041967.201908160941_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.67/linux-headers-4.19.67-041967-generic_4.19.67-041967.201908160941_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 68,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.68/linux-headers-4.19.68-041968_4.19.68-041968.201908250538_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.68/linux-headers-4.19.68-041968-generic_4.19.68-041968.201908250538_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 69,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.69/linux-headers-4.19.69-041969_4.19.69-041969.201908290733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.69/linux-headers-4.19.69-041969-generic_4.19.69-041969.201908290733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 70,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.70/linux-headers-4.19.70-041970_4.19.70-041970.201909060934_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.70/linux-headers-4.19.70-041970-generic_4.19.70-041970.201909060934_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 71,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.71/linux-headers-4.19.71-041971_4.19.71-041971.201909061148_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.71/linux-headers-4.19.71-041971-generic_4.19.71-041971.201909061148_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 72,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.72/linux-headers-4.19.72-041972_4.19.72-041972.201909100635_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.72/linux-headers-4.19.72-041972-generic_4.19.72-041972.201909100635_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 73,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.73/linux-headers-4.19.73-041973_4.19.73-041973.201909160334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.73/linux-headers-4.19.73-041973-generic_4.19.73-041973.201909160334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 74,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.74/linux-headers-4.19.74-041974_4.19.74-041974.201909190833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.74/linux-headers-4.19.74-041974-generic_4.19.74-041974.201909190833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 75,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.75/linux-headers-4.19.75-041975_4.19.75-041975.201909210733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.75/linux-headers-4.19.75-041975-generic_4.19.75-041975.201909210733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 76,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.76/linux-headers-4.19.76-041976_4.19.76-041976.201910010731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.76/linux-headers-4.19.76-041976-generic_4.19.76-041976.201910010731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 77,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.77/linux-headers-4.19.77-041977_4.19.77-041977.201910051424_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.77/linux-headers-4.19.77-041977-generic_4.19.77-041977.201910051424_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 78,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.78/linux-headers-4.19.78-041978_4.19.78-041978.201910071333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.78/linux-headers-4.19.78-041978-generic_4.19.78-041978.201910071333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 79,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.79/linux-headers-4.19.79-041979_4.19.79-041979.201910111332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.79/linux-headers-4.19.79-041979-generic_4.19.79-041979.201910111332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 80,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.80/linux-headers-4.19.80-041980_4.19.80-041980.201910180137_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.80/linux-headers-4.19.80-041980-generic_4.19.80-041980.201910180137_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 81,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.81/linux-headers-4.19.81-041981_4.19.81-041981.201910290545_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.81/linux-headers-4.19.81-041981-generic_4.19.81-041981.201910290545_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 82,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.82/linux-headers-4.19.82-041982_4.19.82-041982.201911060906_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.82/linux-headers-4.19.82-041982-generic_4.19.82-041982.201911060906_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 83,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.83/linux-headers-4.19.83-041983_4.19.83-041983.201911100902_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.83/linux-headers-4.19.83-041983-generic_4.19.83-041983.201911100902_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 84,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.84/linux-headers-4.19.84-041984_4.19.84-041984.201911121836_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.84/linux-headers-4.19.84-041984-generic_4.19.84-041984.201911121836_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 85,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.85/linux-headers-4.19.85-041985_4.19.85-041985.201911201931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.85/linux-headers-4.19.85-041985-generic_4.19.85-041985.201911201931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 86,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.86/linux-headers-4.19.86-041986_4.19.86-041986.201911240340_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.86/linux-headers-4.19.86-041986-generic_4.19.86-041986.201911240340_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 87,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.87/linux-headers-4.19.87-041987_4.19.87-041987.201912010935_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.87/linux-headers-4.19.87-041987-generic_4.19.87-041987.201912010935_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 88,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.88/linux-headers-4.19.88-041988_4.19.88-041988.201912050934_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.88/linux-headers-4.19.88-041988-generic_4.19.88-041988.201912050934_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 89,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.89/linux-headers-4.19.89-041989_4.19.89-041989.201912130343_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.89/linux-headers-4.19.89-041989-generic_4.19.89-041989.201912130343_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 90,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.90/linux-headers-4.19.90-041990_4.19.90-041990.201912172223_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.90/linux-headers-4.19.90-041990-generic_4.19.90-041990.201912172223_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 91,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.91/linux-headers-4.19.91-041991_4.19.91-041991.201912210701_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.91/linux-headers-4.19.91-041991-generic_4.19.91-041991.201912210701_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 92,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.92/linux-headers-4.19.92-041992_4.19.92-041992.201912311134_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.92/linux-headers-4.19.92-041992-generic_4.19.92-041992.201912311134_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 93,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.93/linux-headers-4.19.93-041993_4.19.93-041993.202001041939_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.93/linux-headers-4.19.93-041993-generic_4.19.93-041993.202001041939_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 94,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.94/linux-headers-4.19.94-041994_4.19.94-041994.202001090535_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.94/linux-headers-4.19.94-041994-generic_4.19.94-041994.202001090535_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 95,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.95/linux-headers-4.19.95-041995_4.19.95-041995.202001120738_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.95/linux-headers-4.19.95-041995-generic_4.19.95-041995.202001120738_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 96,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.96/linux-headers-4.19.96-041996_4.19.96-041996.202001141532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.96/linux-headers-4.19.96-041996-generic_4.19.96-041996.202001141532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 97,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.97/linux-headers-4.19.97-041997_4.19.97-041997.202001171458_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.97/linux-headers-4.19.97-041997-generic_4.19.97-041997.202001171458_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 98,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.98/linux-headers-4.19.98-041998_4.19.98-041998.202001230334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.98/linux-headers-4.19.98-041998-generic_4.19.98-041998.202001230334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 99,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.99/linux-headers-4.19.99-041999_4.19.99-041999.202001271433_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.99/linux-headers-4.19.99-041999-generic_4.19.99-041999.202001271433_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 100,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.100/linux-headers-4.19.100-0419100_4.19.100-0419100.202001300442_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.100/linux-headers-4.19.100-0419100-generic_4.19.100-0419100.202001300442_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 101,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.101/linux-headers-4.19.101-0419101_4.19.101-0419101.202002010532_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.101/linux-headers-4.19.101-0419101-generic_4.19.101-0419101.202002010532_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 103,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.103/linux-headers-4.19.103-0419103_4.19.103-0419103.202002111333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.103/linux-headers-4.19.103-0419103-generic_4.19.103-0419103.202002111333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 104,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.104/linux-headers-4.19.104-0419104_4.19.104-0419104.202002150649_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.104/linux-headers-4.19.104-0419104-generic_4.19.104-0419104.202002150649_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 105,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.105/linux-headers-4.19.105-0419105_4.19.105-0419105.202002191935_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.105/linux-headers-4.19.105-0419105-generic_4.19.105-0419105.202002191935_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 106,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.106/linux-headers-4.19.106-0419106_4.19.106-0419106.202002240333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.106/linux-headers-4.19.106-0419106-generic_4.19.106-0419106.202002240333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 107,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.107/linux-headers-4.19.107-0419107_4.19.107-0419107.202002281142_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.107/linux-headers-4.19.107-0419107-generic_4.19.107-0419107.202002281142_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 108,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.108/linux-headers-4.19.108-0419108_4.19.108-0419108.202003051136_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.108/linux-headers-4.19.108-0419108-generic_4.19.108-0419108.202003051136_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 109,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.109/linux-headers-4.19.109-0419109_4.19.109-0419109.202003111523_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.109/linux-headers-4.19.109-0419109-generic_4.19.109-0419109.202003111523_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 110,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.110/linux-headers-4.19.110-0419110_4.19.110-0419110.202003161032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.110/linux-headers-4.19.110-0419110-generic_4.19.110-0419110.202003161032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 111,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.111/linux-headers-4.19.111-0419111_4.19.111-0419111.202003180332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.111/linux-headers-4.19.111-0419111-generic_4.19.111-0419111.202003180332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 112,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.112/linux-headers-4.19.112-0419112_4.19.112-0419112.202003200739_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.112/linux-headers-4.19.112-0419112-generic_4.19.112-0419112.202003200739_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 113,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.113/linux-headers-4.19.113-0419113_4.19.113-0419113.202003250518_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.113/linux-headers-4.19.113-0419113-generic_4.19.113-0419113.202003250518_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 114,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.114/linux-headers-4.19.114-0419114_4.19.114-0419114.202004021037_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.114/linux-headers-4.19.114-0419114-generic_4.19.114-0419114.202004021037_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 19,
			Patch: 115,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.115/linux-headers-4.19.115-0419115_4.19.115-0419115.202004131241_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.19.115/linux-headers-4.19.115-0419115-generic_4.19.115-0419115.202004131241_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20/linux-headers-4.20.0-042000_4.20.0-042000.201812232030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20/linux-headers-4.20.0-042000-generic_4.20.0-042000.201812232030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.1/linux-headers-4.20.1-042001_4.20.1-042001.201901172023_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.1/linux-headers-4.20.1-042001-generic_4.20.1-042001.201901172023_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.2/linux-headers-4.20.2-042002_4.20.2-042002.201901171620_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.2/linux-headers-4.20.2-042002-generic_4.20.2-042002.201901171620_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.3/linux-headers-4.20.3-042003_4.20.3-042003.201901171122_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.3/linux-headers-4.20.3-042003-generic_4.20.3-042003.201901171122_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.4/linux-headers-4.20.4-042004_4.20.4-042004.201901222207_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.4/linux-headers-4.20.4-042004-generic_4.20.4-042004.201901222207_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.5/linux-headers-4.20.5-042005_4.20.5-042005.201901260434_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.5/linux-headers-4.20.5-042005-generic_4.20.5-042005.201901260434_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.6/linux-headers-4.20.6-042006_4.20.6-042006.201901310331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.6/linux-headers-4.20.6-042006-generic_4.20.6-042006.201901310331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.7/linux-headers-4.20.7-042007_4.20.7-042007.201902061234_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.7/linux-headers-4.20.7-042007-generic_4.20.7-042007.201902061234_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.8/linux-headers-4.20.8-042008_4.20.8-042008.201902121544_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.8/linux-headers-4.20.8-042008-generic_4.20.8-042008.201902121544_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.9/linux-headers-4.20.9-042009_4.20.9-042009.201902150331_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.9/linux-headers-4.20.9-042009-generic_4.20.9-042009.201902150331_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.10/linux-headers-4.20.10-042010_4.20.10-042010.201902150516_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.10/linux-headers-4.20.10-042010-generic_4.20.10-042010.201902150516_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.11/linux-headers-4.20.11-042011_4.20.11-042011.201902200535_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.11/linux-headers-4.20.11-042011-generic_4.20.11-042011.201902200535_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.12/linux-headers-4.20.12-042012_4.20.12-042012.201902230431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.12/linux-headers-4.20.12-042012-generic_4.20.12-042012.201902230431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.13/linux-headers-4.20.13-042013_4.20.13-042013.201902270533_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.13/linux-headers-4.20.13-042013-generic_4.20.13-042013.201902270533_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.14/linux-headers-4.20.14-042014_4.20.14-042014.201903051334_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.14/linux-headers-4.20.14-042014-generic_4.20.14-042014.201903051334_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.15/linux-headers-4.20.15-042015_4.20.15-042015.201903100332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.15/linux-headers-4.20.15-042015-generic_4.20.15-042015.201903100332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.16/linux-headers-4.20.16-042016_4.20.16-042016.201903132232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.16/linux-headers-4.20.16-042016-generic_4.20.16-042016.201903132232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 4,
			Minor: 20,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.17/linux-headers-4.20.17-042017_4.20.17-042017.201903190933_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v4.20.17/linux-headers-4.20.17-042017-generic_4.20.17-042017.201903190933_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0/linux-headers-5.0.0-050000_5.0.0-050000.201903032031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0/linux-headers-5.0.0-050000-generic_5.0.0-050000.201903032031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.1/linux-headers-5.0.1-050001_5.0.1-050001.201903100732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.1/linux-headers-5.0.1-050001-generic_5.0.1-050001.201903100732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.2/linux-headers-5.0.2-050002_5.0.2-050002.201903131832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.2/linux-headers-5.0.2-050002-generic_5.0.2-050002.201903131832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.3/linux-headers-5.0.3-050003_5.0.3-050003.201903191333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.3/linux-headers-5.0.3-050003-generic_5.0.3-050003.201903191333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.4/linux-headers-5.0.4-050004_5.0.4-050004.201903231634_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.4/linux-headers-5.0.4-050004-generic_5.0.4-050004.201903231634_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.5/linux-headers-5.0.5-050005_5.0.5-050005.201903271212_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.5/linux-headers-5.0.5-050005-generic_5.0.5-050005.201903271212_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.6/linux-headers-5.0.6-050006_5.0.6-050006.201904030534_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.6/linux-headers-5.0.6-050006-generic_5.0.6-050006.201904030534_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.7/linux-headers-5.0.7-050007_5.0.7-050007.201904052141_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.7/linux-headers-5.0.7-050007-generic_5.0.7-050007.201904052141_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.8/linux-headers-5.0.8-050008_5.0.8-050008.201904170732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.8/linux-headers-5.0.8-050008-generic_5.0.8-050008.201904170732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.9/linux-headers-5.0.9-050009_5.0.9-050009.201904200830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.9/linux-headers-5.0.9-050009-generic_5.0.9-050009.201904200830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.10/linux-headers-5.0.10-050010_5.0.10-050010.201904270832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.10/linux-headers-5.0.10-050010-generic_5.0.10-050010.201904270832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.11/linux-headers-5.0.11-050011_5.0.11-050011.201905020559_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.11/linux-headers-5.0.11-050011-generic_5.0.11-050011.201905020559_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.12/linux-headers-5.0.12-050012_5.0.12-050012.201905040834_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.12/linux-headers-5.0.12-050012-generic_5.0.12-050012.201905040834_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.13/linux-headers-5.0.13-050013_5.0.13-050013.201905051330_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.13/linux-headers-5.0.13-050013-generic_5.0.13-050013.201905051330_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.14/linux-headers-5.0.14-050014_5.0.14-050014.201905080630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.14/linux-headers-5.0.14-050014-generic_5.0.14-050014.201905080630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.15/linux-headers-5.0.15-050015_5.0.15-050015.201905101332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.15/linux-headers-5.0.15-050015-generic_5.0.15-050015.201905101332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.16/linux-headers-5.0.16-050016_5.0.16-050016.201905141431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.16/linux-headers-5.0.16-050016-generic_5.0.16-050016.201905141431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.17/linux-headers-5.0.17-050017_5.0.17-050017.201905161857_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.17/linux-headers-5.0.17-050017-generic_5.0.17-050017.201905161857_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.18/linux-headers-5.0.18-050018_5.0.18-050018.201906091640_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.18/linux-headers-5.0.18-050018-generic_5.0.18-050018.201906091640_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.19/linux-headers-5.0.19-050019_5.0.19-050019.201905251732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.19/linux-headers-5.0.19-050019-generic_5.0.19-050019.201905251732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.20/linux-headers-5.0.20-050020_5.0.20-050020.201905311031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.20/linux-headers-5.0.20-050020-generic_5.0.20-050020.201905311031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 0,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.21/linux-headers-5.0.21-050021_5.0.21-050021.201906040731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.0.21/linux-headers-5.0.21-050021-generic_5.0.21-050021.201906040731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1/linux-headers-5.1.0-050100_5.1.0-050100.201905052130_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1/linux-headers-5.1.0-050100-generic_5.1.0-050100.201905052130_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.1/linux-headers-5.1.1-050101_5.1.1-050101.201905110631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.1/linux-headers-5.1.1-050101-generic_5.1.1-050101.201905110631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.2/linux-headers-5.1.2-050102_5.1.2-050102.201905141830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.2/linux-headers-5.1.2-050102-generic_5.1.2-050102.201905141830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.3/linux-headers-5.1.3-050103_5.1.3-050103.201905161442_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.3/linux-headers-5.1.3-050103-generic_5.1.3-050103.201905161442_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.4/linux-headers-5.1.4-050104_5.1.4-050104.201905220633_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.4/linux-headers-5.1.4-050104-generic_5.1.4-050104.201905220633_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.5/linux-headers-5.1.5-050105_5.1.5-050105.201905251333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.5/linux-headers-5.1.5-050105-generic_5.1.5-050105.201905251333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.6/linux-headers-5.1.6-050106_5.1.6-050106.201905311031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.6/linux-headers-5.1.6-050106-generic_5.1.6-050106.201905311031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.7/linux-headers-5.1.7-050107_5.1.7-050107.201906091704_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.7/linux-headers-5.1.7-050107-generic_5.1.7-050107.201906091704_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.8/linux-headers-5.1.8-050108_5.1.8-050108.201906090832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.8/linux-headers-5.1.8-050108-generic_5.1.8-050108.201906090832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.9/linux-headers-5.1.9-050109_5.1.9-050109.201906111132_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.9/linux-headers-5.1.9-050109-generic_5.1.9-050109.201906111132_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.10/linux-headers-5.1.10-050110_5.1.10-050110.201906151034_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.10/linux-headers-5.1.10-050110-generic_5.1.10-050110.201906151034_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.11/linux-headers-5.1.11-050111_5.1.11-050111.201906171435_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.11/linux-headers-5.1.11-050111-generic_5.1.11-050111.201906171435_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.12/linux-headers-5.1.12-050112_5.1.12-050112.201906190731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.12/linux-headers-5.1.12-050112-generic_5.1.12-050112.201906190731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.13/linux-headers-5.1.13-050113_5.1.13-050113.201906220731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.13/linux-headers-5.1.13-050113-generic_5.1.13-050113.201906220731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.14/linux-headers-5.1.14-050114_5.1.14-050114.201906221030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.14/linux-headers-5.1.14-050114-generic_5.1.14-050114.201906221030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.15/linux-headers-5.1.15-050115_5.1.15-050115.201906250430_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.15/linux-headers-5.1.15-050115-generic_5.1.15-050115.201906250430_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.16/linux-headers-5.1.16-050116_5.1.16-050116.201907031232_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.16/linux-headers-5.1.16-050116-generic_5.1.16-050116.201907031232_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.17/linux-headers-5.1.17-050117_5.1.17-050117.201907231623_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.17/linux-headers-5.1.17-050117-generic_5.1.17-050117.201907231623_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.18/linux-headers-5.1.18-050118_5.1.18-050118.201907231625_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.18/linux-headers-5.1.18-050118-generic_5.1.18-050118.201907231625_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.19/linux-headers-5.1.19-050119_5.1.19-050119.201907231518_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.19/linux-headers-5.1.19-050119-generic_5.1.19-050119.201907231518_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.20/linux-headers-5.1.20-050120_5.1.20-050120.201907260839_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.20/linux-headers-5.1.20-050120-generic_5.1.20-050120.201907260839_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 1,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.21/linux-headers-5.1.21-050121_5.1.21-050121.201907280731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.1.21/linux-headers-5.1.21-050121-generic_5.1.21-050121.201907280731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2/linux-headers-5.2.0-050200_5.2.0-050200.201907231526_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2/linux-headers-5.2.0-050200-generic_5.2.0-050200.201907231526_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.1/linux-headers-5.2.1-050201_5.2.1-050201.201907231201_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.1/linux-headers-5.2.1-050201-generic_5.2.1-050201.201907231201_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.2/linux-headers-5.2.2-050202_5.2.2-050202.201907231250_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.2/linux-headers-5.2.2-050202-generic_5.2.2-050202.201907231250_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.3/linux-headers-5.2.3-050203_5.2.3-050203.201907260838_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.3/linux-headers-5.2.3-050203-generic_5.2.3-050203.201907260838_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.4/linux-headers-5.2.4-050204_5.2.4-050204.201907280731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.4/linux-headers-5.2.4-050204-generic_5.2.4-050204.201907280731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.5/linux-headers-5.2.5-050205_5.2.5-050205.201907310632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.5/linux-headers-5.2.5-050205-generic_5.2.5-050205.201907310632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.6/linux-headers-5.2.6-050206_5.2.6-050206.201908040832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.6/linux-headers-5.2.6-050206-generic_5.2.6-050206.201908040832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.7/linux-headers-5.2.7-050207_5.2.7-050207.201908061831_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.7/linux-headers-5.2.7-050207-generic_5.2.7-050207.201908061831_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.8/linux-headers-5.2.8-050208_5.2.8-050208.201908091630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.8/linux-headers-5.2.8-050208-generic_5.2.8-050208.201908091630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.9/linux-headers-5.2.9-050209_5.2.9-050209.201908160940_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.9/linux-headers-5.2.9-050209-generic_5.2.9-050209.201908160940_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.10/linux-headers-5.2.10-050210_5.2.10-050210.201908251538_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.10/linux-headers-5.2.10-050210-generic_5.2.10-050210.201908251538_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.11/linux-headers-5.2.11-050211_5.2.11-050211.201908290731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.11/linux-headers-5.2.11-050211-generic_5.2.11-050211.201908290731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.12/linux-headers-5.2.12-050212_5.2.12-050212.201909060933_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.12/linux-headers-5.2.12-050212-generic_5.2.12-050212.201909060933_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.13/linux-headers-5.2.13-050213_5.2.13-050213.201909060739_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.13/linux-headers-5.2.13-050213-generic_5.2.13-050213.201909060739_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.14/linux-headers-5.2.14-050214_5.2.14-050214.201909101030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.14/linux-headers-5.2.14-050214-generic_5.2.14-050214.201909101030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.15/linux-headers-5.2.15-050215_5.2.15-050215.201909160732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.15/linux-headers-5.2.15-050215-generic_5.2.15-050215.201909160732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.16/linux-headers-5.2.16-050216_5.2.16-050216.201909190832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.16/linux-headers-5.2.16-050216-generic_5.2.16-050216.201909190832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.17/linux-headers-5.2.17-050217_5.2.17-050217.201909210635_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.17/linux-headers-5.2.17-050217-generic_5.2.17-050217.201909210635_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.18/linux-headers-5.2.18-050218_5.2.18-050218.201910011000_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.18/linux-headers-5.2.18-050218-generic_5.2.18-050218.201910011000_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.19/linux-headers-5.2.19-050219_5.2.19-050219.201910050835_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.19/linux-headers-5.2.19-050219-generic_5.2.19-050219.201910050835_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.20/linux-headers-5.2.20-050220_5.2.20-050220.201910071832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.20/linux-headers-5.2.20-050220-generic_5.2.20-050220.201910071832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 2,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.21/linux-headers-5.2.21-050221_5.2.21-050221.201910111731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.2.21/linux-headers-5.2.21-050221-generic_5.2.21-050221.201910111731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3/linux-headers-5.3.0-050300_5.3.0-050300.201909152230_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3/linux-headers-5.3.0-050300-generic_5.3.0-050300.201909152230_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.1/linux-headers-5.3.1-050301_5.3.1-050301.201909210632_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.1/linux-headers-5.3.1-050301-generic_5.3.1-050301.201909210632_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.2/linux-headers-5.3.2-050302_5.3.2-050302.201910010731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.2/linux-headers-5.3.2-050302-generic_5.3.2-050302.201910010731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.3/linux-headers-5.3.3-050303_5.3.3-050303.201910051039_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.3/linux-headers-5.3.3-050303-generic_5.3.3-050303.201910051039_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.4/linux-headers-5.3.4-050304_5.3.4-050304.201910051526_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.4/linux-headers-5.3.4-050304-generic_5.3.4-050304.201910051526_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.5/linux-headers-5.3.5-050305_5.3.5-050305.201910071830_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.5/linux-headers-5.3.5-050305-generic_5.3.5-050305.201910071830_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.6/linux-headers-5.3.6-050306_5.3.6-050306.201910111731_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.6/linux-headers-5.3.6-050306-generic_5.3.6-050306.201910111731_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.7/linux-headers-5.3.7-050307_5.3.7-050307.201910180652_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.7/linux-headers-5.3.7-050307-generic_5.3.7-050307.201910180652_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.8/linux-headers-5.3.8-050308_5.3.8-050308.201910290940_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.8/linux-headers-5.3.8-050308-generic_5.3.8-050308.201910290940_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.9/linux-headers-5.3.9-050309_5.3.9-050309.201911061337_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.9/linux-headers-5.3.9-050309-generic_5.3.9-050309.201911061337_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.10/linux-headers-5.3.10-050310_5.3.10-050310.201911101133_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.10/linux-headers-5.3.10-050310-generic_5.3.10-050310.201911101133_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.11/linux-headers-5.3.11-050311_5.3.11-050311.201911121635_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.11/linux-headers-5.3.11-050311-generic_5.3.11-050311.201911121635_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.12/linux-headers-5.3.12-050312_5.3.12-050312.201911201137_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.12/linux-headers-5.3.12-050312-generic_5.3.12-050312.201911201137_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.13/linux-headers-5.3.13-050313_5.3.13-050313.201911240840_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.13/linux-headers-5.3.13-050313-generic_5.3.13-050313.201911240840_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.14/linux-headers-5.3.14-050314_5.3.14-050314.201911291040_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.14/linux-headers-5.3.14-050314-generic_5.3.14-050314.201911291040_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.15/linux-headers-5.3.15-050315_5.3.15-050315.201912041733_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.15/linux-headers-5.3.15-050315-generic_5.3.15-050315.201912041733_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.16/linux-headers-5.3.16-050316_5.3.16-050316.201912130343_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.16/linux-headers-5.3.16-050316-generic_5.3.16-050316.201912130343_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.17/linux-headers-5.3.17-050317_5.3.17-050317.201912171609_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.17/linux-headers-5.3.17-050317-generic_5.3.17-050317.201912171609_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 3,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.18/linux-headers-5.3.18-050318_5.3.18-050318.201912181133_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.3.18/linux-headers-5.3.18-050318-generic_5.3.18-050318.201912181133_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4/linux-headers-5.4.0-050400_5.4.0-050400.201911242031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4/linux-headers-5.4.0-050400-generic_5.4.0-050400.201911242031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.1/linux-headers-5.4.1-050401_5.4.1-050401.201911290555_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.1/linux-headers-5.4.1-050401-generic_5.4.1-050401.201911290555_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.2/linux-headers-5.4.2-050402_5.4.2-050402.201912042231_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.2/linux-headers-5.4.2-050402-generic_5.4.2-050402.201912042231_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.3/linux-headers-5.4.3-050403_5.4.3-050403.201912130841_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.3/linux-headers-5.4.3-050403-generic_5.4.3-050403.201912130841_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.4/linux-headers-5.4.4-050404_5.4.4-050404.201912172055_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.4/linux-headers-5.4.4-050404-generic_5.4.4-050404.201912172055_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.5/linux-headers-5.4.5-050405_5.4.5-050405.201912181630_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.5/linux-headers-5.4.5-050405-generic_5.4.5-050405.201912181630_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.6/linux-headers-5.4.6-050406_5.4.6-050406.201912211140_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.6/linux-headers-5.4.6-050406-generic_5.4.6-050406.201912211140_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.7/linux-headers-5.4.7-050407_5.4.7-050407.201912311234_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.7/linux-headers-5.4.7-050407-generic_5.4.7-050407.201912311234_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.8/linux-headers-5.4.8-050408_5.4.8-050408.202001041436_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.8/linux-headers-5.4.8-050408-generic_5.4.8-050408.202001041436_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.9/linux-headers-5.4.9-050409_5.4.9-050409.202001091039_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.9/linux-headers-5.4.9-050409-generic_5.4.9-050409.202001091039_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.10/linux-headers-5.4.10-050410_5.4.10-050410.202001091038_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.10/linux-headers-5.4.10-050410-generic_5.4.10-050410.202001091038_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.11/linux-headers-5.4.11-050411_5.4.11-050411.202001120738_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.11/linux-headers-5.4.11-050411-generic_5.4.11-050411.202001120738_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.12/linux-headers-5.4.12-050412_5.4.12-050412.202001141531_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.12/linux-headers-5.4.12-050412-generic_5.4.12-050412.202001141531_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.13/linux-headers-5.4.13-050413_5.4.13-050413.202001171431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.13/linux-headers-5.4.13-050413-generic_5.4.13-050413.202001171431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.14/linux-headers-5.4.14-050414_5.4.14-050414.202001230832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.14/linux-headers-5.4.14-050414-generic_5.4.14-050414.202001230832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.15/linux-headers-5.4.15-050415_5.4.15-050415.202001261031_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.15/linux-headers-5.4.15-050415-generic_5.4.15-050415.202001261031_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.16/linux-headers-5.4.16-050416_5.4.16-050416.202001300040_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.16/linux-headers-5.4.16-050416-generic_5.4.16-050416.202001300040_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.17/linux-headers-5.4.17-050417_5.4.17-050417.202002011032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.17/linux-headers-5.4.17-050417-generic_5.4.17-050417.202002011032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 18,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.18/linux-headers-5.4.18-050418_5.4.18-050418.202002051737_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.18/linux-headers-5.4.18-050418-generic_5.4.18-050418.202002051737_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 19,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.19/linux-headers-5.4.19-050419_5.4.19-050419.202002111332_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.19/linux-headers-5.4.19-050419-generic_5.4.19-050419.202002111332_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 20,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.20/linux-headers-5.4.20-050420_5.4.20-050420.202002151109_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.20/linux-headers-5.4.20-050420-generic_5.4.20-050420.202002151109_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 21,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.21/linux-headers-5.4.21-050421_5.4.21-050421.202002191431_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.21/linux-headers-5.4.21-050421-generic_5.4.21-050421.202002191431_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 22,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.22/linux-headers-5.4.22-050422_5.4.22-050422.202002240833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.22/linux-headers-5.4.22-050422-generic_5.4.22-050422.202002240833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 23,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.23/linux-headers-5.4.23-050423_5.4.23-050423.202002281329_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.23/linux-headers-5.4.23-050423-generic_5.4.23-050423.202002281329_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 24,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.24/linux-headers-5.4.24-050424_5.4.24-050424.202003051135_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.24/linux-headers-5.4.24-050424-generic_5.4.24-050424.202003051135_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 25,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.25/linux-headers-5.4.25-050425_5.4.25-050425.202003121333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.25/linux-headers-5.4.25-050425-generic_5.4.25-050425.202003121333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 26,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.26/linux-headers-5.4.26-050426_5.4.26-050426.202003180732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.26/linux-headers-5.4.26-050426-generic_5.4.26-050426.202003180732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 27,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.27/linux-headers-5.4.27-050427_5.4.27-050427.202003210835_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.27/linux-headers-5.4.27-050427-generic_5.4.27-050427.202003210835_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 28,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.28/linux-headers-5.4.28-050428_5.4.28-050428.202003250833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.28/linux-headers-5.4.28-050428-generic_5.4.28-050428.202003250833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 29,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.29/linux-headers-5.4.29-050429_5.4.29-050429.202004010635_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 30,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.30/linux-headers-5.4.30-050430_5.4.30-050430.202004021433_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 31,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.31/linux-headers-5.4.31-050431_5.4.31-050431.202004080434_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 4,
			Patch: 32,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.4.32/linux-headers-5.4.32-050432_5.4.32-050432.202004130937_all.deb",
			LinuxHeaderGeneric: "",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5/linux-headers-5.5.0-050500_5.5.0-050500.202001262030_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5/linux-headers-5.5.0-050500-generic_5.5.0-050500.202001262030_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.1/linux-headers-5.5.1-050501_5.5.1-050501.202002011032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.1/linux-headers-5.5.1-050501-generic_5.5.1-050501.202002011032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.2/linux-headers-5.5.2-050502_5.5.2-050502.202002041931_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.2/linux-headers-5.5.2-050502-generic_5.5.2-050502.202002041931_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.3/linux-headers-5.5.3-050503_5.5.3-050503.202002110832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.3/linux-headers-5.5.3-050503-generic_5.5.3-050503.202002110832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.4/linux-headers-5.5.4-050504_5.5.4-050504.202002150446_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.4/linux-headers-5.5.4-050504-generic_5.5.4-050504.202002150446_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 5,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.5/linux-headers-5.5.5-050505_5.5.5-050505.202002191432_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.5/linux-headers-5.5.5-050505-generic_5.5.5-050505.202002191432_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 6,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.6/linux-headers-5.5.6-050506_5.5.6-050506.202002240832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.6/linux-headers-5.5.6-050506-generic_5.5.6-050506.202002240832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 7,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.7/linux-headers-5.5.7-050507_5.5.7-050507.202002281805_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.7/linux-headers-5.5.7-050507-generic_5.5.7-050507.202002281805_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 8,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.8/linux-headers-5.5.8-050508_5.5.8-050508.202003051633_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.8/linux-headers-5.5.8-050508-generic_5.5.8-050508.202003051633_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 9,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.9/linux-headers-5.5.9-050509_5.5.9-050509.202003120738_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.9/linux-headers-5.5.9-050509-generic_5.5.9-050509.202003120738_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 10,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.10/linux-headers-5.5.10-050510_5.5.10-050510.202003180732_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.10/linux-headers-5.5.10-050510-generic_5.5.10-050510.202003180732_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 11,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.11/linux-headers-5.5.11-050511_5.5.11-050511.202003210837_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.11/linux-headers-5.5.11-050511-generic_5.5.11-050511.202003210837_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 12,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.12/linux-headers-5.5.12-050512_5.5.12-050512.202003250833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.12/linux-headers-5.5.12-050512-generic_5.5.12-050512.202003250833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 13,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.13/linux-headers-5.5.13-050513_5.5.13-050513.202003251631_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.13/linux-headers-5.5.13-050513-generic_5.5.13-050513.202003251631_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 14,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.14/linux-headers-5.5.14-050514_5.5.14-050514.202004011032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.14/linux-headers-5.5.14-050514-generic_5.5.14-050514.202004011032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 15,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.15/linux-headers-5.5.15-050515_5.5.15-050515.202004021032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.15/linux-headers-5.5.15-050515-generic_5.5.15-050515.202004021032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 16,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.16/linux-headers-5.5.16-050516_5.5.16-050516.202004080832_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.16/linux-headers-5.5.16-050516-generic_5.5.16-050516.202004080832_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 5,
			Patch: 17,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.17/linux-headers-5.5.17-050517_5.5.17-050517.202004130833_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.5.17/linux-headers-5.5.17-050517-generic_5.5.17-050517.202004130833_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 6,
			Patch: 0,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6/linux-headers-5.6.0-050600_5.6.0-050600.202003292333_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6/linux-headers-5.6.0-050600-generic_5.6.0-050600.202003292333_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 6,
			Patch: 1,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.1/linux-headers-5.6.1-050601_5.6.1-050601.202004011032_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.1/linux-headers-5.6.1-050601-generic_5.6.1-050601.202004011032_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 6,
			Patch: 2,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.2/linux-headers-5.6.2-050602_5.6.2-050602.202004020822_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.2/linux-headers-5.6.2-050602-generic_5.6.2-050602.202004020822_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 6,
			Patch: 3,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.3/linux-headers-5.6.3-050603_5.6.3-050603.202004080837_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.3/linux-headers-5.6.3-050603-generic_5.6.3-050603.202004080837_amd64.deb",
		}, 
		model.KernelVersion{
			Major: 5,
			Minor: 6,
			Patch: 4,
		}: {
			LinuxHeaderAll:     "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.4/linux-headers-5.6.4-050604_5.6.4-050604.202004131234_all.deb",
			LinuxHeaderGeneric: "https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.6.4/linux-headers-5.6.4-050604-generic_5.6.4-050604.202004131234_amd64.deb",
		}, 
	},
}
