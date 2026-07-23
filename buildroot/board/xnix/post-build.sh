#!/bin/sh
set -eu

target_dir=$1
test_public_key=${XNIX_TEST_SSH_PUBLIC_KEY:-}
build_dir=${BUILD_DIR:-}

install_wine_nls_files() {
    [ -n "$build_dir" ] || return 0

    for nls_marker in "$build_dir"/wine-*/nls/l_intl.nls; do
        [ -r "$nls_marker" ] || continue

        nls_source_dir=$(dirname "$nls_marker")
        nls_target_dir="$target_dir/usr/share/wine/nls"
        install -d -m 0755 "$nls_target_dir"
        cp "$nls_source_dir"/*.nls "$nls_target_dir"/
        chmod 0644 "$nls_target_dir"/*.nls
        return 0
    done
}

install_wine_nls_files

[ -z "$test_public_key" ] && exit 0

if [ ! -r "$test_public_key" ]; then
    echo "Xnix SSH test public key is unavailable: $test_public_key" >&2
    exit 1
fi

install -d -m 0700 "$target_dir/root/.ssh"
install -m 0600 "$test_public_key" "$target_dir/root/.ssh/authorized_keys"
