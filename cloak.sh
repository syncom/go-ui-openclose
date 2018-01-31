#!/bin/bash

KERNEL_VER=$(uname -r)
VISIBLE_DRIVE=/dev/nvme0n1
HIDDEN_DRIVE=/dev/nvme0n2
ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TOKEN_FILE=${ROOT_DIR}/authentication_token.key
NVME_KO=/lib/modules/${KERNEL_VER}/kernel/drivers/nvme/host/nvme.ko
NS1_MOUNT_POINT=/mnt/nvme

COMMAND="nvme admin-passthru ${VISIBLE_DRIVE} -o 0xC1 -n 1 "\
"--cdw11=1 --cdw10=8 -l 512 -i ${TOKEN_FILE} -w -s"

echo $COMMAND

eval ${COMMAND} > /dev/null 2>&1 || exit 1
sleep 1
umount ${NS1_MOUNT_POINT} > /dev/null 2>&1  || exit 2
sleep 1
rmmod nvme > /dev/null 2>&1 || exit 3
sleep 1
insmod ${NVME_KO} > /dev/null 2>&1 || exit 4
sleep 1
mount ${VISIBLE_DRIVE} ${NS1_MOUNT_POINT} > /dev/null 2>&1 || exit 5

echo "Data cloaked"

