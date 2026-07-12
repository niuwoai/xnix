#!/bin/sh
set -eu

target_dir=$1
test_public_key=${XNIX_TEST_SSH_PUBLIC_KEY:-}

[ -z "$test_public_key" ] && exit 0

if [ ! -r "$test_public_key" ]; then
    echo "Xnix SSH test public key is unavailable: $test_public_key" >&2
    exit 1
fi

install -d -m 0700 "$target_dir/root/.ssh"
install -m 0600 "$test_public_key" "$target_dir/root/.ssh/authorized_keys"
